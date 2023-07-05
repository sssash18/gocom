package orders

type Repository interface{
	Save(*Order) error
	ByID(string) (*Order,error)
}