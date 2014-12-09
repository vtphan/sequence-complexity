package main

import (
    "fmt"
    "os"
    comp "github.com/vtphan/sequence-complexity"
)

func main(){
    if len(os.Args) != 2 {
        panic("must provide sequence file.")
    }
    idx := new(comp.Index)
    idx.Build(os.Args[1])
    fmt.Printf("%s\tD\t%f\n", os.Args[1], idx.D())
    fmt.Printf("%s\tI\t%f\n", os.Args[1], idx.I())
    for k:=25; k<=200; k+=25 {
        fmt.Printf("%s\tD%d\t%f\n", os.Args[1], k, idx.Dk(k))
    }
}