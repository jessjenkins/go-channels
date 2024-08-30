package main

import (
	"github.com/jessjenkins/go-channels/common"
	"strconv"
)

func main() {
	n := 10

	state := common.NewState()

	numchan := make(chan int)
	go func() {
		defer close(numchan)
		nums := common.GenerateNums(n)
		for _, num := range nums {
			state.Set(1, "…")
			state.Set(1, strconv.Itoa(num))
			numchan <- num
			state.Finish(1)
		}
	}()

	lchan := make(chan string)
	go func() {
		defer close(lchan)
		for num := range numchan {
			state.Set(2, strconv.Itoa(num))
			l := common.ConvertNumToLetter(num)
			state.Set(2, l)
			lchan <- l
			state.Finish(2)
		}
	}()

	uclchan := make(chan string)
	go func() {
		defer close(uclchan)
		for l := range lchan {
			state.Set(3, l)
			ucl := common.UpperCase(l)
			state.Set(3, ucl)
			uclchan <- ucl
			state.Finish(3)
		}
	}()

	joined := ""
	for ucl := range uclchan {
		state.Set(4, joined)
		joined += ucl
		state.Set(4, joined)
		state.Finish(4)
	}
}
