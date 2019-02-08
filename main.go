package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	DATASET_LOC = "/home/eric/files/DATA/datasets/words_dictionary.json"
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
		logrus.Fatal(err)
	}
	dict := map[string]int{}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		logrus.Fatal(err)
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
	logrus.Infof("Created %d shortforms", len(newd))
	scanner := bufio.NewScanner(os.Stdin)
	sc := func() bool {
		fmt.Print("> ")
		return scanner.Scan()
	}
	for sc() {
		text := scanner.Text()

		parts := strings.Split(text, " ")

		if len(parts) >= 3 && parts[0] == "run-command" {
			logrus.Info("Hey whachya doin?")
			if parts[1] == "find" {
				_, yesitdoes := dict[parts[2]]
				if yesitdoes {
					logrus.Info("Yes it does!")
				} else {
					logrus.Warn("Double nope")
				}
			}
			if parts[1] == "short" {
				logrus.Info(makeshort(strings.Join(parts[2:], " ")))
			}
			continue
		}

		for _, part := range parts {
			logrus.Infof("Searching for '%s'", part)

			lis, exists := newd[part]
			if !exists {
				logrus.Warn("nope")
			} else {
				fmt.Println(lis)
			}
		}
	}
}
