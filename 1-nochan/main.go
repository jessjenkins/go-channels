package main

import (
	"github.com/jessjenkins/go-channels/common"
	"strconv"
)

func main() {
	n := 10
	nums := common.GenerateNums(n)

	state := common.NewState()

	joined := ""

	for _, num := range nums {
		state.Set(1, "…")
		state.Set(1, strconv.Itoa(num))
		state.Finish(1)

		state.Set(2, strconv.Itoa(num))
		l := common.ConvertNumToLetter(num)
		state.Set(2, l)
		state.Finish(2)

		state.Set(3, l)
		ucl := common.UpperCase(l)
		state.Set(3, ucl)
		state.Finish(3)

		state.Set(4, joined)
		joined += ucl
		state.Set(4, joined)
		state.Finish(4)
	}
}
