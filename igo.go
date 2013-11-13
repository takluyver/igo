// The igo command starts a Go kernel for IPython.
package main

import (
    "flag"
    "io"
    "log"
    "os"
    igo "github.com/takluyver/igo/igopkg"
)

func main() {
    debug := flag.Bool("debug", false, "Log extra info to stderr")
    flag.Parse()
    if flag.NArg() < 1 {
        log.Fatalln("Need a command line argument for the connection file.")
    }
    var logwriter io.Writer = os.Stderr
    var err error
    if !*debug {
        logwriter, err = os.OpenFile(os.DevNull, os.O_WRONLY, 0666)
        if err != nil {
            log.Fatalln(err)
        }
    }
    igo.RunKernel(flag.Arg(0), logwriter)
}

