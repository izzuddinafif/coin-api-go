package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"emma": {
		AuthToken: "456DEF",
		Username:  "emma",
	},
	"charlie": {
		AuthToken: "789GHI",
		Username:  "charlie",
	},
	"diana": {
		AuthToken: "101JKL",
		Username:  "diana",
	},
	"frank": {
		AuthToken: "202MNO",
		Username:  "frank",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"emma": {
		Coins:    75,
		Username: "emma",
	},
	"charlie": {
		Coins:    150,
		Username: "charlie",
	},
	"diana": {
		Coins:    200,
		Username: "diana",
	},
	"frank": {
		Coins:    50,
		Username: "frank",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
