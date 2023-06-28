package grpc

import (
	"context"
	"github.com/mansoorceksport/hexarch/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	result, err := grpca.api.GetAdd(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{Value: result}

	return ans, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	result, err := grpca.api.GetSubtraction(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{Value: result}

	return ans, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	result, err := grpca.api.GetMultiplication(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{Value: result}

	return ans, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	result, err := grpca.api.GetDivision(req.GetA(), req.GetB())
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{Value: result}

	return ans, nil
}
