package handler

import (
	"net/http"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/service"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/hxxp"
)

type AccountClassificationHandler struct {
	route       *hxxp.Router
	accountType service.AccountClassificationService
}

func NewAccountClassificationHandler(route *hxxp.Router, accountType service.AccountClassificationService) *AccountClassificationHandler {
	return &AccountClassificationHandler{
		route:       route,
		accountType: accountType,
	}
}

func (h *AccountClassificationHandler) RegisterRoutes() {
	h.route.Get("/api/v1/accounting/account-type", h.listAccountType)
}

func (h *AccountClassificationHandler) listAccountType(ctx *hxxp.Context) {
	list, err := h.accountType.ListAccountType(ctx.Ctx)
	if err != nil {
		ctx.Response(http.StatusInternalServerError, hxxp.Response{
			Error:   true,
			Message: "Failed to list account classifications",
		})
		return
	}

	ctx.Response(http.StatusOK, hxxp.Response{
		Error:   false,
		Message: "Success",
		Data:    list,
	})
}
