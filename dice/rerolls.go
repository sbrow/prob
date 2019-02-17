package dice

// Average is a reroll function that rerolls all dice whose rolled side has a Value
// lower than the average.
var Average Reroll = func(d Dice) []bool {
	avgs := make([]float32, len(d))
	reroll := make([]bool, len(d))

	for i, die := range d {
		sides := die.Sides()
		for _, side := range sides {
			avgs[i] += float32(side.Value())
		}
		avgs[i] /= float32(len(sides))
		if float32(Value(die)) < avgs[i] {
			reroll[i] = true
		}
	}

	return reroll
}
