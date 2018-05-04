// Package rolltable enumerates dice probability tables and saves them in csv format.
//
// See https://godoc.org/github.com/sbrow/rolltable/cmd/rolltable for the command.
package rolltable

import (
	"encoding/csv"
	"fmt"
	"github.com/sbrow/dice"
	"go/build"
	"os"
	"path/filepath"
)

// DataDir is the folder in which roll tables will be generated.
var DataDir string = build.Default.GOPATH + "/data/"

// Generate creates the roll table for a single roll of n dice, and saves it in a csv file.
//
// Generate does not generate fresh data if called on a table that has
// already been generated,
//
// Generate loads previously generated data when appropriate. However, it
// cannot tell if data in a file is incomplete. Care must be taken not to feed it
// incomplete tables, as this will cause undesirable results (but will not throw an error).
//
// TODO: Verify loaded data.
func Generate(die dice.Die, n int) [][]string {
	c := make([]<-chan state, len(die.Sides))
	d := 0

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
			c[i] = readState(r, die)
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
	w := csv.NewWriter(file)
	out := [][]string{}
	w.Write([]string{"state", "Sum"})
	for i := range c {
		tmp := []string{}
		for elem := range c[i] {
			w.Write(elem.CSV())
			tmp = append(tmp, elem.CSV()...)
		}
		out = append(out, tmp)
		w.Flush()
	}
	return out
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
		os.Mkdir(folder, 0700)
		file, err := os.Create(filepath.Join(folder, fmt.Sprintf("%v.csv", j+1)))
		if err != nil {
			panic(err)
		}
		w := csv.NewWriter(file)

		//Parse and write input
		for r := range in {
			for i := 0; i < len(r.Next.Sides); i++ {
				rNew := state{
					Current: r.Current + Delim + r.Next.Sides[i],
					Next:    r.Next,
				}
				out <- rNew
				w.Write(rNew.CSV())
			}
			w.Flush()
		}
		close(out)
	}()
	return out
}

// removeData deletes generated data. Called by -force flag.
// if print is true, the names of files we're deleting are output to the console
func DeleteAll(print bool) {
	if print {
		fmt.Println("Removing old data...")
	}
	data, err := os.Open(DataDir)
	if err != nil {
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
		os.RemoveAll(filepath.Join(DataDir, file))
	}
}
