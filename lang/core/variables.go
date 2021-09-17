package core

func DefineVariable(args []string) {
  LexicalScope[args[1]] = args[2]
}