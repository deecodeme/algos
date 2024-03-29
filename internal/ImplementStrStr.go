package main

/*
Time complexity:
Space complexity:
*/
func needlePointerInHaystackNaive(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) || len(haystack) == 0 {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
