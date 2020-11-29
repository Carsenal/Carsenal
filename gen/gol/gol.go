package gol

import (
    "../bitstring"
    "sync"
//    "fmt"
)

type Life struct {
	W, H          uint
	a, b          bitstring.Bitstring
	Current, Past *bitstring.Bitstring
}

func NewLife(w, h uint) *Life {
    l := Life{
        W: w,
        H: h,
    }
    l.a = *bitstring.NewBitstring(w, h)
    l.b = *bitstring.NewBitstring(w, h)
    l.Current = &l.a
    l.Past = &l.b
    return &l
}

func (l *Life) StepCell(x, y uint, wg *sync.WaitGroup) {
    sum := 0
    if l.Past.Get(x-1, y-1) {
        sum++
    }
    if l.Past.Get(x-1, y  ) {
        sum++
    }
    if l.Past.Get(x-1, y+1) {
        sum++
    }
    if l.Past.Get(x  , y-1) {
        sum++
    }
    if l.Past.Get(x  , y+1) {
        sum++
    }
    if l.Past.Get(x+1, y-1) {
        sum++
    }
    if l.Past.Get(x+1, y  ) {
        sum++
    }
    if l.Past.Get(x+1, y+1) {
        sum++
    }
    if sum == 3 || (sum == 2 && l.Past.Get(x, y)) {
        //fmt.Printf("Setting %d, %d to true\n", x, y)
        l.Current.Set(x, y, true)
    } else {
        //fmt.Printf("Setting %d, %d to false\n", x, y)
        l.Current.Set(x, y, false)
    }
    wg.Done()
}

/*
func (l *Life) SetPattern(x, y uint, str []string) {
    var i, j uint
    for i = 0; i < len(str); i++ {
        for j = 0; j < len(str[i]); j++ {
            l.Current.Set
        }
    }
}
*/

func (l *Life) Step() {
    var wg sync.WaitGroup
    var x, y uint
    l.Current, l.Past = l.Past, l.Current
    for x = 0; x < l.W; x++ {
        for y = 0; y < l.H; y++ {
            wg.Add(1)
            go l.StepCell(x, y, &wg)
        }
    }
    wg.Wait()
}

func (l *Life) ToString() string {
    return l.Current.ToString()
}

