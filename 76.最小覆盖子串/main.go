package main

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	tbl := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tbl[t[i]]--
	}

	// detect first mach
	var res string
	for i := 0; i < len(s); i++ {
		if _, ok := tbl[s[i]]; ok {
			tbl[s[i]]++
			ok = func() bool {
				for _, v := range tbl {
					if v < 0 {
						return false
					}
				}
				return true
			}()
			if ok {
				res = s[0 : i+1]
				if len(res) == len(t) {
					return res
				}

				break
			}
		}
	}
	if len(res) == 0 {
		return ""
	}

	for l, r := 0, len(res)-1; ; {
		// move left
		c := s[l]
		l++
		if _, ok := tbl[c]; !ok {
			if len(res) > r-l+1 {
				if len(t) == r-l+1 {
					return s[l : r+1]
				}
				res = s[l : r+1]
			}
		} else {
			tbl[c]--
			if tbl[c] >= 0 {
				if len(res) > r-l+1 {
					if len(t) == r-l+1 {
						return s[l : r+1]
					}
					res = s[l : r+1]
				}
			} else {
				// move right
				for r = r + 1; r < len(s) && s[r] != c; {
					if _, ok := tbl[s[r]]; ok {
						tbl[s[r]]++
					}

					r++
				}
				if r == len(s) {
					return res
				}
				tbl[s[r]]++
				if len(res) > r-l+1 {
					res = s[l : r+1]
				}
			}
		}
	}

}
