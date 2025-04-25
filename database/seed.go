package database

import (
	"fmt"
	"log"
	"time"

	"manuk-pos-backend/helpers"
	"manuk-pos-backend/models/customer"
	"manuk-pos-backend/models/finance"
	"manuk-pos-backend/models/inventory"
	"manuk-pos-backend/models/orders"
	"manuk-pos-backend/models/promotion"
	"manuk-pos-backend/models/purchase"
	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/user"
	"manuk-pos-backend/models/vendor"

	"gorm.io/gorm"
)

func Seed() {
	SeedRoles(DB)
	SeedBranches(DB)
	SeedCategories(DB)
	SeedUsers(DB)
	SeedCustomers(DB)
	SeedSuppliers(DB)
	SeedProductsAndSuppliers(DB)
	SeedDiscounts(DB)
	SeedTaxes(DB)
	SeedLoans(DB)
	SeedCashDrawerData(DB)
	SeedFees(DB)
	SeedPromotions(DB)
	SeedGoodsReceiving(DB)
	SeedPurchaseOrders(DB)
	SeedTransactions(DB)
}

func SeedRoles(db *gorm.DB) {
	// Cek apakah sudah ada data Role
	var roleCount int64
	db.Model(&user.Role{}).Count(&roleCount)

	// Jika belum ada data, seed data awal
	if roleCount == 0 {
		// Seed data Role
		roles := []user.Role{
			{
				Name:        "Admin",
				Description: "Administrator with full access",
			},
			{
				Name:        "Staff",
				Description: "Staff with limited access",
			},
		}

		for _, role := range roles {
			if err := db.Create(&role).Error; err != nil {
				log.Fatalf("Gagal menambahkan data role: %v", err)
			}
		}

		log.Println("Data Role berhasil di-seed!")
	} else {
		log.Println("Data Role sudah ada, tidak perlu di-seed.")
	}
}

func SeedUsers(db *gorm.DB) {
	// Cek apakah sudah ada data User
	var userCount int64
	db.Model(&user.User{}).Count(&userCount)

	// Jika belum ada data, seed data awal
	if userCount == 0 {
		// Seed data User
		usersList := []user.User{
			{
				Username:   "admin",
				Password:   "admin123",
				Name:       "Administrator",
				Email:      "admin@manukpos.com",
				Phone:      "123456789",
				IsActive:   true,
				RoleID:     1,
				BranchID:   1,
				LastLogin:  nil,
				LoginCount: 0,
			},
			{
				Username:   "staff",
				Password:   "staff123",
				Name:       "Staff User",
				Email:      "staff@manukpos.com",
				Phone:      "987654321",
				IsActive:   true,
				RoleID:     2,
				BranchID:   1,
				LastLogin:  nil,
				LoginCount: 0,
			},
		}

		for _, user := range usersList {

			// Hash password
			hashedPassword, err := helpers.HashPassword(user.Password)
			if err != nil {
				log.Fatalf("Gagal meng-hash password: %v", err)
				return
			}

			// set hashed password
			user.Password = hashedPassword

			if err := db.Create(&user).Error; err != nil {
				log.Fatalf("Gagal menambahkan data user: %v", err)
			}
		}

		log.Println("Data User berhasil di-seed!")
	} else {
		log.Println("Data User sudah ada, tidak perlu di-seed.")
	}
}

func SeedBranches(db *gorm.DB) {
	// Cek apakah sudah ada data Branch
	var branchCount int64
	db.Model(&store.Branch{}).Count(&branchCount)

	// Jika belum ada data, seed data awal
	if branchCount == 0 {
		// Seed data Branch
		branches := []store.Branch{
			{
				Code:         "001",
				Name:         "Main Branch",
				Address:      "Jl. Main No. 1, Jakarta",
				Phone:        "021-12345678",
				Email:        "mainbranch@company.com",
				IsMainBranch: true,
				IsActive:     true,
			},
			{
				Code:         "002",
				Name:         "Branch 2",
				Address:      "Jl. Branch No. 2, Jakarta",
				Phone:        "021-87654321",
				Email:        "branch2@company.com",
				IsMainBranch: false,
				IsActive:     true,
			},
			{
				Code:         "003",
				Name:         "Branch 3",
				Address:      "Jl. Branch No. 3, Jakarta",
				Phone:        "021-11223344",
				Email:        "branch3@company.com",
				IsMainBranch: false,
				IsActive:     false, // Contoh branch yang non-aktif
			},
		}

		for _, branch := range branches {
			if err := db.Create(&branch).Error; err != nil {
				log.Fatalf("Gagal menambahkan data branch: %v", err)
			}
		}

		log.Println("Data Branch berhasil di-seed!")
	} else {
		log.Println("Data Branch sudah ada, tidak perlu di-seed.")
	}
}

func SeedCategories(db *gorm.DB) {
	// Cek apakah sudah ada data Category
	var categoryCount int64
	db.Model(&inventory.Category{}).Count(&categoryCount)

	// Jika belum ada data, seed data awal
	if categoryCount == 0 {
		// Seed data Category
		categories := []inventory.Category{
			{
				Name:        "Electronics",
				Code:        "ELEC",
				Description: "Electronics Products",
				Level:       1,
				Path:        "Electronics",
			},
			{
				Name:        "Phones",
				Code:        "PHONES",
				Description: "Smartphones and Accessories",
				Level:       2,
				Path:        "Electronics/Phones",
				ParentID:    intPtr(1), // ParentID refers to the Electronics category (ID = 1)
			},
			{
				Name:        "Laptops",
				Code:        "LAPTOPS",
				Description: "Laptop Computers",
				Level:       2,
				Path:        "Electronics/Laptops",
				ParentID:    intPtr(1), // ParentID refers to the Electronics category (ID = 1)
			},
			{
				Name:        "Fashion",
				Code:        "FASH",
				Description: "Clothing and Apparel",
				Level:       1,
				Path:        "Fashion",
			},
			{
				Name:        "Men's Clothing",
				Code:        "MENS_CLOTHING",
				Description: "Men's Apparel",
				Level:       2,
				Path:        "Fashion/Men's Clothing",
				ParentID:    intPtr(4), // ParentID refers to the Fashion category (ID = 4)
			},
		}

		for _, category := range categories {
			if err := db.Create(&category).Error; err != nil {
				log.Fatalf("Gagal menambahkan data kategori: %v", err)
			}
		}
		log.Println("Data Category berhasil di-seed!")
	} else {
		log.Println("Data Category sudah ada, tidak perlu di-seed.")
	}
}

