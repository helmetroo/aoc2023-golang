package puzzles

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
    Cards string
    Bid int
    Strength int
}

const (
    HighCard = iota
    OnePair
    TwoPair
    ThreeOfAKind
    FullHouse
    FourOfAKind
    FiveOfAKind
)

var CARD_STRENGTHS = map[byte]int {
    '2': 0,
        '3': 1,
        '4': 2,
        '5': 3,
        '6': 4,
        '7': 5,
        '8': 6,
        '9': 7,
        'T': 8,
        'J': 9,
        'Q': 10,
        'K': 11,
        'A': 12,
}

func computeStrength(cards *string) int {
    cardFreqs := map[byte]int{}
    for idx := 0; idx < 5; idx++ {
        cardFreqs[(*cards)[idx]] += 1
    }

    lenCardFreqs := len(cardFreqs)
    switch lenCardFreqs {
    case 2, 3:
        return discernStrength(lenCardFreqs, &cardFreqs)
    case 1:
        return FiveOfAKind
    case 4:
        return OnePair
    default:
        return HighCard
    }
}

func discernStrength(lenCardFreqs int, cardFreqs *map[byte]int) int {
    if lenCardFreqs == 2 {
        // Full house has counts => 2, 3
        // Four of a kind has counts => 1, 4
        for _, strength := range *cardFreqs {
            if (strength == 2 || strength == 3) {
                return FullHouse
            }
        }

        return FourOfAKind
    }

    // Three of a kind has counts => 3, 1, 1
    // Two pair has counts => 2, 2, 1
    for _, strength := range *cardFreqs {
        if strength == 1 {
            continue
        }

        if strength == 2 {
            return TwoPair
        }
    }

    return ThreeOfAKind
}

func parseHands(scanner *bufio.Scanner) ([]Hand, error) {
    hands := []Hand{}

    for scanner.Scan() {
        handLine := scanner.Text()
        splHandLine := strings.Split(handLine, " ")

        cards := splHandLine[0]
        bid, _ := strconv.Atoi(splHandLine[1])
        strength := computeStrength(&cards)

        hands = append(hands, Hand { cards, bid, strength })
    }

    return hands, nil
}

func sortHandsByStrength(hands *[]Hand) {
    sort.Slice(*hands, func(i, j int) bool {
        a, b := (*hands)[i], (*hands)[j]

        if(a.Strength != b.Strength) {
            return a.Strength < b.Strength
        }

        return a.less(&b)
    })
}

func (a *Hand) less(b *Hand) bool {
    for idx := 0; idx < 5; idx++ {
        aCard, bCard := a.Cards[idx], b.Cards[idx]
        aStrength, bStrength := CARD_STRENGTHS[aCard], CARD_STRENGTHS[bCard]
        if aStrength == bStrength {
            continue
        }

        if aStrength < bStrength {
            return true
        }

        return false
    }

    // Never reached
    return false
}

func P7_SolvePartOne(scanner *bufio.Scanner) (string, error) {
    hands, _ := parseHands(scanner)
    sortHandsByStrength(&hands)

    bidRankSum := 0
    for idx, hand := range hands {
        bidRankSum += hand.Bid * (idx + 1)
    }

    return strconv.Itoa(bidRankSum), nil
}

func P7_SolvePartTwo(scanner *bufio.Scanner) (string, error) {
    return "", nil
}
