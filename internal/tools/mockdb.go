package tools

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		Username:  "alex",
		AuthToken: "1234",
	},
	"jane": {
		Username:  "jane",
		AuthToken: "5678",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jane": {
		Coins:    200,
		Username: "jane",
	},
}

func (m *mockDB) GetLoginDetails(username string) *LoginDetails {

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData

}

func (m *mockDB) GetUserCoins(username string) *CoinDetails {

	var coinData = CoinDetails{}

	coinData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &coinData

}

func (m *mockDB) SetupDatabase() error {
	return nil
}
