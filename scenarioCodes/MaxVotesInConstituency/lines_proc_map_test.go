package main

import (
	"reflect"
	"testing"
)

func Test_processData(t *testing.T) {
	type args struct {
		lines []string
	}
	output := map[string]string{}
	output["Aundh"] = "22"
	output["Baner"] = "23"
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{"Test_processData", args{lines: []string{"22, Ravi Pawar,Aundh, 1603", "27, Vinod Chavan, Aundh, 809", "29, Vasant Mahajan, Aundh, 617", "23, Suvarna Kale, Baner, 803", "32, Aarti Patil, Baner, 351", "34, Swaran Bijur, Baner, 352"}}, output, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processData(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("processData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processData() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
