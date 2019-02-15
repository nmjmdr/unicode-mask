package main

import "masker"
import "fmt"

func main() {
  str := "こんにちは世界"
  masked := masker.Mask(str, 3, '*')
  fmt.Println(str, masked)
}
