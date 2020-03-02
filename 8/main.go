package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func hummingDistance(astr, bstr []byte) int {
	if len(astr) != len(bstr) {
		log.Fatal("len should be equal")
	}

	hummingDistanceByte := func(ab, bb byte) int {
		result := 0
		xor := ab ^ bb
		for ; xor > 0; xor >>= 2 {
			switch xor & 3 {
			case 1, 2:
				result++
			case 3:
				result += 2
			}
		}
		return result
	}

	result := 0
	for i := 0; i < len(astr); i++ {
		result += hummingDistanceByte(astr[i], bstr[i])
	}
	return result
}

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

		blocks := make(map[string]bool)
		for i := 16; i < len(line); i += 16 {
			if _, check := blocks[line[i-16:i]]; check {
				fmt.Println("duplicate found", iline, line[i-16:i])
			}
			blocks[line[i-16:i]] = true
		}

		minDistance := math.MaxFloat64
		bline := []byte(line)
		var mink int
		for k := 2; k <= 40; k++ {
			total := 0.0
			times := 0
			for i := k + k; i < len(bline); i += k {
				total += float64(hummingDistance(bline[i-k-k:i-k], bline[i-k:i])) / float64(k)
				times++
			}
			av := total / float64(times)
			fmt.Println(k, av)
			if av < minDistance {
				minDistance = av
				mink = k
			}
		}

		fmt.Println("!###", iline, mink)
	}
}
