package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func timeit(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s duration: %+v\n", name, elapsed)
}

func part1(fd *os.File) (result int, err error) {
	defer timeit(time.Now(), "part1")
	previousValue := 0
	first := true
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		check(err)
		if value > previousValue && !first {
			result++
		}
		previousValue = value
		first = false
	}
	err = scanner.Err()
	check(err)
	return result, err
}

func part2(fd *os.File) (result int, err error) {
	defer timeit(time.Now(), "part2")
	previous := 0
	isFirst := true
	done := false
	queue := list.New()
	scanner := bufio.NewScanner(fd)
	for !done {
		// read values from 3 lines into queue
		for queue.Len() < 3 {
			done = !scanner.Scan()
			if done {
				break
			}
			val, err := strconv.Atoi(scanner.Text())
			check(err)
			queue.PushBack(val)
		}
		err = scanner.Err()
		check(err)
		// sum all queue items and drop the first one
		value := 0
		for e := queue.Front(); e != nil; e = e.Next() {
			value += e.Value.(int)
		}
		queue.Remove(queue.Front())
		if value > previous && !isFirst {
			result++
		}
		previous = value
		isFirst = false
	}
	return result, err
}

func main() {
	defer timeit(time.Now(), "main")
	if len(os.Args) != 2 {
		fmt.Println("please provide a filename argument")
		os.Exit(1)
	}
	filename := os.Args[1]

	fd, err := os.Open(filename)
	defer fd.Close()
	check(err)

	result1, err := part1(fd)
	fmt.Printf("part1 result: %+v\n", result1)

	fd.Seek(0, io.SeekStart)

	result2, err := part2(fd)
	fmt.Printf("part2 result: %+v\n", result2)
}
