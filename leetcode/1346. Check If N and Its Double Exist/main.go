package main

func main() {
}

func checkIfExist(arr []int) bool {
	var x2, d2 int
	var d bool
	for i := range arr {
		x2 = arr[i] << 1
		d2 = arr[i] >> 1
		d = (arr[i] & 1) != 1
		for j := i + 1; j < len(arr); j++ {
			if x2 == arr[j] || (d && d2 == arr[j]) {
				return true
			}
		}
	}
	return false
}
