package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wc := newWordCount()

	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			wordCounter(wc, filename)
			wg.Done()
		}(f)
	}
	wg.Wait()

	wc.Lock()
	for w, c := range wc.words {
		if c > 1 {
			fmt.Printf("Word %s occured %d time\n", w, c)
		}
	}
	wc.Unlock()
}

type WordCount struct {
	sync.Mutex
	words map[string]int
}

func newWordCount() *WordCount {
	return &WordCount{words: map[string]int{}}
}

func (wc *WordCount) Add(word string, n int) {
	wc.Lock()
	defer wc.Unlock()

	count, ok := wc.words[word]
	if !ok {
		wc.words[word] = n
		return
	}
	wc.words[word] = count + n
}

func wordCounter(wc *WordCount, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file %s: %s\n", filename, err)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		word := strings.ToLower(sc.Text())
		wc.Add(word, 1)
	}
}
