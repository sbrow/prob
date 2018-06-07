/*
Package hyper creates hypergeometric distributions. Hypergeometric distributions
are useful for computing the probability of random draws with removal, such as
drawing cards from a deck.

*/
package hyper

import (
	"fmt"

	"github.com/sbrow/prob/combin"
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
	results = fmt.Sprintln(h.String(), "\n", "X = ", dist)
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
		k *= combin.NCR(false, K, dist[i])
		i++
	}

	// TODO Clean up.
	return float64(k) /
		(float64(combin.Product(h.N()-n+1, h.N())) /
			float64(combin.Fact(n)))

}
