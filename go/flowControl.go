package main

import (
  "fmt"
  "math"
  "runtime"
  "time"
)

func testFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
  // Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
	fmt.Println(sum)
}


func initAndPostOptional() {
  sum := 1
  // The init and post statements are optional.
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum) // 1024
}

func forIsWhile() {
  sum := 1
  // Drop the semicolons in initAndPostOptional(): C's "while" is spelled "for" in Go. 
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// Don't run this!!!
func infiniteLoop() {
  // If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
  for {
	}
}

// If
func sqrt(x float64) string {
  if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// If with a short statement
// Variables are only in scope until the end of the if.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in languages like C or C++ is provided automatically in Go.
func testSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
func whenIsSaturday() {
  fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// Switch without a condition is the same as switch true.
// This construct can be a clean way to write long if-then-else chains. 
func hello() {
  t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately.
// Deferred function calls are pushed onto a stack.
func testDefer() {
  fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
  testFor()
  initAndPostOptional()
  forIsWhile()

  fmt.Println("----------Test if")
  fmt.Println(
    sqrt(2),
    sqrt(-4),
  )

  fmt.Println("----------Test if with a short statement")
  fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

  testSwitch()
  whenIsSaturday()

  hello()

  testDefer()
}
