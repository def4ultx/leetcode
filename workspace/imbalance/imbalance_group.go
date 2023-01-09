package main

import "fmt"

// Amazon Academy recently organized a scholarship test on its platform. A total of n students participated in the test. Each student received a unique roll number, i. Each student's rank is stored at rank[i]. Each student gets a unique rank, so rank is a permutation of values 1 through n.

// For improved collaboration, the students are to be divided into groups. Use the following rules to find the imbalance of a group of students.

// A group has students where 1 <= k <= n.
// Groups are formed of students in ranks with consecutive roll numbers, i.e., i, (¡ + 1), .. (1 + K - 1).
// The ranks of the students in a group are then sorted ascending to an array, here named sorted_ rank.
// The imbalance of the group is then defined as the number of students , who are more than 1 rank beneath the rank of the student just ahead of them, 1.e, sorted rank[i]- sorted rank[i - 1] > 1.
// For example, the ranks in a group are [1, 5, 4] so sorted rank = [1, 4, 5).

// 4-1=3, and 3 > 1. This adds 1 to the imbalance.
// 5-4=1, and 1 = 1. This does not add to the imbalance.
// The imbalance is 1.

// Function Description

// Givent the ranks of n students, find the total sum of the imbalance of all possible groups.

// findTotalImbalance has the following parameter:
// int rank[n]: the ranks of each student

// Returns

// long_int: the total sum of imbalance over all possible groups

// Contraints

// 1 <= n < 3 * 10^3
// rank is a permutation of length n
// Input0: [4, 1, 3, 2]

// Output 0: 3

// Explanation 0:
// [1] --> [1], imbalance = 0.
// [2] --> [2], imbalance = 0.
// [3] --> [3], imbalance = 0.
// [4] --> [4], imbalance = 0.
// [4, 1] --> [1, 4], imbalance = 1.
// [1, 3] --> [1, 3], imbalance = 1.
// [3, 2] --> [2, 3], imbalance = 0.
// [4, 1, 3] --> [1, 3, 4], imbalance = 1.
// [1, 3, 2] --> [1, 2, 3], imbalance = 0.
// [4, 1, 3, 2] --> [1, 2, 3, 4], imbalance = 0.

// Summing it up, the total imbalance is 3,

func main() {
	// findImbalance([]int{4, 1, 3, 2})
	// findImbalance([]int{2, 0, 1})
	// findImbalance([]int{1, 0, 2})
	// findImbalance([]int{44, 11, 33, 22})
	// findImbalance([]int{11, 22, 33, 44})
	// findImbalance([]int{44, 33, 22, 11})
	// findImbalance([]int{44, 33, 22, 11, 55})
	// findImbalance([]int{44, 11, 33, 22, 55})
	findImbalance([]int{9, 19, 1, 1, 2, 9})
}

func findImbalance(arr []int) int {
	var total int
	for i := 0; i < len(arr); i++ {

		mapping := make(map[int]struct{})
		valid := make(map[int]struct{})
		for j := i; j < len(arr); j++ {
			mapping[arr[j]] = struct{}{}
			if _, ok := mapping[arr[j]+1]; ok {
				valid[arr[j]+1] = struct{}{}
			}
			if _, ok := mapping[arr[j]-1]; ok {
				valid[arr[j]-1] = struct{}{}
			}

			m, v := len(mapping), len(valid)
			fmt.Println(m, v, m-v, m-v-1, mapping, "|", valid, "i,j", i, j)
			total += m - v - 1
		}
	}
	fmt.Println(total)
	return total
}

// func findImbalance(arr []int) int {

// 	var count int
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println("----")
// 		var imbalance, prev int
// 		min, max := arr[i], arr[i]
// 		mapping := make(map[int]struct{})
// 		for j := i; j < len(arr); j++ {

// 			var c int

// 			// use this logic but count with prev
// 			if _, ok := mapping[arr[j]+1]; !ok && i != j && arr[j] < max-1 {
// 				c++
// 				fmt.Println(i, j, "|", arr[j], arr[j]+1, "at upper", mapping, "prev", prev)
// 			}

// 			if _, ok := mapping[arr[j]-1]; !ok && i != j && arr[j] > min+1 {
// 				c++
// 				fmt.Println(i, j, "|", arr[j], arr[j]-1, "at lower", mapping, "prev", prev)
// 			}

// 			if _, ok := mapping[arr[j]]; ok {
// 				imbalance += prev
// 			} else if c > 0 {

// 				if c > prev {
// 					imbalance += c
// 				} else {
// 					// x := Abs(prev, c)
// 					imbalance += prev + c - 1
// 				}

// 				fmt.Println(imbalance, Abs(prev, c), prev, c)

// 				prev++
// 			}

// 			mapping[arr[j]] = struct{}{}

// 			min = Min(min, arr[j])
// 			max = Max(max, arr[j])
// 		}
// 		count += imbalance
// 	}

// 	fmt.Println("result", count)
// 	return count
// }

// func findImbalance(arr []int) int {

// 	var count int
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println("----")
// 		mapping := make(map[int]struct{})

// 		var localCount int
// 		// var maxFromLast int
// 		min, max := arr[i], arr[i]
// 		for j := i; j < len(arr); j++ {

// 			mapping[arr[j]] = struct{}{}

// 			// var added bool
// 			var c int

// 			if _, ok := mapping[arr[j]+1]; !ok && i != j && arr[j] < max-1 {
// 				// added = true
// 				c++

// 				// fmt.Println("size", j-i, c, j-i-c)
// 				// count++
// 				fmt.Println(i, j, "|", arr[j], arr[j]+1, "at upper", mapping)
// 			}
// 			// } else if arr[j] > max {
// 			// 	fmt.Println(i, j, "|", arr[j], arr[j]-1, "at max", mapping)
// 			// 	count++
// 			// }

// 			if _, ok := mapping[arr[j]-1]; !ok && i != j && arr[j] > min+1 {
// 				// added = true
// 				c++

// 				// fmt.Println("size", j-i, c, j-i-c)
// 				// count++
// 				fmt.Println(i, j, "|", arr[j], arr[j]-1, "at lower", mapping)
// 			}
// 			// } else if arr[j] < min {
// 			// 	fmt.Println(i, j, "|", arr[j], arr[j]-1, "at min", mapping)
// 			// 	count++
// 			// }

// 			min = Min(min, arr[j])
// 			max = Max(max, arr[j])

// 			localCount += c

// 			// if c > 0 {

// 			// 	fmt.Println(localCount, maxFromLast, c)
// 			// 	if maxFromLast > 1 && c == 2 {
// 			// 		localCount = localCount + maxFromLast + 1
// 			// 	} else if maxFromLast == 1 && c == 2 {
// 			// 		localCount = localCount + 2
// 			// 	}
// 			// 	if c > 0 {
// 			// 		localCount = localCount + 1
// 			// 	}

// 			// 	if c > 0 {
// 			// 		maxFromLast++
// 			// 	}

// 			// 	fmt.Println(localCount, maxFromLast, c)
// 			// }

// 			// addthisround := c
// 			// fmt.Println(localCount, maxFromLast, c)
// 			// if c == 2 && maxFromLast >= 2 {
// 			// 	maxFromLast = maxFromLast + c - 1
// 			// } else {
// 			// 	maxFromLast = maxFromLast + c
// 			// }

// 			// if added {
// 			// 	c++
// 			// }
// 		}
// 		fmt.Println("final local count", localCount)
// 		count += localCount
// 	}

// 	fmt.Println("result", count)
// 	return count
// }

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
