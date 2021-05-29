func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x % 10 == 0 && x != 0 {
        return false
    }
    var appendedNum = 0
    for x > appendedNum {
        remaining := x % 10
        x = x / 10
        appendedNum = appendedNum * 10 + remaining
    }
    return x == appendedNum || x == appendedNum / 10
    
}
