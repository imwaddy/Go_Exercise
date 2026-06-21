package main

import "fmt"

func timeConversion() {

	s := "07:10:40AM"

	fmt.Println("Date ", getFormat(s))
}

func getFormat(s string) string {

	hour := int(s[0]-'0')*10 + int(s[1]-'0')

	if s[8] == 'P' {
		if hour != 12 {
			hour += 12
		}
	} else {
		if hour == 12 {
			hour = 0
		}
	}
	return fmt.Sprintf("%02d%s", hour, s[2:8])
}
