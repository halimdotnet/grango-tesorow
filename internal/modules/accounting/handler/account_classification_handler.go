package handler

import (
	"net/http"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/service"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/hxxp"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/hxxp/middleware"
)

type AccountClassificationHandler struct {
	route       *hxxp.Router
	accountType service.AccountClassificationService
}

func (h *AccountClassificationHandler) RegisterRoutes() {
	h.route.Group("/api/v1/accounting", func(accounting *hxxp.Router) {
		accounting.Use(middleware.BearerAuth)

		accounting.Get("/account-type", h.listAccountType)

		accounting.Get("/category", h.ListCategory)
		accounting.Get("/category/{code}", h.FindCategory)
	})
}

func (h *AccountClassificationHandler) listAccountType(ctx *hxxp.Context) {
	list, err := h.accountType.ListAccountType(ctx.Ctx)
	if err != nil {
		ctx.Response(http.StatusInternalServerError, hxxp.Response{
			Error:   true,
			Message: "Failed to list account type",
		})
		return
	}

	if len(list) == 0 {
		ctx.Response(http.StatusNotFound, hxxp.Response{
			Message: "No account type found",
		})
		return
	}

	ctx.Response(http.StatusOK, hxxp.Response{
		Message: "Success",
		Data:    list,
	})
}

func (h *AccountClassificationHandler) ListCategory(ctx *hxxp.Context) {
	list, err := h.accountType.ListCategory(ctx.Ctx)
	if err != nil {
		ctx.Response(http.StatusInternalServerError, hxxp.Response{
			Error:   true,
			Message: "Failed to list account category",
		})
		return
	}

	if len(list) == 0 {
		ctx.Response(http.StatusNotFound, hxxp.Response{
			Message: "No account category found",
		})
		return
	}

	ctx.Response(http.StatusOK, hxxp.Response{
		Message: "Success",
		Data:    list,
	})
}

func (h *AccountClassificationHandler) FindCategory(ctx *hxxp.Context) {
	code := ctx.Param("code")
	if code == "" {
		ctx.Response(http.StatusBadRequest, hxxp.Response{
			Message: "Invalid code",
		})
		return
	}

	data, err := h.accountType.FindCategory(ctx.Ctx, code)
	if err != nil {
		ctx.Response(http.StatusInternalServerError, hxxp.Response{
			Error:   true,
			Message: "Failed to find account category",
		})
		return
	}

	if data == nil {
		ctx.Response(http.StatusNotFound, hxxp.Response{
			Message: "No account category found",
		})
		return
	}

	ctx.Response(http.StatusOK, hxxp.Response{
		Message: "Success",
		Data:    data,
	})
}

func NewAccountClassificationHandler(route *hxxp.Router, accountType service.AccountClassificationService) *AccountClassificationHandler {
	return &AccountClassificationHandler{
		route:       route,
		accountType: accountType,
	}
}
