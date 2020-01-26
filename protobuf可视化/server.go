package main

import (
	pb "awesomeproject"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// func main() {
// 	jsonFilePath := "/Users/tiechengshen/go/src/awesomeProject/user_info.json"
// 	pbFilePath := "/Users/tiechengshen/go/src/awesomeProject/user_info.pb"

// 	buf, err := ioutil.ReadFile(jsonFilePath)
// 	if err != nil {
// 		fmt.Println("Read file err: ", err)
// 		os.Exit(0)
// 	}
// 	userInfo := &user_info.UserInfo{}

// 	if err = jsonpb.UnmarshalString(string(buf), userInfo); err != nil {
// 		fmt.Println("jsonpb UnmarshalString fail: ", err)
// 		os.Exit(0)
// 	}

// 	fmt.Println(proto.MarshalTextString(userInfo))
// 	//fmt.Println("user info pb: ", userInfo.String())

// 	data, err := proto.Marshal(userInfo)
// 	if err != nil {
// 		fmt.Println("proto Marshal fail: ", err)
// 		os.Exit(0)
// 	}

// 	if err = ioutil.WriteFile(pbFilePath, data, os.ModePerm); err != nil {
// 		fmt.Println("Write file err: ", err)
// 	}
// }

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	log.Println("request: ", in.Userinfo)
	return &pb.UserInfoResponse{Message: "hello " + in.Userinfo.String()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)

}
