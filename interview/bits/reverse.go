package main

func main() {

}

func reverseBits(num uint32) uint32 {
	var reversed uint32
	for i := 0; i < 32; i++ {
		reversed <<= 1
		if num&1 == 1 {
			reversed++
		}
		num >>= 1
	}
	return reversed
}
