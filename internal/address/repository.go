package internal

type AddressRepository interface {
	//GetById(ctx context.Context, corpId uint32, ids uint32) (entityPayroll entities.EmployeePayroll, err error)
}

type addressRepository struct {
}

func NewAddressRepository() AddressRepository {
	return &addressRepository{}
}
