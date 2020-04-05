package service

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"time"
)

type RegisterDTO struct {
	vo.UserVO
	Vcode string `json:"vcode"` //短信验证码
}

type LoginDTO struct {
	Phone string `json:"phone"`
	PWD   string `json:"pwd"`
}
type LoginResult struct {
	AccessToken string `json:"access_token"`
}

type AddressPageDTO struct {
	vo.Pageable
	vo.UserAddress
}

type SetDefAddressDTO struct {
	UserId     string `json:"user_id"`
	AdddressId string `json:"adddress_id"`
}

type UserPageDTO struct {
	vo.Pageable
	Id       string `json:"id"`
	Phone    string `json:"phone"`
	Channel  string `json:"channel"`
	Realname string `json:"realname"`

	InvitationCode string `json:"invitation_code"`
	InviterCode    string `json:"inviter_code"`

	TimeStart *time.Time `json:"time_start"`
	TimeEnd   *time.Time `json:"time_end"`
}

type UserService struct {
	Register          func(arg RegisterDTO) error
	Login             func(arg LoginDTO) (LoginResult, error)
	AddressPage       func(arg AddressPageDTO) (vo.PageVO, error)
	AddAddress        func(arg vo.UserAddress) error
	UpdateAddress     func(arg vo.UserAddress) error
	DeleteAddress     func(adddressId string) error
	SetDefaultAddress func(arg SetDefAddressDTO) error
	Find              func(id string) (result vo.UserVO, e error)
	Finds             func(ids []string) (map[string]*vo.UserVO, error)

	FindByPhone  func(phone string) (result vo.UserVO, e error)
	FindByPhones func(phones []string) ([]vo.UserVO, error)

	FindByInvitationCode func(inviter_code string) (vo.UserVO, error)
	Update               func(arg vo.UserVO) error
	UpdatePwd            func(arg vo.UserVO) error

	FindUserAddress func(address_id string) (vo.UserAddressVO, error)

	Page func(arg UserPageDTO) (vo.PageVO, error)
}
