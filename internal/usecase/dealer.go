package usecase

import (
	"database/sql"
	"grpcDemonstracao/internal/entity"
)

type ListAllDealersInputDTO struct {
	IdRevendedor []int
}

type ListAllDealersOutputDTO struct {
	Categoria                       sql.NullString `json:"categoria"`
	Codigo                          int64          `json:"codigo"`
	CodigodeBarra                   sql.NullString `json:"codigodeBarra"`
	Descricao                       sql.NullString `json:"descricao"`
	IdRevendedor                    sql.NullInt64  `json:"idRevendedor"`
	Marca                           sql.NullString `json:"marca"`
	PodeSolicitarNovoCodigoDeBarras bool           `json:"podeSolicitarNovoCodigoDeBarras"`
	PodeSolicitarPermissaoDeVendas  bool           `json:"podeSolicitarPermissaoDeVendas"`
	ProdutoDoRevendedor             bool           `json:"produtoDoRevendedor"`
	Subcategoria                    sql.NullString `json:"subcategoria"`
}

type ListAllDealersUseCase struct {
	DealerRepository entity.DealerRepository
}

func NewListAllDealersUseCase(dealerRepository entity.DealerRepository) *ListAllDealersUseCase {
	return &ListAllDealersUseCase{
		DealerRepository: dealerRepository,
	}
}

func (u *ListAllDealersUseCase) Execute(input *ListAllDealersInputDTO) ([]*ListAllDealersOutputDTO, error) {
	var dealers []*entity.Dealer

	for _, id := range input.IdRevendedor {
		dealer, err := u.DealerRepository.ListAllDealers(id)
		if err != nil {
			return nil, err
		}

		dealers = append(dealers, dealer...)
	}

	var products []*ListAllDealersOutputDTO

	for _, product := range dealers {
		products = append(products, &ListAllDealersOutputDTO{
			Categoria:                       sql.NullString{String: product.Categoria.String, Valid: true}, // Access the String field of product.Categoria
			Codigo:                          product.Codigo,
			CodigodeBarra:                   sql.NullString{String: product.CodigodeBarra.String, Valid: true}, // Access the String field of product.CodigodeBarra
			Descricao:                       sql.NullString{String: product.Descricao.String, Valid: true},     // Access the String field of product.Descricao
			IdRevendedor:                    product.IdRevendedor,
			Marca:                           sql.NullString{String: product.Marca.String, Valid: true}, // Access the String field of product.Marca
			PodeSolicitarNovoCodigoDeBarras: product.PodeSolicitarNovoCodigoDeBarras,
			PodeSolicitarPermissaoDeVendas:  product.PodeSolicitarPermissaoDeVendas,
			ProdutoDoRevendedor:             product.ProdutoDoRevendedor,
			Subcategoria:                    sql.NullString{String: product.Subcategoria.String, Valid: true}, // Access the String field of product.Subcategoria
		})
	}
	return products, nil
}
