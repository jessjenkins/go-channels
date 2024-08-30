package main

import (
	"github.com/jessjenkins/go-channels/common"
	"strconv"
	"sync"
)

const (
	n         = 10
	buffer    = 3
	instances = 3
)

func main() {
	state := common.NewState()

	numchan := make(chan int, buffer)
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

	lchan := make(chan string, buffer)
	go func() {
		defer close(lchan)
		var wg sync.WaitGroup
		for i := 0; i < instances; i++ {
			wg.Add(1)
			go func(instance int) {
				step := 2 + instance
				for num := range numchan {
					state.Set(step, strconv.Itoa(num))
					l := common.ConvertNumToLetter(num)
					state.Set(step, l)
					lchan <- l
					state.Finish(step)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
	}()

	uclchan := make(chan string, buffer)
	go func() {
		defer close(uclchan)
		var wg sync.WaitGroup
		for i := 0; i < instances; i++ {
			wg.Add(1)
			go func(instance int) {
				step := 2 + instances + instance
				for l := range lchan {
					state.Set(step, l)
					ucl := common.UpperCase(l)
					state.Set(step, ucl)
					uclchan <- ucl
					state.Finish(step)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
	}()

	joined := ""
	joinstep := 2 + 2*instances
	for ucl := range uclchan {
		state.Set(joinstep, joined)
		joined += ucl
		state.Set(joinstep, joined)
		state.Finish(joinstep)
	}
}
