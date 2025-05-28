package main

func main() {
	leastBricks([][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	})
}
func leastBricks(wall [][]int) int {
	freq := make(map[int]int)

	for _, ws := range wall {

		idx := 0
		for i := 0; i < len(ws)-1; i++ {
			idx += ws[i]
			freq[idx]++
		}
	}

	// fmt.Println(freq)

	var max int
	for _, v := range freq {
		if v > max {
			max = v
		}
	}

	return len(wall) - max
}

// func leastBricks(wall [][]int) int {
// 	type Range struct {
// 		start, end, count int
// 	}

// 	ranges := make([]Range, 0)
// 	for _, ws := range wall {
// 		start := 0
// 		for _, w := range ws {
// 			size := w*2 - 1

// 			r := Range{
// 				start: start,
// 				end:   start + size,
// 				count: 0,
// 			}
// 			ranges = append(ranges, r)

// 			start = start + size + 1
// 		}
// 	}

// 	sort.Slice(ranges, func(i, j int) bool {
// 		switch {
// 		case ranges[i].start < ranges[j].start:
// 			return true
// 		case ranges[i].start > ranges[j].start:
// 			return false
// 		default:
// 			return ranges[i].end < ranges[j].end
// 		}
// 	})

// 	temp := make([]Range, 0)

// 	var idx int
// 	for len(ranges) < idx+1 {
// 		curr := ranges[idx]
// 		next := ranges[idx+1]

// 		/*
// 			Possible case
// 			Split
// 			=> curr.start != next.start || curr.end != next.end
// 				Split to next.start = curr.end
// 				Put in temporary

// 			Merge
// 			=> curr.start == next.start && curr.end == next.end
// 		*/

// 		if curr.start == next.start && curr.end == next.end {

// 		}

// 		idx++
// 	}
// 	fmt.Println(ranges)

// 	min := math.MaxInt
// 	for _, v := range ranges {
// 		if v.count < min {
// 			min = v.count
// 		}
// 	}
// 	return min
// }

// func remove(slice []int, s int) []int {
// 	return append(slice[:s], slice[s+1:]...)
// }

// func leastBricks(wall [][]int) int {
// 	var width int
// 	for _, v := range wall[0] {
// 		width += v
// 	}

// 	brickCountColumn := make([]int, width*2-1)

// 	fmt.Println(width)
// 	fmt.Println(brickCountColumn)

// 	for _, ws := range wall {
// 		idx := 0
// 		for _, w := range ws {
// 			total := w*2 - 1

// 			for i := idx; i < idx+total; i++ {
// 				brickCountColumn[i]++
// 			}
// 			idx += total + 1
// 		}
// 		fmt.Println(brickCountColumn)
// 	}

// 	fmt.Println(brickCountColumn)

// 	min := math.MaxInt
// 	for _, v := range brickCountColumn {
// 		if v < min {
// 			min = v
// 		}
// 	}
// 	return min
// }

/*
[6,0,0,0,0,0,0,0,0,0,0]


[1,2,2,1]
[3,1,2]
[1,3,2]
[2,4]
[3,1,2]
[1,3,1,1]
*/

/*
[0,1], [1,2]

1	2	3	4	5	6	7	8	9	10	11

1	0	1	1	1	0	1	1	1	0	1
1	1	1	1	1	0	1	0	1	1	1
1	0	1	1	1	1	1	0	1	1	1
1	1	1	0	1	1	1	1	1	1	1
1	1	1	1	1	0	1	0	1	1	1
1	0	1	1	1	1	1	0	1	0	1

6	3	6	5	6	3	6	2	6	4	6
*/
