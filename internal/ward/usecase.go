package internal

//var (
//	getSSTLogger = func(ctx context.Context) *logrus.Entry {
//		return sdk.Logger().
//			WithField("requestTraceID", common.GetRequestTraceId(ctx))
//	}
//)

type WardUC interface {
	//GetByID(
	//	ctx context.Context,
	//	corpID uint, ID uint,
	//) (*entities.Address, error)
}

type wardUC struct {
	wardRepo WardRepository
}

func NewWardUC(
	wardRepo WardRepository,
) WardUC {
	return &wardUC{
		wardRepo: wardRepo}
}
