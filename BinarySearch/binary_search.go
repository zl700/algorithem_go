package BinarySearch

//BinarySearch1 is the recursive way.
//l is the start point of the arr,r is the end point of the arr.
func BinarySearch1(arr []int, target int, l, r int) int {
	m := (l + r) / 2
	if l > r {
		return -1
	}
	if arr[m] < target {
		return BinarySearch1(arr, target, m+1, r)
	} else if arr[m] > target {
		return BinarySearch1(arr, target, l, m-1)
	} else {
		return m
	}

}

//BinarySearch2 is the non-recursive way
func BinarySearch2(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for {
		m := (l + r) / 2
		if l > r {
			return -1
		}
		if arr[m] < target {
			l = m + 1
		} else if arr[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
}
