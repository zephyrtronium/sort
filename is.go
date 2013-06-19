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

// Insertion sort partitions the input into [A, B], where the elements of A
// are in sorted order, then inserts the first element of B into A, such that
// A is kept sorted. This means it performs n comparisons and zero swaps on
// sorted data.
//
// This implementation is stable.
// Average-case comparisons: Ο(n lg n).
// Average-case swaps: Ο(n²).
func IS(data Swapper) {
	doIS(data, 0, data.Len())
}

func doIS(data Swapper, low, high int) {
	for i := low + 1; i < high; i++ {
		if data.Less(i, i-1) {
			j := binsearch(data, i, low, i-1)
			for k := i; k > j; k-- {
				data.Swap(k, k-1)
			}
		}
	}
}

func binsearch(data Swapper, x, low, high int) int {
	len := high - low
	if len == 0 {
		return low
	}
	if len == 1 {
		if data.Less(x, low) {
			return low
		}
		return low + 1
	}
	m := low + len>>1
	if data.Less(x, m) {
		return binsearch(data, x, low, m)
	}
	return binsearch(data, x, m, high)
}

func ISLS(data Swapper) {
	doISLS(data, 0, data.Len())
}

func doISLS(data Swapper, low, high int) {
	for i := low + 1; i < high; i++ {
		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}
