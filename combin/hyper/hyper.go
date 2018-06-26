// Package hyper creates hypergeometric distributions. Hypergeometric distributions
// are useful for computing the probability of random draws with removal, such as
// drawing cards from a deck.
package hyper

import "github.com/sbrow/prob/combin"

// Hyper represents a hypergeometric distribution where k is an array of
// different 'successes' that a sample can have.
type Hyper struct {
	K map[string]int
}

// N returns the population size of the distribution.
func (h *Hyper) N() int {
	size := 0
	for _, k := range h.K {
		size += k
	}
	return size
}

// Sample computes the probability of drawing a certain sample distribution from
// the hyper.
// TODO(sbrow): change to use maps.
func (h Hyper) Sample(items map[string]int) Result {
	return Result{Dist: h.K, Sample: items, PMF: h.PMF(items)}
}

// PMF calculates the probability mass function of the distribution, i.e. the
// probability of drawing this exact distribution.
func (h Hyper) PMF(items map[string]int) float64 {
	n := 0
	k := 1
	for key, value := range h.K {
		n += items[key]
		k *= combin.NCR(false, value, items[key])
	}

	// TODO Clean up.
	return float64(k) /
		(float64(combin.Product(h.N()-n+1, h.N())) /
			float64(combin.Fact(n)))

}

// Result is the results returned by Hyper.Sample.
type Result struct {
	Dist   map[string]int
	Sample map[string]int
	PMF    float64
}
