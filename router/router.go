package router

import (
	"golang-train/app/repository"
	"golang-train/app/service"
	"golang-train/config"
	"golang-train/helper"
	"golang-train/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool, cfg *config.Config) {
	userRepo := repository.NewUserRepository(db)
	mahasiswaRepo := repository.NewMahasiswaRepository(db)
	pekerjaanRepo := repository.NewPekerjaanRepository(db)

	authService := service.NewAuthService(userRepo, cfg.JWTSecretKey, cfg.JWTExpirationHours)
	userService := service.NewUserService(db)
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepo)
	pekerjaanService := service.NewPekerjaanService(pekerjaanRepo)

	authHelper := helper.NewAuthHelper(authService)
	userHelper := helper.NewUserHelper(userService)
	mahasiswaHelper := helper.NewMahasiswaHelper(mahasiswaService)
	pekerjaanHelper := helper.NewPekerjaanHelper(pekerjaanService)

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", authHelper.Register)
	auth.Post("/login", authHelper.Login)

	api.Use(middleware.AuthMiddleware(cfg.JWTSecretKey))

	users := api.Group("/users")
	users.Delete("/:id", userHelper.DeleteUser)

	mahasiswa := api.Group("/mahasiswa")
	mahasiswa.Get("/", mahasiswaHelper.GetAllMahasiswa)
	mahasiswa.Get("/:id", mahasiswaHelper.GetMahasiswaByID)
	mahasiswa.Post("/", middleware.RoleMiddleware("admin"), mahasiswaHelper.CreateMahasiswa)
	mahasiswa.Put("/:id", middleware.RoleMiddleware("admin"), mahasiswaHelper.UpdateMahasiswa)
	mahasiswa.Delete("/:id", middleware.RoleMiddleware("admin"), mahasiswaHelper.DeleteMahasiswa)

	pekerjaan := api.Group("/pekerjaan")
	pekerjaan.Get("/", pekerjaanHelper.GetAllPekerjaan)
	pekerjaan.Get("/deleted", pekerjaanHelper.GetAllPekerjaanDeleted)
	pekerjaan.Get("/:id", pekerjaanHelper.GetPekerjaanByID)
	pekerjaan.Post("/", middleware.RoleMiddleware("admin"), pekerjaanHelper.CreatePekerjaan)
	pekerjaan.Put("/:id", middleware.RoleMiddleware("admin"), pekerjaanHelper.UpdatePekerjaan)
	pekerjaan.Put("/restore/:id", middleware.RoleMiddleware("admin"), pekerjaanHelper.RestorePekerjaan)       // Corrected Restore Route
	pekerjaan.Delete("/:id", middleware.RoleMiddleware("admin"), pekerjaanHelper.DeletePekerjaan)             // Hard Delete Route
	pekerjaan.Delete("/softdel/:id", middleware.RoleMiddleware("admin"), pekerjaanHelper.SoftDeletePekerjaan) // Soft Delete Route
}

