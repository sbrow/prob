package dice

/*
// IntDie represents a die whose sides contain only integer values.
type IntDie []IntSide

func (i IntDie) Sides() []Side {
	sides := make([]Side, len(i))
	for j := range i {
		sides[j] = i[j]
	}
	return sides
}

// NewIntDie returns a new IntDie object with the given sides.
// if given only one side, it will create a die with n sides: [1, n]
func NewIntDie(name string, n ...int) *Die {
	d := new(Die)

	var sides []Side
	if len(n) == 1 {
		sides = make([]Side, n[0])
		for i := 1; i <= n[0]; i++ {
			sides[i-1] = IntSide(i)
		}
	} else {
		sides = make([]Side, len(n))
		for i, num := range n {
			sides[i] = IntSide(num)
		}
	}
	d.Type = DieType{Name: name, Sides: sides}
	return d
}

// D4 returns a standard, four-sided die.
func D4() *Die {
	return NewIntDie("d4", 4)
}

// // D6 returns a standard, six-sided die.
// func D6() *Die {
// 	return NewIntDie("d6", 6)
// }

// D8 returns a standard, eight-sided die.
func D8() *Die {
	return NewIntDie("d8", 8)
}

// D10 returns a standard, ten-sided die.
func D10() *Die {
	return NewIntDie("d10", 10)
}

// TensD10 returns a ten sided die, ranging from 00 to 90.
func TensD10() *Die {
	return NewIntDie("10sd10", 0, 10, 20, 30, 40, 50, 60, 70, 80, 90)
}

// D12 returns a standard, twelve-sided die.
func D12() *Die {
	return NewIntDie("d12", 12)
}

// D20 returns a standard, twenty die.
func D20() *Die {
	return NewIntDie("d20", 20)
}


*/
