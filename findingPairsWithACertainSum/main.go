type FindSumPairs struct {
	arrNums1 []int
	arrNums2 []int
	mpNums2  map[int][]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
//	fmt.Println("!")
	mp := make(map[int][]int)
	//
	for i := range nums2 {
		if _, ok := mp[nums2[i]]; ok {
			mp[nums2[i]] = append(mp[nums2[i]], i)
		} else {
			tmp := make([]int, 0, 10)
			tmp = append(tmp, i)
			mp[nums2[i]] = tmp
		}
	}
	//
	return FindSumPairs{arrNums1: nums1, arrNums2: nums2, mpNums2: mp}
}

func (this *FindSumPairs) Add(index int, val int) {
	num := this.arrNums2[index]
    this.arrNums2[index] += val
	//
	for i := 0; i < len(this.mpNums2[num]); i++ {
		//
		if this.mpNums2[num][i] == index {
			//
			if _, ok := this.mpNums2[num+val]; ok {
				this.mpNums2[num+val] = append(this.mpNums2[num+val], index)

			} else {
				//
				this.mpNums2[num+val] = make([]int, 0, 10)
                this.mpNums2[num+val] = append(this.mpNums2[num+val], index) 
				//
			}
            //
			this.mpNums2[num] = removeAtClear(this.mpNums2[num], index)
			if len(this.mpNums2[num]) == 0 || this.mpNums2[num] == nil {
				delete(this.mpNums2, num)
			}
            //
		}
	}
}
func removeAtClear(arr []int, i int) []int {
	if len(arr) == 1 {
		return nil
	}
	for j := range arr {
		if arr[j] == i {
			i = j
		}
	}
    //
	copy(arr[i:], arr[i+1:])
	arr[len(arr)-1] = 0
	return arr[:len(arr)-1]
}

func (this *FindSumPairs) Count(tot int) int {
	var count int

	for i := range this.arrNums1 {

		if arr, ok := this.mpNums2[tot-this.arrNums1[i]]; ok {
			count += len(arr)
		}
	}
	return count
}
