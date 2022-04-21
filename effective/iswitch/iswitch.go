package iswitch

// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be
// 0 if a==b
// -1 if a<b
// +1 if a>b
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}

	return 0
}
