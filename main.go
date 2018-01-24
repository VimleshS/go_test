package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

func main() {
	sentence := "The quick brown fox jumps over the lazy dog"
	words := Map(sentence)
	rwords := process(words)
	fmt.Println(reduce(rwords))
}

//Map splits sentences into words
//var r =  regexp.MustCompile("[^\\s]+")
func Map(sentence string) []string {
	//return strings.Split(sentence, " ")
	// return strings.Fields(sentence)
	r := regexp.MustCompile("[^\\s]+")
	return r.FindAllString(sentence, -1)
}

func reduce(reverseWords []string) string {
	return strings.Join(reverseWords, " ")
}

var counter int

func process(words []string) []string {
	nosOfWords := len(words)
	buffChannel := make(chan string, nosOfWords)
	//task := new(sync.WaitGroup)
	var task sync.WaitGroup
	task.Add(nosOfWords)

	for _, word := range words {
		go func(word string) {
			defer task.Done()
			buffChannel <- reverse(word)
			counter++
			//for datarace
			//go test -run=Process$ -v -race
			//fmt.Printf("-->Nos of words reversed %d \n", counter)

		}(word)
	}
	task.Wait()
	close(buffChannel)

	rwords := make([]string, 0)
	for rword := range buffChannel {
		rwords = append(rwords, rword)
	}
	return rwords
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func reverse(word string) string {
	//case 1
	// reversedWord := ""
	// for index := len(word) - 1; index >= 0; index-- {
	// 	reversedWord += string(word[index])
	// }
	// return reversedWord

	//case 2
	// var buff bytes.Buffer
	// for index := len(word) - 1; index >= 0; index-- {
	// 	buff.WriteString(string(word[index]))
	// }
	// return buff.String()

	//case 3
	//Not recomended to used in production.
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()
	for index := len(word) - 1; index >= 0; index-- {
		buf.WriteString(string(word[index]))
	}
	return buf.String()
}

type Emp struct {
	name string
}

func nonsense() {
	var emps = []Emp{}
	for i := 0; i < 100; i++ {
		emps = append(emps, Emp{"Me"})
	}
}
