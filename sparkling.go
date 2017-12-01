package sparkling

import (
	"fmt"
	"io"
	"strings"
	"time"
)

const (
	ticks = `▁▂▃▄▅▆▇█`
)

// Sparkling holds the sparklines.
type Sparkling struct {
	out   io.Writer
	lines []Series
}

// New starts a new sparkling.
func New(writer io.Writer) *Sparkling {
	sp := &Sparkling{out: writer}
	return sp
}

// AddSeries adds a new series.
func (sp *Sparkling) AddSeries(data []float64, title string) {
	s := Series{
		data:  data,
		title: title,
		out:   sp.out,
	}
	sp.lines = append(sp.lines, s)
}

// Render renders a set of sparklines.
func (sp *Sparkling) Render() {
	for _, l := range sp.lines {
		l.Draw()
		fmt.Fprintf(sp.out, "\n")
	}
}

// Series holds the data and title for a sparkline.
type Series struct {
	data  []float64
	title string
	out   io.Writer
}

// NewSeries returns a series initialized with data.
func NewSeries(data []float64, title string) *Series {
	s := &Series{
		data:  data,
		title: title,
	}
	return s
}

// Draw draws the line with the corresponding data series.
func (s *Series) Draw() {
	series := s.data
	bars := []rune(ticks)

	min, max := float64(0), float64(0)
	for _, v := range series {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	// use high tick if data is constant
	if min == max {
		bars = []rune(`▅▆`)
	}

	f := (int((max - min)) << 8) / (len(bars) - 1)
	if f < 1 {
		f = 1
	}

	var out string
	for _, v := range series {
		out += string(bars[(int((v-min))<<8)/f])
		fmt.Fprintf(s.out, "\r%s", strings.TrimSpace(fmt.Sprintf("%s %s", s.title, out)))
		time.Sleep(100 * time.Millisecond)
	}
}
