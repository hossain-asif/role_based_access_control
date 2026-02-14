package router

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Router interface {
	Register(r chi.Router)
}

var DomainRegistries = []func(*gorm.DB, chi.Router){
	func(db *gorm.DB, router chi.Router) {
		RegisterRoutes(db, router).Register(router)
	},

	// Add new modules here:
	// role.RegisterRoutes,
}
