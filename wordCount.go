// Author: Mayur Wadekar
// Reading file from remote location and count the words in map[string]int{}
// URL to visit is http://www.gutenberg.org/files/15/text/

/*
	Write a program of wordcount reading files from remote location.
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {

	// For getting file names with URL
	filenames := GetFileNames()

	// initialize regex to remove special character in string
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Errorf("Error while regex compile")
		return
	}

	// map of word count
	MapofValues := map[string]int{}

	// waitgroup sync between goroutines
	var wg sync.WaitGroup

	for _, url := range filenames {
		wg.Add(1)
		go GetFileData(url, &wg, MapofValues, reg)
	}
	wg.Wait()

	fmt.Println(MapofValues)

}

// GetFileNames - returns array of string which contains filenames on remote location
func GetFileNames() (fileNames []string) {
	/* 134 static files so there is hardcoded count is 134
	File pattern is same as
		1.single digit having 00*
		2.double digit having 0**
		3.triple digit having ***
	*/
	for test := 1; test <= 134; test++ {
		var pp string
		// To convert filename in pattern
		if len(strconv.Itoa(test)) == 1 {
			pp = "00" + strconv.Itoa(test)
		} else if len(strconv.Itoa(test)) == 2 {
			pp = "0" + strconv.Itoa(test)
		} else {
			pp = strconv.Itoa(test)
		}
		// append whole filename in filenames slice
		fileNames = append(fileNames, "http://www.gutenberg.org/files/15/text/moby-"+pp+".txt")
	}
	return
}

// GetFileData - This function getting data from remote location and write wordcount on map
func GetFileData(url string, wg *sync.WaitGroup, MapofValues map[string]int, reg *regexp.Regexp) {

	/*
		This get request might giving error if server hitted multiple times.
		Connection might be refused from http://www.gutenberg.org/files/15/text/
		So if you want it to check you can write code to read file from local servers.
	*/
	//rest call
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("Error while reading file from remote ")
		return
	}

	defer resp.Body.Close()

	// reading body of response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Read body: %v", err)
		return
	}

	// initialize bufio scanner to read string
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	// scanner split will use scanwords function which splits string into words
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// remove special character from word
		w := reg.ReplaceAllString(scanner.Text(), "")
		InsertEntryInMap(w, MapofValues)
	}

	// If err while reading scanner
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Read body: %v", err)
		return
	}

	wg.Done()
}

// InsertEntryInMap - entry for each word added to a map
func InsertEntryInMap(word string, MapofValues map[string]int) {
	// added mutex
	var mutex sync.Mutex
	//Lock map
	mutex.Lock()
	// check if key present in map or not
	value, ok := MapofValues[word]
	// if present then count +1
	if ok {
		MapofValues[word] = value + 1
	} else {
		// else assign new key and value in map
		MapofValues[word] = 1
	}
	mutex.Unlock()
}
