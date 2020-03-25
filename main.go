package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// flags
var cpm bool    // toggle CPM or WPM
var list string // input sentence list to be used
var rounds int  // how many sentences to be tested on

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	parseFlags()
	countdown(3)
	clear()
	wpm, errors := playRound(rounds)
	fmt.Println("Result:", resultStats(wpm, errors))
}

func parseFlags() {
	flag.BoolVar(&cpm, "cpm", false, "Use CPM instead of WPM")
	flag.StringVar(&list, "list", "shakespeare", "Input sentence list to be used")
	flag.IntVar(&rounds, "rounds", 3, "How many sentences to be tested on")
	flag.Parse()
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

		if tWPM != 0 { // calculate average wpm
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
	fmt.Println(" " + s)
	start := time.Now() // start the timer
	result := input(":")
	t := time.Now()                 // end the timer
	elapsed := t.Sub(start)         // calculate time elapsed
	wpm := calcWPM(s, elapsed)      // calculate wpm
	errors := calcErrors(s, result) // calculate errors
	return wpm, errors
}

func input(s string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	text, err := reader.ReadString('\n') // get user input
	if err != nil {
		log.Fatal(err)
	}

	return text
}

func resultStats(wpm float64, errorCount int) string {
	unit := "WPM"
	if cpm {
		unit = "CPM"
		wpm *= 5
	}

	return fmt.Sprintf(" %.1f %s | %d error%s", wpm, unit, errorCount, plural(errorCount))
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
	raws, err := ioutil.ReadFile(path + "sentences/" + list)
	if err != nil {
		log.Fatal(err)
	}

	sentences := strings.Split(string(raws), "\n")
	sentence := rand.Intn(len(sentences) - 1) // picks a random sentence from the given file
	return sentences[sentence]

}

func directory() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	exPath := filepath.Dir(ex) + "/" // find program's working directory
	return exPath
}

func clear() {
	print("\033[H\033[2J")
}
