// Primitive version of the hopefully final version

package main

import (
	"os"
	"fmt"
	"bufio"
	"flag"
	e "errors"
	s "strings"
)

var p = fmt.Println	//Does not allow for %s
var sc = fmt.Scanln

func read_file(version string) ([]string, error){

	file, err := os.Open("./bible/"+version)
    if err != nil {
		file, err = os.Open("./../share/bible/"+version)	//In the case of nix derivations where the folders are bin and share
		if err != nil {
			return nil, e.New("Error opening file: ./bible/"+version+ "and ./../share/bible/"+version)
		}
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	var bible []string //Number of verses in the bible
    for scanner.Scan() {	//Goes line by line
        //fmt.Println(scanner.Text())
		//fmt.Println("\n")
		bible = append(bible, scanner.Text())	//Make array with all of the verses
    }

    if err := scanner.Err(); err != nil {
		return nil, e.New("Error reading file:")
    }
	return bible, nil
}

// The name of the book of a specific verse
func first_words(value string, n int) string{	//Int is the first n words

    words := s.Fields(value)

	if len(words) < n {
		return value
	}

	firstN := words[:n]

	return s.Join(firstN, " ")
}

func removeWord(sentence, word string) string {
    parts := s.Fields(sentence)
    filtered := []string{}
    for _, p := range parts {
        if p != word {
            filtered = append(filtered, p)
        }
    }
    return s.Join(filtered, " ")
}

func print_verse(input string, slice []string) (string, error){	//consider taking book and verse as input and not the current "input", this avoids doing first_words thrice
	if len(s.Fields(input)) > 3 {

		return "", e.New("More than 3 words in the book name")
	}

	if len(s.Fields(input)) == 3 {	//Dont want to remove ALL the spaces
		input = s.Replace(input, " ", "", 1)
	}
	for i, quote := range slice {
		tag := first_words(slice[i], 2)
		if tag == input {
			quote = removeWord(quote, first_words(quote, 1))	//cant remove 2 at a time
			quote = removeWord(quote, first_words(quote, 1))
			return quote, nil
		}
	}
	return "", e.New("ERROR, Verse not found")
}

func init() {	//Currently does nothing
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [options] <Book> <Chapter:Verse>\n\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "Options:\n")
        flag.PrintDefaults()
        fmt.Fprintf(os.Stderr, "\nExamples:\n")
        fmt.Fprintf(os.Stderr, "  %s Genesis 1:1\n", os.Args[0])
        //fmt.Fprintf(os.Stderr, "  %s -lang=ES John 3:16\n", os.Args[0])	//lang is not implemented
    }
}

func Run() {
	bible, err := read_file("kjv_preformatted.txt")
	if err != nil {
		p(err)
		return
	}

	p(bible)
}

func main() {
	/*
	Run()
	flag.Bool("read", false, "Read continuously, without other inputs will default to Genesis 1:1")
	flag.Parse()	//Check for any optional flags, right now there arent any

    args := flag.Args()
    if len(args) < 2 {	//error catch
        fmt.Println("Usage: bible [flags] <Book> <Chapter:Verse>")
        flag.PrintDefaults()
        os.Exit(1)
    }

	tag := s.Join(args[0:2], " ")
	quote, err := print_verse(tag, bible)

	if err!= nil {
		p(err, "\n")
		fmt.Printf("Interpreted the input as '%s' and couldn't find it\n", tag)	//Printf and not Println
		return
	}
	p("\n" + quote + "\n")
	*/
}