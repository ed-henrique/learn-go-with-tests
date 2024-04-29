package main

import "fmt"

const (
  french = "French"
  spanish = "Spanish"
  englishHelloPrefix = "Hello, "
  spanishHelloPrefix = "Hola, "
  frenchHelloPrefix = "Bonjour, "
)

func Hello(name, language string) string {
  if name == "" {
    name = "World"
  }

  switch language {
  case spanish:
    return spanishHelloPrefix + name
  case french:
    return frenchHelloPrefix + name
  default:
    return englishHelloPrefix + name
  }
}

func main() {
  fmt.Println(Hello("World", ""))
}
