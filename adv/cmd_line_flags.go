// go build cmd_line_flags.go
// ./cmd_line_flags -word=opt -numb=7 -fork -svar=flag
// ./cmd_line_flags -word=opt
// ./cmd_line_flags -word=opt a1 a2 a3
// ./cmd_line_flags -word=opt a1 a2 a3 -numb=7
// ./cmd_line_flags -h

package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *forkPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}