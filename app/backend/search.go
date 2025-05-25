//Code responsible for searching and returning querys like 1Kings 1:3 and returning said verse

package backend

import (
	"os"
	"fmt"
	"bufio"
	"flag"
	e "errors"
	s "strings"
	u "unicode"	//Right now only used for isdigit()
)

var p = fmt.Println	//Does not allow for %s
var sc = fmt.Scanln

func read_file(version string) ([]string, error){

	file, err := os.Open("./../../bible/"+version)
    if err != nil {
		file, err = os.Open("./../share/bible/"+version)	//In the case of nix derivations where the folders are bin and share
		if err != nil {
			return nil, e.New("Error opening file: ./../../bible/"+version+ "or ./../share/bible/"+version)
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

func get_verse(input string, slice []string) (string, error){	//consider taking book and verse as input and not the current "input", this avoids doing first_words thrice
	if len(s.Fields(input)) > 3 {

		return "", e.New("More than 3 words in the book name")
	}

	if len(input) == 0 {
		return "", e.New("Can't lookup an empty query")
	}

	for i, quote := range slice {
		//Cases like 1Kings	//isdigit only counts base 10 ints, not roman numerals, fracs etc
		switch u.IsDigit(rune(input[0])){
		case false:
			tag := first_words(slice[i], 2)
			if tag == input {
				quote = removeWord(quote, first_words(quote, 1))	//cant remove 2 at a time
				quote = removeWord(quote, first_words(quote, 1))
				return quote, nil
			}
		case true:	//Cases like 1 Kings 2:3
			tag := first_words(slice[i], 3)

			if input[1] != ' ' {	//If the user typed 1Kings instead of 1 Kings for example
				input = input[:1] + " " + input[1:]
			}

			if tag == input {
				quote = removeWord(quote, first_words(quote, 1))	//cant remove 2 at a time
				quote = removeWord(quote, first_words(quote, 1))
				quote = removeWord(quote, first_words(quote, 1))
				return quote, nil
			}

		}
	}
	return "", fmt.Errorf("ERROR, Verse '%s' not found", input)	//This allows %s formatting, e.New doesnt
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

func Search(tag string) (string, error){	//Gets called from other files
	bible, err := read_file("kjv_preformatted.txt")
	if err != nil {
		//p(err)
		return "", err
	}

	quote, err := get_verse(tag, bible)

	if err!= nil {
		//p(err, "\n")
		//quote, _  := get_verse("Genesis 1:1", bible)
		return "", err
	}
	get_verse("1Kings 2:3", bible)
	return quote, nil
}