package entity

type Dealer struct {
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

func NewDealer(categoria string, codigo int, codigodeBarra string, descricao string, idRevendedor int, marca string, podeSolicitarNovoCodigoDeBarras bool, podeSolicitarPermissaoDeVendas bool, produtoDoRevendedor bool, subcategoria string) *Dealer {
	return &Dealer{
		Categoria:                       categoria,
		Codigo:                          codigo,
		CodigodeBarra:                   codigodeBarra,
		Descricao:                       descricao,
		IdRevendedor:                    idRevendedor,
		Marca:                           marca,
		PodeSolicitarNovoCodigoDeBarras: podeSolicitarNovoCodigoDeBarras,
		PodeSolicitarPermissaoDeVendas:  podeSolicitarPermissaoDeVendas,
		ProdutoDoRevendedor:             produtoDoRevendedor,
		Subcategoria:                    subcategoria,
	}
}
