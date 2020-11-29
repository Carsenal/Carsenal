package gen

import (
	"../bitstring"
	"../gol"
	"fmt"
	"strings"
	"os"
)

func MakeSvg(l *gol.Life, rounds uint, filename string, dur, width, height uint) {
	// Generate frame data
	cells := generateStates(l, rounds)
	// Open file for write
	f, _ := os.Create(filename)
	defer f.Close()
	// Write initial junk
	f.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>")
	f.WriteString("<!DOCTYPE svg>")
	f.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" ")
	f.WriteString(fmt.Sprintf("viewbox=\"0 0 %d %d\" ", l.W, l.H)
	f.WriteString(fmt.Sprintf("width=\"%d\" height=\"%d\">", width, height)
	// Write every cell
	for _, cell := range cells {
    	f.WriteString(cell.toSvg(dur))
	}
	// Write closing
	f.WriteString("</svg>")
}

// Struct to hold a state
type cellState struct {
	x, y  uint
	start uint
	dur   uint
}

// Struct to hold cell information
type cell struct {
	states []cellState
}

func generateStates(l *gol.Life, rounds uint) (cells []cell) {
	// Variables
	var i, id uint
	var plot [][]uint
	var born, died bitstring.Bitstring

	// Variable init
	plot = make([][]int, l.W)
	for i = 0; i < l.W; i++ {
		plot[i] = make([]int, l.H)
	}

	// Handle starting
	for coord := range l.Current.List() {
		plot[coord[0]][coord[1]] = len(cells)
		cells.appendNewCell(coord[0], coord[1], 0)
	}

	// Calculate rounds
	for i = 0; i < rounds; i++ {
		// Step
		l.Step()

		// Handle deaths
		died = l.Past.NowOn(l.Current)
		for coord := range died.List() {
			id = plot[coord[0]][coord[1]]
			plot[coord[0]][coord[1]] = -1
			cells[id].setLastDuration(i)
		}

		// Handle births
		born = l.Current.NowOn(l.Past)
		for coord := range born.List() {
			if len(open) > 0 {
				id = open[0]
				plot[coord[0]][coord[1]] = id
				open[0] = ""
				open = open[1:]
				cells[id].addState(coord[0], coord[1], i)
			} else {
				id = len(cells)
				cells.appendNewCell(coord[0], coord[1], i)
			}
		}
	}

	// Set duration for survivors
	for coord := range l.Current.List() {
		id = plot[coord[0]][coord[1]]
		cells[id].setLastDuration(rounds)
	}
	return
}

// Methods for cells
func (arr *[]cell) appendNewCell(x, y, time uint) {
	arr = append(arr, newCell(x, y, time))
}

func newCell(x, y, time uint) *cell {
	c := cell{}
	c.addState(x, y, time)
	return &c
}

func (c *cell) setLastDuration(time uint) {
	index := len(c.states) - 1
	c.states[index].duration = time - c.states[index].start
}

func (c *cell) addState(x, y, time uint) {
	c.states = append(c.states, state{x: x, y: y, start: time})
}

func (c *cell) states(rounds int) []bool {
	status := make([]bool, rounds)
	for i := 0; i < rounds; i++ {
    	status[i] = false
	}
	for _, state := range c.states {
		for i := state.start; i < (state.duration + state.start); i++ {
    		status[i] = true
		}
	}
	return status
}

func (c *cell) coords(rounds int) []bool {
	status := make([]bool, rounds)
	for i := 0; i < rounds; i++ {
    	status[i] = false
	}
	for _, state := range c.states {
		for i := state.start; i < (state.duration + state.start); i++ {
    		status[i] = true
		}
	}
	return status
}

func (c *cell) listOpacity() chan int{
    ch := make(chan bool)
    go func(c *cell) {
        close(ch)
    } (c)
    return ch
}

func (c *cell) toSvg(dur uint) string {
    // Animate opacity
    opacityStr := animateStr(c.listOpacity(), "opacity", dur)
    // Animate x
    xStr := animateStr(c.listX(), "x", dur)
    // Animate y
    yStr := animateStr(c.listY(), "y", dur)
    return fmt.Sprintf(
        "<rect width=\"0.9\" height=\"0.9\"> %s %s %s </rect>",
        opacityStr,
        xStr,
        yStr
    )
}

func animateStr(ch chan int, name string, dur uint) string {
    var strArr []string
    for val := range ch {
        strArr = append(strArr, fmt.Sprintf("%v", val))
    }
	return fmt.Sprintf(
    	"<animate attributeName=\"%s\" values=\"%s\" dur=\"%ds\" repeatCount=\"indefinite\"/>",
    	name, strings.Join(val[], ";")
	)
}

