package restful

import (
	"fmt"

	"github.com/gideonlewis/e-commerce-product-server/pkg/datatypes"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	router *gin.Engine

	// Handler
	csrv *CategoryHandler
	psrv *ProductHandler
}

func NewAPIHandler(
	categoryHandler *CategoryHandler,
	producthandler *ProductHandler,
) *APIHandler {
	return &APIHandler{
		csrv: categoryHandler,
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
		h.router.Use()
		return nil
	}
}

func (h *APIHandler) Start(port int, options ...Options) {
	h.router = gin.New()
	pprof.Register(h.router)
	// global middleware
	h.registerGlobalMiddleware()

	// start with options
	for _, opt := range options {
		opt(h)
	}
}

func (h *APIHandler) registerGlobalMiddleware() {
	h.router.Use(gin.Recovery())
	h.router.Use(gin.Logger())
}

func (h *APIHandler) RegisterAppRouter(port int) Options {
	return func(a *APIHandler) error {
		h.registerAppRouter()
		return h.router.Run(fmt.Sprintf(":%d", port))
	}
}

func (h *APIHandler) registerAppRouter() {
	api := h.router.Group("/api")

	v1 := api.Group("/v1")

	h.registerCategorysRoutes(v1)
	h.registerProductRoutes(v1)
}

func (h *APIHandler) registerCategorysRoutes(group *gin.RouterGroup) {
	category := group.Group("/categories")
	category.GET("", h.csrv.GetList())
}

func (h *APIHandler) registerProductRoutes(group *gin.RouterGroup) {
	product := group.Group("/products")
	product.GET("", h.psrv.GetList())
	product.GET("/:id", h.psrv.GetByID())
}

func (h *APIHandler) RegisterCmsRouter(port int) Options {
	return func(a *APIHandler) error {
		h.registerCmsRouter()
		return h.router.Run(fmt.Sprintf(":%d", port))
	}
}

func (h *APIHandler) registerCmsRouter() {
	api := h.router.Group("/cms/api")
	v1 := api.Group("/v1")

	h.registerCmsCategorysRoutes(v1)
	h.registerCmsProductRoutes(v1)
}

func (h *APIHandler) registerCmsCategorysRoutes(group *gin.RouterGroup) {
	category := group.Group("/categories")
	category.POST("", h.csrv.Create())
}

func (h *APIHandler) registerCmsProductRoutes(group *gin.RouterGroup) {
	product := group.Group("/products")
	product.POST("", h.psrv.Create())
}
