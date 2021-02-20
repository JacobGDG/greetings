package greetings // the package name

import "fmt" //simple string library

func Hello(name string) string { // capital H for externaly callable
  message := fmt.Sprintf("Hello, %v. Welcome!", name) // := declare and assign variable
  return message
}
