package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
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

func strxor(str []byte, key byte) string {
	resb := make([]byte, len(str))
	for i, c := range str {
		d := c ^ key
		if d < 32 || d > 126 {
			return ""
		}
		resb[i] = d
	}
	return strings.ToUpper(string(resb))
}

func calculateScore(a string) float64 {
	score := 0.0
	letters := 0
	for _, c := range a {
		f := letterFreq[byte(c)]
		if f > 0.0 {
			letters++
		}
		score += f

	}

	return score
}

func main() {
	inputb, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}

	max := 0.0
	var key byte
	var sent string
	for k := byte(0); k < byte(128); k++ {
		res := strxor(inputb, k)
		if len(res) == 0 {
			continue
		}
		score := calculateScore(res)
		fmt.Println(score, res)
		if score > max {
			sent = res
			max = score
			key = k
		}
	}
	fmt.Println(max, key, string(key), sent)

	fmt.Println(reflect.TypeOf(0.0))
}
