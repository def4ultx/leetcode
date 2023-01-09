package main

import "fmt"

func main() {
	t := Constructor()
	timeMap := &t
	// timeMap.Set("foo", "bar", 1)  // store the key "foo" and value "bar" along with timestamp = 1.
	// timeMap.Get("foo", 1)         // return "bar"
	// timeMap.Get("foo", 3)         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
	// timeMap.Set("foo", "bar2", 4) // store the key "foo" and value "bar2" along with timestamp = 4.
	// timeMap.Get("foo", 4)         // return "bar2"
	// timeMap.Get("foo", 5)         // return "bar2"

	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)
	timeMap.Get("love", 5)
	timeMap.Get("love", 10)
	timeMap.Get("love", 15)
	timeMap.Get("love", 20)
	timeMap.Get("love", 25)
}

type Point struct {
	timestamp int
	value     string
}

type TimeMap struct {
	data map[string][]Point
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]Point),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	p := Point{
		timestamp: timestamp,
		value:     value,
	}

	_, ok := this.data[key]
	if ok {
		this.data[key] = append(this.data[key], p)
	} else {
		this.data[key] = []Point{p}
	}
	return
}

func (this *TimeMap) Get(key string, timestamp int) string {
	points, ok := this.data[key]
	if !ok {
		return ""
	}

	index := binarySearch(points, timestamp)
	if index == -1 {
		return ""
	}
	fmt.Println(points[index].value, index)
	return points[index].value
}

func binarySearch(points []Point, timestamp int) int {
	l, r := 0, len(points)-1
	for r >= l {
		mid := l + (r-l)/2
		if points[mid].timestamp == timestamp {
			return mid
		} else if points[mid].timestamp > timestamp {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// fmt.Println(l, r)
	if l == 0 {
		return -1
	}
	return l - 1
}
