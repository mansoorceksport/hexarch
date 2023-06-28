package grpc

import (
	"context"
	"github.com/mansoorceksport/hexarch/internal/adapters/app/api"
	"github.com/mansoorceksport/hexarch/internal/adapters/core/arithmetic"
	"github.com/mansoorceksport/hexarch/internal/adapters/framework/left/grpc/pb"
	"github.com/mansoorceksport/hexarch/internal/adapters/framework/right/db"
	"github.com/mansoorceksport/hexarch/internal/ports"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error

	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	// ports
	var dbaseAdapter ports.DbPort
	var arithmeticAdapter ports.ArithmeticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPorts

	dbaseDriver := os.Getenv("DB_DRIVER")
	dbSourceName := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dbSourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}

	arithmeticAdapter = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(arithmeticAdapter, dbaseAdapter)

	grpcAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, grpcAdapter)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}

	return conn
}

func TestAdapter_GetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Fatalf("failed to close GRPC conn: %v", err)
		}
	}(conn)

	client := pb.NewArithmeticServiceClient(conn)
	result, err := client.GetAddition(ctx, &pb.OperationParameters{
		A: 1,
		B: 1,
	})
	if err != nil {
		t.Fatalf("expected: %v got: %v", nil, err)
	}

	require.Equal(t, result.Value, int32(2))
}

func TestAdapter_GetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Fatalf("failed to close GRPC conn: %v", err)
		}
	}(conn)

	client := pb.NewArithmeticServiceClient(conn)
	result, err := client.GetSubtraction(ctx, &pb.OperationParameters{
		A: 1,
		B: 1,
	})
	if err != nil {
		t.Fatalf("expected: %v got: %v", nil, err)
	}

	require.Equal(t, result.Value, int32(0))
}

func TestAdapter_GetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Fatalf("failed to close GRPC conn: %v", err)
		}
	}(conn)

	client := pb.NewArithmeticServiceClient(conn)
	result, err := client.GetMultiplication(ctx, &pb.OperationParameters{
		A: 1,
		B: 1,
	})
	if err != nil {
		t.Fatalf("expected: %v got: %v", nil, err)
	}

	require.Equal(t, result.Value, int32(1))
}

func TestAdapter_GetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Fatalf("failed to close GRPC conn: %v", err)
		}
	}(conn)

	client := pb.NewArithmeticServiceClient(conn)
	result, err := client.GetDivision(ctx, &pb.OperationParameters{
		A: 1,
		B: 1,
	})
	if err != nil {
		t.Fatalf("expected: %v got: %v", nil, err)
	}

	require.Equal(t, result.Value, int32(1))
}
