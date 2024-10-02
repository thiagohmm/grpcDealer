package database

import (
	"database/sql"
	"grpcDemonstracao/internal/entity"
)

type DealerRepository struct {
	Db *sql.DB
}

func NewDealerRepository(db *sql.DB) *DealerRepository {
	return &DealerRepository{
		Db: db,
	}
}

func (r *DealerRepository) ListAllDealers(p_idrevendedor int) ([]*entity.Dealer, error) {
	query := `
	SELECT p.idproduto AS codigo
 , p.descricaoproduto AS descricao
 , m.nomemarca AS Marca
 , ep.CodigoBarras AS codigodebarra
 , em1.nomenivel AS Categoria
 , em2.nomenivel AS Subcategoria
 --, MeuCatalogo
 , r.idrevendedor
, CASE WHEN 
	 (    (pr.idproduto IS NULL) AND
		NOT (op.idproduto IS NOT NULL AND pr1.idproduto IS NOT NULL AND opr.idrevendedor IS NOT NULL) AND
		NOT (cpgi.idproduto IS NOT NULL)
	 ) AND
	 (
	 ( r.idsegmento = 0 AND (ps.idproduto IS NULL AND p.IdNivel1EstrMerc = pm.valor ) ) OR
	 ( r.idsegmento > 0 AND r.idsegmento <> 99 AND (ps.idproduto IS NULL OR ps.qt_seg_dif > 0) ) OR
	 ( r.idsegmento = 99 AND ps.idproduto IS NULL )
	 ) THEN 1 ELSE 0 END AS PodeSolicitarPermissaoDeVendas
, CASE WHEN 
	 ( (pr1.idproduto IS NOT NULL) OR
		 (op1.idproduto IS NOT NULL AND pr1.idproduto IS NOT NULL AND opr1.idrevendedor IS NOT NULL) OR
		 (cpgi.idproduto IS NOT NULL)
	 ) OR
	 ( (-- Revendedor Sem segmento
			r.idsegmento = 0 AND
			((ps.idproduto IS NULL AND p.IdNivel1EstrMerc <> pm.valor) OR
			 (ps.idproduto IS NOT NULL))
		 )
		 OR
		 (-- Revendedor com segmento
			(r.idsegmento > 0 AND r.idsegmento != 99) AND
			(ps.qt_seg_igual > 0)
		 )
		 OR
		 (-- Revendedor com todos os segmentos
			(r.idsegmento = 99 AND ps.idproduto IS NOT NULL)
		 )
	 )
 THEN 1 
 ELSE 0
	END AS PodeSolicitarNovoCodigoDeBarras,
	CASE WHEN pr1.idproduto IS NOT NULL THEN 1 ELSE 0 END AS produto_do_revendedor
FROM produto p
LEFT JOIN revendedor r ON r.idrevendedor = :p_idrevendedor
LEFT JOIN produtorevendedor pr ON pr.idrevendedor = r.idrevendedor AND pr.idproduto = p.idproduto AND pr.StatusProdutoRevendedor = 1
LEFT JOIN produtorevendedor pr1 ON pr1.idrevendedor = r.idrevendedor AND pr1.idproduto = p.idproduto
LEFT JOIN (SELECT DISTINCT idproduto FROM OPTATIVAPRODUTO WHERE statusoptativaproduto = 1) op ON op.idproduto = p.idproduto
LEFT JOIN (SELECT idproduto FROM OPTATIVAPRODUTO GROUP BY idproduto) op1 ON op1.idproduto = p.idproduto
LEFT JOIN (SELECT DISTINCT idrevendedor FROM OPTATIVAREVENDEDOR WHERE statusoptativarevendedor = 1) opr ON opr.idrevendedor = r.idrevendedor
LEFT JOIN (SELECT DISTINCT idrevendedor FROM OPTATIVAREVENDEDOR) opr1 ON opr1.idrevendedor = r.idrevendedor
LEFT JOIN (SELECT cpgi.idproduto
						 , COUNT(CASE WHEN cp.ativo = '1' THEN 'S' ELSE NULL END) AS qt_ativo
					FROM combopromocaogrupoitem cpgi
					LEFT JOIN combopromocaogrupo cpg ON cpgi.idcombopromocaogrupo = cpg.idcombopromocaogrupo
					LEFT JOIN combopromocao cp ON cpg.idcombopromocao = cp.idcombopromocao
				 GROUP BY cpgi.idproduto) cpgi ON cpgi.idproduto = p.idproduto
LEFT JOIN (SELECT idproduto
						 , COUNT(DECODE(r.idsegmento,ps.idsegmento,NULL,'S')) AS qt_seg_dif
						 , COUNT(DECODE(r.idsegmento,ps.idsegmento,'S',NULL)) AS qt_seg_igual
					FROM produtosegmento ps
					LEFT JOIN revendedor r ON r.idrevendedor = :p_idrevendedor
			GROUP BY idproduto) ps ON ps.idproduto = p.idproduto
LEFT JOIN (SELECT codigo, MAX(valor) valor FROM parametro GROUP BY codigo) pm ON pm.codigo = 'ID_ESTRUT_MERC_REGIONAIS'
LEFT JOIN marca m ON m.idmarca = p.idmarca
LEFT JOIN (SELECT ep.IdProduto
						 , ep.CodigoBarras
						 , ROW_NUMBER() OVER (PARTITION BY ep.IdProduto ORDER BY ep.IdEmbalagemProduto) AS nr_seq_ep
					FROM EmbalagemProduto ep
				 WHERE ep.Principal = 1
			 ) ep ON p.IdProduto = ep.IdProduto AND ep.nr_seq_ep = 1
LEFT JOIN EstruturaMercadologica em1 ON em1.IdEstruturaMercadologica = p.idNivel1EstrMerc
LEFT JOIN EstruturaMercadologica em2 ON em2.IdEstruturaMercadologica = p.idNivel2EstrMerc
LEFT JOIN (SELECT DISTINCT op.idproduto, opr.idrevendedor
					FROM solucaooptativa so
				 INNER JOIN optativaproduto op ON so.idsolucaooptativa = op.idsolucaooptativa
				 INNER JOIN optativarevendedor opr ON so.idsolucaooptativa = opr.idsolucaooptativa
				 WHERE so.statussolucaooptativa = 1
					 AND opr.statusoptativarevendedor = 1
			 ) so ON so.idproduto = p.idproduto AND so.idrevendedor = r.idrevendedor
WHERE r.idrevendedor IS NULL
OR ( (p.Ativo=1) AND
		(TRIM(:p_codigo) IS NULL OR TO_CHAR(p.IdProduto) LIKE '%'||TRIM(:p_codigo)||'%') AND
		(TRIM(:p_descricao) IS NULL OR p.DescricaoProduto LIKE '%'||TRIM(:p_descricao)||'%') AND
		(TRIM(:p_marca) IS NULL OR m.NomeMarca LIKE '%'||TRIM(:p_marca)||'%') AND
		(TRIM(:p_codigobarras) IS NULL OR ep.CodigoBarras LIKE '%'||TRIM(:p_codigobarras)||'%') AND
		(TRIM(:p_categoria) IS NULL OR em1.NomeNivel LIKE '%'||TRIM(:p_categoria)||'%') AND
		(TRIM(:p_subcategoria) IS NULL OR em2.NomeNivel LIKE '%'||TRIM(:p_subcategoria)||'%') AND
		(CASE WHEN :p_meucatalogo = 1
					THEN CASE WHEN r.idsegmento > 0
										THEN CASE WHEN pr.idproduto IS NOT NULL
																OR so.idproduto IS NOT NULL 
																OR ps.qt_seg_igual > 0
																OR (r.IdSegmento = 99 AND ps.idproduto IS NOT NULL)
																OR cpgi.qt_ativo > 0
															THEN 1
															ELSE 0
													END
										ELSE  CASE WHEN pr.idproduto IS NOT NULL
																 OR cpgi.qt_ativo > 0
																 OR (ps.idproduto IS NOT NULL OR p.IdNivel1EstrMerc != pm.valor)
															THEN 1
															ELSE 0
													END
								END
					ELSE 1
			END = 1) AND
				(pr.idproduto IS NOT NULL OR NVL(p.foramix,'0') <> '1') AND
			( (p.IdNivel1EstrMerc != pm.valor) OR
				(p.IdNivel1EstrMerc = pm.valor AND pr.idproduto IS NOT NULL) OR
				(TRIM(:p_codigobarras) IS NOT NULL AND ep.CodigoBarras = TRIM(:p_codigobarras))
			)
	)
ORDER BY p.idproduto
`

	rows, err := r.Db.Query(query, sql.Named("p_idrevendedor", p_idrevendedor))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dealers []*entity.Dealer
	for rows.Next() {
		var product entity.Dealer
		if err := rows.Scan(&product.Codigo, &product.Descricao, &product.Marca, &product.CodigodeBarra, &product.Categoria, &product.Subcategoria, &product.IdRevendedor, &product.PodeSolicitarPermissaoDeVendas, &product.PodeSolicitarNovoCodigoDeBarras, &product.ProdutoDoRevendedor); err != nil {
			return nil, err
		}
		dealers = append(dealers, &product)
	}
	return dealers, nil
}
