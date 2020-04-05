package service

type ImageCodeDTO struct {
	Code  string
	Phone string
}

type RedisKVDTO struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SaveRoleDTO struct {
	Map     map[string]*string `json:"map"`
	Account string             `json:"account"`
}

type StoreService struct {
	SaveImageCode  func(arg ImageCodeDTO) error
	CheckImageCode func(arg ImageCodeDTO) error

	ListLPop  func(key string) (string, error)
	ListRPop  func(key string) (string, error)
	ListLPush func(arg RedisKVDTO) error
	ListRPush func(arg RedisKVDTO) error
}
