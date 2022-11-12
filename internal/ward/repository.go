package internal

type WardRepository interface {
	//GetById(ctx context.Context, corpId uint32, ids uint32) (entityPayroll entities.EmployeePayroll, err error)
}

type wardRepository struct {
}

func NewWardRepository() WardRepository {
	return &wardRepository{}
}
