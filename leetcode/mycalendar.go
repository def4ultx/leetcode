package main

func main() {

}

type Interval struct {
	Start int
	End   int
}

type MyCalendar struct {
	data []Interval
}

func Constructor() MyCalendar {
	return MyCalendar{make([]Interval, 0)}
}

func (this *MyCalendar) Book(start int, end int) bool {

}
