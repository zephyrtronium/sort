This will be a drop-in replacement for Go's standard sort package, featuring
multiple sorts, including some that are often faster than the standard.

Currently implemented:

 - Dual-pivot quicksort (DPQS)
 - Heapsort (HS)
 - Insertion sort (two versions; one using binary search to locate the correct
   position (IS) and one using linear search (ISLS))
 - Counting sort (for 8- and 16-bit integer types)

Scheduled:

 - Mergesort
 - Timsort

Tests and benchmarks are included for all implemented sorts except counting
sorts. Benchmarks will also run on the standard sort.Sort() under the
name Shootout (even though it's really an introsort).
