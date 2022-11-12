package internal

type ProvinceRepository interface {
	//GetById(ctx context.Context, corpId uint32, ids uint32) (entityPayroll entities.EmployeePayroll, err error)
}

type provinceRepository struct {
}

func NewProvinceRepository() ProvinceRepository {
	return &provinceRepository{}
}
