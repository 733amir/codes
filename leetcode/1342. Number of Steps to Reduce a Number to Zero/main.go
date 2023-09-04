package main

func main() {
    println(numberOfSteps(14) == 6)
    println(numberOfSteps(8) == 4)
    println(numberOfSteps(123) == 12)
}

func numberOfSteps(num int) int {
    count := 0
    for {
        count += num & 1

        num >>= 1
        if num == 0 {
            break
        }
        count += 1
    }
    return count
}