func SeedCustomers(db *gorm.DB) {
	// Cek apakah sudah ada data Customer
	var count int64
	db.Model(&customer.Customer{}).Count(&count)

	// Jika belum ada data, seed data awal
	if count == 0 {
		customers := []customer.Customer{
			{
				Code:           "CUST0003",
				Name:           "Agus Wijaya",
				Phone:          "081278912345",
				Email:          "agus.wijaya@mail.com",
				Address:        "Jl. Melati No. 22",
				City:           "Surabaya",
				PostalCode:     "60231",
				Birthdate:      timePtr(time.Date(1988, 2, 10, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    1500000,
				CurrentBalance: 120000,
				IsActive:       true,
				Notes:          "Sering beli produk elektronik",
			},
			{
				Code:           "CUST0004",
				Name:           "Rina Kartika",
				Phone:          "082298765432",
				Email:          "rina.kartika@mail.com",
				Address:        "Jl. Anggrek No. 15",
				City:           "Yogyakarta",
				PostalCode:     "55281",
				Birthdate:      timePtr(time.Date(1993, 7, 12, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "premium",
				CreditLimit:    2500000,
				CurrentBalance: 80000,
				IsActive:       true,
				Notes:          "Langganan produk fashion",
			},
			{
				Code:           "CUST0005",
				Name:           "Dewi Lestari",
				Phone:          "081345678912",
				Email:          "dewi.lestari@example.com",
				Address:        "Jl. Kenanga No. 9",
				City:           "Semarang",
				PostalCode:     "50145",
				Birthdate:      timePtr(time.Date(1991, 11, 23, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    1000000,
				CurrentBalance: 25000,
				IsActive:       true,
				Notes:          "Pembelian rutin bulanan",
			},
			{
				Code:           "CUST0006",
				Name:           "Yusuf Hamzah",
				Phone:          "085234567899",
				Email:          "yusuf.hamzah@domain.com",
				Address:        "Jl. Cempaka No. 3",
				City:           "Medan",
				PostalCode:     "20112",
				Birthdate:      timePtr(time.Date(1980, 9, 5, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "vip",
				CreditLimit:    5000000,
				CurrentBalance: 500000,
				IsActive:       true,
				Notes:          "VIP, sering order besar",
			},
			{
				Code:           "CUST0007",
				Name:           "Sari Ningsih",
				Phone:          "087712345678",
				Email:          "sari.ningsih@mail.net",
				Address:        "Jl. Mawar No. 5",
				City:           "Bogor",
				PostalCode:     "16111",
				Birthdate:      timePtr(time.Date(1995, 3, 17, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    800000,
				CurrentBalance: 100000,
				IsActive:       true,
				Notes:          "Aktif di marketplace",
			},
			{
				Code:           "CUST0008",
				Name:           "Tommy Gunawan",
				Phone:          "083876543210",
				Email:          "tommy.gunawan@gmail.com",
				Address:        "Jl. Diponegoro No. 12",
				City:           "Bandung",
				PostalCode:     "40123",
				Birthdate:      timePtr(time.Date(1987, 6, 9, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "premium",
				CreditLimit:    2000000,
				CurrentBalance: 180000,
				IsActive:       true,
				Notes:          "Pembayaran cepat, loyal",
			},
			{
				Code:           "CUST0009",
				Name:           "Maria Fransiska",
				Phone:          "081234567891",
				Email:          "maria.fransiska@outlook.com",
				Address:        "Jl. Kartini No. 7",
				City:           "Solo",
				PostalCode:     "57100",
				Birthdate:      timePtr(time.Date(1992, 4, 25, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    1100000,
				CurrentBalance: 50000,
				IsActive:       true,
				Notes:          "Pembelian teratur",
			},
			{
				Code:           "CUST0010",
				Name:           "Budi Santoso",
				Phone:          "082112345678",
				Email:          "budi.santoso@mail.com",
				Address:        "Jl. Sisingamangaraja No. 30",
				City:           "Makassar",
				PostalCode:     "90123",
				Birthdate:      timePtr(time.Date(1989, 10, 3, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    1200000,
				CurrentBalance: 95000,
				IsActive:       true,
				Notes:          "Biasa order laptop & aksesoris",
			},
			{
				Code:           "CUST0011",
				Name:           "Fitri Handayani",
				Phone:          "085612345678",
				Email:          "fitri.handayani@domain.co.id",
				Address:        "Jl. Gajah Mada No. 18",
				City:           "Malang",
				PostalCode:     "65145",
				Birthdate:      timePtr(time.Date(1990, 12, 8, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    1000000,
				CurrentBalance: 50000,
				IsActive:       true,
				Notes:          "Fashion & aksesoris pelanggan",
			},
			{
				Code:           "CUST0012",
				Name:           "Fajar Maulana",
				Phone:          "081898765432",
				Email:          "fajar.maulana@domain.com",
				Address:        "Jl. Hasanuddin No. 21",
				City:           "Tangerang",
				PostalCode:     "15111",
				Birthdate:      timePtr(time.Date(1994, 1, 5, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "vip",
				CreditLimit:    3500000,
				CurrentBalance: 300000,
				IsActive:       true,
				Notes:          "Sering order dalam jumlah besar",
			},
			{
				Code:           "CUST0013",
				Name:           "Nadia Safira",
				Phone:          "081345678934",
				Email:          "nadia.safira@gmail.com",
				Address:        "Jl. Soekarno-Hatta No. 9",
				City:           "Palembang",
				PostalCode:     "30126",
				Birthdate:      timePtr(time.Date(1996, 9, 13, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    900000,
				CurrentBalance: 85000,
				IsActive:       true,
				Notes:          "Aktif di promo akhir bulan",
			},
			{
				Code:           "CUST0014",
				Name:           "Rizky Firmansyah",
				Phone:          "082198765432",
				Email:          "rizky.firmansyah@protonmail.com",
				Address:        "Jl. Pemuda No. 14",
				City:           "Bekasi",
				PostalCode:     "17113",
				Birthdate:      timePtr(time.Date(1983, 6, 29, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "premium",
				CreditLimit:    1800000,
				CurrentBalance: 200000,
				IsActive:       true,
				Notes:          "Pelanggan tetap mingguan",
			},
			{
				Code:           "CUST0015",
				Name:           "Yuli Andini",
				Phone:          "084298765432",
				Email:          "yuli.andini@mail.org",
				Address:        "Jl. Kalimantan No. 17",
				City:           "Pontianak",
				PostalCode:     "78121",
				Birthdate:      timePtr(time.Date(1997, 12, 30, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    800000,
				CurrentBalance: 100000,
				IsActive:       true,
				Notes:          "Baru bergabung, pembeli aktif",
			},
			{
				Code:           "CUST0016",
				Name:           "Eko Prasetyo",
				Phone:          "081289765432",
				Email:          "eko.prasetyo@domain.id",
				Address:        "Jl. Veteran No. 40",
				City:           "Denpasar",
				PostalCode:     "80231",
				Birthdate:      timePtr(time.Date(1982, 3, 18, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "vip",
				CreditLimit:    5000000,
				CurrentBalance: 600000,
				IsActive:       true,
				Notes:          "Corporate buyer, bulk orders",
			},
			{
				Code:           "CUST0017",
				Name:           "Hana Oktaviani",
				Phone:          "083456789012",
				Email:          "hana.oktaviani@mail.com",
				Address:        "Jl. Bunga Raya No. 8",
				City:           "Balikpapan",
				PostalCode:     "76123",
				Birthdate:      timePtr(time.Date(1990, 10, 11, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    1000000,
				CurrentBalance: 30000,
				IsActive:       true,
				Notes:          "Pembelian via WhatsApp & e-commerce",
			},
		}

		// Simpan data ke database
		for _, customer := range customers {
			if err := db.Create(&customer).Error; err != nil {
				log.Fatalf("Gagal menambahkan data customer: %v", err)
			}
		}
		log.Println("Data Customer berhasil di-seed!")
	} else {
		log.Println("Data Customer sudah ada, tidak perlu di-seed.")
	}
}

func SeedSuppliers(db *gorm.DB) {
	// Cek apakah sudah ada data Supplier
	var supplierCount int64
	db.Model(&vendor.Supplier{}).Count(&supplierCount)

	// Jika belum ada data, seed data awal
	if supplierCount == 0 {
		// Seed data Supplier
		suppliers := []vendor.Supplier{
			{
				Code:          "SUP-TECH01",
				Name:          "Tech Solutions Indonesia",
				ContactPerson: "Dimas Arif",
				Phone:         "081234567890",
				Email:         "contact@techsolutions.co.id",
				Address:       "Jl. Teknologi No. 21, Jakarta Selatan",
				PaymentTerms:  30,
				IsActive:      true,
				Notes:         "Spesialis distributor produk elektronik dan gadget.",
			},
			{
				Code:          "SUP-GADG02",
				Name:          "Gadget Supply Co.",
				ContactPerson: "Siti Rahmawati",
				Phone:         "082234567891",
				Email:         "sales@gadgetsupply.co.id",
				Address:       "Jl. Kemang Timur No. 88, Jakarta Selatan",
				PaymentTerms:  45,
				IsActive:      true,
				Notes:         "Supplier utama untuk aksesoris HP dan perangkat pintar.",
			},
			{
				Code:          "SUP-COMP03",
				Name:          "CompuTech Nusantara",
				ContactPerson: "Andi Prasetyo",
				Phone:         "083334567892",
				Email:         "admin@computech.id",
				Address:       "Jl. Diponegoro No. 56, Bandung",
				PaymentTerms:  30,
				IsActive:      true,
				Notes:         "Distributor resmi laptop dan aksesoris komputer.",
			},
			{
				Code:          "SUP-FASH04",
				Name:          "Urban Style Apparel",
				ContactPerson: "Linda Anggraini",
				Phone:         "084434567893",
				Email:         "support@urbanstyle.id",
				Address:       "Jl. Sudirman No. 12, Yogyakarta",
				PaymentTerms:  60,
				IsActive:      true,
				Notes:         "Penyedia utama pakaian pria dan fashion kasual.",
			},
			{
				Code:          "SUP-HOME05",
				Name:          "Home & Living Essentials",
				ContactPerson: "Rudi Hartono",
				Phone:         "085534567894",
				Email:         "cs@homeessentials.co.id",
				Address:       "Jl. A. Yani No. 8, Surabaya",
				PaymentTerms:  45,
				IsActive:      true,
				Notes:         "Memasok peralatan rumah tangga dan elektronik rumah.",
			},
		}

		for _, supplier := range suppliers {
			if err := db.Create(&supplier).Error; err != nil {
				log.Fatalf("Gagal menambahkan data supplier: %v", err)
			}
		}
		log.Println("Data Supplier berhasil di-seed!")
	} else {
		log.Println("Data Supplier sudah ada, tidak perlu di-seed.")
	}
}

func SeedProductsAndSuppliers(db *gorm.DB) {
	// Cek apakah sudah ada data Product
	var productCount int64
	db.Model(&inventory.Product{}).Count(&productCount)

	// Jika belum ada data, seed data awal
	if productCount == 0 {
		// Seed data Product
		products := []inventory.Product{
			// Phones (CategoryID: 2)
			{CategoryID: 2, SKU: "PHN001", Barcode: "100000000001", Name: "Smartphone Nova X", Description: "Smartphone layar 6.5 inch dengan kamera 64MP dan baterai tahan lama.", BuyingPrice: 2500000, SellingPrice: 3250000, MinStock: 5, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "smartphone,android,novax"},
			{CategoryID: 2, SKU: "PHN002", Barcode: "100000000002", Name: "Charger FastCharge 25W", Description: "Charger USB Type-C dengan fitur fast charging 25W.", BuyingPrice: 120000, SellingPrice: 185000, MinStock: 10, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "charger,accessories,usb"},
			{CategoryID: 2, SKU: "PHN003", Barcode: "100000000003", Name: "Earphone Bluetooth AirSound", Description: "Earphone wireless dengan suara jernih dan koneksi stabil.", BuyingPrice: 275000, SellingPrice: 389000, MinStock: 6, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "earphone,bluetooth,audio"},
			{CategoryID: 2, SKU: "PHN004", Barcode: "100000000004", Name: "Phone Holder Universal", Description: "Holder smartphone serbaguna untuk kendaraan dan meja kerja.", BuyingPrice: 50000, SellingPrice: 89000, MinStock: 8, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "aksesoris,holder,smartphone"},
			{CategoryID: 2, SKU: "PHN005", Barcode: "100000000005", Name: "Tempered Glass Anti Gores", Description: "Pelindung layar tempered glass anti gores untuk berbagai model HP.", BuyingPrice: 30000, SellingPrice: 60000, MinStock: 12, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "pelindung layar,tempered glass,smartphone"},

			// Laptops (CategoryID: 3)
			{CategoryID: 3, SKU: "LPT001", Barcode: "200000000001", Name: "Laptop Ultrabook AirLite 14", Description: "Laptop tipis ringan dengan prosesor i5, RAM 8GB, SSD 512GB.", BuyingPrice: 6500000, SellingPrice: 7590000, MinStock: 3, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "laptop,ultrabook,ssd"},
			{CategoryID: 3, SKU: "LPT002", Barcode: "200000000002", Name: "Mouse Wireless ProClick", Description: "Mouse nirkabel ergonomis dengan sensitivitas tinggi.", BuyingPrice: 90000, SellingPrice: 135000, MinStock: 10, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "mouse,wireless,komputer"},
			{CategoryID: 3, SKU: "LPT003", Barcode: "200000000003", Name: "Laptop Cooling Pad", Description: "Cooling pad dengan kipas ganda untuk menjaga suhu laptop tetap dingin.", BuyingPrice: 70000, SellingPrice: 110000, MinStock: 5, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "cooling pad,aksesori laptop"},
			{CategoryID: 3, SKU: "LPT004", Barcode: "200000000004", Name: "Backpack Laptop UrbanGear", Description: "Tas laptop stylish dengan pelindung air dan ruang penyimpanan luas.", BuyingPrice: 180000, SellingPrice: 275000, MinStock: 4, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "tas laptop,backpack,gear"},
			{CategoryID: 3, SKU: "LPT005", Barcode: "200000000005", Name: "Keyboard Bluetooth Foldable", Description: "Keyboard lipat Bluetooth untuk mobilitas tinggi dan kerja remote.", BuyingPrice: 220000, SellingPrice: 299000, MinStock: 3, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "keyboard,bluetooth,portable"},

			// Men's Clothing (CategoryID: 5)
			{CategoryID: 5, SKU: "MNC001", Barcode: "300000000001", Name: "Kemeja Casual Slim Fit", Description: "Kemeja pria bahan katun dengan potongan slim fit, cocok untuk casual & semi formal.", BuyingPrice: 125000, SellingPrice: 189000, MinStock: 7, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "kemeja,pria,fashion"},
			{CategoryID: 5, SKU: "MNC002", Barcode: "300000000002", Name: "Kaos Polos Premium", Description: "Kaos polos pria bahan combed 30s yang adem dan nyaman dipakai sehari-hari.", BuyingPrice: 75000, SellingPrice: 120000, MinStock: 10, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "kaos,pria,pakaian"},
			{CategoryID: 5, SKU: "MNC003", Barcode: "300000000003", Name: "Celana Chino Stretch", Description: "Celana chino slim fit dengan bahan stretch nyaman untuk aktivitas harian.", BuyingPrice: 170000, SellingPrice: 250000, MinStock: 6, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "celana,chino,pria"},
			{CategoryID: 5, SKU: "MNC004", Barcode: "300000000004", Name: "Jaket Parasut Hoodie", Description: "Jaket pria bahan parasut anti angin dan air, dilengkapi hoodie dan saku.", BuyingPrice: 220000, SellingPrice: 345000, MinStock: 4, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "jaket,pria,outdoor"},
			{CategoryID: 5, SKU: "MNC005", Barcode: "300000000005", Name: "Sabuk Kulit Pria Classic", Description: "Sabuk kulit asli dengan gesper besi, model klasik untuk formal dan casual.", BuyingPrice: 95000, SellingPrice: 145000, MinStock: 5, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "aksesoris,sabuk,pria"},

			// Electronics (Kategori umum) (CategoryID: 1)
			{CategoryID: 1, SKU: "ELEC001", Barcode: "400000000001", Name: "Smart TV 42 Inch 4K UHD", Description: "Smart TV 42 inci dengan resolusi 4K dan fitur YouTube, Netflix built-in.", BuyingPrice: 3400000, SellingPrice: 4250000, MinStock: 2, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "tv,4k,smart"},
			{CategoryID: 1, SKU: "ELEC002", Barcode: "400000000002", Name: "Camera CCTV Wireless 360", Description: "Kamera CCTV nirkabel dengan fitur rotasi 360° dan night vision.", BuyingPrice: 320000, SellingPrice: 450000, MinStock: 3, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "cctv,security,camera"},
			{CategoryID: 1, SKU: "ELEC003", Barcode: "400000000003", Name: "Smart Home Plug WiFi", Description: "Colokan listrik pintar yang bisa dikontrol lewat aplikasi smartphone.", BuyingPrice: 150000, SellingPrice: 225000, MinStock: 4, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "smartplug,rumah pintar,elektronik"},
			{CategoryID: 1, SKU: "ELEC004", Barcode: "400000000004", Name: "Lampu LED Sensor Gerak", Description: "Lampu LED hemat energi dengan sensor gerak otomatis.", BuyingPrice: 50000, SellingPrice: 85000, MinStock: 6, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "lampu,led,otomatis"},
			{CategoryID: 1, SKU: "ELEC005", Barcode: "400000000005", Name: "Smart Air Purifier X2", Description: "Air purifier dengan filter HEPA dan kontrol via aplikasi.", BuyingPrice: 870000, SellingPrice: 999000, MinStock: 3, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "air purifier,smart,hepa"},

			// Fashion umum (CategoryID: 4)
			{CategoryID: 4, SKU: "FSH001", Barcode: "500000000001", Name: "Topi Baseball Unisex", Description: "Topi kasual model baseball cocok untuk pria dan wanita.", BuyingPrice: 45000, SellingPrice: 80000, MinStock: 5, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "topi,aksesoris,unisex"},
			{CategoryID: 4, SKU: "FSH002", Barcode: "500000000002", Name: "Sandal Slide Sporty", Description: "Sandal model slide dengan bahan nyaman dan anti slip.", BuyingPrice: 60000, SellingPrice: 95000, MinStock: 7, IsService: false, IsActive: true, IsFeatured: false, AllowFractions: 0, ImageURL: "", Tags: "sandal,fashion,kasual"},
			{CategoryID: 4, SKU: "FSH003", Barcode: "500000000003", Name: "Jas Hujan Stylish", Description: "Jas hujan pria/wanita model trendy dengan bahan waterproof.", BuyingPrice: 85000, SellingPrice: 135000, MinStock: 6, IsService: false, IsActive: true, IsFeatured: true, AllowFractions: 0, ImageURL: "", Tags: "jas hujan,aksesori,outdoor"},
		}

		for _, product := range products {
			if err := db.Create(&product).Error; err != nil {
				log.Fatalf("Gagal menambahkan data produk: %v", err)
			}
		}
		log.Println("Data Product berhasil di-seed!")

		// Seed data ProductSupplier
		productSuppliers := []inventory.ProductSupplier{
			{ProductID: 3, SupplierID: 1, BuyingPrice: float64Ptr(250000), LeadTime: intPtr(7), MinimumOrderQuantity: 10, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -3))},
			{ProductID: 4, SupplierID: 1, BuyingPrice: float64Ptr(120000), LeadTime: intPtr(5), MinimumOrderQuantity: 15, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -1))},
			{ProductID: 5, SupplierID: 1, BuyingPrice: float64Ptr(275000), LeadTime: intPtr(6), MinimumOrderQuantity: 8, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -10))},
			{ProductID: 6, SupplierID: 2, BuyingPrice: float64Ptr(50000), LeadTime: intPtr(3), MinimumOrderQuantity: 20, IsPrimary: false, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -4))},
			{ProductID: 7, SupplierID: 2, BuyingPrice: float64Ptr(30000), LeadTime: intPtr(4), MinimumOrderQuantity: 30, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, 0))},

			{ProductID: 8, SupplierID: 3, BuyingPrice: float64Ptr(6500000), LeadTime: intPtr(10), MinimumOrderQuantity: 5, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -3, 0))},
			{ProductID: 9, SupplierID: 3, BuyingPrice: float64Ptr(90000), LeadTime: intPtr(2), MinimumOrderQuantity: 25, IsPrimary: false, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -12))},
			{ProductID: 10, SupplierID: 3, BuyingPrice: float64Ptr(70000), LeadTime: intPtr(4), MinimumOrderQuantity: 12, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -8))},
			{ProductID: 11, SupplierID: 3, BuyingPrice: float64Ptr(180000), LeadTime: intPtr(6), MinimumOrderQuantity: 10, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -2))},
			{ProductID: 12, SupplierID: 3, BuyingPrice: float64Ptr(220000), LeadTime: intPtr(5), MinimumOrderQuantity: 10, IsPrimary: false, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -5))},

			{ProductID: 13, SupplierID: 4, BuyingPrice: float64Ptr(125000), LeadTime: intPtr(4), MinimumOrderQuantity: 15, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -7))},
			{ProductID: 14, SupplierID: 4, BuyingPrice: float64Ptr(75000), LeadTime: intPtr(2), MinimumOrderQuantity: 20, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -6))},
			{ProductID: 15, SupplierID: 4, BuyingPrice: float64Ptr(170000), LeadTime: intPtr(7), MinimumOrderQuantity: 8, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -9))},
			{ProductID: 16, SupplierID: 4, BuyingPrice: float64Ptr(220000), LeadTime: intPtr(10), MinimumOrderQuantity: 5, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -11))},
			{ProductID: 17, SupplierID: 4, BuyingPrice: float64Ptr(95000), LeadTime: intPtr(3), MinimumOrderQuantity: 12, IsPrimary: false, LastSupplyDate: timePtr(time.Now().AddDate(0, -3, -1))},

			{ProductID: 18, SupplierID: 5, BuyingPrice: float64Ptr(3400000), LeadTime: intPtr(9), MinimumOrderQuantity: 2, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -10))},
			{ProductID: 19, SupplierID: 5, BuyingPrice: float64Ptr(320000), LeadTime: intPtr(5), MinimumOrderQuantity: 4, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, 0))},
			{ProductID: 20, SupplierID: 5, BuyingPrice: float64Ptr(150000), LeadTime: intPtr(4), MinimumOrderQuantity: 6, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -15))},
			{ProductID: 21, SupplierID: 5, BuyingPrice: float64Ptr(50000), LeadTime: intPtr(3), MinimumOrderQuantity: 15, IsPrimary: false, LastSupplyDate: timePtr(time.Now().AddDate(0, -3, -5))},
			{ProductID: 22, SupplierID: 5, BuyingPrice: float64Ptr(870000), LeadTime: intPtr(7), MinimumOrderQuantity: 4, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -3))},

			// Fashion umum
			{ProductID: 23, SupplierID: 1, BuyingPrice: float64Ptr(45000), LeadTime: intPtr(2), MinimumOrderQuantity: 10, IsPrimary: false, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -6))},
			{ProductID: 24, SupplierID: 2, BuyingPrice: float64Ptr(60000), LeadTime: intPtr(5), MinimumOrderQuantity: 8, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -7))},
			{ProductID: 25, SupplierID: 2, BuyingPrice: float64Ptr(85000), LeadTime: intPtr(6), MinimumOrderQuantity: 6, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -3, 0))},
			{ProductID: 1, SupplierID: 3, BuyingPrice: float64Ptr(130000), LeadTime: intPtr(5), MinimumOrderQuantity: 4, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -2, -15))},
			{ProductID: 2, SupplierID: 4, BuyingPrice: float64Ptr(115000), LeadTime: intPtr(6), MinimumOrderQuantity: 6, IsPrimary: true, LastSupplyDate: timePtr(time.Now().AddDate(0, -1, -20))},
		}

		for _, ps := range productSuppliers {
			if err := db.Create(&ps).Error; err != nil {
				log.Fatalf("Gagal menambahkan data ProductSupplier: %v", err)
			}
		}
		log.Println("Data ProductSupplier berhasil di-seed!")
	} else {
		log.Println("Data Product sudah ada, tidak perlu di-seed.")
	}
}

