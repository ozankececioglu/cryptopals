package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

func calculateScore(arg <-chan byte) float64 {
	letterCounts := make(map[byte]int)
	letters := 0.0
	for c := range arg {
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

func main() {
	dat, err := ioutil.ReadFile("6.txt")
	if err != nil {
		log.Fatal(err)
	}

	raw, err := base64.StdEncoding.DecodeString(string(dat))
	if err != nil {
		log.Fatal(err)
	}

	min := math.MaxFloat64
	var mink int
	for k := 2; k <= 40; k++ {
		total := 0.0
		times := 0
		for i := k + k; i < len(raw); i += k {
			total += float64(hummingDistance(raw[i - k - k:i - k], raw[i - k:i])) / float64(k)
			times++
		}
		av := total / float64(times)
		fmt.Println(k, av)
		if av < min {
			min = av
			mink = k
		}
	}

	fmt.Println(min, mink)

	cipher := make([]byte, mink)
	fail := false
	for i := 0; i < mink; i++ {
		minScore := math.MaxFloat64
		for x := byte(0); x < byte(128); x++ {
			chnl := make(chan byte)
			go func() {
				for j := i; j < len(raw); j += mink {
					chnl <- raw[j] ^ x
				}
				close(chnl)
			}()
			s := calculateScore(chnl)
			if s >= 0.0 && s < minScore {
				minScore = s
				cipher[i] = x
			}
		}
		if cipher[i] == 0 {
			fail = true
			break
		}
	}
	if !fail {
		fmt.Println(string(cipher))
		repeatingXor(raw, cipher)
		fmt.Println(string(raw))
	} else {
		log.Fatal("cipher not found")
	}
}
