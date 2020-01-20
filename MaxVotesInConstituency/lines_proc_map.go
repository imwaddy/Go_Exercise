package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

Max Votes In Constituency

In this problem, you have to modify an existing C++ Program, Java Program, C# Program, Go Program, or Python Program that reads in some text data, and does some processing on it.

The input is being read in from a file called input.txt, in this format:

22, Ravi Pawar, Aundh, 1603
23, Suvarna Kale, Baner, 803
27, Vinod Chavan, Aundh, 809
29, Vasant Mahajan, Aundh, 617
32, Aarti Patil, Baner, 351
34, Swaran Bijur, Baner, 352
Each line consists of 4 fields "Candidate ID," "Name," "Constituency," and "Votes." Here, "Candidate ID" and "Votes" are integers, while "Name" and "Constituency" are strings that do not contain commas or newlines.
22

Currently, the existing program reads the input line by line, then calls processData on this data, and prints the return value to the output. Unfortunately, processData currently does not do anything useful - it just returns an empty data-structure.

You have to modify processData to find the Candidate IDs of the candidate with highest votes in each constituency. Specifically, processData should return a map/dictionary where each key is the name of a constituency, and the value is the Candidate ID of the candidate in that constituency who has the highest votes.

If your program is run with the input given above, it should produce this output:

Aundh: 22
Baner: 23
This is because Candidate ID 22 (Ravi Pawar) has the highest votes in Aundh, and Candidate ID 23 (Suvarna Kale) has the hightest votes in Baner


Please make sure of the following:

Please download this C++ Program, Java Program, C# Program, Go Program, or Python Program, and change it to create your program, and submit your changed program as your answer. If you try to write your own program from scratch, it will take longer, and will probably not work.
All your changes must be inside processData. Do not make any changes to the rest of the program
Make sure that processData returns the correct value
If you're using C, please make sure to use only standard C. Do not use proprietary Microsoft or Turbo-C extensions. Specifically, do not use clrscr, getch, or conio.h.
Do not print anything extra to the output. Any unnecessary printf/println/putchar will result in a program disqualification.
Copy paste the entire updated program in the space below. You must submit a full, working, program. Please compile, run, and check the output of your program before submitting it.

*/

func processData(lines []string) (map[string]string, error) {
	// Modify this function to process `lines` as indicated
	// in the question. At the end, you need to return a map
	// containing the appropriate values.

	// Please create appropriate classes, and use appropriate
	// data structures as necessary.

	// Do not print anything in this function.

	// Submit this entire program (not just this function)
	// as your answer

	// final map
	m := make(map[string]string)

	// temporary map to store values
	tempMap := make(map[string]string)

	// loop through lines
	for _, value := range lines {

		// Splitting line into single value by `,`(comma)
		line := strings.Split(value, ",")

		// third value of slice willl be key of out code
		key := strings.TrimSpace(line[2])

		_, ok := m[key]

		if ok {
			//Tempmap also having same key
			val, _ := tempMap[key]

			var err error
			// Conversion for comparision
			a, err := strconv.Atoi(strings.TrimSpace(line[3]))
			b, err := strconv.Atoi(strings.TrimSpace(val))

			if err != nil {
				return m, err
			}

			if a > b {
				m[line[2]] = strings.TrimSpace(line[0])
			}
		} else {
			// Assign value vs key
			tempMap[key] = strings.TrimSpace(line[3])
			m[key] = strings.TrimSpace(line[0])

		}
	}

	return m, nil
}

func main() {
	// Open input file
	fin, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	// initialize scanner
	scanner := bufio.NewScanner(fin)

	var lines []string

	// Scan all lines and append as string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	ret_value, err := processData(lines)
	if err != nil {
		panic(err)
	}

	// Write output file
	fout, err := os.Create("output.txt")

	defer fout.Close()

	for k, v := range ret_value {
		_, err = fmt.Fprintln(fout, k, ": ", v)
	}
}
