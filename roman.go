package roman

func Roman(s string) int {

	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	mincount := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]
		sign := 1
		if temp < res {
			mincount++
			//左减数字必须为一位
			if mincount > 1 {
				return 0
			}
			//左减的数字有限制，仅限于I、X、C
			if (s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X')) || (s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C')) || (s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M')) {
				sign = -1
			} else if !(s[i] == s[i+1] && s[i] == s[i+2]) {
				//右加数字不可连续超过三位
				return 0
			}

		} else {
			mincount = 0
		}
		res += sign * temp
	}

	return res
}
