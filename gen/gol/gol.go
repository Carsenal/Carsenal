package gol

import (
	"../bitstring"
	"sync"
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
	if l.Past.Get(x-1, y) {
		sum++
	}
	if l.Past.Get(x-1, y+1) {
		sum++
	}
	if l.Past.Get(x, y-1) {
		sum++
	}
	if l.Past.Get(x, y+1) {
		sum++
	}
	if l.Past.Get(x+1, y-1) {
		sum++
	}
	if l.Past.Get(x+1, y) {
		sum++
	}
	if l.Past.Get(x+1, y+1) {
		sum++
	}
	if sum == 3 || (sum == 2 && l.Past.Get(x, y)) {
		l.Current.Set(x, y, true)
	} else {
		l.Current.Set(x, y, false)
	}
	wg.Done()
}

func (l *Life) SetPattern(x, y uint, str string) {
	start := x
	for _, c := range str {
		if c == '\n' {
			x = start
			y++
		} else {
			if c == '.' {
				l.Current.Set(x, y, false)
			} else {
				l.Current.Set(x, y, true)
			}
			x++
		}
	}
}

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
