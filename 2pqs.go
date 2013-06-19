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

// Threshold below which to switch from DPQS to IS.
const isThresh = 32

// Dual-pivot quicksort chooses two pivots P and Q and recursively sorts the
// list into [x < P, P <= x <= Q, Q < x].
//
// This implementation is unstable.
// Average comparisons: Ο(n lg n).
// Average swaps: Ο(n lg n).
// Auxiliary space: Ω(log3 n) (recursion).
func DPQS(data Swapper) {
	doDPQS(data, 0, data.Len()-1)
}

func doDPQS(data Swapper, low, high int) {
	if high-low < isThresh {
		doISLS(data, low, high+1)
		return
	}

	_, p, q, _ := getPivots(data, low, high)
	if data.Less(p, q) {
		// Partition:
		// low is p
		// (low, r1) < p
		// p <= [r1, r2) <= q
		// [r2, r3] is not yet partitioned
		// q < (r3, high)
		// high is q
		r1, r3 := low+1, high-1
		data.Swap(p, low)
		data.Swap(q, high)
		p, q, low, high = low, high, p, q
		for data.Less(r1, p) {
			r1++
		}
		for data.Less(q, r3) {
			r3--
		}
		r2 := r1
		for r2 <= r3 {
			if data.Less(r2, p) {
				data.Swap(r2, r1)
				r1++
				r2++
			} else if data.Less(q, r2) {
				data.Swap(r2, r3)
				r3--
			} else {
				r2++
			}
		}
		data.Swap(p, r1-1)
		data.Swap(q, r3+1)
		p, q, low, high = low, high, p, q
		doDPQS(data, low, r1-2)
		doDPQS(data, r1, r3)
		doDPQS(data, r3+2, high)
	} else {
		// P == Q
		r1, r3 := low, high
		r2 := r1
		for r2 <= r3 {
			if data.Less(r2, p) {
				data.Swap(r2, r1)
				r1++
				r2++
			} else if data.Less(p, r2) {
				data.Swap(r2, r3)
				r3--
			} else {
				r2++
			}
		}
		doDPQS(data, low, r1-1)
		doDPQS(data, r3+1, high)
	}
}

func getPivots(data Swapper, low, high int) (int, int, int, int) {
	sev := (high - low + 1) / 7
	mid := (high + low) >> 1
	ml, mu := mid-sev, mid+sev
	l, u := ml-sev, mu+sev
	// Mini-sort insertion-style on these five positions.
	if data.Less(ml, l) {
		data.Swap(ml, l)
	}
	if data.Less(mid, ml) {
		data.Swap(mid, ml)
		if data.Less(ml, l) {
			data.Swap(ml, l)
		}
	}
	if data.Less(mu, mid) {
		data.Swap(mu, mid)
		if data.Less(mid, ml) {
			data.Swap(mid, ml)
			if data.Less(ml, l) {
				data.Swap(ml, l)
			}
		}
	}
	if data.Less(u, mu) {
		data.Swap(u, mu)
		if data.Less(mu, mid) {
			data.Swap(mu, mid)
			if data.Less(mid, ml) {
				data.Swap(mid, ml)
				if data.Less(ml, l) {
					data.Swap(ml, l)
				}
			}
		}
	}
	// Second and fourth positions make good pivots.
	return l, ml, mu, u
}
