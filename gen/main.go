package main

import (
    "./gol"
    "./gen"
    "fmt"
)

func main() {
    l := gol.NewLife(10, 8)
    l.Current.Set(4, 4, true)
    l.Current.Set(5, 4, true)
    l.Current.Set(6, 4, true)
    fmt.Printf("%s\n", l.ToString())
    fmt.Println("Starting gen")
    gen.MakeSvg(l, 10)
}

