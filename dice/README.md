# dice
[![GoDoc](https://godoc.org/github.com/sbrow/dice?status.svg)](https://godoc.org/github.com/sbrow/dice) [![Build Status](https://travis-ci.org/sbrow/dice.svg?branch=master)](https://travis-ci.org/sbrow/dice) [![Coverage Status](https://coveralls.io/repos/github/sbrow/dice/badge.svg?branch=master)](https://coveralls.io/github/sbrow/dice?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/sbrow/dice)](https://goreportcard.com/report/github.com/sbrow/dice)

Package dice provides functions for simulating and analyzing dice probabilities.


### Monte Carlo

For Monte Carlo simulations, i.e. Rolling dice a large number of times to get
the approximate result, you can use any of: Roll(Dice), Dice.Roll() or
Die.Roll().

Die.Roll rolls a single die, Roll(Dice) and Dice.Roll() roll multiple dice.

Keep in mind that programs will always return the same dice rolls unless you
provide a different seed every time, (see "random.Seed").


### Enumeration

If you wish to determine exact results of a die roll, you can enumerate a roll
table by calling NewTable(DiceToRoll).

TODO: Add Table.Verify()

## Installation
```bash
$ go get -u github.com/sbrow/prob/dice
```
