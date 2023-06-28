package main

import (
	"github.com/mansoorceksport/hexarch/internal/adapters/app/api"
	"github.com/mansoorceksport/hexarch/internal/adapters/core/arithmetic"
	"github.com/mansoorceksport/hexarch/internal/adapters/framework/left/grpc"
	"github.com/mansoorceksport/hexarch/internal/adapters/framework/right/db"
	"github.com/mansoorceksport/hexarch/internal/ports"
	"log"
	"os"
)

func main() {
	var err error
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
	defer dbaseAdapter.CloseDbConnection()

	arithmeticAdapter = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(arithmeticAdapter, dbaseAdapter)

	grpcAdapter = grpc.NewAdapter(appAdapter)

	grpcAdapter.Run()

	//arithPort := arithmetic.NewAdapter()
	//dbPort, err := db.NewAdapter("", "")

	//apiPort := api.NewAdapter(arithPort, dbPort)
	//grpc := grpc2.NewAdapter(apiPort)
	//grpc.Run()
}
