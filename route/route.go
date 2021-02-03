package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-gorm-demo/controller"
	"net/http"
)

func InitRouter(app *iris.Application) {
	pathUrl := "/api"
	mvc.New(app.Party(pathUrl + "/user")).Handle(controller.NewUserController())
	//app.Use(middleware.GetJWT().Serve)
	mvc.New(app.Party(pathUrl + "/book")).Handle(controller.NewBookController())
}

func CrossAccess11(net http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		net.ServeHTTP(w, r)
	})
}

func CrossAccess(ctx iris.Context) {
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", "*")
}
