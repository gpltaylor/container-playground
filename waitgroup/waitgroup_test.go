package main

/*
	WaitGroup is a struct that allows to wait for a collection of goroutines to finish.
	https://gobyexample.com/waitgroups
*/
import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func TestWorker_JAGBE01(t *testing.T) {
	var wg sync.WaitGroup
	i := 0

	for ; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

	wg.Wait()
	require.Equal(t, i, 5)
}

func TestWithoutWoker_JAGBE02(t *testing.T) {
	i := 0
	for ; i < 5; i++ {
		worker(i)
	}

	require.Equal(t, i, 5)
}

func TestWithoutWoker_JAGBE03(t *testing.T) {
	i := 0
	for ; i < 5; i++ {
		go worker(i)
	}

	require.Equal(t, i, 5)
}
