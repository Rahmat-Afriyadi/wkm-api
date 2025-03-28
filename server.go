package main

import (
	"fmt"
	"time"
	"wkm/config"
	"wkm/middleware"

	"wkm/controller"
	"wkm/repository"
	"wkm/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/robfig/cron/v3"
)

var (
	gormDBWkm, conn = config.GetConnection()
	gormECardplus, connECardplus = config.GetConnectionECardPlus()
	gormDBWkmTest, connTest = config.GetConnectionTest()
	connUser, sqlConnUser = config.GetConnectionUser()
	connGormAsuransi, sqlConnGormAsuransi = config.NewAsuransiGorm()

	eCardplusRepository repository.ECardplusRepository = repository.NewECardplusRepository(conn, gormDBWkm ,connECardplus, gormECardplus)
	eCardplusService    service.ECardplusService       = service.NewECardplusService(eCardplusRepository)
	eCardplusController controller.ECardplusController = controller.NewECardplusController(eCardplusService)
	
	tr3Repository repository.Tr3Repository = repository.NewTr3nRepository(conn, gormDBWkm,connGormAsuransi)
	tr3Service    service.Tr3Service       = service.NewTr3Service(tr3Repository)
	tr3Controller controller.Tr3Controller = controller.NewTr3Controller(tr3Service, eCardplusService)

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

	mstMtrRepository repository.MstMtrRepository = repository.NewMstMtrRepository(connGormAsuransi)
	mstMtrService    service.MstMtrService       = service.NewMstMtrService(mstMtrRepository)
	mstMtrController controller.MstMtrController = controller.NewMstMtrController(mstMtrService)

	merkRepository repository.MerkRepository = repository.NewMerkRepository(connGormAsuransi)
	merkService    service.MerkService       = service.NewMerkService(merkRepository)
	merkController controller.MerkController = controller.NewMerkController(merkService)

	vendorRepository repository.VendorRepository = repository.NewVendorRepository(connGormAsuransi)
	vendorService    service.VendorService       = service.NewVendorService(vendorRepository)
	vendorController controller.VendorController = controller.NewVendorController(vendorService)

	transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(connGormAsuransi)
	transaksiService    service.TransaksiService       = service.NewTransaksiService(transaksiRepository)
	transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService)

	tglMerahRepository repository.TglMerahRepository = repository.NewTglMerahRepository(gormDBWkm)
	tglMerahService    service.TglMerahService       = service.NewTglMerahService(tglMerahRepository)
	tglMerahController controller.TglMerahController = controller.NewTglMerahController(tglMerahService)

	extendBayarRepository repository.ExtendBayarRepository = repository.NewExtendBayarRepository(gormDBWkm)
	extendBayarService    service.ExtendBayarService       = service.NewExtendBayarService(extendBayarRepository)
	extendBayarController controller.ExtendBayarController = controller.NewExtendBayarController(extendBayarService)

	ticketSupportRepository repository.TicketSupportRepository = repository.NewTicketSupportRepository(conn)
	ticketSupportService    service.TicketSupportService       = service.NewTicketSupportService(ticketSupportRepository)
	ticketSupportController controller.TicketSupportController = controller.NewTicketSupportController(ticketSupportService)

	mstRepository repository.MstRepository = repository.NewMstRepository(gormDBWkm)
	mstService    service.MstService       = service.NewMstService(userRepository,mstRepository)
	mstController controller.MstController = controller.NewMstController(mstService)

	customerMtrRepository repository.CustomerMtrRepository = repository.NewCustomerMtrRepository(connTest,gormDBWkmTest,connGormAsuransi)
	customerMtrService    service.CustomerMtrService       = service.NewCustomerMtrService(customerMtrRepository)
	customerMtrController controller.CustomerMtrController = controller.NewCustomerMtrController(customerMtrService)

	asuransiPARepository repository.AsuransiPARepository = repository.NewAsuransiPARepository(conn)
	asuransiPAService    service.AsuransiPAService       = service.NewAsuransiPAService(asuransiPARepository)
	asuransiPAController controller.AsuransiPAController = controller.NewAsuransiPAController(asuransiPAService)
)

