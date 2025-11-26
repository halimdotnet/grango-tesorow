package app

import "github.com/halimdotnet/grango-tesorow/internal/modules/accounting"

func (a *App) modules() {
	accountingModule := accounting.NewModule(a.Pg, a.Router, a.Logger)
	accountingModule.Register()
}
