
package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftProducts := make([]int, n)
	rightProducts := make([]int, n)
	product := 1
	for i := 0; i < n; i++ {
		product *= nums[i]
		leftProducts[i] = product
	}
	product = 1
	for i := n - 1; i >= 0; i-- {
		product *= nums[i]
		rightProducts[i] = product
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = 1
		if i-1 >= 0 {
			result[i] *= leftProducts[i-1]
		}
		if i+1 < n {
			result[i] *= rightProducts[i+1]
		}
	}
	return result
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}
	return result
}
