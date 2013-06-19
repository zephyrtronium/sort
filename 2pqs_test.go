/*
Copyright (c) 2013 Branden J Brown

This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

   1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.

   2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.

   3. This notice may not be removed or altered from any source
   distribution.
*/

package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

type tracker struct {
	d    []int
	c, s int
}

func (d *tracker) Less(i, j int) bool {
	d.c++
	return d.d[i] < d.d[j]
}
func (d *tracker) Swap(i, j int) {
	d.s++
	d.d[i], d.d[j] = d.d[j], d.d[i]
}
func (d *tracker) Len() int {
	return len(d.d)
}
func (d *tracker) String() string {
	return fmt.Sprintf("%d elements %d compares %d swaps", len(d.d), d.c, d.s)
}

func TestDPQS(t *testing.T) {
	for i := 0; i < 100; i++ {
		d := &tracker{d: make([]int, 33+0*rand.Intn(9999))}
		for j := range d.d {
			d.d[j] = rand.Int()
		}
		DPQS(d)
		if !sort.IsSorted(d) {
			for _, v := range d.d {
				t.Logf("%16x\n", v)
			}
			t.Fatal("DPQS failed on slice of length", len(d.d), d)
		}
	}
}

func BenchmarkRandomDPQS10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 10000)
		for j := range d {
			d[j] = rand.Int()
		}
		b.StartTimer()
		DPQS(d)
	}
}

func BenchmarkSortedDPQS10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 10000)
		for j := range d {
			d[j] = j
		}
		b.StartTimer()
		DPQS(d)
	}
}

func BenchmarkReverseDPQS10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 10000)
		for j := range d {
			d[j] = len(d) - j
		}
		b.StartTimer()
		DPQS(d)
	}
}

func BenchmarkSawtoothDPQS10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 10000)
		x, y := 0, 1
		for j := range d {
			d[j] = x
			x++
			if x == y {
				x = 0
				y++
			}
		}
		b.StartTimer()
		DPQS(d)
	}
}
