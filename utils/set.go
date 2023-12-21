package utils

import (
    "strconv"
)

func ToNumSet(nums string) map[int]interface{} {
    numSet := map[int]interface{}{}

    numIndices := NUMBER_REGEX.FindAllStringIndex(nums, -1)
    for _, indices := range numIndices {
        firstDigit, lastDigit := indices[0], indices[1]
        num, _ := strconv.Atoi(nums[firstDigit:lastDigit])
        numSet[num] = nil
    }

    return numSet
}

func Intersect[K comparable](first *map[K]interface{}, second *map[K]interface{}) map[K]interface{} {
    intersection := map[K]interface{}{}

    // Iterate over the smaller one (the result set won't be any bigger than it!)
    if len(*first) > len(*second) {
        first, second = second, first
    }

    for num := range *first {
        if _, exists := (*second)[num]; exists {
            intersection[num] = nil
        }
    }

    return intersection
}
