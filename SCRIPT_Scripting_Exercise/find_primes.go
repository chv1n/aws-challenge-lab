package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// isPrime ตรวจสอบว่าตัวเลขเป็นจำนวนเฉพาะหรือไม่
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var primes []string
	lower, upper := 1, 250
	filename := "results.txt"

	// หาจำนวนเฉพาะ
	for i := lower; i <= upper; i++ {
		if isPrime(i) {
			primes = append(primes, strconv.Itoa(i))
		}
	}

	// รวมผลลัพธ์
	content := strings.Join(primes, "\n")

	// เขียนลงไฟล์
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Printf("Prime numbers between %d and %d:\n", lower, upper)
	fmt.Println(strings.Join(primes, ", "))

	// หา Absolute Path
	absPath, _ := filepath.Abs(filename)
	scriptPath, _ := filepath.Abs("find_primes.go")

	fmt.Printf("\n[Success] Results saved to: %s\n", absPath)
	fmt.Printf("[Note] Script location: %s\n", scriptPath)
}