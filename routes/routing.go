package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pioniry/controllers"
	"pioniry/helpers"
	"pioniry/middlewares"
)

type Routing struct {
	user              controllers.UserController
	userLogin         controllers.SignInController
	pegawai           controllers.PegawaiController
	pelatihanInternal controllers.PelatihanInternalController
	pelatihanUnit     controllers.PelatihanUnit
	dokumen           controllers.DokumenController
}

func (r Routing) GetRoutes() *echo.Echo {
	e := echo.New()

	//Setting Validator
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}

	//Middleware
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	//e.Use(middleware.Recover())

	//Login
	e.POST("/signin/", r.userLogin.SignIn)
	e.GET("/signout/", r.userLogin.SignOut)

	//Group ROUTE
	g := e.Group("/backend")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &helpers.Claims{},
		SigningKey:              []byte(helpers.GetJWTSecret()),
		TokenLookup:             "cookie:access-token", // "<source>:<name>"
		ErrorHandlerWithContext: helpers.JWTErrorChecker,
	}))
	//Middleware refresh token
	g.Use(middlewares.TokenRefresherMiddleware)

	//Pelatihan Internal
	g.POST("/pelatihan/internal/", r.pelatihanInternal.Create)
	g.PUT("/pelatihan/internal/:id/", r.pelatihanInternal.Update)

	g.POST("/pelatihan/internal/dokumen/", r.dokumen.Create)
	g.DELETE("/pelatihan/internal/dokumen/:id/", r.dokumen.Delete)

	//Pelatihan Unit
	g.POST("/pelatihan/unit/", r.pelatihanUnit.Create)

	//Pegawai
	g.GET("/pegawai/", r.pegawai.GetPegawai)
	g.GET("/pegawai/", r.pegawai.SearchPegawai)

	p := e.Group("/api/v1")
	p.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &helpers.Claims{},
		SigningKey:              []byte(helpers.GetJWTSecret()),
		TokenLookup:             "cookie:access-token", // "<source>:<name>"
		ErrorHandlerWithContext: helpers.JWTErrorChecker,
	}))

	//Middleware refresh token
	p.Use(middlewares.TokenRefresherMiddleware)

	//User Route
	p.GET("/user/:id/", r.user.GetUser)
	p.PUT("/user/:id/", r.user.Update)
	p.DELETE("/user/:id/", r.user.Delete)
	p.GET("/user/", r.user.GetAllUser)
	p.POST("/user/", r.user.Create)

	return e
}
