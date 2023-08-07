package main

// use hash table
//func isAnagram(s string, t string) bool {
//	cnt := make(map[byte]int, len(s))
//
//	for i := range s {
//		cnt[s[i]] += 1
//	}
//
//	for i := range t {
//		cnt[t[i]] -= 1
//	}
//
//	for _, v := range cnt {
//		if v != 0 {
//			return false
//		}
//	}
//
//	return true
//}

// user array to simulate a hash table
func isAnagram(s string, t string) bool {
	cnt1, cnt2 := [26]byte{}, [26]byte{}

	for i := 0; i < len(s); i++ {
		cnt1[s[i]-'a'] += 1
	}

	for i := 0; i < len(t); i++ {
		cnt2[t[i]-'a'] += 1
	}

	return cnt1 == cnt2
}
