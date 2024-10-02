package entity

type DealerRepository interface {
	ListAllDealers(p_idrevendedor int) ([]*Dealer, error)
}
