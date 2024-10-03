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
			Categoria:                       dealer.Categoria,
			Codigo:                          int32(dealer.Codigo),
			CodigodeBarra:                   dealer.CodigodeBarra,
			Descricao:                       dealer.Descricao,
			IdRevendedor:                    int32(dealer.IdRevendedor),
			Marca:                           dealer.Marca,
			PodeSolicitarNovoCodigoDeBarras: dealer.PodeSolicitarNovoCodigoDeBarras,
			PodeSolicitarPermissaoDeVendas:  dealer.PodeSolicitarPermissaoDeVendas,
			ProdutoDoRevendedor:             dealer.ProdutoDoRevendedor,
			Subcategoria:                    dealer.Subcategoria,
		}
	}

	return &pb.ListProductsResponse{
		Products: products,
	}, nil
}
