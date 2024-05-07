package main

import (
	"database/sql"
	"wkm/config"
	"wkm/middleware"

	"wkm/controller"
	"wkm/repository"
	"wkm/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	conn          *sql.DB                  = config.GetConnection()
	tr3Repository repository.Tr3Repository = repository.NewTr3nRepository(conn)
	tr3Service    service.Tr3Service       = service.NewTr3Service(tr3Repository)
	tr3Controller controller.Tr3Controller = controller.NewTr3Controller(tr3Service)

	kerjaRepository repository.KerjaRepository = repository.NewKerjanRepository(conn)
	kerjaService    service.KerjaService       = service.NewKerjaService(kerjaRepository)
	kerjaController controller.KerjaController = controller.NewKerjaController(kerjaService)

	leasRepository repository.LeasRepository = repository.NewLeasnRepository(conn)
	leasService    service.LeasService       = service.NewLeasService(leasRepository)
	leasController controller.LeasController = controller.NewLeasController(leasService)

	userRepository repository.UserRepository = repository.NewUserRepository(conn)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService)

	asuransiRepository repository.AsuransiRepository = repository.NewAsuransiRepository(conn)
	asuransiService    service.AsuransiService       = service.NewAsuransiService(asuransiRepository)
	asuransiController controller.AsuransiController = controller.NewAsuransiController(asuransiService)

	kodeposRepository repository.KodeposRepository = repository.NewKodeposRepository(conn)
	kodeposService    service.KodeposService       = service.NewKodeposService(kodeposRepository)
	kodeposController controller.KodeposController = controller.NewKodeposController(kodeposService)
)

func main() {
	defer conn.Close()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${method} | ${path} | ${ip} | ${queryParams} |${latency} | ${body}\n\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Indonesia/Jakarta",
		Done: func(c *fiber.Ctx, logString []byte) {
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	auth := app.Group("/auth")
	auth.Post("/login", authController.SignInUser)
	auth.Post("/refresh-token", authController.RefreshAccessToken)
	auth.Post("/logout", authController.LogoutUser)
	auth.Get("/generate-password", authController.GeneratePassword)

	app.Post("/export-data-wa-blast", middleware.DeserializeUser, tr3Controller.ExportDataWaBlast)
	app.Post("/edit-jenis-bayar", middleware.DeserializeUser, tr3Controller.EditJenisBayar)
	app.Get("/leas/master-data", middleware.DeserializeUser, leasController.MasterData)
	app.Get("/kerja/master-data", middleware.DeserializeUser, kerjaController.MasterData)

	app.Get("/asuransi/master-data", middleware.DeserializeUser, asuransiController.MasterData)
	app.Post("/asuransi/update", middleware.DeserializeUser, asuransiController.UpdateAsuransi)
	app.Get("/asuransi/:no_msn", middleware.DeserializeUser, asuransiController.FindAsuransiByNoMsn)

	app.Get("/kodepos/master-data", middleware.DeserializeUser, kodeposController.MasterData)
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte("DE6ED21B4E643161949DFCE42DABC")},
	// 	ErrorHandler: func(c *fiber.Ctx, err error) error {
	// 		return c.Status(401).JSON(fiber.Map{"error_description": "Token has Expired"})
	// 	},
	// }))

	app.Listen(":3001")
}
