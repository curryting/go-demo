package main

import (
	"bufio"
	"context"
	"fmt"
	pb "github.com/grpc-example/route"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000, // 乘10的7次方
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

func runSecond(client pb.RouteGuideClient) {
	serverStream, err := client.ListFeature(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{
			Latitude:  312978870, // 乘10的7次方
			Longitude: 121503457,
		},
		Hi: &pb.Point{
			Latitude:  311416130, // 乘10的7次方
			Longitude: 121424904,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		feature, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(feature)
	}
}

func runThird(client pb.RouteGuideClient) {
	clientStream, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	points := []*pb.Point{
		{Latitude: 313374060, Longitude: 121358540},
		{Latitude: 311034130, Longitude: 121598790},
		{Latitude: 310235000, Longitude: 121437403},
	}

	for id, point := range points {
		if err := clientStream.Send(point); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("send finished, id: ", id)
		time.Sleep(time.Second)
	}
	summary, err1 := clientStream.CloseAndRecv()
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Println(summary)
}

// 从标准流中读一个integer存到target中
func readIntFromCommandLine(reader *bufio.Reader, target *int32) {
	_, err := fmt.Fscanf(reader, "%d\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runFourth(client pb.RouteGuideClient) {
	stream, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// this goroutine listen to the server stream
	go func() {
		feature, err2 := stream.Recv()
		if err2 != nil {
			log.Fatalln(err2)
		}
		fmt.Println("recommend:", feature)
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		var mode int32
		request := pb.RecommendationRequest{
			Point: new(pb.Point),
		}
		fmt.Print("Enter Recommendation Mode (0 for farthest, 1 for the nearest)")
		readIntFromCommandLine(reader, &mode)
		fmt.Print("Enter Latitude: ")
		readIntFromCommandLine(reader, &request.Point.Latitude)
		fmt.Print("Enter Longitude: ")
		readIntFromCommandLine(reader, &request.Point.Longitude)
		request.Mode = pb.RecommendationMode(mode)

		if err := stream.Send(&request); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// grpc拨号请求
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock()) // WithInsecure 忽略证书验证		WithBlock-阻塞, 拨号成功才会往下走
	if err != nil {
		log.Fatalln("client conn failed")
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	runFourth(client)
}
