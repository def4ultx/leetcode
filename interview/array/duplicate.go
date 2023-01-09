package main

func main() {

}

func containsDuplicate(nums []int) bool {
	mapping := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := mapping[v]; ok {
			return true
		}
		mapping[v] = struct{}{}
	}
	return false
}
