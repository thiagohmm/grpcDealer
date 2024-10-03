package usecase

import "grpcDemonstracao/internal/entity"

type ListAllDealersInputDTO struct {
	IdRevendedor []int
}

type ListAllDealersOutputDTO struct {
	Categoria                       string
	Codigo                          int
	CodigodeBarra                   string
	Descricao                       string
	IdRevendedor                    int
	Marca                           string
	PodeSolicitarNovoCodigoDeBarras bool
	PodeSolicitarPermissaoDeVendas  bool
	ProdutoDoRevendedor             bool
	Subcategoria                    string
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
			Categoria:                       product.Categoria,
			Codigo:                          product.Codigo,
			CodigodeBarra:                   product.CodigodeBarra,
			Descricao:                       product.Descricao,
			IdRevendedor:                    product.IdRevendedor,
			Marca:                           product.Marca,
			PodeSolicitarNovoCodigoDeBarras: product.PodeSolicitarNovoCodigoDeBarras,
			PodeSolicitarPermissaoDeVendas:  product.PodeSolicitarPermissaoDeVendas,
			ProdutoDoRevendedor:             product.ProdutoDoRevendedor,
			Subcategoria:                    product.Subcategoria,
		})
	}
	return products, nil
}
