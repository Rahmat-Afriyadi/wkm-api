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
	conn                                  *sql.DB = config.GetConnection()
	connUser, sqlConnUser                         = config.GetConnectionUser()
	connGormAsuransi, sqlConnGormAsuransi         = config.NewAsuransiGorm()

	tr3Repository repository.Tr3Repository = repository.NewTr3nRepository(conn)
	tr3Service    service.Tr3Service       = service.NewTr3Service(tr3Repository)
	tr3Controller controller.Tr3Controller = controller.NewTr3Controller(tr3Service)

	kerjaRepository repository.KerjaRepository = repository.NewKerjanRepository(conn)
	kerjaService    service.KerjaService       = service.NewKerjaService(kerjaRepository)
	kerjaController controller.KerjaController = controller.NewKerjaController(kerjaService)

	leasRepository repository.LeasRepository = repository.NewLeasnRepository(conn)
	leasService    service.LeasService       = service.NewLeasService(leasRepository)
	leasController controller.LeasController = controller.NewLeasController(leasService)

	userRepository repository.UserRepository = repository.NewUserRepository(connUser)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService)

	asuransiRepository repository.AsuransiRepository = repository.NewAsuransiRepository(connGormAsuransi)
	asuransiService    service.AsuransiService       = service.NewAsuransiService(asuransiRepository, userRepository)
	asuransiController controller.AsuransiController = controller.NewAsuransiController(asuransiService)

	approvalRepository repository.ApprovalRepository = repository.NewApprovalRepository(connGormAsuransi)
	approvalService    service.ApprovalService       = service.NewApprovalService(approvalRepository)
	approvalController controller.ApprovalController = controller.NewApprovalController(approvalService)

	kodeposRepository repository.KodeposRepository = repository.NewKodeposRepository(connGormAsuransi)
	kodeposService    service.KodeposService       = service.NewKodeposService(kodeposRepository)
	kodeposController controller.KodeposController = controller.NewKodeposController(kodeposService)

	dlrRepository repository.DlrRepository = repository.NewDlrRepository(connGormAsuransi)
	dlrService    service.DlrService       = service.NewDlrService(dlrRepository)
	dlrController controller.DlrController = controller.NewDlrController(dlrService)

	produkRepository repository.ProdukRepository = repository.NewProdukRepository(connGormAsuransi)
	produkService    service.ProdukService       = service.NewProdukService(produkRepository)
	produkController controller.ProdukController = controller.NewProdukController(produkService)

	otrRepository repository.OtrRepository = repository.NewOtrRepository(connGormAsuransi)
	otrService    service.OtrService       = service.NewOtrService(otrRepository)
	otrController controller.OtrController = controller.NewOtrController(otrService)
)

func main() {

	defer conn.Close()
	defer sqlConnUser.Close()
	defer sqlConnGormAsuransi.Close()

	app := fiber.New(fiber.Config{})

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
	auth.Post("/login", authController.SignInUserAsuransi)
	auth.Post("/refresh-token", authController.RefreshAccessTokenAsuransi)
	auth.Post("/reset-password", middleware.DeserializeUser, authController.ResetPassword)
	auth.Post("/logout", authController.LogoutUser)
	auth.Get("/generate-password", authController.GeneratePassword)

	app.Post("/export-data-wa-blast", middleware.DeserializeUser, tr3Controller.ExportDataWaBlast)
	app.Post("/edit-jenis-bayar", middleware.DeserializeUser, tr3Controller.EditJenisBayar)
	app.Get("/leas/master-data", middleware.DeserializeUser, leasController.MasterData)
	app.Get("/kerja/master-data", middleware.DeserializeUser, kerjaController.MasterData)

	app.Post("/approval/update", middleware.DeserializeUser, approvalController.Update)
	app.Get("/mokita/token", middleware.DeserializeUser, approvalController.MokitaToken)
	app.Post("/mokita/update/token", middleware.DeserializeUser, approvalController.MokitaUpdateToken)
	app.Get("/otr/mst-produk", middleware.DeserializeUser, otrController.OtrMstProduk)
	app.Get("/otr/otr-na", middleware.DeserializeUser, otrController.OtrMstNa)
	app.Post("/otr/create-otr", middleware.DeserializeUser, otrController.CreateOtr)

	app.Post("/asuransi/export-report-asuransi", middleware.DeserializeUser, asuransiController.ExportReportAsuransi)
	app.Get("/asuransi/master-data/:status", middleware.DeserializeUser, asuransiController.MasterData)
	app.Get("/asuransi/master-approval", middleware.DeserializeUser, asuransiController.ListApprovalTransaksi)
	app.Get("/asuransi/master-approval-count", middleware.DeserializeUser, asuransiController.ListApprovalTransaksiCount)
	app.Get("/asuransi/detail-approval/:idTrx", middleware.DeserializeUser, asuransiController.DetailApprovalTransaksi)
	app.Get("/asuransi/master-data-count/:status", middleware.DeserializeUser, asuransiController.MasterDataCount)
	app.Get("/asuransi/rekap-by-status-kduser", middleware.DeserializeUser, asuransiController.RekapByStatusKdUser)
	app.Get("/asuransi/master-data-rekap", middleware.DeserializeUser, asuransiController.MasterDataRekapTele)
	app.Get("/asuransi/master-alasan-pending", middleware.DeserializeUser, asuransiController.MasterAlasanPending)
	app.Get("/asuransi/master-alasan-tdk-berminat", middleware.DeserializeUser, asuransiController.MasterAlasanTdkBerminat)
	app.Get("/asuransi/rekap-by-status-tele", middleware.DeserializeUser, asuransiController.RekapByStatus)
	app.Get("/asuransi/rekap-by-status-leader-tele", middleware.DeserializeUser, asuransiController.RekapByStatusLt)
	app.Post("/asuransi/update", middleware.DeserializeUser, asuransiController.UpdateAsuransi)
	app.Post("/asuransi/update/berminat", asuransiController.UpdateAsuransiBerminat)
	app.Post("/asuransi/update/batal-bayar", asuransiController.UpdateAsuransiBatalBayar)
	app.Post("/asuransi/update-ambil-asuransi", middleware.DeserializeUser, asuransiController.UpdateAmbilAsuransi)
	app.Get("/asuransi/:no_msn", middleware.DeserializeUser, asuransiController.FindAsuransiByNoMsn)

	app.Get("/kodepos/master-data", middleware.DeserializeUser, kodeposController.MasterData)
	app.Get("/dealer/master-data", middleware.DeserializeUser, dlrController.MasterData)
	app.Get("/produk/master-data", middleware.DeserializeUser, produkController.MasterData)

	// a := asuransiRepository.RincianByAlasanPendingKdUser("2024-05-01", "2024-05-30")
	// fmt.Println("ini data yaa guys yaa ", a)
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte("DE6ED21B4E643161949DFCE42DABC")},
	// 	ErrorHandler: func(c *fiber.Ctx, err error) error {
	// 		return c.Status(401).JSON(fiber.Map{"error_description": "Token has Expired"})
	// 	},
	// }))

	app.Listen(":3001")
}
