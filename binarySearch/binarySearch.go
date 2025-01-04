package binarySearch

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (high + low) / 2
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
