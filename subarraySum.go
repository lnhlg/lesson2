package lesson2

//SubarraySum: 统计和为指定目标值的子数组数量
/*parameter
nums: 整型数组
target: 目标值
return: 符合条件的子数组数量
*/
func SubarraySum(nums []int, target int) (ans int) {

	mp, colNum := map[int]int{0: 1}, len(nums)
	for i, pre := 0, 0; i < colNum; i++ {

		pre += nums[i]
		if _, ok := mp[pre-target]; ok {

			ans += mp[pre-target]
		}

		mp[pre]++
	}

	return
}
