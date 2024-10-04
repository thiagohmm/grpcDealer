package service

import (
	"context"

	"grpcDemonstracao/internal/infra/grpc/pb"
	"grpcDemonstracao/internal/usecase"
)

type ListProductsService struct {
	pb.UnimplementedListProductsServiceServer
	ListProductsUseCase usecase.ListAllDealersUseCase
}

func (s *ListProductsService) ListProducts(ctx context.Context, in *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	dtoIN := usecase.ListAllDealersInputDTO{
		IdRevendedor: make([]int, len(in.Idrevendedor)),
	}

	output, err := s.ListProductsUseCase.Execute(&dtoIN)
	if err != nil {
		return nil, err
	}

	products := make([]*pb.Product, len(output))
	for i, dealer := range output {
		products[i] = &pb.Product{
			Categoria:                       dealer.Categoria.String, // Access the value of dealer.Categoria
			Codigo:                          int32(dealer.Codigo),
			CodigodeBarra:                   dealer.CodigodeBarra.String,
			Descricao:                       dealer.Descricao.String,
			IdRevendedor:                    int32(dealer.IdRevendedor.Int64), // Convert sql.NullInt64 to int32
			Marca:                           dealer.Marca.String,              // Access the value of dealer.Marca
			PodeSolicitarNovoCodigoDeBarras: dealer.PodeSolicitarNovoCodigoDeBarras,
			PodeSolicitarPermissaoDeVendas:  dealer.PodeSolicitarPermissaoDeVendas,
			ProdutoDoRevendedor:             dealer.ProdutoDoRevendedor,
			Subcategoria:                    dealer.Subcategoria.String,
		}
	}

	return &pb.ListProductsResponse{
		Products: products,
	}, nil
}
