package routes

import (
	"go-fiber-starter/app/controllers"
	"go-fiber-starter/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// PUBLIC ROUTES
	// Root route renders the index page (previously only available at /import-demo)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/import-demo", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	apiRoute := app.Group("/api", middlewares.JwtMiddleware)

	// AUTH ROUTES
	auth := apiRoute.Group("/auth")
	authController := new(controllers.AuthController)
	auth.Post("/login", authController.Login)
	auth.Post("/register", authController.Register)

	// IMPORT ROUTES
	importGroup := apiRoute.Group("/import")
	importController := new(controllers.ImportController)
	importGroup.Get("/logs", importController.ImportLogAll)
	importGroup.Get("/logs/:guid", importController.ImportLogDetail)
	importGroup.Post("/products", importController.ImportProducts)
	importGroup.Get("/products/stream", importController.ImportProductsStream)

	// USER ROUTES
	userGroup := apiRoute.Group("/user")
	userController := new(controllers.UserController)
	userGroup.Get("/", userController.List)
	userGroup.Post("/", userController.Add)
	userGroup.Put("/:guid", userController.Update)
	userGroup.Delete("/:guid", userController.Delete)

	// PRODUCT ROUTES
	productGroup := apiRoute.Group("/product")
	productController := new(controllers.ProductController)
	productGroup.Get("/", productController.List)
	productGroup.Get("/:guid", productController.Detail)
	productGroup.Post("/", productController.Add)
	productGroup.Put("/:guid", productController.Update)
	productGroup.Delete("/:guid", productController.Delete)

	// NEWS ROUTES
	newsGroup := apiRoute.Group("/news")
	newsController := new(controllers.NewsController)
	newsGroup.Get("/", newsController.List)
	newsGroup.Get("/:guid", newsController.Detail)
	newsGroup.Post("/", newsController.Add)
	newsGroup.Put("/:guid", newsController.Update)
	newsGroup.Delete("/:guid", newsController.Delete)

	// UPLOAD ROUTES
	uploadController := new(controllers.UploadController)
	apiRoute.Post("/cdn", uploadController.CDN)

	// ADDITIONAL ENTITY ROUTES

	// ADMISSION MODE
	admissionModeGroup := apiRoute.Group("/admission-mode")
	admissionModeController := new(controllers.AdmissionModeController)
	admissionModeGroup.Get("/", admissionModeController.List)
	admissionModeGroup.Get("/:id", admissionModeController.Detail)
	admissionModeGroup.Post("/", admissionModeController.Add)
	admissionModeGroup.Put("/:id", admissionModeController.Update)
	admissionModeGroup.Delete("/:id", admissionModeController.Delete)

	// ASSISTANCE CATEGORY
	assistanceCatGroup := apiRoute.Group("/assistance-category")
	assistanceCatController := new(controllers.AssistanceCategoryController)
	assistanceCatGroup.Get("/", assistanceCatController.List)
	assistanceCatGroup.Get("/:id", assistanceCatController.Detail)
	assistanceCatGroup.Post("/", assistanceCatController.Add)
	assistanceCatGroup.Put("/:id", assistanceCatController.Update)
	assistanceCatGroup.Delete("/:id", assistanceCatController.Delete)

	// ASSISTANCE TYPE
	assistanceTypeGroup := apiRoute.Group("/assistance-type")
	assistanceTypeController := new(controllers.AssistanceTypeController)
	assistanceTypeGroup.Get("/", assistanceTypeController.List)
	assistanceTypeGroup.Get("/:id", assistanceTypeController.Detail)
	assistanceTypeGroup.Post("/", assistanceTypeController.Add)
	assistanceTypeGroup.Put("/:id", assistanceTypeController.Update)
	assistanceTypeGroup.Delete("/:id", assistanceTypeController.Delete)

	// ASSISTED THROUGH
	assistedThroughGroup := apiRoute.Group("/assisted-through")
	assistedThroughController := new(controllers.AssistedThroughController)
	assistedThroughGroup.Get("/", assistedThroughController.List)
	assistedThroughGroup.Get("/:id", assistedThroughController.Detail)
	assistedThroughGroup.Post("/", assistedThroughController.Add)
	assistedThroughGroup.Put("/:id", assistedThroughController.Update)
	assistedThroughGroup.Delete("/:id", assistedThroughController.Delete)

	// PROGRAM
	programGroup := apiRoute.Group("/program")
	programController := new(controllers.ProgramController)
	programGroup.Get("/", programController.List)
	programGroup.Get("/:id", programController.Detail)
	programGroup.Post("/", programController.Add)
	programGroup.Put("/:id", programController.Update)
	programGroup.Delete("/:id", programController.Delete)

	// CLIENT
	clientGroup := apiRoute.Group("/client")
	clientController := new(controllers.ClientController)
	clientGroup.Get("/", clientController.List)
	clientGroup.Get("/:guid", clientController.Detail)
	clientGroup.Post("/", clientController.Add)
	clientGroup.Put("/:guid", clientController.Update)
	clientGroup.Delete("/:guid", clientController.Delete)

	// BENEFICIARY
	beneficiaryGroup := apiRoute.Group("/beneficiary")
	beneficiaryController := new(controllers.BeneficiaryController)
	beneficiaryGroup.Get("/", beneficiaryController.List)
	beneficiaryGroup.Get("/:guid", beneficiaryController.Detail)
	beneficiaryGroup.Post("/", beneficiaryController.Add)
	beneficiaryGroup.Put("/:guid", beneficiaryController.Update)
	beneficiaryGroup.Delete("/:guid", beneficiaryController.Delete)

	// CLIENT BENEFICIARY
	cbGroup := apiRoute.Group("/client-beneficiary")
	cbController := new(controllers.ClientBeneficiaryController)
	cbGroup.Get("/", cbController.List)
	cbGroup.Post("/", cbController.Add)
	cbGroup.Delete("/:client_id/:beneficiary_id", cbController.Delete)

	// TRANSACTION
	transactionGroup := apiRoute.Group("/transaction")
	transactionController := new(controllers.TransactionController)
	transactionGroup.Get("/", transactionController.List)
	transactionGroup.Get("/:id", transactionController.Detail)
	transactionGroup.Post("/", transactionController.Add)
	transactionGroup.Put("/:id", transactionController.Update)
	transactionGroup.Delete("/:id", transactionController.Delete)

	// ASSISTANCE
	assistanceGroup := apiRoute.Group("/assistance")
	assistanceController := new(controllers.AssistanceController)
	assistanceGroup.Get("/", assistanceController.List)
	assistanceGroup.Get("/:id", assistanceController.Detail)
	assistanceGroup.Post("/", assistanceController.Add)
	assistanceGroup.Put("/:id", assistanceController.Update)
	assistanceGroup.Delete("/:id", assistanceController.Delete)
}
