package cards

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []PlayingCard{
		{"2", "c"},
		{"3", "c"},
		{"4", "c"},
		{"5", "c"},
		{"6", "c"},
		{"7", "c"},
		{"8", "c"},
		{"9", "c"},
		{"T", "c"},
		{"J", "c"},
		{"Q", "c"},
		{"K", "c"},
		{"A", "c"},
		{"2", "d"},
		{"3", "d"},
		{"4", "d"},
		{"5", "d"},
		{"6", "d"},
		{"7", "d"},
		{"8", "d"},
		{"9", "d"},
		{"T", "d"},
		{"J", "d"},
		{"Q", "d"},
		{"K", "d"},
		{"A", "d"},
		{"2", "h"},
		{"3", "h"},
		{"4", "h"},
		{"5", "h"},
		{"6", "h"},
		{"7", "h"},
		{"8", "h"},
		{"9", "h"},
		{"T", "h"},
		{"J", "h"},
		{"Q", "h"},
		{"K", "h"},
		{"A", "h"},
		{"2", "s"},
		{"3", "s"},
		{"4", "s"},
		{"5", "s"},
		{"6", "s"},
		{"7", "s"},
		{"8", "s"},
		{"9", "s"},
		{"T", "s"},
		{"J", "s"},
		{"Q", "s"},
		{"K", "s"},
		{"A", "s"},
	}
	for _, tt := range tests {
		gotDeck := New()
		t.Run(fmt.Sprintf("containsOne(%s)", tt.String()), func(t *testing.T) {
			if got, ok := gotDeck.cards[tt]; !ok {
				t.Errorf(`New() does not contain "%s"`, tt)
			} else if got != 1 {
				t.Errorf("New()[%v] = %v, want %v", tt, got, 1)
			}
		})
	}
	t.Run("contains52UniqueCards", func(t *testing.T) {
		want := 52
		gotDeck := New()
		if got := len(gotDeck.cards); got != want {
			t.Errorf("len(New()) = %v, want %v", got, want)
		}
	})
}

func TestPlayingCardDeck_Remove(t *testing.T) {
	//type args struct { card PlayingCard }
	tests := []struct {
		name    string
		p       PlayingCardDeck
		args    []PlayingCard
		want    PlayingCardDeck
		wantErr bool
	}{
		{"removesOneCard", *New(), []PlayingCard{{"A", "s"}}, *New(), false},
		{"removesMultipleCards", *New(), []PlayingCard{{"A", "s"}, {"Q", "d"}}, *New(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Remove(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlayingCardDeck.Draw() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayingCardDeck.Remove(%v) = %v, want %v", tt.args, got, tt.want)
			}
			for _, card := range tt.args {
				if _, ok := got.cards[card]; ok {
					t.Errorf("PlayingCardDeck.Remove(%v) = failed to remove %s", tt.args, card)
				}
			}
		})
	}
}

func TestPlayingCardDeck_Size(t *testing.T) {
	deck2, _ := New().Draw(PlayingCard{"A", "s"})
	type fields struct {
		cards map[PlayingCard]int
		odds  *DeckOdds
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"", fields{cards: New().cards}, 52},
		{"", fields{cards: deck2.cards}, 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PlayingCardDeck{
				cards: tt.fields.cards,
				odds:  tt.fields.odds,
			}
			if got := p.Size(); got != tt.want {
				t.Errorf("PlayingCardDeck.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
