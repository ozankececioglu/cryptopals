package main

import (
	"encoding/hex"
	"log"
)

const (
	input    = "1c0111001f010100061a024b53535009181c"
	xor      = "686974207468652062756c6c277320657965"
	expected = "746865206b696420646f6e277420706c6179"
)

func str_xor(inp string, key string) string {
	if len(input) != len(key) {
		log.Fatal("input length is not equal to key length")
	}
	inpb, err := hex.DecodeString(inp)
	if err != nil {
		log.Fatal(err)
	}
	keyb, err := hex.DecodeString(key)
	if err != nil {
		log.Fatal(err)
	}
	resb := make([]byte, len(keyb))
	for i := 0; i < len(inpb); i++ {
		resb[i] = keyb[i] ^ inpb[i]
	}
	return hex.EncodeToString(resb)
}

func main() {
	res := str_xor(input, xor)
	if expected != res {
		log.Fatal("Not equal, expected: %v, result: %v", expected, res)
	} else {
		log.Fatal("Yeah")
	}
}
