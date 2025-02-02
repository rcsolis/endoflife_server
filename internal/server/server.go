package server

import (
	"context"
	"log"

	api "github.com/rcsolis/endoflife_server/internal/apicall"
	pb "github.com/rcsolis/endoflife_server/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server is the gRPC server
type Server struct {
	pb.UnimplementedCycleServiceServer
}

/**
 * NewServer creates a new server
 * @return *Server
 */
func newServer() *Server {
	return &Server{}
}

func RegisterGrpcServer(grpcServer *grpc.Server) {
	pb.RegisterCycleServiceServer(grpcServer, newServer())
}

/**
 * GetAllLanguages fetches all the languages from the API
 * @param context.Context
 * @param *pb.Empty
 * @return *pb.AllLanguagesResponse
 * @return error
 */
func (s *Server) GetAllLanguages(ctx context.Context, req *pb.Empty) (*pb.AllLanguagesResponse, error) {
	log.Println("GetAllLanguages function was invoked")

	// get from api
	langs, err := api.GetAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error fetching data: %v", err)
	}

	// return response
	var allLangs []*pb.Language
	for _, lang := range langs {
		allLangs = append(allLangs, &pb.Language{Name: lang})
	}
	return &pb.AllLanguagesResponse{Languages: allLangs}, nil
}

/**
 * GetAllVersions fetches all the versions of a language from the API and streams the data
 * @param *pb.Language
 * @param pb.CycleService_GetAllVersionsServer
 * @return error
 */
func (s *Server) GetAllVersions(req *pb.Language, stream pb.CycleService_GetAllVersionsServer) error {
	log.Println("GetAllVersions function was invoked")

	// Get from api
	details, err := api.GetAllDetails(req.GetName())
	if err != nil {
		return status.Errorf(codes.Internal, "Error fetching data: %v", err)
	}

	for _, item := range details {
		cycle := &pb.Cycle{
			Cycle:           item.Cycle,
			ReleaseDate:     item.ReleaseDate,
			Eol:             item.Eol,
			Latest:          item.Latest,
			Link:            item.Link,
			Lts:             item.Lts,
			Support:         item.Support,
			Discontinued:    item.Discontinued,
			ExtendedSupport: item.ExtendedSupport,
		}
		if err := stream.Send(cycle); err != nil {
			log.Println("Error sending data: ", item)
			return status.Errorf(codes.Internal, "Error sending data: %v", err)
		}
	}
	return nil
}

/**
 * GetDetails fetches the details of a specific version of a language from the API
 * @param context.Context
 * @param *pb.DetailsRequest
 * @return *pb.Cycle
 * @return error
 */
func (s *Server) GetDetails(ctx context.Context, req *pb.DetailsRequest) (*pb.Cycle, error) {
	log.Println("GetDetails function was invoked")
	// Get from api
	details, err := api.GetCycleDetails(req.GetName(), req.GetVersion())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error fetching data: %v", err)
	}
	cycle := &pb.Cycle{
		Cycle:           details.Cycle,
		ReleaseDate:     details.ReleaseDate,
		Eol:             details.Eol,
		Latest:          details.Latest,
		Link:            details.Link,
		Lts:             details.Lts,
		Support:         details.Support,
		Discontinued:    details.Discontinued,
		ExtendedSupport: details.ExtendedSupport,
	}
	return cycle, nil
}
