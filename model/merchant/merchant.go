package merchant
import(
	"github.com/daffashafwan/taskmajoo/business/merchants"
)
type Merchants struct {
	Id        int `gorm:"primaryKey"`
	User_id      int
	Merchant_name  string
}

func (user *Merchants) ToDomain() merchants.Domain {
	return merchants.Domain{
		Id:        user.Id,
		UserId:      user.User_id,
		MerchantName:  user.Merchant_name,
	}
}

func FromDomain(domain merchants.Domain) Merchants {
	return Merchants{
		Id:        domain.Id,
		User_id:      domain.UserId,
		Merchant_name:  domain.MerchantName,
	}
}