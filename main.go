package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	countdown(3)
	clear()
	wpm, errors := playRound(3)
	fmt.Println("Result:", resultStats(wpm, errors))
}

func countdown(length int) {
	for i := length; i > 0; i-- {
		clear()
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func playRound(rounds int) (float64, int) {
	var tWPM float64 = 0
	var tErrors int = 0

	for i := 0; i < rounds; i++ {
		clear()
		fmt.Println(resultStats(tWPM, tErrors), "\n\n\n")
		wpm, errors := ttest(getSentence())

		if tWPM != 0 {
			tWPM = (tWPM + wpm) / 2
		} else {
			tWPM = wpm
		}
		tErrors += errors
	}

	return tWPM, tErrors
}

func plural(n int) string {
	if n == 1 {
		return ""
	} else {
		return "s"
	}
}

func ttest(s string) (float64, int) {
	start := time.Now()
	fmt.Println(" " + s)
	result := input(":")
	t := time.Now()
	elapsed := t.Sub(start)
	wpm := calcWPM(s, elapsed)
	errors := calcErrors(s, result)
	return wpm, errors
}

func input(s string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return text
}

func resultStats(wpm float64, errorCount int) string {
	return fmt.Sprintf(" %.1f WPM | %d error%s", wpm, errorCount, plural(errorCount))
}

func calcWPM(s string, elapsed time.Duration) float64 {
	chars := strings.Count(s, "")                      // count the number of characters in the string (sentence)
	cps := float64(chars) / float64(elapsed.Seconds()) // calculate the characters per second
	cpm := cps * 60                                    // calculate characters per minute
	wpm := cpm / 5                                     // convert cpm to wpm (5 characters per word)

	return wpm
}

func calcErrors(expected string, result string) int {
	expectedWords := strings.Split(expected, " ")
	resultWords := strings.Split(result, " ")
	errorCount := 0

	for i := 0; i < len(expectedWords); i++ {
		if len(resultWords) <= i {
			return errorCount + (len(expectedWords) - len(resultWords))
		} else if expectedWords[i] != resultWords[i] {
			errorCount++
		}
	}

	return errorCount - 1
}

func getSentence() string {
	path := directory()
	raws, err := ioutil.ReadFile(path + "sentences/shakespeare")
	if err != nil {
		log.Fatal(err)
	}

	sentences := strings.Split(string(raws), "\n")
	sentence := rand.Intn(len(sentences) - 1)
	return sentences[sentence]

}

func directory() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	exPath := filepath.Dir(ex) + "/"
	return exPath
}

func clear() {
	print("\033[H\033[2J")
}
