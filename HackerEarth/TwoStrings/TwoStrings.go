package main

import (
	"fmt"
	"math"
)

func main() {

	var cnt int
	fmt.Scanf("%d", &cnt)
	if cnt < 1 || cnt > 100 {
		return
	}

	for i := 0; i < cnt; i++ {

		var str1, str2 string
		fmt.Scanf("%s%s", &str1, &str2)
		if len(str1) < 1 || len(str2) < 1 || float64(len(str1)) > math.Pow10(5) || float64(len(str2)) > math.Pow10(5) {
			return
		}

		str1Map := generateMap(str1)
		str2Map := generateMap(str2)

		if calculateForString1(str2Map, str1) == "YES" && calculateForString2(str1Map, str2) == "YES" {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func generateMap(str string) map[byte]byte {
	str1Map := make(map[byte]byte, 0)
	for p := 0; p < len(str); p++ {
		str1Map[str[p]] = str[p]
	}
	return str1Map
}

func calculateForString2(str1Map map[byte]byte, str2 string) string {
	for i := 0; i < len(str2); i++ {
		_, ok := str1Map[str2[i]]
		if !ok {
			return "NO"
		}
	}
	return "YES"
}

func calculateForString1(str2Map map[byte]byte, str1 string) string {
	for i := 0; i < len(str1); i++ {
		_, ok := str2Map[str1[i]]
		if !ok {
			return "NO"
		}
	}
	return "YES"
}
