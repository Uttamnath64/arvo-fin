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
			// Income (use appropriate category icons)
			{Name: "Salary", Type: commonType.TransactionTypeIncome, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 16},        // 💰
			{Name: "Freelance", Type: commonType.TransactionTypeIncome, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 17},     // 🛍️
			{Name: "Investments", Type: commonType.TransactionTypeIncome, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 24},   // ✈️
			{Name: "Interest", Type: commonType.TransactionTypeIncome, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 30},      // 🎁
			{Name: "Gifts", Type: commonType.TransactionTypeIncome, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 30},         // 🎁
			{Name: "Rental Income", Type: commonType.TransactionTypeIncome, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 23}, // ⛽

			// Expense
			{Name: "Groceries", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 18},      // 🛒
			{Name: "Rent", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 16},           // 💰
			{Name: "Utilities", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 34},      // 🛠️
			{Name: "Transportation", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 22}, // 🚗
			{Name: "Dining Out", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 19},     // 🍔
			{Name: "Entertainment", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 25},  // 🎬
			{Name: "Healthcare", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 28},     // 🏥
			{Name: "Education", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 31},      // 📚
			{Name: "Subscriptions", Type: commonType.TransactionTypeExpense, SourceId: 1, SourceType: commonType.UserTypeAdmin, AvatarId: 27},  // 🎵
		}
		for _, c := range categories {
			if err := db.FirstOrCreate(&c, models.Category{Name: c.Name, Type: c.Type, SourceId: c.SourceId, SourceType: c.SourceType}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
