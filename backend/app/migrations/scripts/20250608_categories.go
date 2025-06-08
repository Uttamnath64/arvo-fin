package script

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

func categories(container *storage.Container) error {
	return RunOnce("20250608_categories", container.Config.ReadWriteDB, func(db *gorm.DB) error {
		categories := []models.Category{

			{Name: "Salary", Type: commonType.TransactionTypeIncome, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Freelance", Type: commonType.TransactionTypeIncome, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Investments", Type: commonType.TransactionTypeIncome, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Interest", Type: commonType.TransactionTypeIncome, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Gifts", Type: commonType.TransactionTypeIncome, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Rental Income", Type: commonType.TransactionTypeIncome, SourceID: 1, SourceType: commonType.UserTypeAdmin},

			// Expense Categories
			{Name: "Groceries", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Rent", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Utilities", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Transportation", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Dining Out", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Entertainment", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Healthcare", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Education", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
			{Name: "Subscriptions", Type: commonType.TransactionTypeExpense, SourceID: 1, SourceType: commonType.UserTypeAdmin},
		}
		for _, c := range categories {
			if err := db.FirstOrCreate(&c, models.Category{Name: c.Name, Type: c.Type, SourceID: c.SourceID, SourceType: c.SourceType}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
