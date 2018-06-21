package dice

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDice(t *testing.T) {
	dicetests := []struct {
		die   Die
		name  string
		sides []string
	}{
		{*NewDie("D3", "1", "A", "3"), "D3", []string{"1", "A", "3"}},
		{*NewDie("D4", "1", "2", "3", "4"), "D4", []string{"1", "2", "3", "4"}},
		{D4(), "d4", []string{"1", "2", "3", "4"}},
		{D6(), "d6", []string{"1", "2", "3", "4", "5", "6"}},
		{D8(), "d8", []string{"1", "2", "3", "4", "5", "6", "7", "8"}},
		{D10(), "d10", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}},
		{D12(), "d12", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}},
		{D20(), "d20", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12",
			"13", "14", "15", "16", "17", "18", "19", "20"}},
	}
	for i, tt := range dicetests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			want := tt.sides
			got := tt.die.Sides
			if !reflect.DeepEqual(got, want) {
				t.Errorf("wanted: %s\ngot: %s\n", want, got)
			}
		})
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		dice Dice
		out  string
	}{
		{Dice{D6(), D6(), D4()}, "2d6+1d4"},
	}
	for _, tt := range tests {
		t.Run(tt.out, func(t *testing.T) {
			got := Table{Dice: tt.dice}.Name()
			if got != tt.out {
				t.Errorf("wanted: %s\ngot: %s\n", tt.out, got)
			}
		})
	}
}

var tabletests = []struct {
	in   *Table
	dice Dice
	out  Table
}{
	{nil, Dice{D6()}, Table{
		Data: [][]string{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}, {"6"}},
		Dice: Dice{D6()}}},
	{nil, Dice{D6(), D6()}, Table{
		Data: [][]string{
			{"1+1"}, {"1+2"}, {"1+3"}, {"1+4"}, {"1+5"}, {"1+6"},
			{"2+1"}, {"2+2"}, {"2+3"}, {"2+4"}, {"2+5"}, {"2+6"},
			{"3+1"}, {"3+2"}, {"3+3"}, {"3+4"}, {"3+5"}, {"3+6"},
			{"4+1"}, {"4+2"}, {"4+3"}, {"4+4"}, {"4+5"}, {"4+6"},
			{"5+1"}, {"5+2"}, {"5+3"}, {"5+4"}, {"5+5"}, {"5+6"},
			{"6+1"}, {"6+2"}, {"6+3"}, {"6+4"}, {"6+5"}, {"6+6"},
		}, Dice: Dice{D6(), D6()}}},
	{&Table{Dice: Dice{D6()}, Data: [][]string{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}, {"6"}}},
		Dice{D6(), D6()}, Table{
			Data: [][]string{
				{"1+1"}, {"1+2"}, {"1+3"}, {"1+4"}, {"1+5"}, {"1+6"},
				{"2+1"}, {"2+2"}, {"2+3"}, {"2+4"}, {"2+5"}, {"2+6"},
				{"3+1"}, {"3+2"}, {"3+3"}, {"3+4"}, {"3+5"}, {"3+6"},
				{"4+1"}, {"4+2"}, {"4+3"}, {"4+4"}, {"4+5"}, {"4+6"},
				{"5+1"}, {"5+2"}, {"5+3"}, {"5+4"}, {"5+5"}, {"5+6"},
				{"6+1"}, {"6+2"}, {"6+3"}, {"6+4"}, {"6+5"}, {"6+6"},
			}, Dice: Dice{D6(), D6()}}},
}

func TestNewTable(t *testing.T) {
	DeleteData()
	for i, tt := range tabletests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := NewTable(tt.dice)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(*got, tt.out) {
				t.Fail()
				// t.Errorf("\nwanted: %s\ngot: %s\n", tt.out, *got)
			}
		})
	}
}

func TestGenTable(t *testing.T) {
	for i, tt := range tabletests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got, err := generateTable(nil, tt.dice)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(*got, tt.out) {
				t.Errorf("wanted: %s\ngot: %s\n", tt.out, *got)
			}
		})
	}
}
func TestLoadTable(t *testing.T) {
	for i, tt := range tabletests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			tbl, err := generateTable(nil, tt.dice)
			if err != nil {
				t.Fatal(err)
			}
			tbl.save()
			got, err := loadTable(tt.dice)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(*got, tt.out) {
				t.Errorf("wanted: %s\ngot: %s\n", tt.out, *got)
			}

		})
	}
}

func TestDeleteTables(t *testing.T) {
	var err error
	for i := 0; i < 3; i++ {
		_, err = os.Create(filepath.Join(DataDir, fmt.Sprintf("test_%d", i)))
		if err != nil {
			t.Fatal(err)
		}
		DeleteData()
		var dir *os.File
		if dir, err = os.Open(DataDir); err != nil {
			t.Fatal(err)
		}
		var files []string
		if files, err = dir.Readdirnames(-1); err != nil {
			t.Fatal(err)
		}
		if len(files) > 0 {
			t.Error("DeleteTables did not delete data.")
		}

	}
}

var gentests = []struct {
	roll [][]byte
	die  Die
	out  [][]byte
}{
	{bytes.Split([]byte("1\n"), []byte(",")), D6(),
		bytes.Split([]byte("1+1\n1+2\n1+3\n1+4\n1+5\n1+6\n"), []byte{'\n'})},
}

func TestGenRoll(t *testing.T) {
	for i, tt := range gentests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			var buff bytes.Buffer
			for _, line := range tt.roll {
				buff.Write(line)
			}
			out := generateRoll(buff, tt.die)
			got := bytes.Split(out.Bytes(), []byte("\n"))
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("wanted: %s\ngot: %s\n", tt.out, got)
			}
		})
	}
}

func TestDiceRoll(t *testing.T) {
	want := []string{"6", "4"}
	dice := New(D6(), D6())
	got := dice.Roll()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %s\ngot: %s\n", want, got)
	}
}

func TestDiceSwap(t *testing.T) {
	want := New(D8(), D6(), D4())
	got := New(D4(), D6(), D8())
	if reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %s\ngot: %s\n", want, got)
	}
	got.Swap(0, 2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %s\ngot: %s\n", want, got)
	}
}

func TestDiceSort(t *testing.T) {
	want := New(D8(), D6(), D4())
	got := New(D4(), D6(), D8())
	if reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %s\ngot: %s\n", want, got)
	}
	got.Sort()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %s\ngot: %s\n", want, got)
	}
}
