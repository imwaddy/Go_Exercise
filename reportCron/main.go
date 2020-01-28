package main

import (
	"Go_Exercise/reportCron/cronjobs"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
)

var mainpackagevariable string

func main() {
	fmt.Println("Main Package variable=", mainpackagevariable)
	go InitCornJobs()
	// Wait for output only and do not press any key for a while
	// otherwise you will exit from main as aftre user input given
	// as if main exits goroutine will exit
	a := 0
	fmt.Scanf("%d", &a)
}

//InitCornJobs schedules functions for execution at a specified time
func InitCornJobs() {
	c := cron.New()
	fmt.Println("Cron Init at :", time.Now())
	c.AddFunc("1 * * * * *", cronjobs.CountEverythingService)

	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
