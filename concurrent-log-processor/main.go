package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type LogDetails struct {
	ServiceName             string    `json:"service_name"`
	TotalRequests           int       `json:"total_requests"`
	TotalResponseTimeInMs   float64   `json:"total_response_time_ms"`
	AverageResponseTimeInMs float64   `json:"average_response_time_ms"`
	MaximumResponseTimeInMs float64   `json:"maximum_response_time_ms"`
	P95Latency              float64   `json:"p95_latency"`
	ResponseTimesArray      []float64 `json:"response_times_array"`
}

func main() {
	// var rows int
	// fmt.Printf("Enter no of log rows:%d", &rows)
	// var reader bytes.Buffer
	// reader.ReadFrom(r)

	var logString = `100 auth 120
	101 payment 300
	102 auth 150
	103 auth 200
	104 payment 250
	105 inventory 80
	106 auth 110
	107 payment 400
	108 inventory 90
	109 inventory 100
	110 payment 350
	111 auth 170`

	logs := strings.Split(logString, "\n")

	response := make(map[string]LogDetails, 0)

	for _, log := range logs {
		row := strings.Split(log, " ")
		rTime, _ := strconv.ParseFloat(row[2], 64)
		val, ok := response[row[1]]
		if ok {
			val.TotalResponseTimeInMs = getTotalResponseTime(val.TotalResponseTimeInMs, row[2])
			val.MaximumResponseTimeInMs = getMaxResponseTime(val.MaximumResponseTimeInMs, row[2])
			val.TotalRequests += 1
			val.ResponseTimesArray = append(val.ResponseTimesArray, rTime)
			// val.ResponseTimesArray =

			response[row[1]] = val
		} else {
			object := LogDetails{
				ServiceName:             row[1],
				TotalRequests:           1,
				TotalResponseTimeInMs:   rTime,
				AverageResponseTimeInMs: 0,
				MaximumResponseTimeInMs: rTime,
				P95Latency:              0,
				ResponseTimesArray:      append(val.ResponseTimesArray, rTime),
			}
			response[row[1]] = object
		}
	}

	for _, log := range response {
		sort.Float64s(log.ResponseTimesArray)
		var p95 = math.Ceil(0.95*float64(log.TotalRequests)) - 1
		fmt.Printf("%s %d %.2f %.2f %.2f\n", log.ServiceName, log.TotalRequests, getAverageResponseTime(log), log.MaximumResponseTimeInMs, log.ResponseTimesArray[int(p95)])
	}
}

func getTotalResponseTime(RespTime float64, respTime string) float64 {
	rTime, _ := strconv.ParseFloat(respTime, 64)
	return RespTime + rTime
}

func getMaxResponseTime(maxRespTime float64, val string) float64 {
	newRespTime, _ := strconv.ParseFloat(val, 64)
	if maxRespTime < newRespTime {
		return newRespTime
	}
	return maxRespTime
}

func getAverageResponseTime(log LogDetails) float64 {
	return log.TotalResponseTimeInMs / float64(log.TotalRequests)
}
