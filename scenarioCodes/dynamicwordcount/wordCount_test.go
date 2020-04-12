// Author: Mayur Wadekar
// Reading file from remote location and count the words in map[string]int{}
// URL to visit is http://www.gutenberg.org/files/15/text/

/*
	Write a program of wordcount reading files from remote location.
*/

package main

import (
	"fmt"
	"reflect"
	"regexp"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"main function"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestGetFileNames(t *testing.T) {
	tests := []struct {
		name          string
		wantFileNames []string
	}{
		{"Files Test", GetFileNames()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFileNames := GetFileNames(); !reflect.DeepEqual(gotFileNames, tt.wantFileNames) {
				t.Errorf("GetFileNames() = %v, want %v", gotFileNames, tt.wantFileNames)
			}
		})
	}
}

func TestGetFileData(t *testing.T) {
	type args struct {
		url         string
		wg          *sync.WaitGroup
		MapofValues map[string]int
		reg         *regexp.Regexp
	}
	var wg sync.WaitGroup
	wg.Add(1)
	MapofValues := map[string]int{}
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		fmt.Errorf("Error while regex compile")
		return
	}
	tests := []struct {
		name string
		args args
	}{
		{"URL 1", args{
			url:         "http://www.gutenberg.org/files/15/text/moby-001.txt",
			wg:          &wg,
			MapofValues: MapofValues,
			reg:         reg,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetFileData(tt.args.url, tt.args.wg, tt.args.MapofValues, tt.args.reg)
		})
	}
}
