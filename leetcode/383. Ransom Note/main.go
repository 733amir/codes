package main

func main() {
    println(canConstruct("a", "a") == true)
    println(canConstruct("aa", "ab") == false)
    println(canConstruct("aa", "aab") == true)
}

func canConstruct(ransomNote string, magazine string) bool {
    letters := [27]int{0}

    if len(ransomNote) > len(magazine) {
        return false
    }

    for i := range ransomNote {
        letters[ransomNote[i] - 'a'] += 1
    }

    for i := range magazine {
        letters[magazine[i] - 'a'] -= 1
    }

    for i := range letters {
        if letters[i] > 0 {
            return false
        }
    }

    return true
}
