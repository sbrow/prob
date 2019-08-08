package cards

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeckOdds_Draw(t *testing.T) {
	type args struct {
		hands []Hander
	}
	tests := []struct {
		args     args
		wantOdds DeckOdds
		wantErr  bool
	}{
		{args{[]Hander{Hand{{"A", "h"}: 1}}}, DeckOdds{nil, 1, 52}, false},
		{args{[]Hander{Hand{{"A", "h"}: 1, {"Q", "d"}: 1}}}, DeckOdds{nil, 1, 1326}, false},
		{args{[]Hander{PlayingCard{"A", "h"}, PlayingCard{"Q", "d"}}}, DeckOdds{nil, 1, 2652}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.args.hands), func(t *testing.T) {
			d := newDeckOdds(New())
			gotOdds, err := d.Draw(tt.args.hands...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeckOdds.Draw() error = %v, wantErr %v", err, tt.wantErr)
				t.Errorf("%v", gotOdds.deck.cards)
				return
			}
			tt.wantOdds.deck = gotOdds.deck
			if !reflect.DeepEqual(gotOdds, tt.wantOdds) {
				t.Errorf("DeckOdds.Draw() = %v, want %v", gotOdds, tt.wantOdds)
			}
		})
	}
}
