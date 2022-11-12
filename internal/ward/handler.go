package internal

//var (
//	getHandlerLogger = func(ctx context.Context) *logrus.Entry {
//		return sdk.Logger().
//			WithField("requestTraceID", common.GetRequestTraceId(ctx))
//	}
//
//	wViper = wrapviper.GetViper()
//)

type WardHandler interface {
	//GetDetailByID() func(*gin.Context)
}

type wardHandler struct {
	wardUC WardUC
}

func NewWardHandler(
	wardUC WardUC,
) WardHandler {
	return &wardHandler{wardUC: wardUC}
}
