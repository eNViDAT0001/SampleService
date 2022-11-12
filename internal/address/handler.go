package internal

//var (
//	getHandlerLogger = func(ctx context.Context) *logrus.Entry {
//		return sdk.Logger().
//			WithField("requestTraceID", common.GetRequestTraceId(ctx))
//	}
//
//	wViper = wrapviper.GetViper()
//)

type AddressHandler interface {
	//GetDetailByID() func(*gin.Context)
}

type addressHandler struct {
	addressUC AddressUC
}

func NewAddressHandler(
	addressUC AddressUC,
) AddressHandler {
	return &addressHandler{addressUC}
}
