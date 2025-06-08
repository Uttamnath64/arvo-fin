package script

import (
	"github.com/Uttamnath64/arvo-fin/app/models"
	"github.com/Uttamnath64/arvo-fin/app/storage"
	"gorm.io/gorm"
)

func currencies(container *storage.Container) error {
	return RunOnce("20250608_currencies", container.Config.ReadWriteDB, func(db *gorm.DB) error {
		currencies := []models.Currency{
			{Code: "INR", Name: "Indian Rupee", Symbol: "₹"},
			{Code: "USD", Name: "United States Dollar", Symbol: "$"},
			{Code: "EUR", Name: "Euro", Symbol: "€"},
			{Code: "GBP", Name: "British Pound Sterling", Symbol: "£"},
			{Code: "JPY", Name: "Japanese Yen", Symbol: "¥"},
			{Code: "CNY", Name: "Chinese Yuan", Symbol: "¥"},
			{Code: "AUD", Name: "Australian Dollar", Symbol: "A$"},
			{Code: "CAD", Name: "Canadian Dollar", Symbol: "C$"},
			{Code: "CHF", Name: "Swiss Franc", Symbol: "CHF"},
			{Code: "SGD", Name: "Singapore Dollar", Symbol: "S$"},
			{Code: "NZD", Name: "New Zealand Dollar", Symbol: "NZ$"},
			{Code: "HKD", Name: "Hong Kong Dollar", Symbol: "HK$"},
			{Code: "KRW", Name: "South Korean Won", Symbol: "₩"},
			{Code: "SEK", Name: "Swedish Krona", Symbol: "kr"},
			{Code: "NOK", Name: "Norwegian Krone", Symbol: "kr"},
			{Code: "DKK", Name: "Danish Krone", Symbol: "kr"},
			{Code: "ZAR", Name: "South African Rand", Symbol: "R"},
			{Code: "BRL", Name: "Brazilian Real", Symbol: "R$"},
			{Code: "RUB", Name: "Russian Ruble", Symbol: "₽"},
			{Code: "MXN", Name: "Mexican Peso", Symbol: "$"},
			{Code: "THB", Name: "Thai Baht", Symbol: "฿"},
			{Code: "MYR", Name: "Malaysian Ringgit", Symbol: "RM"},
			{Code: "IDR", Name: "Indonesian Rupiah", Symbol: "Rp"},
			{Code: "PHP", Name: "Philippine Peso", Symbol: "₱"},
			{Code: "VND", Name: "Vietnamese Dong", Symbol: "₫"},
			{Code: "AED", Name: "UAE Dirham", Symbol: "د.إ"},
			{Code: "SAR", Name: "Saudi Riyal", Symbol: "﷼"},
			{Code: "TRY", Name: "Turkish Lira", Symbol: "₺"},
			{Code: "PLN", Name: "Polish Zloty", Symbol: "zł"},
			{Code: "HUF", Name: "Hungarian Forint", Symbol: "Ft"},
			{Code: "CZK", Name: "Czech Koruna", Symbol: "Kč"},
			{Code: "ILS", Name: "Israeli New Shekel", Symbol: "₪"},
			{Code: "EGP", Name: "Egyptian Pound", Symbol: "E£"},
			{Code: "NGN", Name: "Nigerian Naira", Symbol: "₦"},
			{Code: "PKR", Name: "Pakistani Rupee", Symbol: "₨"},
			{Code: "BDT", Name: "Bangladeshi Taka", Symbol: "৳"},
			{Code: "LKR", Name: "Sri Lankan Rupee", Symbol: "Rs"},
			{Code: "KES", Name: "Kenyan Shilling", Symbol: "KSh"},
			{Code: "GHS", Name: "Ghanaian Cedi", Symbol: "₵"},
			{Code: "UAH", Name: "Ukrainian Hryvnia", Symbol: "₴"},
			{Code: "ARS", Name: "Argentine Peso", Symbol: "$"},
			{Code: "CLP", Name: "Chilean Peso", Symbol: "$"},
			{Code: "COP", Name: "Colombian Peso", Symbol: "$"},
			{Code: "TWD", Name: "New Taiwan Dollar", Symbol: "NT$"},
			{Code: "MAD", Name: "Moroccan Dirham", Symbol: "د.م."},
		}
		for _, c := range currencies {
			if err := db.FirstOrCreate(&c, models.Currency{Code: c.Code}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