func SeedDiscounts(db *gorm.DB) {
	// Cek apakah sudah ada data Discount
	var count int64
	db.Model(&promotion.Discount{}).Count(&count)

	// Jika belum ada data, seed data awal
	if count == 0 {
		discounts := []promotion.Discount{
			{
				CategoryID:    1,
				ProductID:     2,
				CustomerID:    1,
				Name:          "Diskon Musim Panas",
				Code:          "SUMMER2025",
				Description:   "Diskon spesial untuk produk musim panas.",
				DiscountType:  "percentage",
				DiscountValue: 10.00,                                                    // 10% diskon
				MinPurchase:   100000,                                                   // Minimal pembelian
				MaxDiscount:   50000,                                                    // Maksimum diskon
				StartDate:     time.Now().Format("2006-01-02"),                          // Gunakan time yang benar
				EndDate:       time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02"), // berlaku selama 30 hari
				UsageLimit:    100,
				UsageCount:    0,
				IsActive:      true,
				AppliesTo:     "all_products", // atau "specific_product", sesuaikan
			},
			{
				CategoryID:    2,
				ProductID:     1,
				CustomerID:    2,
				Name:          "Diskon Akhir Tahun",
				Code:          "YEAR_END2025",
				Description:   "Diskon untuk menyambut tahun baru.",
				DiscountType:  "flat",
				DiscountValue: 50000, // Potongan 50,000
				MinPurchase:   50000,
				MaxDiscount:   50000,
				StartDate:     time.Now().Format("2006-01-02"),
				EndDate:       time.Now().Add(15 * 24 * time.Hour).Format("2006-01-02"),
				UsageLimit:    200,
				UsageCount:    0,
				IsActive:      true,
				AppliesTo:     "specific_product",
			},
		}

		// Simpan data ke database
		for _, discount := range discounts {
			if err := db.Create(&discount).Error; err != nil {
				log.Fatalf("Gagal menambahkan data discount: %v", err)
			}
		}
		log.Println("Data Discount berhasil di-seed!")
	} else {
		log.Println("Data Discount sudah ada, tidak perlu di-seed.")
	}
}

