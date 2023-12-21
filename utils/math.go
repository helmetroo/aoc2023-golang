package utils

type UInt interface {
    uint | uint8 | uint16 | uint32 | uint64
}

func Lcm[T UInt](arr *[]T) T {
    size := len(*arr)
    curLcm := (*arr)[0]
    idx := 1

    for idx < size {
        a, b := curLcm, (*arr)[idx]
        curLcm = (a * b) / Gcd(a, b)
        idx++
    }

    return curLcm
}

func Gcd[T UInt](a, b T) T {
    if b == 0 {
        return a
    }

    return Gcd(b, a % b)
}
