package dice

import (
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

/*
func TestNewDatabase(t *testing.T) {
	tests := []struct {
		name    string
		DB      *Database
		wantErr bool
	}{
		{"a", DB, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.DB
			got, err := NewDatabase()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDatabase() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("wanted: %v\ngot: %v\n", want, got)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		d Die
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"d6", args{d: D6()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DB.RegisterDieType(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/

/*
func TestNewTable(t *testing.T) {
	type args struct {
		d Die
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"d6", args{d: new(D6), n: 1}, false},
		// {"d6", args{d: D6(), n: 2}, false},
		// {"d6", args{d: D6(), n: 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DB.Mut.Lock()
			defer DB.Mut.Unlock()
			if err := DB.NewTable(tt.args.d, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("NewTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/

func TestDatabase_Roll(t *testing.T) {
	type args struct {
		dice    *Dice
		fn      Reroll
		rerolls int
	}
	tests := []struct {
		name string
		args args
		want Table
	}{
		{"1d6", args{New(new(D6)), nil, 0}, Table{Dice: []Die{new(D6)}, Rows: []Row{
			{[]int{1}, 1},
			{[]int{2}, 1},
			{[]int{3}, 1},
			{[]int{4}, 1},
			{[]int{5}, 1},
			{[]int{6}, 1},
		}}},
	}
	DB.Mut.Lock()
	defer DB.Mut.Unlock()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DB.Roll(tt.args.dice, tt.args.fn, tt.args.rerolls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.Roll() = %v, want %v", got, tt.want)
			}
		})
	}
}
