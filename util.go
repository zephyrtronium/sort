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

import "sort"

type (
	// Interface for sortable data structures.
	Comparator interface {
		Len() int
		Less(i, j int) bool
	}

	// Interface for sorts that swap. (Same as built-in sort.Interface.)
	Swapper interface {
		Comparator
		Swap(i, j int)
	}

	// Interface for sorts that insert.
	Inserter interface {
		Comparator
		// Insert the item at position j into position i without rearranging
		// items between. The item that was at i prior to insertion should be
		// at i+1 afterward.
		Insert(i, j int)
	}
)

// Create an Interface from an Inserter.
type SwapByInsert struct {
	Inserter
}

// Swaps i and j with two inserts.
func (s SwapByInsert) Swap(i, j int) {
	if i != j {
		s.Insert(i, j)
		if i < j {
			s.Insert(j+1, i+1)
		} else {
			s.Insert(j, i+1)
		}
	}
}

// Create an Inserter from an Interface.
type InsertBySwap struct {
	Swapper
}

// Insert the item at position j into position i with |j-i| swaps.
func (s InsertBySwap) Insert(i, j int) {
	if i < j {
		for j >= i {
			s.Swap(j-1, j)
			j--
		}
	} else if i > j {
		for j <= i {
			s.Swap(j+1, j)
			j++
		}
	}
}

type ranje struct {
	Swapper
	low, high int
}

// Sort [from, to) of on.
func Range(on Swapper, from, to int) Swapper {
	return ranje{on, from, to}
}
func (r ranje) Len() int           { return r.high - r.low }
func (r ranje) Less(i, j int) bool { return r.Swapper.Less(i-r.low, j-r.low) }
func (r ranje) Swap(i, j int)      { r.Swapper.Swap(i-r.low, j-r.low) }

// Proxies to make this a drop-in replacement for built-in sort.

// Same as a Swapper. Drop-in replacement for built-in sort.Interface.
type Interface interface {
	Swapper
}

// Sort in descending order. (This is a proxy to built-in sort.Reverse().)
func Reverse(of Swapper) Swapper { return sort.Reverse(of) }

// Determine whether data is sorted.
func IsSorted(data Comparator) bool {
	n := data.Len()
	for i := 1; i < n; i++ {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

//TODO: finish.
