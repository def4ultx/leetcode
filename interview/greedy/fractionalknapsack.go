package main

import (
	"fmt"
	"sort"
)

func main() {
	factional([]int{60, 100, 120}, []int{10, 20, 30}, 50)
}

type Item struct {
	weight int
	price  int
	ratio  float32
}

type ItemList []Item

func (l ItemList) Len() int           { return len(l) }
func (l ItemList) Less(i, j int) bool { return l[i].ratio < l[j].ratio }
func (l ItemList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func factional(prices, weights []int, weight int) int {

	lenght := len(prices)
	items := make(ItemList, lenght)
	for i := 0; i < lenght; i++ {
		items[i] = Item{
			weight: weights[i],
			price:  prices[i],
			ratio:  float32(prices[i]) / float32(weights[i]),
		}
	}
	sort.Reverse(items)
	fmt.Println(items)

	// sort.Slice(items, func(i, j int) bool { return items[i].ratio < items[j].ratio })

	totalWeight, totalPrice := 0, 0

	for i := 0; i < len(items) && totalWeight < weight; i++ {
		if totalWeight+items[i].weight <= weight {
			totalPrice += items[i].price
			totalWeight += items[i].weight
		} else {
			remainingWeight := weight - totalWeight
			remainingPrice := float32(remainingWeight) / (float32(items[i].weight)) * float32(items[i].price)
			totalWeight += totalWeight
			totalPrice += int(remainingPrice)
		}
	}

	fmt.Println(totalPrice)
	return totalPrice
}

// Input:
// Items as (value, weight) pairs
// arr[] = {{60, 10}, {100, 20}, {120, 30}}
// Knapsack Capacity, W = 50;

// Output:
// Maximum possible value = 240
// by taking items of weight 10 and 20 kg and 2/3 fraction
// of 30 kg. Hence total price will be 60+100+(2/3)(120) = 240
