package internal

//var (
//	getSSTLogger = func(ctx context.Context) *logrus.Entry {
//		return sdk.Logger().
//			WithField("requestTraceID", common.GetRequestTraceId(ctx))
//	}
//)

type DistrictUC interface {
	//GetByID(
	//	ctx context.Context,
	//	corpID uint, ID uint,
	//) (*entities.Address, error)
}

type districtUC struct {
	districtSto DistrictRepository
}

func NewDistrictUC(
	districtSto DistrictRepository,
) DistrictUC {
	return &districtUC{
		districtSto: districtSto}
}
