package _001__Two_Sum

func twoSum(nums []int, target int) []int {
	// index是用来保存map[值]下标的序列号
	var indexs = make(map[int]int, len(nums))

	for i, v := range nums {
		// 计算另一个值
		another := target - v
		// 如果查询序列里有记录那个值
		if _, ok := indexs[another]; ok {
			return []int{indexs[another], i}
		}
		// 否则设置当前值的下标
		indexs[v] = i
	}
	return nil
}