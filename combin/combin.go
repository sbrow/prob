package combin

import (
	"fmt"
)

type combin struct {
	numer []int
	denom int
}

func Combin(n, r int) *combin {
	return &combin{[]int{n, n - r + 1}, r}
}

func (c *combin) String() string {
	return fmt.Sprintf("(%d...%d) / %d!", c.numer[0], c.numer[1], c.denom)
}

func (c *combin) Float64() float64 {
	return float64(Product(c.numer[1], c.numer[0])) / float64(Fact(c.denom))
}

func (c *combin) Int() int {
	return int(c.Float64())
}

func (c *combin) Div(div combin) (quot float64) {
	return c.Float64() / div.Float64()
}

func (c *combin) Mult(mult ...combin) (prod int) {
	prod = c.Int()
	for _, factor := range mult {
		prod *= factor.Int()
	}
	return prod
}
