package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/vohumana/vohumana-gointerpreter/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
    fmt.Printf("Feel free to type in commands")

    repl.Start(os.Stdin, os.Stdout)
}