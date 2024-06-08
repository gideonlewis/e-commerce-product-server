package handler

import (
	"fmt"

	"github.com/gideonlewis/e-commerce-product-server/internal/adapters/middleware"
	"github.com/gideonlewis/e-commerce-product-server/internal/config"
	"github.com/gideonlewis/e-commerce-product-server/internal/pkg/datatypes"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	router *gin.Engine

	// Handler
	psrv *ProductHandler
}

func NewAPIHandler(
	producthandler *ProductHandler) *APIHandler {
	return &APIHandler{
		psrv: producthandler,
	}
}

type Options func(*APIHandler) error

func (h *APIHandler) WithSetMode(mode string) Options {
	return func(a *APIHandler) error {
		if mode != string(datatypes.StageTypeDev) {
			gin.SetMode(gin.ReleaseMode)
		}

		return nil
	}
}

func (h *APIHandler) WithMiddleware(middleware gin.HandlerFunc) Options {
	return func(a *APIHandler) error {
		h.router.Use(middleware)
		return nil
	}
}

func (h *APIHandler) Start(options ...Options) error {
	h.router = gin.New()
	pprof.Register(h.router)
	// global middleware
	h.registerGlobalMiddleware()

	// start with options
	for _, opt := range options {
		opt(h)
	}

	// declare router
	h.registerRouter()

	return h.router.Run(fmt.Sprintf(":%d", config.Server.Port))
}

func (h *APIHandler) registerRouter() {
	api := h.router.Group("/api")

	v1 := api.Group("/v1")
	v1.Use(middleware.TokenAuthMiddleware(config.Api.JWTSecret))

	h.registerProductRoutes(v1)
}

func (h *APIHandler) registerGlobalMiddleware() {
	h.router.Use(gin.Recovery())
	h.router.Use(gin.Logger())
}

func (h *APIHandler) registerProductRoutes(group *gin.RouterGroup) {
	_ = group.Group("/products")
}
