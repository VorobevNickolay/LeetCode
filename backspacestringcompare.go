package LeetCode

import (
	"fmt"
	"unicode/utf8"
)

func backspaceCompare(s string, t string) bool {
	sPointer := utf8.RuneCountInString(s) - 1
	tPointer := utf8.RuneCountInString(t) - 1
	sDel := 0
	tDel := 0
	for true {
		if sPointer >= 0 {
			for s[sPointer] == '#' || sDel > 0 {
				if s[sPointer] == '#' {
					sDel += 2
				}
				sDel--
				sPointer--
				if sPointer < 0 {
					break
				}
			}
		}
		if tPointer >= 0 {
			for t[tPointer] == '#' || tDel > 0 {
				if t[tPointer] == '#' {
					tDel += 2
				}
				tDel--
				tPointer--
				if tPointer < 0 {
					break
				}
			}
		}
		fmt.Println(sPointer, tPointer)
		if sPointer < 0 || tPointer < 0 {
			if sPointer < 0 && tPointer < 0 {
				return true
			} else if sPointer < 0 {
				if t[tPointer] != '#' {
					return false
				}
				for j := tPointer; j >= 0; j-- {
					if t[j] == '#' {
						tDel += 2
					}
				}
				return tPointer < tDel
			} else {
				if s[sPointer] != '#' {
					return false
				}
				for j := sPointer; j >= 0; j-- {
					if s[j] == '#' {
						sDel += 2
					}
				}
				return sPointer < sDel
			}
		}
		if s[sPointer] != t[tPointer] {
			return false
		}
		sPointer--
		tPointer--
	}
	if tPointer < 0 && sPointer < 0 {
		return true
	}
	return false
}