func SeedTaxes(db *gorm.DB) {
	// Cek apakah sudah ada data Tax
	var taxCount int64
	db.Model(&finance.Tax{}).Count(&taxCount)

	// Jika belum ada data, seed data awal
	if taxCount == 0 {
		// Seed data Tax
		taxes := []finance.Tax{
			{
				Name:      "Pajak Pertambahan Nilai (PPN)",
				Rate:      10.00, // Contoh rate 10%
				IsDefault: 1,     // Menandakan ini adalah pajak default
				IsActive:  true,
			},
			{
				Name:      "Pajak Penghasilan (PPh)",
				Rate:      5.00, // Contoh rate 5%
				IsDefault: 0,    // Menandakan ini bukan pajak default
				IsActive:  true,
			},
			{
				Name:      "Pajak Daerah",
				Rate:      3.00, // Contoh rate 3%
				IsDefault: 0,    // Menandakan ini bukan pajak default
				IsActive:  true,
			},
		}

		for _, tax := range taxes {
			if err := db.Create(&tax).Error; err != nil {
				log.Fatalf("Gagal menambahkan data tax: %v", err)
			}
		}
		log.Println("Data Tax berhasil di-seed!")
	} else {
		log.Println("Data Tax sudah ada, tidak perlu di-seed.")
	}
}

