package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type LogDetails struct {
	ServiceName             string
	TotalRequests           int
	TotalResponseTimeInMs   float64
	MaximumResponseTimeInMs float64
	ResponseTimesArray      []float64
}

type parsedLog struct {
	serviceName  string
	responseTime float64
}

func parseAndSend(line string, ch chan<- parsedLog, wg *sync.WaitGroup) {
	defer wg.Done()
	row := strings.Split(line, " ")
	if len(row) < 3 {
		return
	}
	responseTime, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return
	}
	ch <- parsedLog{serviceName: row[1], responseTime: responseTime}
}

func main() {
	fmt.Println("Enter the logs (Press Ctrl+D to end input):")
	scanner := bufio.NewScanner(os.Stdin)

	ch := make(chan parsedLog)
	var wg sync.WaitGroup

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		wg.Add(1)
		go parseAndSend(line, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	response := make(map[string]*LogDetails)
	for p := range ch {
		val, ok := response[p.serviceName]
		if ok {
			val.TotalResponseTimeInMs += p.responseTime
			if p.responseTime > val.MaximumResponseTimeInMs {
				val.MaximumResponseTimeInMs = p.responseTime
			}
			val.TotalRequests++
			val.ResponseTimesArray = append(val.ResponseTimesArray, p.responseTime)
		} else {
			response[p.serviceName] = &LogDetails{
				ServiceName:             p.serviceName,
				TotalRequests:           1,
				TotalResponseTimeInMs:   p.responseTime,
				MaximumResponseTimeInMs: p.responseTime,
				ResponseTimesArray:      []float64{p.responseTime},
			}
		}
	}

	fmt.Printf("%-12s %10s %12s %10s %10s\n", "SERVICE", "REQUESTS", "AVG(ms)", "MAX(ms)", "P95(ms)")
	fmt.Println(strings.Repeat("-", 58))

	for _, log := range response {
		sort.Float64s(log.ResponseTimesArray)
		p95idx := int(math.Ceil(0.95*float64(log.TotalRequests))) - 1
		avg := log.TotalResponseTimeInMs / float64(log.TotalRequests)
		fmt.Printf("%-12s %10d %12.2f %10.2f %10.2f\n",
			log.ServiceName, log.TotalRequests, avg,
			log.MaximumResponseTimeInMs, log.ResponseTimesArray[p95idx])
	}
}
