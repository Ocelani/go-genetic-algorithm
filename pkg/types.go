package pkg

import "strings"

var (
	corpus = strings.Split("abcdefghijklmnopqrstuvwxyz ", "")
	target = strings.Split("software release", "")
)

// Strings is a slice of strings.
type Strings []string
