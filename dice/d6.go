package dice

import (
	"fmt"
	"math/rand"
)

// IntSide is a Side that is represented purely as a number.
type IntSide int

func (i IntSide) String() string {
	return fmt.Sprint(int(i))
}

func (i IntSide) Value() int {
	return int(i)
}

func (i IntSide) Side() Side {
	return Side(i)
}

type D6 struct {
	Upside *int
}

func (d D6) Name() string {
	return "d6"
}

func (d *D6) Roll() Side {
	up := rand.Intn(len(d.Sides()))
	d.SetUp(&up)
	return d.Sides()[*d.Up()]
}

func (d *D6) SetUp(side *int) {
	d.Upside = side
}

func (d D6) Sides() []Side {
	n := 6
	sides := make([]Side, n)
	for i := 0; i < n; i++ {
		sides[i] = IntSide(i + 1)
	}
	return sides
}

func (d *D6) Up() *int {
	return d.Upside
}

type D8 struct {
	D6
}

func (d D8) Sides() []Side {
	return append(d.D6.Sides(), IntSide(7), IntSide(8))
}

type D10 struct {
	D8
}

func (d D10) Sides() []Side {
	return append(d.D6.Sides(), IntSide(9), IntSide(10))
}
