package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	DATASET_LOC = "words_dictionary.json"
)

func makeshort(fullword string) string {
	short := ""
	for _, letter := range fullword {
		if strings.Contains("ikmoqsvwxyz", string(letter)) {
			continue
		}
		short += string(letter)
	}
	return short
}

func main() {
	data, err := ioutil.ReadFile(DATASET_LOC)
	if err != nil {
		log.Fatal(err)
	}
	dict := map[string]int{}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		log.Fatal(err)
	}
	newd := map[string][]string{}
	for fullword := range dict {
		short := makeshort(fullword)
		prevList, exists := newd[short]
		if !exists {
			prevList = []string{}
		}
		prevList = append(prevList, fullword)
		newd[short] = prevList
	}
	log.Printf("Created %d shortforms", len(newd))
	scanner := bufio.NewScanner(os.Stdin)
	sc := func() bool {
		fmt.Fprint(os.Stderr, "> ")
		return scanner.Scan()
	}
	for sc() {
		text := scanner.Text()

		parts := strings.Split(text, " ")

		if len(parts) >= 3 && parts[0] == "run-command" {
			if parts[1] == "find" {
				_, yesitdoes := dict[parts[2]]
				if yesitdoes {
					log.Println("Yes it does!")
				} else {
					log.Println("Double nope")
				}
			}
			if parts[1] == "short" {
				fmt.Println(makeshort(strings.Join(parts[2:], " ")))
			}
			continue
		}

		for _, part := range parts {
			log.Printf("Searching for '%s'\n", part)

			lis, exists := newd[part]
			if !exists {
				log.Println("nope")
			} else {
				fmt.Println(lis)
			}
		}
	}
}
