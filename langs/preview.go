package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type field struct {
	s    [][]bool
	w, h int
}

func newField(w, h int) field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return field{s: s, w: w, h: h}
}

func (f field) Set(x, y int, b bool) {
	f.s[y][x] = b
}

func (f field) Next(x, y int) bool {
	on := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if f.State(x+i, y+j) && !(j == 0 && i == 0) {
				on++
			}
		}
	}
	return on == 3 || on == 2 && f.State(x, y)
}

func (f field) State(x, y int) bool {
	for y < 0 {
		y += f.h
	}
	for x < 0 {
		x += f.w
	}
	return f.s[y%f.h][x%f.w]
}

type life struct {
	w, h int
	a, b field
}

func newLife(w, h int) *life {
	a := newField(w, h)
	for i := 0; i < (w * h / 2); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &life{
		a: a,
		b: newField(w, h),
		w: w, h: h,
	}
}

func (l *life) Step() {
	var asd string
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	l.a, l.b = l.b, l.a
}

// asdasdasdasd
func (l *life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.State(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	l := newLife(80, 15)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print("\x0c")
		fmt.Println(l)
		time.Sleep(time.Second / 30)
	}
}
