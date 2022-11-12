package internal

//var (
//	getSSTLogger = func(ctx context.Context) *logrus.Entry {
//		return sdk.Logger().
//			WithField("requestTraceID", common.GetRequestTraceId(ctx))
//	}
//)

type ProvinceUC interface {
	//GetByID(
	//	ctx context.Context,
	//	corpID uint, ID uint,
	//) (*entities.Address, error)
}

type provinceUC struct {
	provinceSto ProvinceRepository
}

func NewProvinceUC(
	provinceSto ProvinceRepository,
) ProvinceRepository {
	return &provinceUC{
		provinceSto: provinceSto}
}
