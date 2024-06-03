package external

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientSuit struct {
	suite.Suite
	client *Client
}

func TestClientSuit(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	suite.Run(t, new(ClientSuit))
}

func (suite *ClientSuit) SetupSuite() {
	suite.client = NewClient("https://www.oklink.com", "baeffec3-4eb6-4346-9dc3-f1bc8f4345cd")
}

func (suite *ClientSuit) TestGetUserBtcBalance() {
	endpoint := "/api/v5/explorer/address/address-summary"
	method := "GET"
	params := map[string]string{
		"chainShortName": "btc",
		"address":        "bc1p3g3ly2z2lm2vwmgvvvssv97ttasaqpyms6l6p82qhgledrpqa7lqx4csda",
	}
	var result map[string]interface{}
	err := suite.client.SendRequestAndParseJSON(method, endpoint, params, &result)
	suite.NoError(err)
	fmt.Println("Response:", result)
}

// 获取用户的 ordinals nft BtcName NFT余额
func (suite *ClientSuit) TestGetUserOrdinalNftBalance() {
	endpoint := "/api/v5/explorer/inscription/address-token-list"
	method := "GET"
	params := map[string]string{
		"chainShortName": "btc",
		"protocolType":   "ordinals_nft",
		"symbol":         "BtcName",
		"address":        "bc1p3g3ly2z2lm2vwmgvvvssv97ttasaqpyms6l6p82qhgledrpqa7lqx4csda",
		"limit":          "10",
	}
	var result map[string]interface{}
	err := suite.client.SendRequestAndParseJSON(method, endpoint, params, &result)
	suite.NoError(err)
	fmt.Println("Response:", result)
}

// btcname最新交易记录
func (suite *ClientSuit) TestGetUserOrdinalNftTransfer() {
	endpoint := "/api/v5/explorer/inscription/token-transaction-list"
	method := "GET"
	params := map[string]string{
		"chainShortName": "btc",
		"protocolType":   "ordinals_nft",
		"symbol":         "BtcName",
		"limit":          "20",
		//"startTime":      "",
		//"endTime":        "",
	}
	var result map[string]interface{}
	err := suite.client.SendRequestAndParseJSON(method, endpoint, params, &result)
	suite.NoError(err)
	fmt.Println("Response:", result)
}

// 查询铭文列表
// 查询铭文列表，按照InscriptionNumber倒序排列，最多返回10000条数据，支持BTC链的BRC-20、Ordinals NFT
func (suite *ClientSuit) TestGetUserOrdinalNftList() {
	endpoint := "/api/v5/explorer/inscription/inscription-list"
	method := "GET"
	params := map[string]string{
		"chainShortName": "btc",
		"protocolType":   "ordinals_nft",
		"symbol":         "BtcName",
		"limit":          "100",
	}
	var result map[string]interface{}
	err := suite.client.SendRequestAndParseJSON(method, endpoint, params, &result)
	suite.NoError(err)
	fmt.Println("Response:", result)
}

func (suite *ClientSuit) TestGetUserBrc20List() {
	endpoint := "/api/v5/explorer/inscription/address-inscription-list"
	method := "GET"
	params := map[string]string{
		"chainShortName": "btc",
		"protocolType":   "brc20",
		"address":        "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d",
		"limit":          "20",
	}
	var result map[string]interface{}
	err := suite.client.SendRequestAndParseJSON(method, endpoint, params, &result)
	suite.NoError(err)
	fmt.Println("Response:", result)
}

// -------------------------------------nftscan-------------------------
func (suite *ClientSuit) TestGetUserOrdinalNft() {
	suite.client = NewClient("https://btcapi.nftscan.com/api/btc/account/own/bc1p3g3ly2z2lm2vwmgvvvssv97ttasaqpyms6l6p82qhgledrpqa7lqx4csda", "D5y8mrZd4un91RjKUDWfsJVN")
	method := "GET"
	params := map[string]string{
		"show_attribute": "false",
	}
	var result map[string]interface{}
	err := suite.client.SendRequestAndParseJSON(method, "", params, &result)
	suite.NoError(err)
	fmt.Println("Response:", result)
}
