// TODO: https://blog.couchbase.com/using-graphql-with-golang-and-a-nosql-database/

package main

import (
  "fmt"
  "time"
  "math/rand"
  "math"
)

// Go's Declaration Syntax (compared to C): https://blog.golang.org/gos-declaration-syntax
func add(a, b int) int { // same as "x int, y int"
  return a + b
}

func cate(a string, b string) string {
  return a + b
}

// Multiple results
func swap(a, b string) (string, string) {
  return b, a
}

// Named return values
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return // Naked return.
}

func main() {
  fmt.Println("Hello,", "世界")
  fmt.Println("The time is", time.Now())

  // TODO: rand.Seed()
  fmt.Println("My favorite number is", rand.Intn(10))

  // In Go, a name is exported if it begins with a capital letter.
  fmt.Println(math.Pi)

  // Printf
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

  fmt.Println(add(55, 45))
  fmt.Println(cate("Hehe", "Haha"))

  a, b := swap("a", "b")
  fmt.Println(a, b)

  fmt.Println(split(17))
}
