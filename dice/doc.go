/*
Package dice provides functions for simulating and analyzing dice probabilities.

Monte Carlo

For Monte Carlo simulations, i.e. Rolling dice a large number of times
to get the approximate result,
you can use any of: Roll(Dice), Dice.Roll() or Die.Roll().

Die.Roll rolls a single die,
Roll(Dice) and Dice.Roll() roll multiple dice.

Keep in mind that programs will always return the same dice rolls unless you provide
a different seed every time, (see "random.Seed").

Enumeration

If you wish to determine exact results of a die roll, you can enumerate a roll table
by calling NewTable(DiceToRoll).

TODO: Add Table.Verify()
*/
package dice
