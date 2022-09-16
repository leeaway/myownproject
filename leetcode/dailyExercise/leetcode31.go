package dailyExercise

func nextPermutation(nums []int)  {
	n := len(nums)
	if n==1 {
		return
	}
	index := n-2
	for index >= 0 && nums[index]>=nums[index+1]{
		index--
	}
	if index >= 0 {
		r := n-1
		for nums[r] <= nums[index] {
			r--
		}
		nums[index],nums[r] = nums[r],nums[index]
	}
	reverse(nums[index+1:])
}

func reverse(arr []int) {
	for i,j := 0,len(arr)-1;i<j;i,j = i+1,j-1 {
		arr[i],arr[j] = arr[j],arr[i]
	}
}
