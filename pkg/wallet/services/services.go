package services

func NewServices() *Services {
	return &Services{
		WalletSvc:      NewWalletService(),
		TransactionSvc: NewTransactionService(),
	}
}
