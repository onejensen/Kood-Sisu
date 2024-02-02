package sprint

func Payout(amount int, denominations []int) (payout []int) {

	denomins := SortIntegerTable(denominations)

	for _, denom := range denomins {
		count := amount / denom
		for i := 0; i < count; i++ {
			payout = append(payout, denom)
		}
		amount %= denom
	}
	if amount == 0 {
		return payout
	}
	return []int{}
}

func SortIntegerTable(table []int) []int {
	n := len(table)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] < table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	return table
}
