/*
Package hyper creates hypergeometric distributions. Hypergeometric distributions
are useful for computing the probability of random draws with removal, such as
drawing cards from a deck.

*/
package hyper // import "github.com/sbrow/combin/hyper"

import (
	"fmt"
	"github.com/sbrow/combin"
)

/*
Hyper represents a hypergeometric distribution where k is an array of
different 'successes' that a sample can have.
*/
type Hyper struct {
	K map[string]int
}

// String returns the distribution as a string, printing h.K and h.N().
func (h *Hyper) String() string {
	K := "{"
	for k, v := range h.K {
		K += fmt.Sprintf(" %s:%d,", k, v)
	}
	K = K[:len(K)-1] + " }"
	return fmt.Sprintf("K = %s, N = %d", K, h.N())
}

// N returns the population size of the distribution.
func (h *Hyper) N() (N int) {
	for _, k := range h.K {
		N += k
	}
	return N
}

/*
Sample computes the probability of drawing a certain sample distribution from
the hyper.
TODO: change to use maps.
*/
func (h *Hyper) Sample(dist ...int) (results string) {
	results = fmt.Sprintln(h.String())
	results += fmt.Sprintln("X = ", dist, "\n")
	results += fmt.Sprintf("P(X = x) = %f", h.PMF(dist))
	return results
}

/*
PMF calculates the probability mass function of the distribution, i.e. the
probability of drawing this exact distribution.
*/
func (h *Hyper) PMF(dist []int) float64 {
	n, i := 0, 0
	k := 1
	for _, K := range h.K {
		n += dist[i]
		k *= combin.NCR(K, dist[i])
		i++
	}
	total := combin.Combin(h.N(), n)
	return float64(k) / total.Float64()
}
