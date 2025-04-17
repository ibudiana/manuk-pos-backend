package database

import (
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
				Code:           "CUST0001",
				Name:           "John Doe",
				Phone:          "081234567890",
				Email:          "john.doe@example.com",
				Address:        "Jl. Raya No. 10",
				City:           "Jakarta",
				PostalCode:     "12345",
				Birthdate:      timePtr(time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "regular",
				CreditLimit:    1000000.00,
				CurrentBalance: 50000.00,
				IsActive:       true,
				Notes:          "Regular customer",
			},
			{
				Code:           "CUST0002",
				Name:           "Jane Smith",
				Phone:          "082345678901",
				Email:          "jane.smith@example.com",
				Address:        "Jl. Merdeka No. 5",
				City:           "Bandung",
				PostalCode:     "67890",
				Birthdate:      timePtr(time.Date(1985, 8, 20, 0, 0, 0, 0, time.UTC)),
				JoinDate:       time.Now(),
				CustomerType:   "premium",
				CreditLimit:    2000000.00,
				CurrentBalance: 150000.00,
				IsActive:       true,
				Notes:          "Premium customer with higher credit limit",
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
				Code:          "SUP001",
				Name:          "Supplier A",
				ContactPerson: "John Doe",
				Phone:         "081234567890",
				Email:         "suppliera@example.com",
				Address:       "Jl. Raya No. 1, Jakarta",
				PaymentTerms:  30,
				IsActive:      true,
				Notes:         "Reliable supplier for electronics.",
			},
			{
				Code:          "SUP002",
				Name:          "Supplier B",
				ContactPerson: "Jane Smith",
				Phone:         "082345678901",
				Email:         "supplierb@example.com",
				Address:       "Jl. Merdeka No. 2, Bandung",
				PaymentTerms:  45,
				IsActive:      true,
				Notes:         "Supplier for office supplies.",
			},
			{
				Code:          "SUP003",
				Name:          "Supplier C",
				ContactPerson: "Michael Johnson",
				Phone:         "083456789012",
				Email:         "supplierc@example.com",
				Address:       "Jl. Pahlawan No. 3, Surabaya",
				PaymentTerms:  60,
				IsActive:      true,
				Notes:         "Supplier for furniture products.",
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
			{
				CategoryID:     1, // Asumsikan ada kategori dengan ID 1
				SKU:            "SKU0001",
				Barcode:        "123456789012",
				Name:           "Product A",
				Description:    "Description of Product A",
				BuyingPrice:    100.00,
				SellingPrice:   150.00,
				MinStock:       5,
				IsService:      false,
				IsActive:       true,
				IsFeatured:     false,
				AllowFractions: 0,
				ImageURL:       "http://example.com/productA.jpg",
				Tags:           "tag1,tag2",
			},
			{
				CategoryID:     2, // Asumsikan ada kategori dengan ID 1
				SKU:            "SKU0002",
				Barcode:        "987654321098",
				Name:           "Product B",
				Description:    "Description of Product B",
				BuyingPrice:    50.00,
				SellingPrice:   75.00,
				MinStock:       3,
				IsService:      false,
				IsActive:       true,
				IsFeatured:     true,
				AllowFractions: 0,
				ImageURL:       "http://example.com/productB.jpg",
				Tags:           "tag3,tag4",
			},
		}

		for _, product := range products {
			if err := db.Create(&product).Error; err != nil {
				log.Fatalf("Gagal menambahkan data produk: %v", err)
			}
		}
		log.Println("Data Product berhasil di-seed!")

		// Seed data ProductSupplier
		productSuppliers := []inventory.ProductSupplier{
			{
				ProductID:            1, // Asumsikan ProductID 1 sudah ada
				SupplierID:           1, // Asumsikan SupplierID 1 sudah ada
				BuyingPrice:          float64Ptr(95.00),
				LeadTime:             intPtr(7),
				MinimumOrderQuantity: 10,
				IsPrimary:            true,
				LastSupplyDate:       timePtr(time.Now().AddDate(0, -1, 0)), // 1 bulan lalu
			},
			{
				ProductID:            2, // Asumsikan ProductID 2 sudah ada
				SupplierID:           2, // Asumsikan SupplierID 2 sudah ada
				BuyingPrice:          float64Ptr(45.00),
				LeadTime:             intPtr(10),
				MinimumOrderQuantity: 20,
				IsPrimary:            false,
				LastSupplyDate:       timePtr(time.Now().AddDate(0, -2, 0)), // 2 bulan lalu
			},
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
