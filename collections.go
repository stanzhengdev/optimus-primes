// We often need our programs to perform operations on
// collections of data, like selecting all items that
// satisfy a given predicate or mapping all items to a new
// collection with a custom function.

// In some languages it's idiomatic to use [generic](http://en.wikipedia.org/wiki/Generic_programming)
// data structures and algorithms. Go does not support
// generics; in Go it's common to provide collection
// functions if and when they are specifically needed for
// your program and data types.

// Here are some example collection functions for slices
// of `strings`. You can use these examples to build your
// own functions. Note that in some cases it may be
// clearest to just inline the collection-manipulating
// code directly, instead of creating and calling a
// helper function.
// https://play.golang.org/p/3PNdke3Wia

package optimusprime

// Returns the first index of the target string `t`, or
// -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Returns `true` if the target string t is in the
// slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Returns `true` if one of the strings in the slice
// satisfies the predicate `f`.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Returns `true` if all of the strings in the slice
// satisfy the predicate `f`.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Returns a new slice containing the results of applying
// the function `f` to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
