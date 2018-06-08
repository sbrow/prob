package dice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRoll(t *testing.T) {
	want := "2+4"
	got := Roll(D4(), D6())
	if got != want {
		t.Errorf("wanted: %s\ngot: %s\n", want, got)
	}
}

var dicetests = []struct {
	Die   Die
	Sides []string
}{
	{*New("D3", "1", "A", "3"), []string{"1", "A", "3"}},
	{*New("D4", "4"), []string{"1", "2", "3", "4"}},
	{D4(), []string{"1", "2", "3", "4"}},
	{D6(), []string{"1", "2", "3", "4", "5", "6"}},
	{D8(), []string{"1", "2", "3", "4", "5", "6", "7", "8"}},
	{D10(), []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}},
	{D12(), []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}},
	{D20(), []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12",
		"13", "14", "15", "16", "17", "18", "19", "20"}},
}

func TestDice(t *testing.T) {
	for i, tt := range dicetests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			want := tt.Sides
			got := tt.Die.Sides
			if !reflect.DeepEqual(got, want) {
				t.Errorf("wanted: %s\ngot: %s\n", want, got)
			}
		})
	}
}

func TestRollTable(t *testing.T) {
	DeleteTables(false)
	got := Table(D4(), 1)
	fmt.Println(got)
	got = Table(D4(), 2)
	fmt.Println(got)
	got = Table(D4(), 3)
	fmt.Println(got)
	got = Table(D4(), 3)
	fmt.Println(got)
}
