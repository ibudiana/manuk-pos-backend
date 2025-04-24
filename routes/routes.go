package routes

import (
	"manuk-pos-backend/controllers"
	"manuk-pos-backend/controllers/financecontroller"
	"manuk-pos-backend/controllers/inventorycontroller"
	"manuk-pos-backend/controllers/ordercontroller"
	"manuk-pos-backend/controllers/promotioncontroller"
	"manuk-pos-backend/controllers/purchasecontroller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	apiRoutes := router.Group("/api")
	{
		// Auth Routes
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/register", controllers.RegisterUser)
			authRoutes.POST("/login", controllers.LoginUser)
		}

		// inventories Routes
		inventoryRoutes := apiRoutes.Group("/inventories")
		{
			// Product Routes
			productRoutes := inventoryRoutes.Group("/products")
			{
				productRoutes.POST("", inventorycontroller.CreateProduct)
				productRoutes.GET("", inventorycontroller.GetProducts)

				productRoutes.GET("/:id", inventorycontroller.GetProductByID)
				productRoutes.PATCH("/:id", inventorycontroller.UpdateProductByID)
				productRoutes.DELETE("/:id", inventorycontroller.DeleteProductByID)

				// Product Category Routes
				productRoutes.GET("/categories", inventorycontroller.GetProductCategories)
				productRoutes.POST("/categories", inventorycontroller.CreateProductCategory)

				productRoutes.GET("/categories/:id", inventorycontroller.GetProductCategoryByID)
				productRoutes.PATCH("/categories/:id", inventorycontroller.UpdateProductCategoryByID)
				productRoutes.DELETE("/categories/:id", inventorycontroller.DeleteProductCategoryByID)

				// Product Supplier Routes
				productRoutes.POST("/:id/suppliers", inventorycontroller.AddProductSuppliersByIdProduct)
				productRoutes.PATCH("/:id/suppliers", inventorycontroller.UpdateProductSupplierByIdProduct)
				productRoutes.DELETE("/:id/suppliers", inventorycontroller.DeleteProductSupplierByIdProduct)
			}

		}

		// Order Routes
		orderRoutes := apiRoutes.Group("/orders")
		{
			orderRoutes.POST("", ordercontroller.CreateTransactionOrder)
			orderRoutes.GET("", ordercontroller.GetTransactionOrders)

			// Fee Routes
			feeRoutes := orderRoutes.Group("/fees")
			{
				feeRoutes.POST("", ordercontroller.CreateFee)
				feeRoutes.GET("", ordercontroller.GetFees)

				feeRoutes.GET("/:id", ordercontroller.GetFeeByID)
				feeRoutes.PATCH("/:id", ordercontroller.UpdateFeeByID)
				feeRoutes.DELETE("/:id", ordercontroller.DeleteFeeByID)
			}
		}

		// Customer Routes
		customerRoutes := apiRoutes.Group("/customers")
		{
			customerRoutes.POST("", controllers.CreateCustomer)
			customerRoutes.GET("", controllers.GetCustomers)

			customerRoutes.GET("/:id", controllers.GetCustomerByID)
			customerRoutes.PATCH("/:id", controllers.UpdateCustomerByID)
			customerRoutes.DELETE("/:id", controllers.DeleteCustomerByID)
		}

		// Vendor Supplier Routes
		supplierRoutes := apiRoutes.Group("/vendor/suppliers")
		{
			supplierRoutes.POST("", controllers.CreateSupplier)
			supplierRoutes.GET("", controllers.GetSuppliers)

			supplierRoutes.GET("/:id", controllers.GetSupplierByID)
			supplierRoutes.PATCH("/:id", controllers.UpdateSupplierByID)
			supplierRoutes.DELETE("/:id", controllers.DeleteSupplierByID)
		}

		// User Routes
		userRoutes := apiRoutes.Group("/users")
		{
			userRoutes.POST("", controllers.CreateUser)
			userRoutes.GET("", controllers.GetUsers)

			userRoutes.GET("/:id", controllers.GetUserByID)
			userRoutes.PATCH("/:id", controllers.UpdateUserByID)
			userRoutes.DELETE("/:id", controllers.DeleteUserByID)
		}

		// Role Routes
		roleRoutes := apiRoutes.Group("/roles")
		{
			roleRoutes.POST("", controllers.CreateRole)
			roleRoutes.GET("", controllers.GetRoles)
			roleRoutes.GET("/:id", controllers.GetRoleByID)
			roleRoutes.PATCH("/:id", controllers.UpdateRoleByID)
			roleRoutes.DELETE("/:id", controllers.DeleteRoleByID)
		}

		// Branch Routes
		branchRoutes := apiRoutes.Group("/store/branches")
		{
			branchRoutes.POST("", controllers.CreateBranch)
			branchRoutes.GET("", controllers.GetBranches)

			branchRoutes.GET("/:id", controllers.GetBranchByID)
			branchRoutes.PATCH("/:id", controllers.UpdateBranchByID)
			branchRoutes.DELETE("/:id", controllers.DeleteBranchByID)
		}

		// Finance Routes
		financeRoutes := apiRoutes.Group("/finance")
		{
			taxRoutes := financeRoutes.Group("/taxes")
			{
				taxRoutes.POST("", financecontroller.CreateTax)
				taxRoutes.GET("", financecontroller.GetTaxes)

				taxRoutes.GET("/:id", financecontroller.GetTaxByID)
				taxRoutes.PATCH("/:id", financecontroller.UpdateTaxByID)
				taxRoutes.DELETE("/:id", financecontroller.DeleteTaxByID)
			}

			loanRoutes := financeRoutes.Group("/loans")
			{
				loanRoutes.POST("", financecontroller.CreateLoan)
				loanRoutes.GET("", financecontroller.GetLoans)

				loanRoutes.GET("/:id", financecontroller.GetLoanByID)
				loanRoutes.PATCH("/:id", financecontroller.UpdateLoanByID)
				loanRoutes.DELETE("/:id", financecontroller.DeleteLoanByID)
			}

			cashDrawerRoutes := financeRoutes.Group("/cash-drawers")
			{
				cashDrawerRoutes.POST("", financecontroller.CreateCashDrawer)
				cashDrawerRoutes.GET("", financecontroller.GetCashDrawers)

				cashDrawerRoutes.GET("/:id", financecontroller.GetCashDrawerByID)
				cashDrawerRoutes.PATCH("/:id", financecontroller.UpdateCashDrawerByID)
				cashDrawerRoutes.DELETE("/:id", financecontroller.DeleteCashDrawerByID)
			}
		}

		// Promotions Routes
		promotionRoutes := apiRoutes.Group("/promotions")
		{
			promotionRoutes.POST("", promotioncontroller.CreatePromotion)
			promotionRoutes.GET("", promotioncontroller.GetPromotions)

			promotionRoutes.GET("/:id", promotioncontroller.GetPromotionByID)
			promotionRoutes.PATCH("/:id", promotioncontroller.UpdatePromotionByID)
			promotionRoutes.DELETE("/:id", promotioncontroller.DeletePromotionByID)

			// Discount Routes
			discountRoutes := promotionRoutes.Group("/discounts")
			{
				discountRoutes.POST("", promotioncontroller.CreateDiscount)
				discountRoutes.GET("", promotioncontroller.GetDiscounts)

				discountRoutes.GET("/:id", promotioncontroller.GetDiscountByID)
				discountRoutes.PATCH("/:id", promotioncontroller.UpdateDiscountByID)
				discountRoutes.DELETE("/:id", promotioncontroller.DeleteDiscountByID)
			}
		}

		// Promotions Routes
		purchaseRoutes := apiRoutes.Group("/purchases")
		{
			purchaseRoutes.POST("", purchasecontroller.CreatePurchaseOrder)
			purchaseRoutes.GET("", purchasecontroller.GetPurchaseOrders)

			purchaseRoutes.GET("/:id", purchasecontroller.GetPurchaseOrderByID)
			purchaseRoutes.PATCH("/:id", purchasecontroller.UpdatePurchaseOrderByID)
			purchaseRoutes.DELETE("/:id", purchasecontroller.DeletePurchaseOrderByID)

			// Good Recieve Routes
			goodReceiveRoutes := purchaseRoutes.Group("/good-receivings")
			{
				goodReceiveRoutes.POST("", purchasecontroller.CreateGoodsReceiving)
				goodReceiveRoutes.GET("", purchasecontroller.GetGoodsReceivings)

				goodReceiveRoutes.GET("/:id", purchasecontroller.GetGoodsReceivingByID)
				goodReceiveRoutes.PATCH("/:id", purchasecontroller.UpdateGoodsReceivingByID)
				goodReceiveRoutes.DELETE("/:id", purchasecontroller.DeleteGoodsReceivingByID)
			}
		}

	}

	return router

}
