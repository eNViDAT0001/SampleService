package internal

//var (
//	getHandlerLogger = func(ctx context.Context) *logrus.Entry {
//		return sdk.Logger().
//			WithField("requestTraceID", common.GetRequestTraceId(ctx))
//	}
//
//	wViper = wrapviper.GetViper()
//)

type ProvinceHandler interface {
	//GetDetailByID() func(*gin.Context)
}

type provinceHandler struct {
	provinceUC ProvinceUC
}

func NewProvinceHandler(
	provinceUC ProvinceUC,
) ProvinceHandler {
	return &provinceHandler{provinceUC: provinceUC}
}
