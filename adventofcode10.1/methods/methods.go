package methods

import "fmt"

func CalculateFirst(str string) (s string) {
	res := ""
	count := 1

	for k, i := 0, 0; i < len(str); i, k = i+1, k+i+1 {
		if str[i] != str[k] {

			i = k
			break
		}
		if str[i] == str[i] {
			count++
		}

		fmt.Println(count, string(str[i-1]))
		count = 1

	}
	return res
}
