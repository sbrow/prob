package dice

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
)

// DataDir is the default directory to store Tables in.
var DataDir = filepath.Join(os.Getenv("GOPATH"), "data")

// Delim specifies the delimiter to use for dice in a roll.
var Delim = "+"

// DeleteData deletes the contents of DataDir.
func DeleteData() error {
	dir, err := os.Open(DataDir)
	switch {
	case os.IsNotExist(err):
		makeDataDir()
		fallthrough
	case err != nil:
		return err
	}

	files, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	dir.Close()

	for _, file := range files {
		if err := os.RemoveAll(filepath.Join(DataDir, file)); err != nil {
			log.Println(err)
		}
	}
	return nil
}

// Table holds a roll table for a set of dice.
//
// TODO(sbrow): Add Table.Verify
//
// BUG(sbrow): Add Table.Verify
//
type Table struct {
	Dice         // The Dice in the table.
	Rolls []roll // The possibility space of the dice when each is rolled once.
}

// NewTable returns a dice table from the given dice.
//
// If the table already exists in DataDir, it is loaded.
//
// If the table does not already exist in DataDir, NewTable looks for parent tables until
// it finds one or hits rock bottom, then generates the remaining data.
//
// For example, if we call with an empty DataDir:
// 		OneD6 := NewTable(D6())
// The table will be generated from scratch. If we then call:
// 		ThreeD6 := NewTable(D6(), D6(), D6())
// NewTable will start by loading the saved version of OneD6 from disk,
// and then generate the other two Dice.
func NewTable(dice Dice) (*Table, error) {
	var t *Table
	var err error

	// Look for generated tables until we find one.
	var i int
main:
	for i = len(dice); i > 0; i-- {
		t, err = loadTable(dice[:i])
		switch {
		case err == nil:
			break main
		case os.IsNotExist(err):
			continue
		default:
			return nil, err
		}
	}
	// Once we've found a saved table, (or run out of dice),
	// generate all remaining data.
	if t, err = generateTable(t, dice[i:]); err != nil {
		return nil, err
	}
	t.Save()
	return t, nil
}

// Name returns the name of the table, which is determined by its Dice.
func (t *Table) Name() string {
	return t.Dice.Name()
}

// generateRoll returns all possible die rolls, given a list of previous
// die rolls and the next die to roll.
func generateRoll(prev bytes.Buffer, d Die) bytes.Buffer {
	var err error
	// var line []byte
	var out bytes.Buffer

	del := []byte(Delim)
	length := prev.Len()

	for err != io.EOF {
		// Read a line.
		line, err := prev.ReadBytes('\n')
		if err == io.EOF && length > 0 {
			break
		}
		for _, s := range d.Sides {
			// Only add the delimiter if there is previous data to append.
			if len(line) > 1 {
				out.Write(line[:len(line)-1])
				out.Write(del)
			}
			out.WriteString(s)
			out.WriteByte('\n')
		}
	}
	return out
}

// generateTable generates a Table from Dice.
//
// TODO(sbrow): Add concurrency.
func generateTable(t *Table, dice Dice) (*Table, error) {
	var rolls bytes.Buffer

	if t != nil {
		for _, roll := range t.Rolls {
			rolls.WriteString(roll.String())
			rolls.WriteByte('\n')
		}
	}

	for _, d := range dice {
		rolls = generateRoll(rolls, d)
	}

	data, err := csv.NewReader(&rolls).ReadAll()
	if err != nil {
		return nil, err
	}
	sort.Strings(data[0])

	t = &Table{}
	t.Rolls = make([]roll, len(data))
	for i, line := range data {
		t.Rolls[i] = roll(line)
	}
	t.Dice = append(t.Dice, dice...)
	return t, nil
}

func makeDataDir() {
	if err := os.MkdirAll(DataDir, 0700); err != nil {
		panic(err)
	}
}

// loadTable loads a table from generated data. If the table does not exist,
// os.IsNotExist is returned.
func loadTable(dice Dice) (*Table, error) {
	t := &Table{Dice: dice}

	// Start by making sure the data folder exists.
	makeDataDir()

	// Check to see if this table has been saved.
	f, err := os.Open(filepath.Join(DataDir, t.Name()+".csv"))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// If it has, load its data.
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	t.Rolls = make([]roll, len(data))

	for i, line := range data {
		t.Rolls[i] = roll(line)
	}
	return t, nil
}

// Save writes the table to disk in csv format.
func (t *Table) Save() error {
	f, err := os.Create(filepath.Join(DataDir, t.Name()+".csv"))
	if err != nil {
		return err
	}
	defer f.Close()
	for _, roll := range t.Rolls {
		f.WriteString(roll.String())
		f.Write([]byte{'\n'})
	}
	return nil
}
