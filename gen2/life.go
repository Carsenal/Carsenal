package life

import (
    ""
)

type Life struct {
	w, h, r uint
	a, b    [][]bool
	c, o    *[][]bool
}

func NewLife(w, h uint) *Life {

}

func (g *[][]bool) GetCell (x, y uint) bool {
    // TODO: overflow check
    return g.c[x][y]
}

func (g *[][]bool) ToString () string {
    var str string
}

