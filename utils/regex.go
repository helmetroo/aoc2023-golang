package utils

import (
    "regexp"
)

var NUMBER_REGEX = regexp.MustCompile(`(\d+)`)
var NUMBER_WITH_NEG_REGEX = regexp.MustCompile(`-?\d+`)

