package utils

func ValidOffset(r, c, rows, cols int) bool {
    return r >= 0 && r <= rows - 1 &&
        c >= 0 && c <= cols - 1
}
