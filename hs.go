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

// Heapsort builds a heap of the data, then repeatedly chooses the root and
// rebuilds the heap.
//
// Heapsort is unstable.
// Average comparisons: Ο(n lg n).
// Average swaps: Ο(n lg n).
func HS(data Swapper) {
	doHS(data, 0, data.Len())
}

func doHS(data Swapper, a, b int) {
	len := b - a
	for i := (len - 1) >> 1; i >= 0; i-- {
		heapify(data, i, len, a)
	}
	for i := len - 1; i >= 0; i-- {
		data.Swap(a, a+i)
		heapify(data, 0, i, a)
	}
}

func heapify(data Swapper, root, high, first int) {
	for {
		child := 2*root + 1
		if child < high {
			if child+1 < high && data.Less(first+child, first+child+1) {
				child++
			}
			if !data.Less(first+root, first+child) {
				return
			}
			data.Swap(first+root, first+child)
			root = child
		} else {
			return
		}
	}
}
