package main

import (
    "fmt"
    "os"
    igo "github.com/takluyver/igo/igopkg"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Need a command line argument for the connection file.")
        os.Exit(1)
    }
    igo.RunKernel(os.Args[1])
}

