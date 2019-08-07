package cards

import "testing"

func TestPlayingCard_String(t *testing.T) {
	tests := []struct {
		name string
		p    PlayingCard
		want string
	}{
		{"works", PlayingCard{"A", "s"}, "As"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("PlayingCard.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
