package main

import (
	"fmt"
	"log"

	"github.com/LutfiyaAinurrahmanP/boilerplate_fiber_restful_api/internal/config"
	"github.com/LutfiyaAinurrahmanP/boilerplate_fiber_restful_api/internal/handlers"
	"github.com/LutfiyaAinurrahmanP/boilerplate_fiber_restful_api/internal/models"
	"github.com/LutfiyaAinurrahmanP/boilerplate_fiber_restful_api/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/boilerplate_fiber_restful_api/internal/services"
	"github.com/LutfiyaAinurrahmanP/boilerplate_fiber_restful_api/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// ============================================
	// 1. LOAD CONFIGURATION
	// ============================================
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
	log.Println("✅ Configuration loaded successfully")

	// ============================================
	// 2. INITIALIZE DATABASE CONNECTION
	// ============================================
	db, err := database.NewMySQLConnection(cfg.GetDSN())
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	log.Println("✅ Database connection established")

	// ============================================
	// 3. RUN DATABASE MIGRATIONS
	// ============================================
	migrator := database.NewMigrator(db)

	modelsToMigrate := []interface{}{
		&models.User{}, // Model User dengan field role
		&models.TokenBlacklist{}, // Token blacklist untuk logout
	}

	if err := migrator.RunMigrations(modelsToMigrate...); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	// ============================================
	// 4. INITIALIZE DEPENDENCIES
	// ============================================
	log.Println("🔧 Initializing application dependencies...")

	// Repository Layer
	userRepo := repositories.NewUserRepository(db)
	tokenRepo := repositories.NewTokenRepository(db)

	// Service Layer
	authService := services.NewAuthService(userRepo, tokenRepo, cfg.JWTSecret)
	userService := services.NewUserService(userRepo)

	// Handler Layer
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	log.Println("✅ Dependencies initialized successfully")

	// ============================================
	// 5. SETUP FIBER APPLICATION
	// ============================================
	app := fiber.New(fiber.Config{
		AppName:           cfg.AppName,
		Prefork:           false,
		CaseSensitive:     false,
		StrictRouting:     false,
		ServerHeader:      cfg.AppName,
		EnablePrintRoutes: cfg.AppEnv == "development", // Print routes di development

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.Printf("❌ Error: %v", err)
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	// ============================================
	// 6. REGISTER GLOBAL MIDDLEWARES
	// ============================================
	log.Println("🔧 Registering global middlewares...")

	app.Use(recover.New(recover.Config{
		EnableStackTrace: cfg.AppEnv == "development",
	}))

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
		MaxAge:           3600,
	}))

	log.Println("✅ Middlewares registered successfully")

	// ============================================
	// 7. REGISTER ROUTES
	// ============================================
	log.Println("🔧 Registering application routes...")

	routeConfig := &RouteConfig{
		AuthHandler: authHandler,
		UserHandler: userHandler,
		JWTSecret:   cfg.JWTSecret,
		TokenRepo: tokenRepo,
	}

	SetupRoutes(app, routeConfig)

	log.Println("✅ Routes registered successfully")

	// ============================================
	// 8. START SERVER
	// ============================================
	port := fmt.Sprintf(":%s", cfg.AppPort)

	log.Println("========================================")
	log.Printf("🚀 Server starting on http://localhost%s", port)
	log.Printf("📝 Environment: %s", cfg.AppEnv)
	log.Println("========================================")
	log.Println("📚 Available Endpoints:")
	log.Println("")
	log.Println("   🔓 Public (No Auth):")
	log.Println("   - POST /auth/register")
	log.Println("   - POST /auth/login")
	log.Println("")
	log.Println("   🔒 Auth Required:")
	log.Println("   - POST /auth/logout")
	log.Println("")
	log.Println("   👑 Admin Only:")
	log.Println("   - GET    /admin/dashboard")
	log.Println("   - GET    /admin/user (list active users)")
	log.Println("   - GET    /admin/user/deleted (list deleted users)")
	log.Println("   - POST   /admin/user/create")
	log.Println("   - GET    /admin/user/:id")
	log.Println("   - PUT    /admin/user/update/:id")
	log.Println("   - DELETE /admin/user/:id (soft delete)")
	log.Println("   - DELETE /admin/user/permanent/:id (hard delete)")
	log.Println("   - POST   /admin/user/restore/:id (restore)")
	log.Println("")
	log.Println("   👤 User Only:")
	log.Println("   - GET /user/dashboard")
	log.Println("   - GET /user/profile")
	log.Println("   - PUT /user/profile/update")
	log.Println("========================================")
	log.Println("Press Ctrl+C to shutdown server")

	if err := app.Listen(port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
