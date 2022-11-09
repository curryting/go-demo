package main

import (
	"context"
	pb "github.com/grpc-example/route"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math"
	"net"
	"time"
)

// RouteGuideServer 实现RouteGuide服务
type RouteGuideServer struct {
	features []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

// GetFeature GetFeature
func (s *RouteGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if proto.Equal(feature.Location, point) { // proto.Equal 检测两个message是否一样
			return feature, nil
		}
	}
	return nil, nil
}

// check if a point is inside a rectangle  检查点是否在矩形内
func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

// ListFeature ListFeature
func (s *RouteGuideServer) ListFeature(rectangle *pb.Rectangle, stream pb.RouteGuide_ListFeatureServer) error {
	for _, feature := range s.features {
		if inRange(feature.Location, rectangle) {
			if err := stream.Send(feature); err != nil {
				log.Fatalln(err)
			}
		}
	}
	return nil
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// calcDistance calculates the distance between two points using the "haversine" formula.
// The formula is based on http://mathforum.org/library/drmath/view/51879.html.
// 通过两个点的经度和纬度，计算出他们的长度(m)
func calcDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

// RecordRoute RecordRoute
func (s *RouteGuideServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	startTime := time.Now()
	var pointCount, distance int32
	var prevPoint *pb.Point

	for {
		point, err := stream.Recv()
		if err == io.EOF {
			// conclude a route summary
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:  pointCount,
				Distance:    distance,
				ElapsedTime: int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		pointCount++ // point的个数
		if prevPoint != nil {
			distance += calcDistance(prevPoint, point)
		}
		prevPoint = point
	}
}

func (s *RouteGuideServer) recommendOnce(request *pb.RecommendationRequest) (*pb.Feature, error) {
	var nearest, farthest *pb.Feature
	var nearestDistance, farthestDistance int32

	for _, feature := range s.features {
		distance := calcDistance(feature.Location, request.Point)
		if nearest == nil || distance < nearestDistance {
			nearestDistance = distance
			nearest = feature
		}
		if farthest == nil || distance > farthestDistance {
			farthestDistance = distance
			farthest = feature
		}
	}
	if request.Mode == pb.RecommendationMode_GetFarthest {
		return farthest, nil
	} else {
		return nearest, nil
	}
}

// Recommend 1
func (s *RouteGuideServer) Recommend(stream pb.RouteGuide_RecommendServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		recommend, err1 := s.recommendOnce(request)
		if err1 != nil {
			return err1
		}
		return stream.Send(recommend)
	}
}

func NewServer() *RouteGuideServer {
	return &RouteGuideServer{
		features: []*pb.Feature{
			{Name: "上海交通大学", Location: &pb.Point{
				Latitude:  310235000, // 乘10的7次方
				Longitude: 121437403,
			}},
			{Name: "复旦大学", Location: &pb.Point{
				Latitude:  312978870, // 乘10的7次方
				Longitude: 121503457,
			}},
			{Name: "华东理工大学", Location: &pb.Point{
				Latitude:  311416130, // 乘10的7次方
				Longitude: 121424904,
			}},
		},
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("Listen server failed")
	}
	grpcServer := grpc.NewServer()                       // 创建一个grpc服务
	pb.RegisterRouteGuideServer(grpcServer, NewServer()) // 注册路线指南服务 到 grpc服务 上面
	log.Fatalln(grpcServer.Serve(lis))                   // 接受一个连接
}
