package main

import (
    "./gol"
    "./gen"
    "fmt"
)

func main() {
    // Setup simulation
    fmt.Println("Preparing simulation")
    l := gol.NewLife(256, 64)
    //l.SetPattern(32, 32, "xxx\n\n\n\nxxx")
    l.SetRle(4, 4, "./c.rle")

    // Generate
    gen.MakeSvg(l, "../pattern.svg", 1872, 10, 1000, 800)
    fmt.Println("Done")
}


