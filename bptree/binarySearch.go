package bptree

// 传入的keys要求有序
func BinarySearch(keys []int, searchKey int) (int, int) {
	var count int
	var low, mid int
	var high = len(keys) - 1
	//mid := 0
	for {
		count++
		if low > high {
			break
		}
		mid = (low + high) / 2
		if keys[mid] == searchKey {
			return keys[mid], count
		} else if keys[mid] > searchKey {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, -1
}

//func main() {
//	keys := []int{10, 11, 12, 16, 20, 21, 22, 40, 50, 51, 52, 55, 56, 57, 60, 69, 77, 78, 79, 80, 85, 86, 87}
//	res, count := BinarySearch(keys, 60)
//	fmt.Println("res:", res, "count:", count)
//}
