package LeetCode

import "unicode"

func letterCasePermutation(s string) []string {
	var ss []string
	ss = append(ss, "")
	for _, chr := range s {
		if unicode.IsLetter(chr) {
			for i := range ss {
				ss = append(ss, ss[i])
				ss[i] += string(unicode.ToLower(chr))
			}
			for i := len(ss) / 2; i < len(ss); i++ {
				ss[i] += string(unicode.ToUpper(chr))
			}
		} else {
			for i := range ss {
				ss[i] += string(chr)
			}
		}
	}
	return ss
}
