package main

import (
	"fmt"
	"testing"
)

func TestGetLineCombinations(t *testing.T) {
  tests := []struct{
    line string
    expected int
  }{
    {
      line: "???.### 1,1,3",
      expected: 1,
    },
    {
      line: ".??..??...?##. 1,1,3",
      expected: 4,
    },
    {
      line: "?#?#?#?#?#?#?#? 1,3,1,6",
      expected: 1,
    },
    {
      line: "????.#...#... 4,1,1",
      expected: 1,
    },
    {
      line: "????.######..#####. 1,6,5",
      expected: 4,
    },
    {
      line: "?###???????? 3,2,1",
      expected: 10,
    },
    {
      line: "???.###????.###????.###????.###????.### 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3",
      expected: 1,
    },
    {
      line: ".??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##. 1,1,3,1,1,3,1,1,3,1,1,3,1,1,3",
      expected: 16384,
    },
    {
      line: "?#?#?#?#?#?#?#???#?#?#?#?#?#?#???#?#?#?#?#?#?#???#?#?#?#?#?#?#???#?#?#?#?#?#?#? 1,3,1,6,1,3,1,6,1,3,1,6,1,3,1,6,1,3,1,6",
      expected: 1,
    },
    {
      line: "????.#...#...?????.#...#...?????.#...#...?????.#...#...?????.#...#... 4,1,1,4,1,1,4,1,1,4,1,1,4,1,1",
      expected: 16,
    },
    {
      line: "????.######..#####.?????.######..#####.?????.######..#####.?????.######..#####.?????.######..#####. 1,6,5,1,6,5,1,6,5,1,6,5,1,6,5",
      expected: 2500,
    },
    {
      line: "?###??????????###??????????###??????????###??????????###???????? 3,2,1,3,2,1,3,2,1,3,2,1,3,2,1",
      expected: 506250,
    },
  }

  for i, test := range tests {
    t.Run(fmt.Sprintf("Line %d", i + 1), func(t *testing.T) {
      pattern, counts := GetPatternAndCounts(test.line)
      output := GetLineCombinations(pattern, counts)
      if output != test.expected {
        t.Errorf("Expected %v, received %v", test.expected, output)
      }
    })
  }
}

func TestParts(t *testing.T) {
  sample := `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`
  tests := map[string]struct{
    fn func(string)int
    expected int
  }{
    "Part One": {
      fn: PartOne,
      expected: 21,
    },
    "Part Two": {
      fn: PartTwo,
      expected: 525152,
    },
  }
  for name, test := range tests {
    t.Run(name, func(t *testing.T) {
      output := test.fn(sample)
      if output!= test.expected {
        t.Errorf("Expected %v, received %v", test.expected, output)
      }
    })
  }
}
