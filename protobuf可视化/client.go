package main

import (
	pb "awesomeProject"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50002"
)

func main() {
	jsonFilePath := "/Users/tiechengshen/go/src/awesomeProject/user_info.json"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	buf, err := ioutil.ReadFile(jsonFilePath)
	userinfo := &pb.UserInfo{}
	if err = jsonpb.UnmarshalString(string(buf), userinfo); err != nil {
		fmt.Println("jsonpb UnmarshalString fail: ", err)
		os.Exit(0)
	}
	r, err := c.SayHello(context.Background(), &pb.UserInfoRequest{Userinfo: userinfo})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.Message)
}
