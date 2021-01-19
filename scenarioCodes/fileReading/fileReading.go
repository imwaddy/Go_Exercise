// You have to read a file and each row of file contains the number. We have to read a file concurrently and have to append the sum of numbers into the file. If first file contains the filename then also calculate the sum and add with the first file ex.

// file.txt
// 10
// 20
// file2.txt
// 30

// file2.txt
// 10

// output
// file.txt
// 10
// 20
// 30
// 70

// Sum of file.txt is 60 and file2.txt is 10 total 70 should write into file.txt

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var sum int
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go readText(readFile("file.txt"), &wg, &mutex)
	wg.Wait()
	fmt.Println("Final Sum: ", sum)
}

func readFile(filename string) (file *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error while reading file ", err)
	}
	return file
}

func readText(file *os.File, wg *sync.WaitGroup, mutex *sync.Mutex) {
	sc := bufio.NewScanner(file)
	defer wg.Done()
	for sc.Scan() {
		var text = sc.Text()
		i, e := strconv.Atoi(text)
		if e != nil && strings.Contains(text, "txt") {
			wg.Add(1)
			go readText(readFile(text), wg, mutex)
		} else {
			mutex.Lock()
			sum += i
			mutex.Unlock()
		}
	}
}
