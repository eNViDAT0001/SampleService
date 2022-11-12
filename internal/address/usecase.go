package internal

//var (
//	getSSTLogger = func(ctx context.Context) *logrus.Entry {
//		return sdk.Logger().
//			WithField("requestTraceID", common.GetRequestTraceId(ctx))
//	}
//)

type AddressUC interface {
	//GetByID(
	//	ctx context.Context,
	//	corpID uint, ID uint,
	//) (*entities.Address, error)
}

type addressUC struct {
	addressSto AddressRepository
}

func NewAddressUC(
	addressSto AddressRepository,
) AddressUC {
	return &addressUC{
		addressSto: addressSto}
}
