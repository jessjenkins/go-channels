package common

import (
	"fmt"
	"maps"
	"slices"
	"sync"
)

type State struct {
	mu     sync.Mutex
	values map[int]string
}

func NewState() *State {
	state := State{values: make(map[int]string)}

	return &state
}

func (s *State) Set(step int, value string) {
	s.Update(step, value)
	s.Print()
}

func (s *State) Finish(step int) {
	s.Update(step, ".")
}

func (s *State) Update(step int, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[step] = value
}

func (s *State) Print() {
	s.mu.Lock()
	defer s.mu.Unlock()
	line := ""

	for _, step := range slices.Sorted(maps.Keys(s.values)) {
		line += fmt.Sprintf("%-4s", s.values[step])
	}
	fmt.Println(line)
}
