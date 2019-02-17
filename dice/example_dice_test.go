package dice

import "fmt"

func ExampleDie_Roll() {
	fmt.Println(new(D6).Roll())
	// Output:
	// 4
}

// func ExampleDice_Roll() {
// 	fmt.Println(new(D6).Roll())
// 	fmt.Println(New(D6, D6).Roll())
// 	// Output:
// 	// [6]
// 	// [6 2]
// }

// func ExampleCopy() {
// 	dice2D6 := Copy(new(D6), 2)
// 	fmt.Println(dice2D6)
// 	// Output:
// 	// [ d6 d6 ]
// }
