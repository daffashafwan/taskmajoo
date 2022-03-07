package user
import(
	"github.com/daffashafwan/taskmajoo/business/users"
)
type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	User_name  string
	Password  string
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.User_name,
		Password:  user.Password,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:        domain.Id,
		Name:      domain.Name,
		Password:  domain.Password,
		User_name:  domain.Username,
	}
}