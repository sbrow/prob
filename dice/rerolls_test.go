package dice

/*
func TestAverage(t *testing.T) {
	tests := []struct {
		name string
		args RolledDice
		want []bool
	}{
		{"1d6=1", RolledDice{{D6(), 1}}, []bool{true}},
		{"1d6=2", RolledDice{{D6(), 2}}, []bool{true}},
		{"1d6=3", RolledDice{{D6(), 3}}, []bool{true}},
		{"1d6=4", RolledDice{{D6(), 4}}, []bool{false}},
		{"1d6=5", RolledDice{{D6(), 5}}, []bool{false}},
		{"1d6=6", RolledDice{{D6(), 6}}, []bool{false}},
		{"2d6=4+3", RolledDice{{D6(), 4}, {D6(), 3}}, []bool{false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			got := Average(tt.args)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("wanted: %v\ngot: %v\n", want, got)
			}
		})
	}
}
*/
