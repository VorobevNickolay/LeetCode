package LeetCode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := range nums1 {
			nums1[i] = nums2[i]
		}
	}
	nums3 := nums1[:m]
	n--
	m--
	pointer1 := len(nums1) - 1
	for pointer1 >= 0 {
		if n < 0 {
			nums1[pointer1] = nums3[m]
			m--
		} else if m < 0 {
			nums1[pointer1] = nums2[n]
			n--
		} else if nums3[m] > nums2[n] {
			nums1[pointer1] = nums3[m]
			m--
		} else {
			nums1[pointer1] = nums2[n]
			n--
		}
		pointer1--
	}
}