func SeedLoans(db *gorm.DB) {
	// Cek apakah sudah ada data Loan
	var loanCount int64
	db.Model(&finance.Loan{}).Count(&loanCount)

	if loanCount == 0 {
		loans := []finance.Loan{
			{
				CustomerID:        1,
				LoanAmount:        10000000,
				InterestRate:      5.5,
				LoanTerm:          12,
				InstallmentAmount: 875000,
				RemainingAmount:   8750000,
				StartDate:         "2025-01-01",
				DueDate:           "2025-12-31",
				Status:            "active",
				Notes:             "Loan for working capital",
			},
			{
				CustomerID:        2,
				LoanAmount:        5000000,
				InterestRate:      3.5,
				LoanTerm:          6,
				InstallmentAmount: 850000,
				RemainingAmount:   2550000,
				StartDate:         "2025-03-01",
				DueDate:           "2025-08-31",
				Status:            "active",
				Notes:             "Emergency loan",
			},
		}

		for _, loan := range loans {
			if err := db.Create(&loan).Error; err != nil {
				log.Fatalf("Gagal menambahkan data Loan: %v", err)
			}
		}
		log.Println("Data Loan berhasil di-seed!")
	} else {
		log.Println("Data Loan sudah ada, tidak perlu di-seed.")
	}

	// Cek apakah sudah ada data LoanPayment
	var paymentCount int64
	db.Model(&finance.LoanPayment{}).Count(&paymentCount)

	if paymentCount == 0 {
		loanPayments := []finance.LoanPayment{
			{
				LoanID:        1,
				PaymentAmount: 1250000,
				Notes:         "First installment",
			},
			{
				LoanID:        2,
				PaymentAmount: 850000,
				Notes:         "First payment",
			},
		}

		for _, payment := range loanPayments {
			if err := db.Create(&payment).Error; err != nil {
				log.Fatalf("Gagal menambahkan data LoanPayment: %v", err)
			}
		}
		log.Println("Data LoanPayment berhasil di-seed!")
	} else {
		log.Println("Data LoanPayment sudah ada, tidak perlu di-seed.")
	}
}

