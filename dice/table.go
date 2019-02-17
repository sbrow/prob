package dice

import (
	"os"
	"path/filepath"
)

// DataDir is the default directory to store Tables in.
//
// DEPRECATED
var DataDir = filepath.Join(os.Getenv("GOPATH"), "data")

// Delim specifies the delimiter to use for dice in a roll.
var Delim = "+"
