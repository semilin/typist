package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"time"
	"io/ioutil"
	"path/filepath"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	countdown(3); clear()
	wpm, errors := playRound(3)
	fmt.Printf("Result: %g WPM | %d error(s)\n",wpm,errors)
}

func countdown(length int) {
	for i:=length;i>0;i-- {
		clear()
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func playRound(rounds int) (float64, int) {
	tWPM := float64(0)
	tErrors := int(0)
	
	for i:=0;i<rounds;i++ {
		clear()
		fmt.Printf(" %g WPM | %d error(s)\n\n\n",tWPM,tErrors)
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

func ttest(s string) (float64, int) {
	start := time.Now()
	fmt.Println(" " + s)
	result := input(":")
	t := time.Now()
	elapsed := t.Sub(start)
	wpm := calcWPM(s, elapsed)
	errors := calcErrors(s,result)
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

func calcWPM(s string, elapsed time.Duration) float64 {
	wpm := float64((float64(strings.Count(s,"")) / float64(elapsed.Seconds()) * 12))
	return wpm
}

func calcErrors(expected string, result string) int {
	owords := strings.Split(expected," ")
	rwords := strings.Split(result," ")
	e := 0

	for i:=0;i<len(owords);i++ {
		if len(rwords) <= i {
			return e + (len(owords) - len(rwords))
		} else if owords[i] != rwords[i] {
			e++
		} 
	}

	return e-1
}

func getSentence() string {
	path := directory()
	raws, err := ioutil.ReadFile(path + "sentences/shakespeare")
	if err != nil {log.Fatal(err)}
	
	sentences := strings.Split(string(raws),"\n")
	sentence := rand.Intn(len(sentences)-1)
	return sentences[sentence]
	
	
}

func directory() string{
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