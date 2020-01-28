package cronjobs

import "fmt"

var otherpackagevariable string

func CountEverythingService() {
	fmt.Println("Other Package variable=", otherpackagevariable)
	return
}
