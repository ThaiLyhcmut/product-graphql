package resolver

import controller "ThaiLy/graph/controllers"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	accountController controller.AccountController
	categoryResolver  controller.CategoryController
}
