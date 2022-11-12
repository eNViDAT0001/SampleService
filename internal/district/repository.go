package internal

type DistrictRepository interface {
	//GetById(ctx context.Context, corpId uint32, ids uint32) (entityPayroll entities.EmployeePayroll, err error)
}

type districtRepository struct {
}

func NewDistrictRepository() DistrictRepository {
	return &districtRepository{}
}
