package lesson2

//FindShortestSubArray: 查找与指定数组拥有同样度的最短子数组
/*parameter
num: 整型数组
return: 符合条件的子数组的长度
*/
func FindShortestSubArray(nums []int) int {

	cntMap := make(map[int][]int)
	for i := range nums {

		if _, has := cntMap[nums[i]]; has {

			cntMap[nums[i]][0]++
			cntMap[nums[i]][2] = i
		} else {

			cntMap[nums[i]] = make([]int, 3)
			cntMap[nums[i]][0] = 1
			cntMap[nums[i]][1] = i
			cntMap[nums[i]][2] = i
		}
	}

	max, len := 0, 0
	for k := range cntMap {

		if cntMap[k][0] > max {
			max, len = cntMap[k][0], cntMap[k][2]-cntMap[k][1]+1
		} else if cntMap[k][0] == max {
			len = min(len, cntMap[k][2]-cntMap[k][1]+1)
		}
	}

	return len
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
