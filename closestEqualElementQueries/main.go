func solveQueries(arr []int, queries []int) []int {
	hash := make(map[int][]int)
	lenArr := len(arr)
	for i := 0; i < lenArr; i++ {
		hash[arr[i]] = append(hash[arr[i]], i)
	}

	for i := 0; i < len(queries); i++ {
		bucket := hash[arr[queries[i]]]
		lenBucket := len(bucket)
		indexBucket := queries[i]
		// search index
		pos := -1
		l, r := 0, lenBucket-1
		if l == r {
            queries[i] = -1 
            continue
        }
        for l <= r {
			mid := l + (r-l)/2
           
			if bucket[mid] == indexBucket {
				
                pos = mid
				switch pos {
				case lenBucket-1:
					pos = min((bucket[pos] - bucket[pos-1]),(lenArr - bucket[pos]) + bucket[0])
				case 0:
					pos = min(bucket[pos+1] - bucket[pos], lenArr - bucket[lenBucket-1]+bucket[pos])
				default:
                    pos = min(bucket[pos+1] - bucket[pos], bucket[pos] - bucket[pos-1])
                }
				break
			}
            if bucket[mid] > indexBucket {
				r = mid-1
			} else {
				l = mid+1
			}
		}
		queries[i] = pos
		//


	}
	return queries
}
