//go:build integrationtestv5asset

package integrationtestv5asset

import (
	"testing"

	"github.com/shaykinanton/bybit/v2"
	"github.com/shaykinanton/bybit/v2/integrationtest/testhelper"
	"github.com/stretchr/testify/require"
)

func TestGetInternalTransferRecords(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	limit := 1
	res, err := client.V5().Asset().GetInternalTransferRecords(bybit.V5GetInternalTransferRecordsParam{
		Limit: &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-asset-get-internal-transfer-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetDepositRecords(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	limit := 1
	res, err := client.V5().Asset().GetDepositRecords(bybit.V5GetDepositRecordsParam{
		Limit: &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-asset-get-deposit-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetSubDepositRecords(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	limit := 1
	res, err := client.V5().Asset().GetSubDepositRecords(bybit.V5GetSubDepositRecordsParam{
		SubMemberID: "1462488",
		Limit:       &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-asset-get-sub-deposit-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetInternalDepositRecords(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	limit := 1
	res, err := client.V5().Asset().GetInternalDepositRecords(bybit.V5GetInternalDepositRecordsParam{
		Limit: &limit,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-asset-get-internal-deposit-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetWithdrawalRecords(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	limit := 1
	typ := bybit.WithdrawTypeAll
	res, err := client.V5().Asset().GetWithdrawalRecords(bybit.V5GetWithdrawalRecordsParam{
		Limit:        &limit,
		WithdrawType: &typ,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-asset-get-withdrawal-records.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetCoinInfo(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	coin := bybit.CoinBTC
	res, err := client.V5().Asset().GetCoinInfo(bybit.V5GetCoinInfoParam{
		Coin: &coin,
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-asset-get-coin-info.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}

func TestGetAllCoinsBalance(t *testing.T) {
	client := bybit.NewTestClient().WithAuthFromEnv()
	res, err := client.V5().Asset().GetAllCoinsBalance(bybit.V5GetAllCoinsBalanceParam{
		AccountType: bybit.AccountTypeUnified,
		Coins:       []bybit.Coin{bybit.CoinBTC},
	})
	require.NoError(t, err)
	{
		goldenFilename := "./testdata/v5-asset-get-all-coins-balance.json"
		testhelper.Compare(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
		testhelper.UpdateFile(t, goldenFilename, testhelper.ConvertToJSON(res.Result))
	}
}
