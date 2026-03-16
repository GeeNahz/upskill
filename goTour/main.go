package main

import (
	"fmt"
	"runtime"
)

func Sqrt(x float64) float64 {
	z := x / 2
	u := 0.0
	for y := 1; z != u; y++ {
		u = z                    // Update u to value of z
		z -= (z*z - x) / (2 * z) // Update value of z
		// fmt.Printf("Value for z is: %f ", z)
		// fmt.Printf("for itr %d\n", y)
	}
	return z
}
func GoRuntime() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func main() {
	fmt.Println(Sqrt(25))
	GoRuntime()
}
