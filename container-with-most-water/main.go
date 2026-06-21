package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type WaterHeights struct {
	ID      int   `json:"id"`
	Heights []int `json:"heights"`
}

func LoadTasks(path string) ([]WaterHeights, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error while reading file %+v", err)
		return []WaterHeights{}, err
	}

	var heights []WaterHeights
	err = json.Unmarshal(data, &heights)
	if err != nil {
		log.Printf("Error while unmarshaling %+v", err)
		return []WaterHeights{}, err
	}
	return heights, err
}

func maxWater(heights []int) int {
	i := 0
	j := len(heights) - 1
	var area int
	for i < j {
		calculatedArea := min(heights[i], heights[j]) * (j - i)
		if calculatedArea > area {
			area = calculatedArea
		}
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return area
}

func main() {
	waterHeights, err := LoadTasks("input/input.json")
	if err != nil {
		return
	}

	for _, height := range waterHeights {
		fmt.Println("Calculated Area for ID [", height.ID, "] is:", maxWater(height.Heights))
	}
}
