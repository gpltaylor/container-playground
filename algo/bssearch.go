package main

func binarySearch(arr []int, needle int) bool {
	low := 0
	hi := len(arr)

	mid := hi / 2

	if arr[mid] == needle {
		return true
	}

	if arr[mid] < needle {
		low = mid + 1
	} else {
		hi = mid
	}

	return binarySearch(arr[low:hi], needle)
}
