package main

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {

	totalGas := 0
	totalCost := 0
	for i := range gas {
		totalGas += gas[i]
		totalCost += cost[i]
	}

	if totalGas < totalCost {
		return -1
	}

	total := 0
	start := -1
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		if start == -1 {
			start = i
		}
		if total < 0 {
			total = 0
			start = -1
		}
	}
	return start
}
