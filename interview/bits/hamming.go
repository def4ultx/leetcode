package main

func main() {

}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		rem := num % 2
		count += int(rem)
		num = num / 2
	}
	return count
}
