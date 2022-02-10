package repositories

type Data struct {
	Id    int64
	Title string
	Type  string
}

type StockRepository interface {
	FindByID(ID int) (*Data, error)
}
