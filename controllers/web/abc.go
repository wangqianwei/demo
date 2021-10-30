package web

import "demo/controllers"

type AbcController struct {
	controllers.MainController
}


func A(a *AbcController) {
	//a.Ctx.Redirect();
}