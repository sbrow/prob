//go:generate sh -c "godoc2md -template ./.doc.template github.com/sbrow/prob > README.md"

// Package prob contains a variety of tools for computing combinatorics and
// probability problems.
//
// Cmd
//
// The cmd sub-package contains executables for quickly computing specific things- combinations, permutations, etc.
//
// Combin
//
// The combin sub-package contains a library for using combinatorics in more sophisticated environments.
//
// Dice
//
// The dice sub-package contains a library for rolling dice as well as generating and evaluating roll tables.
package prob
