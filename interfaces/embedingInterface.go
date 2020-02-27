package main

import "fmt"

// AttendanceCalculator ...
type AttendanceCalculator interface {
	Calculate()
}

// DisplayAttendance ...
type DisplayAttendance interface {
	DisplayAttendancePercentage()
}

// StudentOperations ...
type StudentOperations interface {
	AttendanceCalculator
	DisplayAttendance
}

// Student ...
type Student struct {
	Name                 string
	AttendanceDays       int
	AttendancePercentage float32
}

// DisplayAttendancePercentage ...
func (s *Student) DisplayAttendancePercentage() {
	fmt.Println(s.Name, " has", s.AttendancePercentage, "% attendance")
}

// Calculate ...
func (s *Student) Calculate() {
	// forget about formula it uses just for purpose
	s.AttendancePercentage = float32(s.AttendanceDays*30) / float32(100)
}

func main() {
	s := Student{
		Name:           "Mayur",
		AttendanceDays: 24,
	}
	s.Calculate()
	s.DisplayAttendancePercentage()
}
