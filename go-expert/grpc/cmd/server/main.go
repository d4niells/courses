package main

import (
	"database/sql"
	"net"

	"github.com/d4niells/goexpert-grpc/internal/database"
	"github.com/d4niells/goexpert-grpc/internal/pb"
	"github.com/d4niells/goexpert-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../../db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
