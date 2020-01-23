package main

import "fmt"
import "log"
import "encoding/hex"

const (
  input = "128379ad12"
  expected = "EoN5rRI="
  inputWiki = "Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure."
  expectedWiki = "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlzIHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2YgdGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZGVsaWdodCBpbiB0aGUgY29udGludWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdlLCBleGNlZWRzIHRoZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4="

  encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
  )

func base64(c byte) byte {
  return encodeStd[c]
}

func main() {
  result := encode([]byte(inputWiki))
  fmt.Println(string(result))
  if string(result) != expectedWiki {
    log.Fatal("Not equal!")
  } else {
    fmt.Println("Yeah!")
  }
}

func main2() {
  decoded, err := hex.DecodeString(input)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(decoded)

  result := encode(decoded)
  fmt.Println(string(result))
}

func min(i, j int) int {
  if i < j {
    return i
  }
  return j
}

func encode(b[] byte) string {
  result := ""
  l := len(b)
  fmt.Println("length", l)
  for i := 0; i < l; i += 3 {
    var k uint64 = 1
    m := min((l - i), 3)
    for j := 0; j < m; j++ {
      k = k << 8
      k |= uint64(b[i + j])
    }

    if m == 1 {
      k = k << 4
    } else if m == 2 {
      k = k << 2
    }

    n := ""
    for ; k > 1; {
      n = string(base64(byte(k & 63))) + n
      k = k >> 6
    }

    result += n
    if m == 1 {
      result += "=="
    } else if m == 2{
      result += "="
    }
  }
  fmt.Println(result)
  return result
}