func SeedCashDrawerData(db *gorm.DB) {
	// Seed CashDrawer
	var drawerCount int64
	db.Model(&finance.CashDrawer{}).Count(&drawerCount)

	if drawerCount == 0 {
		cashDrawers := []finance.CashDrawer{
			{
				UserID:         1,
				BranchID:       1,
				OpeningTime:    "2025-04-16 08:00:00",
				OpeningAmount:  500000,
				Status:         "open",
				SyncStatus:     "pending",
				Notes:          "Morning shift",
				ExpectedAmount: 500000,
			},
			{
				UserID:         2,
				BranchID:       1,
				OpeningTime:    "2025-04-16 14:00:00",
				OpeningAmount:  400000,
				Status:         "open",
				SyncStatus:     "pending",
				Notes:          "Afternoon shift",
				ExpectedAmount: 400000,
			},
		}

		for _, drawer := range cashDrawers {
			if err := db.Create(&drawer).Error; err != nil {
				log.Fatalf("Gagal menambahkan data CashDrawer: %v", err)
			}
		}
		log.Println("Data CashDrawer berhasil di-seed!")
	} else {
		log.Println("Data CashDrawer sudah ada, tidak perlu di-seed.")
	}

	// Seed CashDrawerTransaction
	var txCount int64
	db.Model(&finance.CashDrawerTransaction{}).Count(&txCount)

	if txCount == 0 {
		transactions := []finance.CashDrawerTransaction{
			{
				CashDrawerID:    1,
				UserID:          1,
				TransactionType: "sale",
				Amount:          250000,
				Notes:           "First sale of the day",
				SyncStatus:      "pending",
			},
			{
				CashDrawerID:    1,
				UserID:          1,
				TransactionType: "refund",
				Amount:          -50000,
				Notes:           "Refund to customer",
				SyncStatus:      "pending",
			},
			{
				CashDrawerID:    2,
				UserID:          2,
				TransactionType: "sale",
				Amount:          100000,
				Notes:           "Afternoon transaction",
				SyncStatus:      "pending",
			},
		}

		for _, tx := range transactions {
			if err := db.Create(&tx).Error; err != nil {
				log.Fatalf("Gagal menambahkan data CashDrawerTransaction: %v", err)
			}
		}
		log.Println("Data CashDrawerTransaction berhasil di-seed!")
	} else {
		log.Println("Data CashDrawerTransaction sudah ada, tidak perlu di-seed.")
	}
}

