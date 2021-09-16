package core

var Keywords = map[string]func(args []string) {
	"def": DefineVariable,
}