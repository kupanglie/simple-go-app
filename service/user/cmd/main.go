package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	grpc_user "github.com/kupanglie/simple-go-app/grpc/user"
	internal_user_grpc "github.com/kupanglie/simple-go-app/service/user/internal/grpc"
	"github.com/kupanglie/simple-go-app/service/user/internal/repository"
	"github.com/kupanglie/simple-go-app/service/user/internal/usecase"
	"google.golang.org/grpc"
)

const (
	SERVICE_USER_PORT = "8900"
)

func main() {
	ctx := context.Background()

	if err := InitGRPC(ctx, SERVICE_USER_PORT); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func InitGRPC(ctx context.Context, servicePort string) error {
	listen, err := net.Listen("tcp", ":"+servicePort)
	if err != nil {
		return err
	}

	srv := grpc.NewServer()

	db, err := sql.Open("mysql", "root:root@tcp(0.0.0.0:3307)/warung_alfin_user")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	userRP := repository.NewUserRP(db)
	userUC := usecase.NewUserUC(userRP)
	userSrv := internal_user_grpc.NewUserGRPC(userUC)

	grpc_user.RegisterUserHandlerServer(srv, userSrv)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")

			srv.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Printf("starting gRPC server on port %s...", servicePort)
	return srv.Serve(listen)
}
