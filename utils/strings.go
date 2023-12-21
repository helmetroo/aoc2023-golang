package utils

func CharToInt(char byte) int {
    return int(char) - '0'
}

// In-place string reverse
// https://stackoverflow.com/a/10030772
func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
