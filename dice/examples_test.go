package dice

import "fmt"

func ExampleDelim() {
	Delim = "+"
	dice := Dice{D6(), D6()}
	fmt.Println(Roll(dice...))
	// Output: [6 6]
}

func ExampleRoll() {
	fmt.Println(Roll(D6(), D6()))
	// Output: [2 1]
}

func ExampleDice_Name() {
	dice := Dice{D8(), D4(), D10(), D6(), D4(), D6()}
	fmt.Print(dice.Name())
	// Output: 2d6+2d4+1d8+1d10
}
