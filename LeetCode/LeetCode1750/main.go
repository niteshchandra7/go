package main

func minimumLength(s string) int {
	first := 0
	last := len(s) - 1
	for first < last {
		if s[first] == s[last] {
			char := s[first]
			for first <= last && s[first] == char {
				first++
			}
			for first <= last && s[last] == char {
				last--
			}

		} else {
			return last - first + 1
		}
	}
	if first == last {
		return 1
	}
	return 0
}

func main() {

}
