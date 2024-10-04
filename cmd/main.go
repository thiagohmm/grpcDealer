package main

import (
	"grpcDemonstracao/config"
	"grpcDemonstracao/internal/infra/database"
	"grpcDemonstracao/internal/infra/grpc/pb"
	"grpcDemonstracao/internal/infra/grpc/service"
	"grpcDemonstracao/internal/repository"
	"grpcDemonstracao/internal/usecase"

	"net"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg, err := config.LoadConfig("../../.env")

	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Conecta ao banco de dados

	db, err := database.ConectarBanco(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	dealerUsecase := usecase.NewListAllDealersUseCase(repository.NewDealerRepository(db))

	server := grpc.NewServer(grpc.MaxSendMsgSize(100*1024*1024), grpc.MaxRecvMsgSize(100*1024*1024))

	dealergrpService := service.ListProductsService{
		ListProductsUseCase: *dealerUsecase,
	}

	pb.RegisterListProductsServiceServer(server, &dealergrpService)

	reflection.Register(server)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := server.Serve(listen); err != nil {
		panic(err)
	}

}
