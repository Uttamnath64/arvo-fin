package script

import (
	commonType "github.com/Uttamnath64/arvo-fin/app/common/types"
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

func avatars(container *storage.Container) error {
	return RunOnce("20250608_avatars", container.Config.ReadWriteDB, func(db *gorm.DB) error {
		avatars := []models.Avatar{

			// AvatarTypeDefault
			{Name: "Star", Icon: "⭐", Type: commonType.AvatarTypeDefault},
			{Name: "Sparkles", Icon: "✨", Type: commonType.AvatarTypeDefault},
			{Name: "Fire", Icon: "🔥", Type: commonType.AvatarTypeDefault},
			{Name: "Heart", Icon: "❤️", Type: commonType.AvatarTypeDefault},
			{Name: "Sun", Icon: "☀️", Type: commonType.AvatarTypeDefault},
			{Name: "Moon", Icon: "🌙", Type: commonType.AvatarTypeDefault},
			{Name: "Globe", Icon: "🌍", Type: commonType.AvatarTypeDefault},

			// AvatarTypeUser
			{Name: "Person", Icon: "🧑", Type: commonType.AvatarTypeUser},
			{Name: "Technologist", Icon: "🧑‍💻", Type: commonType.AvatarTypeUser},
			{Name: "Student", Icon: "🎓", Type: commonType.AvatarTypeUser},
			{Name: "Chef", Icon: "🧑‍🍳", Type: commonType.AvatarTypeUser},
			{Name: "Artist", Icon: "🧑‍🎨", Type: commonType.AvatarTypeUser},
			{Name: "Musician", Icon: "🧑‍🎤", Type: commonType.AvatarTypeUser},
			{Name: "Doctor", Icon: "🧑‍⚕️", Type: commonType.AvatarTypeUser},
			{Name: "Police", Icon: "🧑‍✈️", Type: commonType.AvatarTypeUser},
			{Name: "Guard", Icon: "🛡️", Type: commonType.AvatarTypeUser},

			// AvatarTypeCategory
			{Name: "Money Bag", Icon: "💰", Type: commonType.AvatarTypeCategory},
			{Name: "Shopping Bag", Icon: "🛍️", Type: commonType.AvatarTypeCategory},
			{Name: "Cart", Icon: "🛒", Type: commonType.AvatarTypeCategory},
			{Name: "Burger", Icon: "🍔", Type: commonType.AvatarTypeCategory},
			{Name: "Pizza", Icon: "🍕", Type: commonType.AvatarTypeCategory},
			{Name: "Sushi", Icon: "🍣", Type: commonType.AvatarTypeCategory},
			{Name: "Car", Icon: "🚗", Type: commonType.AvatarTypeCategory},
			{Name: "Fuel", Icon: "⛽", Type: commonType.AvatarTypeCategory},
			{Name: "Travel", Icon: "✈️", Type: commonType.AvatarTypeCategory},
			{Name: "Movie", Icon: "🎬", Type: commonType.AvatarTypeCategory},
			{Name: "Game", Icon: "🎮", Type: commonType.AvatarTypeCategory},
			{Name: "Music", Icon: "🎵", Type: commonType.AvatarTypeCategory},
			{Name: "Hospital", Icon: "🏥", Type: commonType.AvatarTypeCategory},
			{Name: "Gym", Icon: "🏋️", Type: commonType.AvatarTypeCategory},
			{Name: "Gift", Icon: "🎁", Type: commonType.AvatarTypeCategory},
			{Name: "Book", Icon: "📚", Type: commonType.AvatarTypeCategory},
			{Name: "Pet", Icon: "🐶", Type: commonType.AvatarTypeCategory},
			{Name: "Plant", Icon: "🪴", Type: commonType.AvatarTypeCategory},
			{Name: "Tools", Icon: "🛠️", Type: commonType.AvatarTypeCategory},

			// AvatarTypePortfolio
			{Name: "Bank", Icon: "🏦", Type: commonType.AvatarTypePortfolio},
			{Name: "Wallet", Icon: "👛", Type: commonType.AvatarTypePortfolio},
			{Name: "Briefcase", Icon: "💼", Type: commonType.AvatarTypePortfolio},
			{Name: "Chart", Icon: "📈", Type: commonType.AvatarTypePortfolio},
			{Name: "Piggy Bank", Icon: "🐷", Type: commonType.AvatarTypePortfolio},
			{Name: "Safe", Icon: "🧱", Type: commonType.AvatarTypePortfolio},
			{Name: "Gem", Icon: "💎", Type: commonType.AvatarTypePortfolio},
			{Name: "Gold Coin", Icon: "🪙", Type: commonType.AvatarTypePortfolio},
			{Name: "Stock Market", Icon: "🏛️", Type: commonType.AvatarTypePortfolio},
			{Name: "Cash", Icon: "💵", Type: commonType.AvatarTypePortfolio},
			{Name: "Credit", Icon: "💳", Type: commonType.AvatarTypePortfolio},
			{Name: "Loan", Icon: "🤝", Type: commonType.AvatarTypePortfolio},

			// AvatarTypeRegularPayment
			{Name: "SIP", Icon: "📅", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Mutual Fund", Icon: "📊", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Loan EMI", Icon: "💸", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Installment", Icon: "🔁", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Insurance", Icon: "🛡️", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Recurring Deposit", Icon: "💰", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Auto Save", Icon: "🤖", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Gold Saving", Icon: "🥇", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Stock SIP", Icon: "📈", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Emergency Fund", Icon: "🚨", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Piggy Save", Icon: "🐷", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Credit Payment", Icon: "💳", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Debt Payoff", Icon: "📉", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Random Save", Icon: "🎯", Type: commonType.AvatarTypeRegularPayment},
			{Name: "Stool Saving", Icon: "🪑", Type: commonType.AvatarTypeRegularPayment},
		}
		for _, a := range avatars {
			if err := db.FirstOrCreate(&a, models.Avatar{Name: a.Name, Type: a.Type}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
