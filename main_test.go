package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	//For setup and tearDown approch use something like below
	fmt.Println("Do all setup")
	m.Run()
	fmt.Println("Do the teardown")
}

func TestMap(t *testing.T) {
	slicedWord := Map("The quick brown fox jumps over the lazy dog")
	if len(slicedWord) != 9 {
		t.Log("Test failed, nos of words returned are incorrect")
		t.Fail()
	}
}

func TestMapForBlankSentence(t *testing.T) {
	words := Map("")
	if len(words) != 0 {
		t.Error("TestMapForBlankSentence failed")
	}
}
func TestReverse(t *testing.T) {
	t.Parallel()
	if reverse("Hello") != "olleH" {
		t.Fatal("Reverse is incorrect")
	}
}

func ExampleReverse() {
	fmt.Println(reverse("help"))
	//Output: pleh
}

func TestReverseForMultipleInput(t *testing.T) {
	if testing.Short() {
		t.Skip("TestReverseForMultipleInput skipped")
	}
	tests := []struct {
		input  string
		output string
	}{
		{"Hello", "olleH"},
		{"benchmarks", "skramhcneb"},
		{"provide", "edivorp"},
		{"flag", "galf"},
	}
	for _, test := range tests {
		if reverse(test.input) != test.output {
			t.Fatal("TestReverseForMultipleInput failed")
		}
	}
}

// go test -run=TestReverseForMultipleInputWithSubTest
// or to run a single test
// go test -run=TestReverseForMultipleInputWithSubTest/input_hello -v
func TestReverseForMultipleInputWithSubTest(t *testing.T) {
	if testing.Short() {
		t.Skip("TestReverseForMultipleInputWithSubTest skipped")
	}
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"input hello", "Hello", "olleH"},
		{"input benchmarks", "benchmarks", "skramhcneb"},
		{"input provide", "provide", "edivorp"},
		{"input flag", "flag", "galf"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if reverse(test.input) != test.output {
				t.Fatal("TestReverseForMultipleInputWithSubTest failed")
			}
		})
	}
}

func TestProcessForHello(t *testing.T) {
	words := []string{"Hello"}
	reversedWords := process(words)
	if len(reversedWords) != 1 {
		t.Fatal("ProcessForHello failed to return expected word count")
	}
	if reversedWords[0] != "olleH" {
		t.Fatal("ProcessForHello failed to reverse word")
	}
}

func TestProcess(t *testing.T) {
	words := []string{"There", "is", "no", "flag", "you", "can", "provide", ",",
		"that", "will", "run", "only", "benchmarks"}
	reversedWords := process(words)
	t.Log(reversedWords)
}

// http://stackoverflow.com/questions/16161142/how-to-test-only-one-benchmark-function
//  go test -bench=Map$ -run=^$
func BenchmarkMap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Map("There is no flag you can provide, that will run only benchmarks")
	}
}

// https://github.com/bradfitz/talk-yapc-asia-2015/blob/master/talk.md
func BenchmarkReverse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		reverse("Reverse")
	}
}

//for block profile
func BenchmarkProcess(b *testing.B) {
	b.StopTimer()
	words := []string{"There", "is", "no", "flag", "you", "can", "provide", ",",
		"that", "will", "run", "only", "benchmarks"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		process(words)
	}
}

func BenchmarkNonsense(b *testing.B) {
	if testing.Verbose() {
		b.Skip("BenchmarkNonsense skipped")
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		nonsense()
	}
}

func BenchmarkReduce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		words := []string{"olleH", "skramhcneb"}
		b.StopTimer()
		reduce(words)
	}

	//Contention benchmark
	/*
		words := []string{"olleH", "skramhcneb" }
		b.SetParallelism(30)
		b.RunParallel(func (pb *testing.PB){
			for pb.Next(){
				reduce(words)
			}
		})
	*/
}
