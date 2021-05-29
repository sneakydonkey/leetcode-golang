func reverse(x int) int {
    num := 0
    var sign int
    if x >= 0 {
        sign = 1
    } else {
        sign = -1
        x = -x
    }
    for x != 0 {
        mod := x % 10
        x = x / 10
        num = num * 10 + mod
    }
    num = num * sign
    if num > math.MaxInt32 || num < math.MinInt32 {
        return 0
    }
    return num   
}
