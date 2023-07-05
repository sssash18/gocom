package orders

type Repository interface {
	Save(*Order) error
	ByID(ID) (*Order, error)
}
