/*
  https://leetcode.com/problems/powx-n/
*/

func myPow(x float64, n int) float64 {
    var mmap = make(map[int]float64)
    return runMyPow(x, n, &mmap)
}

func runMyPow(x float64, n int, mmap *map[int]float64) float64 {
    if n == 1 {
        return x
    }
    if n == 0 {
        return 1
    }
    if (n == -1) {
        return 1.0 / x
    }
    if val, ok := (*mmap)[n]; ok {
        return val
    }
    res := runMyPow(x, n / 2, mmap) * runMyPow(x, n / 2, mmap) * runMyPow(x, n - n / 2 - n / 2, mmap)
    (*mmap)[n] = res
    return res
}
