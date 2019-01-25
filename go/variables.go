package main

import (
  "fmt"
  "math"
  "math/cmplx"
)

var c, python, java bool
// Variables declared without an explicit initial value are given their zero value:
//    0 for numeric types,
//    false for the boolean type, and
//    "" (the empty string) for strings.

// Variables with initializers
var i, j int = 1, 2

func shortVariableDeclaration() {
  k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java) // 1 2 3 true false no!
}

func basicTypes() {
  var (
  	ToBe   bool       = false
    i      int        = 100 // The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int.
  	MaxInt uint64     = 1<<64 - 1
    bt     byte       = 255 // alias for uint8
    r      rune       = 32 // alias for int32, represents a Unicode code point
    f32    float32    = 0.0001
    f64    float64    = 0.00000000001
  	z64    complex64 = -5i
  	z128   complex128 = cmplx.Sqrt(-5 + 12i)
  )
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", bt, bt)
	fmt.Printf("Type: %T Value: %v\n", r, r)
	fmt.Printf("Type: %T Value: %v\n", f32, f32)
	fmt.Printf("Type: %T Value: %v\n", f64, f64)
	fmt.Printf("Type: %T Value: %v\n", z64, z64)
	fmt.Printf("Type: %T Value: %v\n", z128, z128)
}

func typeConversions() {
  var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // If removing float64 here, error: cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt.
	var z uint = uint(f)
	fmt.Println(x, y, z) // 3 4 5
}

func typeInference() {
  i := 42           // int
  f := 3.142        // float64
  g := 0.867 + 0.5i // complex128
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", g, g)
}

func constants() {
  const World = "世界" // Constants cannot be declared using the := syntax.
	fmt.Println("Hello", World)
  const Pi = 3.14
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func numericConstants() {
  const (
  	// Create a huge number by shifting a 1 bit left 100 places.
  	// In other words, the binary number that is 1 followed by 100 zeroes.
  	Big = 1 << 100
  	// Shift it right again 99 places, so we end up with 1<<1, or 2.
  	Small = Big >> 99
  )
  
  // Numeric constants are high-precision values.
  // An untyped constant takes the type needed by its context. 
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func main() {
	var b int
	fmt.Println(b, c, python, java) // 0 false false false

  var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java) // 0 2 true false no!

  shortVariableDeclaration()

  basicTypes()

  typeConversions()

  typeInference()

  constants()

  numericConstants()
}