func SeedFees(db *gorm.DB) {
	// Cek apakah sudah ada data Fee
	var feeCount int64
	db.Model(&orders.Fee{}).Count(&feeCount)

	// Jika belum ada data, seed data awal
	if feeCount == 0 {
		// Seed data Fee
		fees := []orders.Fee{
			{
				Name:      "Biaya Pengiriman",
				FeeType:   "Shipping",
				FeeValue:  5000.00, // Contoh biaya pengiriman
				IsDefault: 1,       // Menandakan ini adalah biaya default
				IsActive:  true,
			},
			{
				Name:      "Biaya Layanan",
				FeeType:   "Service",
				FeeValue:  10000.00, // Contoh biaya layanan
				IsDefault: 0,        // Menandakan ini bukan biaya default
				IsActive:  true,
			},
			{
				Name:      "Biaya Penanganan",
				FeeType:   "Handling",
				FeeValue:  2000.00, // Contoh biaya penanganan
				IsDefault: 0,       // Menandakan ini bukan biaya default
				IsActive:  true,
			},
		}

		for _, fee := range fees {
			if err := db.Create(&fee).Error; err != nil {
				log.Fatalf("Gagal menambahkan data fee: %v", err)
			}
		}
		log.Println("Data Fee berhasil di-seed!")
	} else {
		log.Println("Data Fee sudah ada, tidak perlu di-seed.")
	}
}

func SeedPromotions(db *gorm.DB) {
	// Cek apakah sudah ada data Promotion
	var count int64
	db.Model(&promotion.Promotion{}).Count(&count)

	if count > 0 {
		log.Println("Data Promotion sudah ada, tidak perlu di-seed.")
		return
	}

	// Seed Promotion
	promos := []promotion.Promotion{
		{
			Name:        "Diskon Awal Tahun",
			Description: "Diskon untuk pembelian di awal tahun",
			PromoType:   "percentage",
			StartDate:   "2025-01-01",
			EndDate:     "2025-01-31",
			IsActive:    true,
		},
		{
			Name:        "Promo Beli 2 Gratis 1",
			Description: "Beli 2 produk tertentu gratis 1",
			PromoType:   "product",
			StartDate:   "2025-02-01",
			EndDate:     "2025-02-28",
			IsActive:    true,
		},
	}

	if err := db.Create(&promos).Error; err != nil {
		log.Fatalf("Gagal menambahkan data Promotion: %v", err)
	}

	// Seed PromotionRule
	rules := []promotion.PromotionRule{
		{
			PromotionID:   promos[0].ID,
			MinAmount:     100000,
			DiscountType:  "percentage",
			DiscountValue: 10.0,
		},
		{
			PromotionID:  promos[1].ID,
			MinQuantity:  2,
			DiscountType: "product",
		},
	}

	if err := db.Create(&rules).Error; err != nil {
		log.Fatalf("Gagal menambahkan data PromotionRule: %v", err)
	}

	// Seed PromotionProduct
	promoProducts := []promotion.PromotionProduct{
		{
			PromotionID: promos[1].ID,
			ProductID:   1, // pastikan ID produk 1 sudah ada
			IsTrigger:   true,
			Quantity:    2,
		},
		{
			PromotionID: promos[1].ID,
			ProductID:   2, // pastikan ID produk 2 sudah ada
			IsTarget:    true,
			Quantity:    1,
		},
	}

	if err := db.Create(&promoProducts).Error; err != nil {
		log.Fatalf("Gagal menambahkan data PromotionProduct: %v", err)
	}

	log.Println("Data Promotion, PromotionRule, dan PromotionProduct berhasil di-seed!")
}

