package router

import (
	"golang-train/app/repository"
	"golang-train/app/service"
	"golang-train/config"
	"golang-train/helper"
	"golang-train/middleware"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, db *mongo.Database, cfg *config.Config) {
	// Initialize repositories with MongoDB database
	userRepo := repository.NewUserRepository(db)
	mahasiswaRepo := repository.NewMahasiswaRepository(db)
	pekerjaanRepo := repository.NewPekerjaanRepository(db)
	alumniRepo := repository.NewAlumniRepository(db)
	fotoRepo := repository.NewFotoRepository(db)             // Repositori baru
	sertifikatRepo := repository.NewSertifikatRepository(db) // Repositori baru

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg.JWTSecretKey, cfg.JWTExpirationHours)
	userService := service.NewUserService(db)
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepo)
	pekerjaanService := service.NewPekerjaanService(pekerjaanRepo)
	alumniService := service.NewAlumniService(alumniRepo)
	mediaService := service.NewMediaService(fotoRepo, sertifikatRepo) // Servis baru

	// Initialize helpers
	authHelper := helper.NewAuthHelper(authService)
	userHelper := helper.NewUserHelper(userService)
	mahasiswaHelper := helper.NewMahasiswaHelper(mahasiswaService)
	pekerjaanHelper := helper.NewPekerjaanHelper(pekerjaanService)
	alumniHelper := helper.NewAlumniHelper(alumniService)
	
	// MEMPERBAIKI TYPO: Menggunakan 'mediaService' bukan 'mediaHelper'
	mediaHelper := helper.NewMediaHelper(mediaService) // Helper baru

	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/register", authHelper.Register)
	auth.Post("/login", authHelper.Login)

	// Authenticated routes
	api.Use(middleware.AuthMiddleware(cfg.JWTSecretKey))

	// User routes
	users := api.Group("/users")
	users.Delete("/:id", userHelper.DeleteUser)
	users.Put("/restore/:id", middleware.RoleMiddleware("admin"), userHelper.RestoreUser)

	// Rute Media (Foto & Sertifikat) baru
	// Menggunakan parameter mediaType untuk membedakan antara 'foto' dan 'sertifikat'
	// Hanya mengizinkan 'foto' atau 'sertifikat' sebagai mediaType
	users.Post("/:id/:mediaType(foto|sertifikat)", mediaHelper.UploadMedia)
	users.Get("/:id/:mediaType(foto|sertifikat)", mediaHelper.GetMedia)

	// Mahasiswa routes
	mahasiswa := api.Group("/mahasiswa")
	mahasiswa.Get("/", mahasiswaHelper.GetAllMahasiswa)
	mahasiswa.Get("/:id", mahasiswaHelper.GetMahasiswaByID)
	mahasiswa.Post("/", middleware.RoleMiddleware("admin"), mahasiswaHelper.CreateMahasiswa)
	mahasiswa.Put("/:id", middleware.RoleMiddleware("admin"), mahasiswaHelper.UpdateMahasiswa)
	mahasiswa.Delete("/:id", middleware.RoleMiddleware("admin"), mahasiswaHelper.DeleteMahasiswa)

	// Pekerjaan routes
	pekerjaan := api.Group("/pekerjaan")
	pekerjaan.Get("/", pekerjaanHelper.GetAllPekerjaan)
	pekerjaan.Get("/deleted", pekerjaanHelper.GetAllPekerjaanDeleted)
	pekerjaan.Get("/:id", pekerjaanHelper.GetPekerjaanByID)
	pekerjaan.Post("/", middleware.RoleMiddleware("admin"), pekerjaanHelper.CreatePekerjaan)
	pekerjaan.Put("/:id", middleware.RoleMiddleware("admin"), pekerjaanHelper.UpdatePekerjaan)
	pekerjaan.Put("/restore/:id", middleware.RoleMiddleware("admin"), pekerjaanHelper.RestorePekerjaan)
	pekerjaan.Delete("/:id", middleware.RoleMiddleware("admin"), pekerjaanHelper.DeletePekerjaan)
	pekerjaan.Delete("/softdel/:id", middleware.RoleMiddleware("admin"), pekerjaanHelper.SoftDeletePekerjaan)

	// Alumni routes (were missing from router)
	alumni := api.Group("/alumni")
	alumni.Get("/", alumniHelper.GetAllAlumni)
	alumni.Get("/:id", alumniHelper.GetAlumniByID)
	alumni.Post("/", middleware.RoleMiddleware("admin"), alumniHelper.CreateAlumni)
	alumni.Put("/:id", middleware.RoleMiddleware("admin"), alumniHelper.UpdateAlumni)
	alumni.Delete("/:id", middleware.RoleMiddleware("admin"), alumniHelper.DeleteAlumni)
}
