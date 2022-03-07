package outlet
import(
	"github.com/daffashafwan/taskmajoo/business/outlets"
)
type Outlets struct {
	Id        int `gorm:"primaryKey"`
	Outlet_name      string
}

func (user *Outlets) ToDomain() outlets.Domain {
	return outlets.Domain{
		Id:        user.Id,
		OutletName: user.Outlet_name,
	}
}

func FromDomain(domain outlets.Domain) Outlets {
	return Outlets{
		Id:        domain.Id,
		Outlet_name: domain.OutletName,
	}
}