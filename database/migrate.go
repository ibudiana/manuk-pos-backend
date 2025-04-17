package database

import (
	"fmt"

	"manuk-pos-backend/models/customer"
	"manuk-pos-backend/models/finance"
	"manuk-pos-backend/models/inventory"
	"manuk-pos-backend/models/orders"
	"manuk-pos-backend/models/promotion"
	"manuk-pos-backend/models/purchase"
	"manuk-pos-backend/models/store"
	"manuk-pos-backend/models/sync"
	"manuk-pos-backend/models/user"
	"manuk-pos-backend/models/vendor"
)

func MigrateTables() {
	err := DB.AutoMigrate(
		&customer.Customer{},

		&finance.CashDrawer{},
		&finance.CashDrawerTransaction{},
		&finance.Loan{},
		&finance.LoanPayment{},
		&finance.Tax{},

		&orders.Transaction{},
		&orders.TransactionItem{},
		&orders.Fee{},
		&orders.Payment{},
		&orders.Points{},
		&orders.PointsHistory{},

		&inventory.Category{},
		&inventory.Inventory{},
		&inventory.InventoryTransaction{},
		&inventory.InventoryTransfer{},
		&inventory.InventoryTransferItem{},
		&inventory.Product{},
		&inventory.ProductSupplier{},
		&inventory.StockOpname{},
		&inventory.StockOpname{},

		&promotion.Discount{},
		&promotion.Promotion{},
		&promotion.PromotionRule{},
		&promotion.PromotionProduct{},

		&purchase.GoodsReceiving{},
		&purchase.GoodsReceivingItem{},
		&purchase.PurchaseOrder{},
		&purchase.PurchaseOrderItem{},

		&store.Branch{},

		&sync.SyncLog{},
		&sync.SyncDetail{},

		&user.User{},
		&user.Role{},

		&vendor.Supplier{},
	)
	if err != nil {
		panic("Gagal migrasi database: " + err.Error())
	}

	fmt.Println("Migrasi berhasil dilakukan!")
}
