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
	"math/rand"
	"sort"
	"testing"
)

func TestIS(t *testing.T) {
	for i := 0; i < 10; i++ {
		// fewer elements for is because it's much slower for random input
		d := &tracker{d: make([]int, 1+rand.Intn(999))}
		for j := range d.d {
			d.d[j] = rand.Int()
		}
		t.Log(d)
		IS(d)
		if !sort.IsSorted(d) {
			t.Fatal("IS failed on slice of length", len(d.d), d)
		}
	}
}

func BenchmarkRandomIS1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 1000)
		for j := range d {
			d[j] = rand.Int()
		}
		b.StartTimer()
		IS(d)
	}
}

func BenchmarkSortedIS1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 1000)
		for j := range d {
			d[j] = j
		}
		b.StartTimer()
		IS(d)
	}
}

func BenchmarkReverseIS1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 1000)
		for j := range d {
			d[j] = len(d) - j
		}
		b.StartTimer()
		IS(d)
	}
}

func BenchmarkSawtoothIS1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 1000)
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
		IS(d)
	}
}

func TestISLS(t *testing.T) {
	for i := 0; i < 10; i++ {
		// fewer elements for is because it's much slower for random input
		d := &tracker{d: make([]int, 1+rand.Intn(999))}
		for j := range d.d {
			d.d[j] = rand.Int()
		}
		t.Log(d)
		ISLS(d)
		if !sort.IsSorted(d) {
			t.Fatal("ISLS failed on slice of length", len(d.d), d)
		}
	}
}

func BenchmarkRandomISLS1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 1000)
		for j := range d {
			d[j] = rand.Int()
		}
		b.StartTimer()
		ISLS(d)
	}
}

func BenchmarkSortedISLS1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 1000)
		for j := range d {
			d[j] = j
		}
		b.StartTimer()
		ISLS(d)
	}
}

func BenchmarkReverseISLS1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 1000)
		for j := range d {
			d[j] = len(d) - j
		}
		b.StartTimer()
		ISLS(d)
	}
}

func BenchmarkSawtoothISLS1k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		d := make(sort.IntSlice, 1000)
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
		ISLS(d)
	}
}
