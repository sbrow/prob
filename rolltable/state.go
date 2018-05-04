package rolltable

import (
	"fmt"
	"github.com/sbrow/dice"
	"strconv"
	"strings"
)

// Delim is the delimiter that separates dice in a roll.
const Delim = "+"

// state is the current state of a number of dice that have been rolled.
// Current is the state of dice that have already been rolled.
// Next is the possibility space of the next die to be rolled.
type state struct {
	Current string
	Next    dice.Die
}

// Sum calculates the sum of *state.current.
func (s *state) Sum() int {
	sum := 0
	vals := strings.Split(s.Current, Delim)
	for _, char := range vals {
		dieVal, _ := strconv.Atoi(char)
		sum += dieVal
	}
	return sum
}

// String represents the state as a string.
func (s *state) String() string {
	return fmt.Sprintf("%s,%d", s.Current, s.Sum())
}

// CSV Formats state for writing to a csv file.
func (s *state) CSV() []string {
	return []string{s.Current, strconv.Itoa(s.Sum())}
}

// statesToChan makes a channel from given States
//
// Deprecated:
func statesToChan(states ...state) <-chan state {
	out := make(chan state)
	go func() {
		for _, s := range states {
			out <- s
		}
		close(out)
	}()
	return out
}

// readState loads States from a csv file into a channel.
func readState(strs [][]string, die dice.Die) <-chan state {
	out := make(chan state)
	go func() {
		for _, s := range strs {
			out <- state{Current: s[0], Next: die}
		}
		close(out)
	}()
	return out
}
