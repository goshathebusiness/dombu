package services

func NewServices() *Services {
	return &Services{
		UserSvc: NewUserService(),
	}
}
