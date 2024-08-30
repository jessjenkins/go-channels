package common

import (
	"math/rand/v2"
	"strings"
	"time"
)

func GenerateNums(n int) []int {
	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return nums
}

func ConvertNumToLetter(n int) string {
	FakeDelay()
	ln := (n-1)%26 + 1
	char := string(rune(96 + ln))
	return char
}

func UpperCase(l string) string {
	FakeDelay()
	return strings.ToUpper(l)
}

func FakeDelay() {
	time.Sleep(time.Millisecond * time.Duration(100+rand.IntN(500)))
}
