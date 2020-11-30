package gol

import (
    "../bitstring"
    "sync"
    "unicode"
    "bufio"
    "os"
    "fmt"
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

func (l *Life) SetRle(x, y uint, filename string) {
    fmt.Printf("Parsing file %s\n", filename)
    var w, h int
    var data string
    f, err := os.Open(filename)
    if err != nil {
        fmt.Printf("Err: %v\n", err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        if scanner.Text()[0] == '#' {
            continue
        } else if scanner.Text()[0] == 'x' {
            contine
        } else {
            data += scanner.Text()
        }
    }
    fmt.Printf("w: %d, h: %d\n", w, h)

    // Parse data str
    var coeff uint
    coeff = 0
    start := x
    for _, c := range data {
        if unicode.IsDigit(c) {
            coeff = coeff*10 + uint(c - '0')
        } else if c == 'b' {
            if coeff == 0 {
                coeff = 1
            }
            dest := x + coeff
            for ; x < dest; x++ {
                //l.Current.Set(x, y, false)
            }
            coeff = 0
        } else if c == 'o' {
            if coeff == 0 {
                coeff = 1
            }
            dest := x + coeff
            for ; x < dest; x++ {
                l.Current.Set(x, y, true)
            }
            coeff = 0
        } else if c == '$' {
            x = start
            y++
            coeff = 0
        } else if c == '!' {
            // done
            fmt.Printf("Resultant:\n")
            fmt.Println(l.ToString())
            return
        } else {
            fmt.Printf("Unrecognized char '%s'\n", c)
        }
    }

    fmt.Printf("Resultant:\n")
    fmt.Println(l.ToString())
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

