package main

import (
  "fmt"
  "strings"
  "math"
)

func testPointer() {
  i, j := 42, 2701

  p := &i         // point to i
  fmt.Println(*p) // read i through the pointer
  *p = 21         // set i through the pointer. Known as "dereferencing" or "indirecting".
  fmt.Println(i)  // see the new value of i

  p = &j         // point to j
  *p = *p / 37   // divide j through the pointer
  fmt.Println(j) // see the new value of j

  fmt.Println("Unlike C, Go has no pointer arithmetic.")
}

func testStruct() {
  type Vertex struct {
    X, y int;
  }
  fmt.Println(Vertex{0, 0})
  fmt.Println(Vertex{4, 5})

  v := Vertex{1, 2}
  fmt.Println(v.X) // Uppercase works.
  fmt.Println(v.y) // Lowercase also works here.

  p := &v
  fmt.Printf("%T\n", p)
  // The language permits us to write just p.X instead of (*p).X. There is no p->X.
  fmt.Printf("%v\n", p.X)

  fmt.Println("Struct Literals:")
  var (
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 5} // {5 0}
    v3 = Vertex{} // {0 0}
    pl = &v1
    pl2 = &Vertex{1, 2} // has type *Vertex
  )
  fmt.Println(v1, v2, v3, pl, pl2)
  v1.X = 100
}

func testArrays() {
  // An array's length is part of its type, so arrays cannot be resized.
  var a [2]string
  a[0] = "Hello"
  a[1] = "world"
  fmt.Println(a, a[1])

  b := [10]int{1,2,3}
  fmt.Println(b)

  c := [3]int{} // Must specify an expression.
  fmt.Println(c)

  //----- Slice
  fmt.Println("A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.")
  var s []int = b[0:3] // The type []T is a slice with elements of type T. This selects a half-open range: [0, 3).
  fmt.Println(s) // [1 2 3]

  //-----
  fmt.Println("A slice does not store any data.")
  // Changing the elements of a slice modifies the corresponding elements of its underlying array.
  names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	n1 := names[0:2]
	n2 := names[1:3]
	fmt.Println(n1, n2)
	n1[0] = "XXX"
	fmt.Println(n1, n2)
	fmt.Println(names)

  //----- A slice literal is like an array literal without the length.
  // This is an array literal:
  // [3]bool{true, true, false}
  // And this creates the same array as above, then builds a slice that references it:
  // []bool{true, true, false}
	sl := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sl)

  //----- Slice defaults
  aa := [5]int{1,2,3,4,5}
  // these slice expressions are equivalent:
  fmt.Println(aa[0:5])
  fmt.Println(aa[:5])
  fmt.Println(aa[0:])
  fmt.Println(aa[:])
  fmt.Printf("%T\n", aa[:]) // []int

  // Nil slices
  var sn []int
	fmt.Println(sn, len(sn), cap(sn)) // [] 0 0
	if sn == nil {
		fmt.Println("nil!") // nil!
	}
}

//----- Slice length and capacity
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func testSliceCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]
	
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []
	printSlice(s[:]) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=6 [5 7]

  // panic: runtime error: slice bounds out of range
	// s = s[:7]
}

func testMake() {
  // This is how you create dynamically-sized arrays.

  a := make([]int, 5)  // len(a)=5
  printSlice(a)
  // To specify a capacity, pass a third argument to make:
  b := make([]int, 0, 5) // len(b)=0, cap(b)=5
  printSlice(b)
  b = b[:cap(b)] // len(b)=5, cap(b)=5
  b = b[1:]      // len(b)=4, cap(b)=4
}

func slicesOfSlices() {
// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func testAppend() {
  // If the backing array of s is too small to fit all the given values a bigger array will be allocated.
  var s []int
	printSlice(s) // len=0 cap=0 []

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s) // len=1 cap=1 [0]

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s) // len=2 cap=2 [0 1]

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s) // len=5 cap=6 [0 1 2 3 4]
	
	s = s[:2]
	printSlice(s) // len=2 cap=6 [0 1]
}

func testRange() {
  var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

  pow = make([]int, 10)
	for i := range pow { // Can omit v
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow { // Can omit i
		fmt.Printf("%d\n", value)
	}
}

func testMap() {
  type Vertex struct {
  	Lat, Long float64
  }
  var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

  // The zero value of a map is nil. A nil map has no keys, nor can keys be added.
  // var mzz map[string]Vertex
  // mzz["Try"] = Vertex{1,1} // panic: assignment to entry in nil map

  mz := map[string]Vertex{}
  fmt.Println(mz) // map[]
  mz["Try"] = Vertex{1,1}
  fmt.Println(mz) // map[Try:{1 1}]

  //----- Map literals
  var ml = map[string]Vertex{
	  "Bell Labs": Vertex{
	  	40.68433, -74.39967,
	  },
	  "Google": Vertex{
	  	37.42202, -122.08408,
	  },
  }
	fmt.Println(ml)
  // Or
  var ml2 = map[string]Vertex{
  	"Bell Labs": {40.68433, -74.39967},
  	"Google":    {37.42202, -122.08408},
  }
	fmt.Println(ml2)
}

func mutateMaps() {
  m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// Functions are values too.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
  testPointer()
  testStruct()
  testArrays()
  testSliceCap()
  testMake()
  slicesOfSlices()
  testAppend()
  testRange()
  testMap()
  mutateMaps()

  // Function values
  hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13
	fmt.Println(compute(hypot)) // 5
	fmt.Println(compute(math.Pow)) // 81

  // Function closure
  pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	} // 45 -90
}
