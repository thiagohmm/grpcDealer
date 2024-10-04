package entity

import "database/sql"

type Dealer struct {
	Categoria                       sql.NullString
	Codigo                          int64
	CodigodeBarra                   sql.NullString
	Descricao                       sql.NullString
	IdRevendedor                    sql.NullInt64
	Marca                           sql.NullString
	PodeSolicitarNovoCodigoDeBarras bool
	PodeSolicitarPermissaoDeVendas  bool
	ProdutoDoRevendedor             bool
	Subcategoria                    sql.NullString
}

func NewDealer(categoria string, codigo int64, codigodeBarra string, descricao string, idRevendedor int64, marca string, podeSolicitarNovoCodigoDeBarras bool, podeSolicitarPermissaoDeVendas bool, produtoDoRevendedor bool, subcategoria string) *Dealer {
	return &Dealer{
		Categoria:                       sql.NullString{String: categoria, Valid: true},
		Codigo:                          codigo,
		CodigodeBarra:                   sql.NullString{String: codigodeBarra, Valid: true},
		Descricao:                       sql.NullString{String: descricao, Valid: true},
		IdRevendedor:                    sql.NullInt64{Int64: idRevendedor, Valid: true},
		Marca:                           sql.NullString{String: marca, Valid: true},
		PodeSolicitarNovoCodigoDeBarras: podeSolicitarNovoCodigoDeBarras,
		PodeSolicitarPermissaoDeVendas:  podeSolicitarPermissaoDeVendas,
		ProdutoDoRevendedor:             produtoDoRevendedor,
		Subcategoria:                    sql.NullString{String: subcategoria, Valid: true},
	}
}
