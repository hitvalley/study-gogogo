package main;

import "fmt";

func printBoolean(data bool) {
  fmt.Printf("\nprintBoolean(data bool)\n")
  fmt.Printf("args: %v, %T\n", data, data)
  fmt.Printf("%%t(%v): %t\n", data, data);
}

func printInteger(data int) {
  fmt.Printf("\nprintInteger(data int)\n")
  fmt.Printf("args: %v, %T\n", data, data)

  fmt.Printf("%%b(%v): %b\n", data, data)
  fmt.Printf("%%c(%v): %c\n", data, data)
  fmt.Printf("%%q(%v): %q\n", data, data)
  fmt.Printf("%%x(%v): %x\n", data, data)
  fmt.Printf("%%X(%v): %X\n", data, data)
  fmt.Printf("%%U(%v): %U\n", data, data)
}

func printFloat(data float32) {
  fmt.Printf("\nprintFloat(data int)\n")
  fmt.Printf("args: %v, %T\n", data, data)

  fmt.Printf("%%b(%v): %b\n", data, data)
}

func main() {
  printBoolean(true)
  printBoolean(false)
  printInteger(43)
  printFloat(43.33342)
}

