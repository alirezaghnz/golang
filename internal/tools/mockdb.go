package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Alireza": {
		AuthToken: "123ABC",
		Username:  "alireza",
	},
	"Amir": {
		AuthToken: "456DEF",
		Username:  "amir",
	},
	"Hossein": {
		AuthToken: "789GHI",
		Username:  "hossein",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alireza": {
		Coins:   100,
		Username: "alireza",
	},
	"amir": {
		Coins:   200,
		Username: "amir",
	},
	"hossein": {
		Coins:   300,
		Username: "hossein",
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