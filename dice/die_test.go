package dice

import (
	"reflect"
	"testing"
)

func TestDie_Roll(t *testing.T) {
	type args struct {
		typ Die
		up  int
	}
	tests := []struct {
		name string
		d    Die
		want args
	}{
		{"1d6", new(D6), args{new(D6), 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Roll()
			want := tt.want.typ
			up := &tt.want.up
			want.SetUp(up)
			got := tt.d
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Die.Roll() = %v, want %v", got, tt.want)
			}
		})
	}
}
