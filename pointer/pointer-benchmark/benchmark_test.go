package main

import "testing"

type SmallStruct struct {
	A, B, C int
}

type LargeStruct struct {
	Data [10_000_000]int
}

func BenchmarkSmallStructValue(b *testing.B) {
	s := SmallStruct{1, 2, 3}
	for i := 0; i < b.N; i++ {
		ProcessSmallValue(s)
	}
}

func BenchmarkSmallStructPointer(b *testing.B) {
	s := SmallStruct{1, 2, 3}
	for i := 0; i < b.N; i++ {
		ProcessSmallPointer(&s)
	}
}

func BenchmarkLargeStructValue(b *testing.B) {
	l := LargeStruct{}
	for i := 0; i < b.N; i++ {
		ProcessLargeValue(l)
	}
}

func BenchmarkLargeStructPointer(b *testing.B) {
	l := LargeStruct{}
	for i := 0; i < b.N; i++ {
		ProcessLargePointer(&l)
	}
}

func ProcessSmallValue(s SmallStruct) int {
	return s.A + s.B + s.C
}

func ProcessSmallPointer(s *SmallStruct) int {
	return s.A + s.B + s.C
}

func ProcessLargeValue(l LargeStruct) int {
	var sum int
	for _, v := range l.Data {
		sum += v
	}
	return sum
}

func ProcessLargePointer(l *LargeStruct) int {
	var sum int
	for _, v := range l.Data {
		sum += v
	}
	return sum
}
