package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("8.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(dat), "\n")
	for iline, line := range lines {
		if len(line) == 0 {
			continue
		}
		fmt.Println(len(line))

		blocks := make(map[string]bool)
		for i := 16; i < len(line); i += 16 {
			if _, check := blocks[line[i-16:i]]; check {
				fmt.Println("duplicate found", iline, line[i-16:i])
			}
			blocks[line[i-16:i]] = true
		}
	}
}
