package main

import (
	"encoding/hex"
	"fmt"
	"log"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/protobuf/proto"
)

func main() {
	// s := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(s)
	// index := r.Intn(2)
	// fmt.Println(index)

	name := "Ball"
	book := &pb.HelloRequest{Name: name}
	// ...

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	// if err := ioutil.WriteFile("fname.bin", out, 0644); err != nil {
	// 	log.Fatalln("Failed to write address book:", err)
	// }

	fmt.Println(hex.Dump(out))
}
