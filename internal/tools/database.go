package tools

//Database Collections

type LoginDetails struct {
	Username  string
	AuthToken string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err = database.SetupDatabase()

	if err != nil {
		return nil, err
	}

	return &database, nil
}
