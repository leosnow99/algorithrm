package dp

import (
	"fmt"
	"sort"
)

type EvenLoop struct {
	Len int
	Wid int
}

func NewEvenLoop(len, wid int) *EvenLoop {
	return &EvenLoop{
		Len: len,
		Wid: wid,
	}
}

type EvenLoopSlice []*EvenLoop

func (s EvenLoopSlice) Len() int {
	return len(s)
}

func (s EvenLoopSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less 以信封的长度作为排序的base
func (s EvenLoopSlice) Less(i, j int) bool {
	if s[i].Len == s[j].Len {
		return s[i].Wid < s[j].Wid
	}
	return s[i].Len < s[j].Len
}

func getSortedEnvelopes(matrix [][]int) []*EvenLoop {
	results := make([]*EvenLoop, len(matrix))
	for i := 0; i < len(matrix); i++ {
		results[i] = NewEvenLoop(matrix[i][0], matrix[i][1])
	}

	sort.Sort(EvenLoopSlice(results))
	return results
}

func TestSortedEnvelopes() {
	envelopes := getSortedEnvelopes([][]int{{4, 4}, {8, 6}, {2, 2}})
	for _, envelope := range envelopes {
		fmt.Printf("result: %+v \n", envelope)
	}
}

func maxEnvelopes(matrix [][]int) int {
	envelopes := getSortedEnvelopes(matrix)
	ends := make([]int, len(envelopes))
	ends[0] = envelopes[0].Wid
	right, l, r, m := 0, 0, 0, 0

	for i := 1; i < len(envelopes); i++ {
		l = 0
		for r = right; l < r; {
			m = l + (r-l)>>1
			if ends[m] < envelopes[m].Wid {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		if right < l {
			right = l
		}
		ends[l] = envelopes[i].Wid
	}
	return right + 1
}

