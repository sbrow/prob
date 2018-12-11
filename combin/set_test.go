package combin

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want string
	}{
		{"{{a,b}}", []interface{}{[]string{"a", "b"}}, "a,b\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet(tt.args...); got.String() != tt.want {
				t.Errorf("NewSet() = \"%v\", want %v", got, tt.want)
			}
		})
	}
}
