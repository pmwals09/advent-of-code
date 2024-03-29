package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Range struct {
	Start int
	End   int
}

func (r Range) Contains(n int) bool {
	return n >= r.Start && n < r.End
}

func NewRange(startString string, lengthString string) Range {
	r := Range{}
	r.Start, _ = strconv.Atoi(startString)
	length, _ := strconv.Atoi(lengthString)
	r.End = r.Start + length
	return r
}

type MaterialRange struct {
	FromRange Range
	ToRange   Range
}

func NewMaterialRange(s string) MaterialRange {
	parts := strings.Split(s, " ")
	return MaterialRange{
		FromRange: NewRange(parts[1], parts[2]),
		ToRange:   NewRange(parts[0], parts[2]),
	}
}

func (m MaterialRange) Translate(n int) int {
	delta := n - m.FromRange.Start
	return m.ToRange.Start + delta
}

type MaterialMap struct {
	FromType string
	ToType   string
	Ranges   []MaterialRange
}

func NewMaterialMap(group string) MaterialMap {
	res := MaterialMap{}
	lines := strings.Split(group, "\n")
	for i, l := range lines {
		if i == 0 {
			front, _, _ := strings.Cut(l, " ")
			parts := strings.Split(front, "-")
			res.FromType = parts[0]
			res.ToType = parts[2]
		} else {
			res.Ranges = append(res.Ranges, NewMaterialRange(l))
		}
	}
	return res
}

func main() {
	input := getInput()
	fmt.Println("Part One:", PartOne(string(input)))
	fmt.Println("Part Two:", PartTwo(string(input)))
}

func getInput() []byte {
	ba, err := os.ReadFile("./day-05-input.txt")
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %v", err))
	}
	return ba
}

type PathState struct {
	CurrentPoint int
	CurrentStage string
}

func PartOne(input string) int {
	minLocation := math.MaxInt
	seeds, maps := NewGroups(input)
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for i := range seeds {
		wg.Add(1)
		go func(idx int) {
		Maps:
			for _, materialMap := range maps {
				for _, r := range materialMap.Ranges {
					if r.FromRange.Contains(seeds[idx]) {
						seeds[idx] = r.Translate(seeds[idx])
						continue Maps
					}
				}
			}
			ch <- seeds[idx]
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for s := range ch {
		if s < minLocation {
			minLocation = s
		}
	}

	return minLocation
}

func PartTwo(input string) int {
	minLocation := math.MaxInt
	seeds, maps := NewGroupRanges(input)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := range seeds {
		wg.Add(1)
		go func(batch []int) {
			for idx := range batch {
			Maps:
				for _, materialMap := range maps {
					for _, r := range materialMap.Ranges {
						if r.FromRange.Contains(batch[idx]) {
							batch[idx] = r.Translate(batch[idx])
							continue Maps
						}
					}
				}
				if batch[idx] < minLocation {
					mu.Lock()
					minLocation = batch[idx]
					mu.Unlock()
				}
			}
			wg.Done()
		}(seeds[i])
	}
	wg.Wait()

	return minLocation
}

func NewGroups(input string) ([]int, []MaterialMap) {
	paragraphs := strings.Split(input, "\n\n")
	seeds := []int{}
	maps := []MaterialMap{}
	for i, group := range paragraphs {
		if i == 0 {
			GroupToSeeds(group, &seeds)
		} else {
			maps = append(maps, NewMaterialMap(group))
		}
	}
	return seeds, maps
}

func GroupToSeeds(group string, seeds *[]int) {
	_, seedStr, _ := strings.Cut(group, ": ")
	seedStrs := strings.Split(seedStr, " ")
	for _, s := range seedStrs {
		n, _ := strconv.Atoi(s)
		*seeds = append(*seeds, n)
	}
}

func NewGroupRanges(input string) ([][]int, []MaterialMap) {
	paragraphs := strings.Split(input, "\n\n")
	seeds := [][]int{}
	maps := []MaterialMap{}
	for i, group := range paragraphs {
		if i == 0 {
			GroupToSeedRanges(group, &seeds)
		} else {
			n := NewMaterialMap(group)
			maps = append(maps, n)
		}
	}
	return seeds, maps
}

func GroupToSeedRanges(group string, seeds *[][]int) {
	_, seedStr, _ := strings.Cut(group, ": ")
	parts := strings.Split(seedStr, " ")
	for i := 0; i < len(parts); i += 2 {
		start, _ := strconv.Atoi(parts[i])
		length, _ := strconv.Atoi(parts[i+1])
		batch := []int{}
		for j := start; j < start+length; j++ {
			batch = append(batch, j)
		}
		*seeds = append(*seeds, batch)
	}
}
