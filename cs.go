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

func CS8U(data []uint8) {
	counts := [256]int{}
	for _, v := range data {
		counts[int(v)]++
	}
	n := 0
	for v, i := range counts {
		for i >= 0 {
			data[n] = uint8(v)
			n++
			i--
		}
	}
}

func CS8S(data []int8) {
	counts := [256]int{}
	for _, v := range data {
		counts[int(v)+128]++
	}
	n := 0
	for v, i := range counts {
		for i >= 0 {
			data[n] = int8(v - 128)
			n++
			i--
		}
	}
}

func CS16U(data []uint16) {
	counts := [65536]int{}
	for _, v := range data {
		counts[int(v)]++
	}
	n := 0
	for v, i := range counts {
		for i >= 0 {
			data[n] = uint16(v)
			n++
			i--
		}
	}
}

func CS16S(data []int16) {
	counts := [65536]int{}
	for _, v := range data {
		counts[int(v)+32768]++
	}
	n := 0
	for v, i := range counts {
		for i >= 0 {
			data[n] = int16(v - 32768)
			n++
			i--
		}
	}
}
