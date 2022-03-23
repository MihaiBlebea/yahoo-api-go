package yahooapi_test

import (
	"testing"

	yahooapi "github.com/MihaiBlebea/yahoo-api-go"
)

func getClient() *yahooapi.Client {
	client := yahooapi.NewClient(true)
	client.SetRequestClient(&yahooapi.JsonRequest{})

	return client
}

func TestIncomeStatementHistoryQuarterly(t *testing.T) {
	client := getClient()

	res, err := client.IncomeStatementHistoryQuarterly("AAPL")
	if err != nil {
		t.Error(err)
	}

	statement := res[0]
	if statement.CostOfRevenue != 2644000000 {
		t.Errorf("expected period to be 2644000000, got %d", statement.CostOfRevenue)
	}
}

func TestEarningsHistory(t *testing.T) {
	client := getClient()

	res, err := client.EarningsHistory("AAPL")
	if err != nil {
		t.Error(err)
	}

	history := res[0]
	if history.Period != "-4q" {
		t.Errorf("expected period to be -4q, got %s", history.Period)
	}
}

func TestBalnceSheetHistoryQuarterly(t *testing.T) {
	client := getClient()

	res, err := client.BalanceSheetHistoryQuarterly("AAPL")
	if err != nil {
		t.Error(err)
	}

	balanceSheet := res[0]
	if balanceSheet.CapitalSurplus != 10385000000 {
		t.Errorf("expected capital surplus to be 10385000000, got %d", balanceSheet.CapitalSurplus)
	}
}
