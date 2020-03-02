package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

var letterFreq = map[byte]float64{
	' ': 0.18316857525301761,
	'E': 0.10217877079731477,
	'T': 0.07509993977246371,
	'A': 0.06553070586681289,
	'O': 0.06200554045501006,
	'N': 0.05703083740730461,
	'I': 0.05734255240415464,
	'S': 0.05326267384545168,
	'R': 0.04971999259117901,
	'H': 0.04862209247937962,
	'L': 0.03356165502675509,
	'D': 0.03352273770645298,
	'U': 0.02295200399609547,
	'C': 0.02265088364428036,
	'M': 0.02017270369335810,
	'F': 0.01971808878271046,
	'W': 0.01689613959174190,
	'G': 0.01635866071660768,
	'P': 0.01503115599710262,
	'Y': 0.01469954633196355,
	'B': 0.01270765661547618,
	'V': 0.00788048153834165,
	'K': 0.00569167117510226,
	'X': 0.00149808323785044,
	'J': 0.00114405442279746,
	'Q': 0.00088093022255462,
	'Z': 0.00059793009448438,
}

func calculateScore(arg []byte, key byte) float64 {
	letterCounts := make(map[byte]int)
	letters := 0.0
	for _, c := range arg {
		c ^= key
		if c < byte(32) && c != '\n' && c != '\r' && c != '	' || c > byte(126) {
			return -1.0
		}
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		letterCounts[c]++
		letters += 1.0
	}

	score := 0.0
	for k, f := range letterFreq {
		expected := f * letters
		diff := float64(letterCounts[k]) - expected
		score += (diff * diff) / letters
	}

	return score
}

func main() {
	dat, err := ioutil.ReadFile("4.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(dat), "\n")
	var bestLine string
	bestScore := math.MaxFloat64

	for _, line := range lines {
		inputb, err := hex.DecodeString(line)
		if err != nil {
			log.Fatal(err)
		}

		min := math.MaxFloat64
		var key byte
		for k := byte(32); k < byte(127); k++ {
			score := calculateScore(inputb, k)
			if score >= 0.0 && score < min {
				min = score
				key = k
			}
		}

		if min < bestScore {
			bestScore = min
			for i := range inputb {
				inputb[i] ^= key
			}
			bestLine = string(inputb)
			fmt.Println(bestScore, bestLine)
		}
	}
	fmt.Println(bestLine)
}
