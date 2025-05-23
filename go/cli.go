// Primitive version of the hopefully final version

package main

import (
	"os"
	"fmt"
	"bufio"
	e "errors"
	s "strings"
)

var p = fmt.Println
var sc = fmt.Scanln

// The name of the book of a specific verse
func first_words(value string, n int) string{	//Int is the first n words

    words := s.Fields(value)

	if len(words) < n {
		return value
	}

	firstN := words[:n]

	return s.Join(firstN, " ")
}



func print_verse(input string, slice []string) (string, error){
	if len(s.Fields(input)) > 3 {

		return "More than 3 words in the book name", e.New("More than 3 words in the book name")
	}

	if len(s.Fields(input)) == 3 {	//Dont want to remove ALL the spaces
		input = s.Replace(input, " ", "", 1)
	}
	for i := range slice {
		tag := first_words(slice[i], 2)
		if tag == input {
			return slice[i], nil
		}
	}
	return "Verse not found", e.New("Verse not found")
}

func main() {

    file, err := os.Open("./../bible/kjv_preformatted.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
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
        fmt.Println("Error reading file:", err)
    }


	var book, chapter string
	tag := book + " " + chapter
	//p(print_verse(tag, bible))	//p to spot any errors
	//p(first_words("Revelation 22:21	The grace", 2))

	var param string
	if len(os.Args) == 2 {
		param = os.Args[1]
		p(print_verse(param)[0])
	} else if len(os.Args) < 2 {
		p("Usage: vible-cli <parameter>\n")
		p("Type a verse to lookup e.g '1Kings 1:2 (exactly that format)'")
		return
	} else {
		fmt.Println("Type a verse to lookup e.g '1Kings 1:2 (exactly that format)'")
		fmt.Scan(&book, &chapter)
	}
}