func SeedGoodsReceiving(db *gorm.DB) {
	// Cek apakah sudah ada data GoodsReceiving
	var count int64
	db.Model(&purchase.GoodsReceiving{}).Count(&count)

	if count > 0 {
		log.Println("Data GoodsReceiving sudah ada, tidak perlu di-seed.")
		return
	}

	// Seeder tergantung entitas lain (pastikan data Supplier, Branch, User, Product sudah ada)
	gr := purchase.GoodsReceiving{
		ReferenceNumber: "GR-20250401-001",
		SupplierID:      1, // pastikan ID 1 ada
		BranchID:        1, // pastikan ID 1 ada
		UserID:          1, // pastikan ID 1 ada
		ReceivingDate:   time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
		Subtotal:        500000,
		TaxAmount:       50000,
		DiscountAmount:  10000,
		ShippingCost:    20000,
		OtherCosts:      0,
		GrandTotal:      560000,
		Notes:           strPtr("Penerimaan barang pertama"),
	}

	if err := db.Create(&gr).Error; err != nil {
		log.Fatalf("Gagal membuat GoodsReceiving: %v", err)
	}

	items := []purchase.GoodsReceivingItem{
		{
			GRID:            gr.ID,
			ProductID:       1, // pastikan ID produk 1 ada
			Quantity:        10,
			UnitPrice:       50000,
			DiscountPercent: 0,
			TaxPercent:      10,
			Subtotal:        500000,
			BatchNumber:     strPtr("BN-001"),
			ExpiryDate:      strPtr("2025-12-31"),
			Notes:           strPtr("Barang segar"),
		},
	}

	if err := db.Create(&items).Error; err != nil {
		log.Fatalf("Gagal membuat GoodsReceivingItem: %v", err)
	}

	log.Println("Data GoodsReceiving dan GoodsReceivingItem berhasil di-seed.")
}

func SeedPurchaseOrders(db *gorm.DB) {
	// Cek apakah sudah ada data PurchaseOrder
	var count int64
	db.Model(&purchase.PurchaseOrder{}).Count(&count)

	if count > 0 {
		log.Println("Data PurchaseOrder sudah ada, tidak perlu di-seed.")
		return
	}

	// Seeder tergantung entitas lain (pastikan data Supplier, Branch, User, Product sudah ada)
	po := purchase.PurchaseOrder{
		PONumber:       "PO-20250401-001",
		SupplierID:     1, // pastikan ID 1 ada
		BranchID:       1, // pastikan ID 1 ada
		UserID:         1, // pastikan ID 1 ada
		PODate:         time.Date(2025, 4, 1, 10, 0, 0, 0, time.UTC),
		Subtotal:       500000,
		TaxAmount:      50000,
		DiscountAmount: 10000,
		ShippingCost:   20000,
		OtherCosts:     0,
		GrandTotal:     560000,
		Status:         "draft",
		Notes:          strPtr("Pembelian barang untuk cabang utama"),
		PaymentTerms:   intPtr(30), // jika ada term pembayaran
	}

	if err := db.Create(&po).Error; err != nil {
		log.Fatalf("Gagal membuat PurchaseOrder: %v", err)
	}

	items := []purchase.PurchaseOrderItem{
		{
			POID:            po.ID,
			ProductID:       1, // pastikan ID produk 1 ada
			Quantity:        10,
			UnitPrice:       50000,
			DiscountPercent: 0,
			TaxPercent:      10,
			Subtotal:        500000,
			Notes:           strPtr("Produk utama dalam pembelian"),
		},
	}

	if err := db.Create(&items).Error; err != nil {
		log.Fatalf("Gagal membuat PurchaseOrderItem: %v", err)
	}

	log.Println("Data PurchaseOrder dan PurchaseOrderItem berhasil di-seed.")
}

func SeedTransactions(db *gorm.DB) {
	var transactionCount int64
	db.Model(&orders.Transaction{}).Count(&transactionCount)

	if transactionCount == 0 {
		// Asumsi 10 customer pertama punya ID 1 - 10
		for i := 1; i <= 10; i++ {
			invoiceNum := "INV-20240425-00" + fmt.Sprint(i)
			transaction := orders.Transaction{
				CustomerID:       i,
				UserID:           1, // Asumsi user id 1 exist
				BranchID:         1, // Asumsi branch id 1 exist
				DiscountID:       1,
				TaxID:            1,
				FeeID:            1,
				InvoiceNumber:    invoiceNum,
				InvoiceDate:      time.Now(),
				TransactionDate:  time.Now(),
				DueDate:          time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
				Subtotal:         300000,
				DiscountAmount:   15000,
				TaxAmount:        15000,
				FeeAmount:        5000,
				ShippingCost:     10000,
				GrandTotal:       315000,
				AmountPaid:       320000,
				AmountReturned:   5000,
				PaymentStatus:    "paid",
				PointsEarned:     30,
				PointsUsed:       0,
				Notes:            "Transaksi seed pelanggan ke-" + fmt.Sprint(i),
				Status:           "completed",
				ReferenceID:      1,
				ShippingAddress:  "Alamat kirim customer " + fmt.Sprint(i),
				ShippingTracking: "TRK2024" + fmt.Sprint(1000+i),
				SyncStatus:       "synced",
			}

			if err := db.Create(&transaction).Error; err != nil {
				log.Fatalf("Gagal menyimpan transaksi %d: %v", i, err)
			}

			items := []orders.TransactionItem{
				{
					TransactionID:   transaction.ID,
					ProductID:       1,
					Quantity:        2,
					UnitPrice:       100000,
					OriginalPrice:   110000,
					DiscountPercent: 5,
					DiscountAmount:  10000,
					TaxPercent:      5,
					TaxAmount:       10000,
					Subtotal:        190000,
					Notes:           "Barang A",
					SyncStatus:      "synced",
				},
				{
					TransactionID:   transaction.ID,
					ProductID:       2,
					Quantity:        1,
					UnitPrice:       100000,
					OriginalPrice:   100000,
					DiscountPercent: 5,
					DiscountAmount:  5000,
					TaxPercent:      5,
					TaxAmount:       5000,
					Subtotal:        95000,
					Notes:           "Barang B",
					SyncStatus:      "synced",
				},
			}

			for _, item := range items {
				if err := db.Create(&item).Error; err != nil {
					log.Fatalf("Gagal menyimpan item transaksi %d: %v", i, err)
				}
			}
		}

		log.Println("10 transaksi berhasil di-seed!")
	} else {
		log.Println("Data transaksi sudah ada, tidak perlu di-seed.")
	}
}

// Helper function to create pointer of time for birthdate field
func timePtr(t time.Time) *time.Time {
	return &t
}

// Helper function to create pointer of float64 for buying price
func float64Ptr(f float64) *float64 {
	return &f
}

// Helper function to create pointer of int for lead time
func intPtr(i int) *int {
	return &i
}

func strPtr(s string) *string {
	return &s
}
