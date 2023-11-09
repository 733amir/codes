package main

func main() {

}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
        return false
    }

    inc := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}

		if inc {
			if arr[i] < arr[i+1] {
				continue
			}
			inc = false
            if i == 0 {
                return false
            }
		}

		if arr[i] < arr[i+1] {
			return false
		}
	}
	return !inc
}
