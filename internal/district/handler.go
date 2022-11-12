package internal

//var (
//	getHandlerLogger = func(ctx context.Context) *logrus.Entry {
//		return sdk.Logger().
//			WithField("requestTraceID", common.GetRequestTraceId(ctx))
//	}
//
//	wViper = wrapviper.GetViper()
//)

type DistrictHandler interface {
	//GetDetailByID() func(*gin.Context)
}

type districtHandler struct {
	districtUC DistrictUC
}

func NewDistrictHandler(
	districtUC DistrictUC,
) DistrictHandler {
	return &districtHandler{districtUC: districtUC}
}
