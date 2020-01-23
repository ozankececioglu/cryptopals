package main

import "fmt"
import "log"
import "encoding/hex"

const (
  input = "128379ad12"
  expected = "EoN5rRKT"
  encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
  )

func base64(c byte) byte {
  return encodeStd[c]
}

func main() {
  decoded, err := hex.DecodeString(input)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(decoded)

  encode(decoded)
}

func min(i, j int) int {
  if i < j {
    return i
  }
  return j
}

func encode(b[] byte) {
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
}