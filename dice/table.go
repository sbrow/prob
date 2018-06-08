package dice

import (
	"encoding/csv"
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// DataDir is the folder in which roll tables will be generated.
var DataDir = filepath.Join(build.Default.GOPATH, "data/")

func init() {
	_, err := os.Open(DataDir)
	if os.IsNotExist(err) {
		os.MkdirAll(DataDir, 0777)
	}
}

// Table loads the roll table for a single roll of n dice.
//
// Table does not generate fresh data if called on a table that has
// already been generated,
//
// Table loads previously generated data when appropriate. However, it
// cannot tell if data in a file is incomplete. Care must be taken not to feed it
// incomplete tables, as this will cause undesirable results (but will not throw an error).
//
// TODO: Verify loaded data.
func Table(die Die, n int) [][]string {
	c := make([]<-chan state, len(die.Sides))
	var d int

	// If the table already exists, stop executing.
	f, err := os.Open(filepath.Join(DataDir, fmt.Sprintf("%v%v.csv", n, die.Name)))
	if !os.IsNotExist(err) {
		r := csv.NewReader(f)
		out, err := r.ReadAll()
		if err != nil {
			panic(err)
		}
		return out
	}

	// If the previous table does not exist, generate all recursions from scratch.
	_, err = os.Open(filepath.Join(DataDir, fmt.Sprintf("%v%v/", n-1, die.Name)))
	if os.IsNotExist(err) {
		d = 1
		for i, side := range die.Sides {
			c[i] = statesToChan(state{Current: side, Next: die})
		}
	} else {
		d = n - 1
		for i := range die.Sides {
			name := filepath.Join(DataDir, fmt.Sprintf("%v%v", n-1, die.Name),
				fmt.Sprintf("%v.csv", i+1))
			file, err := os.Open(name)
			if err != nil {
				panic(err)
			}
			r, err := csv.NewReader(file).ReadAll()
			if err != nil {
				panic(err)
			}
			c[i] = readState(r[1:], die)
		}
	}

	for d < n {
		for i := range die.Sides {
			c[i] = rollDie(c[i], d, i, die.Name)
		}
		d++
	}

	file, err := os.Create(filepath.Join(DataDir, fmt.Sprintf("%d%s.csv", n, die.Name)))
	if err != nil {
		panic(err)
	}
	out := [][]string{}
	w := csv.NewWriter(file)
	if err := w.Write([]string{"State", "Sum"}); err != nil {
		log.Println("error:", err)
	}
	for i := range c {
		tmp := []string{}
		for elem := range c[i] {
			if err := w.Write(elem.CSV()); err != nil {
				log.Println(err)
				continue
			}
			tmp = append(tmp, elem.CSV()...)
		}
		out = append(out, tmp)
		w.Flush()
	}
	return append([][]string{{"State", "Sum"}}, out...)
}

// rollDie is a pipeline function that takes a node on a roll tree and generates
// each of that node's children.
// Nodes are passed in on one channel and passed out on another.
func rollDie(in <-chan state, n int, j int, name string) <-chan state {
	out := make(chan state)
	folderName := fmt.Sprintf("%v%v", n+1, name)
	folder := filepath.Join(DataDir, folderName)

	go func() {
		// Set up our folders, files, and Writers.
		switch err := os.Mkdir(folder, 0700); {
		case os.IsExist(err):
			fallthrough
		case err == nil:
		default:
			panic(err)
		}
		file, err := os.Create(filepath.Join(folder, fmt.Sprintf("%v.csv", j+1)))
		if err != nil {
			panic(err)
		}
		w := csv.NewWriter(file)
		if err := w.Write([]string{"State", "Sum"}); err != nil {
			log.Println("error:", err)
		}

		//Parse and write input
		for r := range in {
			for i := 0; i < len(r.Next.Sides); i++ {
				rNew := state{
					Current: r.Current + Delim + r.Next.Sides[i],
					Next:    r.Next,
				}
				out <- rNew
				if err := w.Write(rNew.CSV()); err != nil {
					log.Println(err)
				}
			}
			w.Flush()
		}
		close(out)
	}()
	return out
}

// DeleteTables deletes generated data. Called by -force flag.
// if print is true, the names of files we're deleting are output to the console
func DeleteTables(print bool) {
	if print {
		fmt.Println("Removing old data...")
	}
	data, err := os.Open(DataDir)
	if os.IsNotExist(err) {
		os.MkdirAll(DataDir, 0777)
		DeleteTables(print)
		return
	} else if err != nil {
		panic(err)
	}
	contents, err := data.Readdirnames(-1)
	if err != nil {
		panic(err)
	}
	for _, file := range contents {
		if print {
			fmt.Println("Deleting", filepath.Join(DataDir, file))
		}
		if err := os.RemoveAll(filepath.Join(DataDir, file)); err != nil {
			log.Println(err)
		}
	}
}

// Delim is the delimiter that separates dice in a roll.
const Delim = "+"

// state is the current state of a number of dice that have been rolled.
// Current is the state of dice that have already been rolled.
// Next is the possibility space of the next die to be rolled.
type state struct {
	Current string
	Next    Die
}

// Sum calculates the sum of *state.current.
func (s *state) Sum() int {
	sum := 0
	vals := strings.Split(s.Current, Delim)
	for _, char := range vals {
		dieVal, err := strconv.Atoi(char)
		if err != nil {
			continue
		}
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
func readState(strs [][]string, die Die) <-chan state {
	out := make(chan state)
	go func() {
		for _, s := range strs {
			out <- state{Current: s[0], Next: die}
		}
		close(out)
	}()
	return out
}
