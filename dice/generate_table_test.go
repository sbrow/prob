package dice

/*
import (
	"sync"
	"testing"
)

func BenchmarkGenTable(b *testing.B) {
	dice := Dice{D6(), D6(), D6(), D6(), D6(), D6(), D6(), D6()}
	var wg sync.WaitGroup
	b.Run("8d6", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			t, err := generateTable(nil, dice)
			if err != nil {
				b.Error(err)
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				t.Save()
			}()
		}
		wg.Wait()
	})
	b.ResetTimer()
	b.Run("4d6", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			t, err := generateTable(nil, dice[:4])
			if err != nil {
				b.Error(err)
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				t.Save()
			}()
		}
		wg.Wait()
	})
}
*/