func main() {


	defer conn.Close()
	defer sqlConnUser.Close()
	defer sqlConnGormAsuransi.Close()

	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	defer scheduler.Stop()
	// "menit jam hari"
	_, err := scheduler.AddFunc("15 18 * * *", func() {
		otrRepository.ListApi()
		fmt.Println("running task ", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println("cron error ", err)
	}

	scheduler.Start()

	app := fiber.New(fiber.Config{})
	app.Static("/uploads", "./uploads")

	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${method} | ${path} | ${ip} | ${queryParams} |${latency} \n\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Indonesia/Jakarta",
		Done: func(c *fiber.Ctx, logString []byte) {
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, http://localhost:3000, http://192.168.70.17:3000",
		AllowHeaders: "Origin, Content-Type, Accept,  Access-Control-Allow-Origin, Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		panic("something went wrong")
	})

	auth := app.Group("/auth")
	auth.Post("/login", authController.SignInUserAsuransi)
	auth.Post("/login/admin", authController.SignInUserAsuransi)
	auth.Post("/refresh-token", authController.RefreshAccessTokenAsuransi)
	auth.Post("/reset-password", middleware.DeserializeUser, authController.ResetPassword)
	auth.Post("/logout", authController.LogoutUser)
	auth.Get("/generate-password", authController.GeneratePassword)

	app.Post("/export-data-wa-blast", middleware.DeserializeUser, tr3Controller.ExportDataWaBlast)
	app.Post("/search-no-msn-by-wa", tr3Controller.SearchNoMsnByWa)
	app.Post("/edit-jenis-bayar", middleware.DeserializeUser, tr3Controller.EditJenisBayar)
	app.Get("/leas/master-data", middleware.DeserializeUser, leasController.MasterData)
	app.Get("/kerja/master-data", middleware.DeserializeUser, kerjaController.MasterData)
	app.Get("/kerja/master-data/choices", middleware.DeserializeUser, kerjaController.MasterDataChoices)

	app.Post("/approval/update", middleware.DeserializeUser, approvalController.Update)
	app.Get("/mokita/token", middleware.DeserializeUser, approvalController.MokitaToken)
	app.Post("/mokita/update/token", middleware.DeserializeUser, approvalController.MokitaUpdateToken)

	app.Get("/otr/mst-produk", middleware.DeserializeUser, otrController.OtrMstProduk)
	app.Get("/otr/otr-na", middleware.DeserializeUser, otrController.OtrMstNa)
	app.Get("/otr/master-data", middleware.DeserializeUser, otrController.MasterData)
	app.Get("/otr/master-data-count", middleware.DeserializeUser, otrController.MasterDataCount)
	app.Get("/otr/detail-otr/:id", middleware.DeserializeUser, otrController.DetailOtr)
	app.Post("/otr/detail-otr-by-no-mtr", middleware.DeserializeUser, otrController.DetailOtrByNoMtr)
	app.Post("/otr/create-otr", middleware.DeserializeUser, otrController.CreateOtr)
	app.Post("/otr/update-otr", middleware.DeserializeUser, otrController.UpdateOtr)

	app.Get("/mst-mtr/master-data", middleware.DeserializeUser, mstMtrController.MasterData)
	app.Get("/mst-mtr/master-data-count", middleware.DeserializeUser, mstMtrController.MasterDataCount)
	app.Get("/mst-mtr/detail-mst-mtr/:id", middleware.DeserializeUser, mstMtrController.DetailMstMtr)
	app.Post("/mst-mtr/create-mst-mtr", middleware.DeserializeUser, mstMtrController.CreateMstMtr)
	app.Post("/mst-mtr/update-mst-mtr", middleware.DeserializeUser, mstMtrController.UpdateMstMtr)

	app.Get("/mst-user-ts", mstController.ListClientUser)
	app.Get("/mst-agama", mstController.MasterAgama)
	app.Get("/mst-alasan-void-konfirmasi", mstController.AlasanVoidKonfirmasi)
	app.Get("/mst-pendidikan", mstController.MasterPendidikan)
	app.Get("/mst-tujuan-pakai", mstController.MasterTujuPak)
	app.Get("/mst-keluar-bln", mstController.MasterKeluarBln)
	app.Get("/mst-aktivitas-jual", mstController.MasterAktivitasJual)
	app.Get("/mst-kodepos", middleware.DeserializeUser, kodeposController.MasterDataAll)
	app.Post("/mst-script/create", middleware.DeserializeUser, mstController.CreateScript)
	app.Post("/mst-script/update/:id", middleware.DeserializeUser, mstController.UpdateScript)
	app.Get("/mst-script/detail/:id", mstController.ViewScript)
	app.Get("/mst-script/all", mstController.ListAllScript)
	app.Get("/mst-script/active", mstController.MasterScript)
	app.Get("/mst-get-state/:id",middleware.DeserializeUser, mstController.GetState)
	app.Post("/mst-update-state",middleware.DeserializeUser, mstController.UpdateState)
	
	app.Get("/mst-alasan-tdk-membership/:tipe", middleware.DeserializeUser, mstController.MasterAlasanTdkMembership)
	app.Get("/mst-produk-membership", middleware.DeserializeUser, mstController.MasterProdukMembership)
	app.Get("/mst-promo-transfer", middleware.DeserializeUser, mstController.MasterPromoTransfer)
	app.Get("/mst-hobbies", middleware.DeserializeUser, mstController.MasterHobbies)

	app.Post("/asuransi/export-report-asuransi", middleware.DeserializeUser, asuransiController.ExportReportAsuransi)
	app.Post("/asuransi/export-report-asuransi-telesales", middleware.DeserializeUser, asuransiController.ExportReportAsuransiTele)
	app.Get("/asuransi/master-data/:status", middleware.DeserializeUser, asuransiController.MasterData)
	app.Get("/asuransi/master-produk", middleware.DeserializeUser, asuransiController.AsuransiMstProduk)
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
	app.Get("/kodepos/master-data-1", middleware.DeserializeUser, kodeposController.MasterData1)
	app.Get("/dealer/master-data", middleware.DeserializeUser, dlrController.MasterData)

	app.Get("/produk/master-data", middleware.DeserializeUser, produkController.MasterData)
	app.Get("/produk/master-data-count", middleware.DeserializeUser, produkController.MasterDataCount)
	app.Get("/produk/detail-produk/:id", middleware.DeserializeUser, produkController.DetailMstMtr)
	app.Post("/produk/create-produk", middleware.DeserializeUser, produkController.Create)
	app.Post("/produk/update-produk", middleware.DeserializeUser, produkController.Update)
	app.Post("/produk/upload-logo", middleware.DeserializeUser, produkController.UploadLogo)
	app.Delete("/produk/delete-manfaat/:id", middleware.DeserializeUser, produkController.DeleteManfaat)
	app.Delete("/produk/delete-syarat/:id", middleware.DeserializeUser, produkController.DeleteSyarat)
	app.Delete("/produk/delete-paket/:id", middleware.DeserializeUser, produkController.DeletePaket)

	app.Get("/transaksi/master-data", middleware.DeserializeUser, transaksiController.MasterData)
	app.Get("/transaksi/master-data-count", middleware.DeserializeUser, transaksiController.MasterDataCount)
	app.Get("/transaksi/detail-transaksi/:id", middleware.DeserializeUser, transaksiController.DetailMstMtr)
	app.Post("/transaksi/create-transaksi", middleware.DeserializeUser, transaksiController.Create)
	app.Post("/transaksi/update-transaksi", middleware.DeserializeUser, transaksiController.Update)
	app.Post("/transaksi/upload-dokumen", middleware.DeserializeUser, transaksiController.UploadDokumen)
	app.Post("/transaksi/import-excell", middleware.DeserializeUser, transaksiController.ImportExcell)

	app.Get("/vendor/master-data", middleware.DeserializeUser, vendorController.MasterData)
	app.Get("/vendor/master-data-count", middleware.DeserializeUser, vendorController.MasterDataCount)
	app.Get("/vendor/detail-vendor/:id", middleware.DeserializeUser, vendorController.DetailMstMtr)
	app.Post("/vendor/create-vendor", middleware.DeserializeUser, vendorController.Create)
	app.Post("/vendor/update-vendor", middleware.DeserializeUser, vendorController.Update)

	app.Get("/tgl-merah/master-data", middleware.DeserializeUser, tglMerahController.MasterData)
	app.Get("/tgl-merah/master-data-count", middleware.DeserializeUser, tglMerahController.MasterDataCount)
	app.Get("/tgl-merah/min-tgl-bayar", middleware.DeserializeUser, tglMerahController.MinTglBayar)
	app.Get("/tgl-merah/detail-tgl-merah/:id", middleware.DeserializeUser, tglMerahController.DetailTglMerah)
	app.Post("/tgl-merah/create-tgl-merah", middleware.DeserializeUser, tglMerahController.Create)
	app.Post("/tgl-merah/update-tgl-merah", middleware.DeserializeUser, tglMerahController.Update)
	app.Post("/tgl-merah/upload-excel", middleware.DeserializeUser, tglMerahController.UploadDokumen)
	app.Delete("/tgl-merah/delete/:id", middleware.DeserializeUser, tglMerahController.Delete)

	app.Get("/extend-bayar/master-data", middleware.DeserializeUser, extendBayarController.MasterData)
	app.Get("/extend-bayar/master-data-count", middleware.DeserializeUser, extendBayarController.MasterDataCount)
	app.Get("/extend-bayar/detail-extend-bayar/:id", middleware.DeserializeUser, extendBayarController.DetailExtendBayar)
	app.Post("/extend-bayar/create-extend-bayar", middleware.DeserializeUser, extendBayarController.Create)
	app.Post("/extend-bayar/update-extend-bayar", middleware.DeserializeUser, extendBayarController.UpdateFa)
	app.Post("/extend-bayar/upload-excel", middleware.DeserializeUser, extendBayarController.UploadDokumen)
	app.Post("/extend-bayar/update-extend-bayar/lf", middleware.DeserializeUser, extendBayarController.UpdateLf)
	app.Post("/extend-bayar/update-extend-bayar/approval-lf", middleware.DeserializeUser, extendBayarController.UpdateApprovalLf)
	app.Delete("/extend-bayar/delete/:id", middleware.DeserializeUser, extendBayarController.Delete)

	app.Post("/faktur-3/input-bayar", middleware.DeserializeUser, tr3Controller.UpdateInputBayar)
	app.Post("/ecardplus/input-bayar", middleware.DeserializeUser, eCardplusController.InputBayarEMembership)
	app.Post("/faktur-3/input-bayar/asuransi-pa", middleware.DeserializeUser, tr3Controller.UpdateInputBayarAsuransiPA)
	app.Post("/faktur-3/input-bayar/asuransi-mtr", middleware.DeserializeUser, tr3Controller.UpdateInputBayarAsuransiMtr)
	app.Post("/faktur-3/search/will-bayar", middleware.DeserializeUser, tr3Controller.WillBayar)
	app.Post("/export-data-pembayaran", middleware.DeserializeUser, tr3Controller.ExportPembayaranRenewal)

	app.Get("/merk/master-data/:jenisKendaraan", middleware.DeserializeUser, merkController.MasterData)

	app.Post("/data-renewal", middleware.DeserializeUser, tr3Controller.DataRenewal)
	app.Post("/export-data-renewal", tr3Controller.ExportDataRenewal)
	app.Post("/export-data-plat-plus", tr3Controller.ExportDataPlatinumPlus)

	app.Post("/ticket-support/add", middleware.DeserializeUser, ticketSupportController.CreateTicketSupport)
	app.Post("/ticket-support/edit/:no_ticket", middleware.DeserializeUser, ticketSupportController.EditTicketSupport)
	app.Get("/ticket-support/view/:no_ticket", middleware.DeserializeUser, ticketSupportController.ViewTicketSupport)
	app.Get("/ticket-support/user-ticket", middleware.DeserializeUser, ticketSupportController.ListTicketUser)
	app.Get("/ticket-support/ticket-queue", middleware.DeserializeUser, ticketSupportController.ListTicketQueue)
	app.Get("/ticket-support/it-ticket", middleware.DeserializeUser, ticketSupportController.ListTicketIT)
	app.Post("/export-rekap-ticket", middleware.DeserializeUser, ticketSupportController.ExportDataTiketSupport)
	app.Get("/mst-it-support", middleware.DeserializeUser, ticketSupportController.ListItSupport)

	app.Get("/customer-mtr/self-count", middleware.DeserializeUser, customerMtrController.SelfCount)
	app.Get("/customer-mtr/all-status-master-data", middleware.DeserializeUser, customerMtrController.AllStatusMasterData)
	app.Get("/customer-mtr/all-status-master-data-count", middleware.DeserializeUser, customerMtrController.AllStatusMasterDataCount)
	app.Get("/customer-mtr/master-data", middleware.DeserializeUser, customerMtrController.MasterData)
	app.Get("/customer-mtr/master-data-count", middleware.DeserializeUser, customerMtrController.MasterDataCount)
	app.Get("/customer-mtr/master-data-balikan", middleware.DeserializeUser, customerMtrController.MasterDataBalikan)
	app.Get("/customer-mtr/master-data-balikan-count", middleware.DeserializeUser, customerMtrController.MasterDataBalikanCount)
	app.Get("/customer-mtr/list-ambil-data", middleware.DeserializeUser, customerMtrController.ListAmbilData)
	app.Post("/customer-mtr/ambil-data", middleware.DeserializeUser, customerMtrController.AmbilData)
	app.Get("/customer-mtr/show/:no_msn", middleware.DeserializeUser, customerMtrController.Show)
	app.Get("/customer-mtr/show-all/:no_msn/:from", middleware.DeserializeUser, customerMtrController.ShowAll)
	app.Get("/customer-mtr/show-balikan/:no_msn", middleware.DeserializeUser, customerMtrController.ShowBalikan)
	app.Get("/customer-mtr/rekap-tele", middleware.DeserializeUser, customerMtrController.RekapTele)
	app.Post("/customer-mtr/update", middleware.DeserializeUser, customerMtrController.Update)
	app.Get("/customer-mtr/list-berminat-membership", middleware.DeserializeUser, customerMtrController.ListBerminatMembership)
	app.Get("/customer-mtr/list-berminat-asuransi-pa", middleware.DeserializeUser, customerMtrController.ListDataAsuransiPA)
	app.Get("/customer-mtr/list-berminat-asuransi-mtr", middleware.DeserializeUser, customerMtrController.ListDataAsuransiMtr)
	app.Post("/customer-mtr/export-rekap-tele", middleware.DeserializeUser, customerMtrController.ExportRekapTele)

	app.Post("asuransi-pa/create", middleware.DeserializeUser, asuransiPAController.CreateAsuransiPA)
	app.Post("asuransi-pa/update/:id", middleware.DeserializeUser, asuransiPAController.UpdateAsuransiPA)
	app.Post("/customer-mtr/update/oke", middleware.DeserializeUser, customerMtrController.UpdateOkeMembership)
	app.Listen(":3001")
}
