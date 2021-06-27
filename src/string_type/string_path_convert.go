package string_type

import sets "github.com/deckarep/golang-set"
import "github.com/ahrtr/gocontainer/list"

func getNexts(words []string) map[string][]string {
	dict := sets.NewSet(words)
	nexts := make(map[string][]string)

	for _, w := range words {
		nexts[w] = getNext(w, dict)
	}
	return nexts
}

func getNext(word string, dict sets.Set) (res []string) {
	chs := []byte(word)
	for cur := byte('a'); cur < 'z'; cur++ {
		for i := 0; i < len(chs); i++ {
			if chs[i] != cur {
				tmp := chs[i]
				chs[i] = cur
				if dict.Contains(string(chs)) {
					res = append(res, string(chs))
				}
				chs[i] = tmp
			}
		}
	}
	return
}

func getDistance(start string, nexts map[string][]string) map[string]int {
	var distance = make(map[string]int)
	distance[start] = 0
	var queue = list.NewLinkedList()
	var sets = sets.NewSet()
	sets.Add(start)
	for !queue.IsEmpty() {
		v, _ := queue.Get(0)
		cur := v.(string)
		for _, str := range nexts[cur] {
			if !sets.Contains(str) {
				distance[str] = distance[cur] + 1
				queue.Add(str)
				sets.Add(str)
			}
		}
	}
	return distance
}

func getShortestPaths(cur, to string, nexts map[string][]string, distance map[string]int, solution []string,
	res [][]string) {
	solution = append(solution, cur)
	if to == cur {
		res = append(res, solution)
	} else {
		for _, next := range nexts[cur] {
			if distance[next] == distance[cur]+1 {
				getShortestPaths(next, to, nexts, distance, solution, res)
			}
		}
	}
	solution = solution[:len(solution)-1]
}

func findMinPaths(start, to string, lists []string) [][]string {
	lists = append(lists, start)
	nexts := getNexts(lists)
	distance := getDistance(start, nexts)
	var pathList []string
	var res [][]string
	getShortestPaths(start, to, nexts, distance, pathList, res)
	return res
}
