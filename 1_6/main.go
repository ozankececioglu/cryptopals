package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

const (
	cipher_6_1 = "DO NOT OPEN THE DOOR!"
	cipher_6_2 = "ASCII Table and Description"
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

func calculateScore(arg []byte, key byte, start, step int) float64 {
	letterCounts := make(map[byte]int)
	for i := start; i < len(arg); i += step {
		c := arg[i] ^ key
		if c < byte(32) && c != '\n' && c != '\r' && c != '	' || c > byte(126) {
			return -1.0
		}
		if c >= 'a' && c <= 'z' {
			c -= 32
		}
		letterCounts[c]++
	}

	score := 0.0
	letters := math.Ceil((float64(len(arg)) - float64(start)) / float64(step))
	for k, f := range letterFreq {
		expected := f * letters
		diff := float64(letterCounts[k]) - expected
		score += (diff * diff) / letters
	}

	return score
}

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

func repeatingXor(input, key []byte) []byte {
	for i := 0; i < len(input); i++ {
		input[i] ^= key[i%len(key)]
	}
	return input
}

func main2() {
	inp := ""
	keyword := ""
	out := base64.StdEncoding.EncodeToString(repeatingXor([]byte(inp), []byte(keyword)))
	ioutil.WriteFile("6_4.txt", []byte(out), 777)
}

func main() {
	dat, err := ioutil.ReadFile("6_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	raw, err := base64.StdEncoding.DecodeString(string(dat))

	minDistance := math.MaxFloat64
	var mink int
	for k := 2; k <= 40; k++ {
		total := 0.0
		times := 0
		for i := k + k; i < len(raw); i += k {
			total += float64(hummingDistance(raw[i-k-k:i-k], raw[i-k:i])) / float64(k)
			times++
		}
		av := total / float64(times)
		fmt.Println(k, av)
		if av < minDistance {
			minDistance = av
			mink = k
		}
	}

	fmt.Println("key", mink, minDistance)

	cipher := make([]byte, mink)
	fail := false
	for i := 0; i < mink; i++ {
		minScore := math.MaxFloat64
		fmt.Println("--------", i)
		for x := byte(32); x < byte(128); x++ {
			s := calculateScore(raw, x, i, mink)
			if s >= 0.0 && s < minScore {
				if x > 32 {
					fmt.Println("!###", string(x), s)
				} else {
					fmt.Println("!###", x, s)
				}

				minScore = s
				cipher[i] = x
			}
		}
		if cipher[i] == 0 {
			fail = true
			log.Fatal("cipher not found", i)
			break
		}
	}
	if !fail {
		fmt.Println(string(cipher))
		fmt.Println(string(repeatingXor(raw, cipher)))
	} else {

	}
}
