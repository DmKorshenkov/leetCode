package main

import "fmt"

func main() {
	groupThePeople([]int{3, 3, 3, 3, 1, 3, 3, 1, 2, 2})

}

func groupThePeople(groupSizes []int) [][]int {
	mp := make(map[int][]int)
	//	fmt.Println(groupSizes)

	for i := range groupSizes {
		mp[groupSizes[i]] = append(mp[groupSizes[i]], i)
	}
	//	fmt.Println(mp)
	res := make([][]int, 0, len(mp))

	for key := range mp {
		fmt.Println(key, mp[key])
		tmp := make([]int, 0, key)
		for i := 0; i < len(mp[key]); i++ {
			tmp = append(tmp, mp[key][i])
			if len(tmp) == key {
				res = append(res, tmp)
				//				tmp = nil
				tmp = make([]int, 0, key)
			}
		}
	}

	fmt.Println(res)
	return nil
}
