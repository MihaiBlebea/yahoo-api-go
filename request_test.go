package yahooapi_test

import (
	"testing"

	yahooapi "github.com/MihaiBlebea/yahoo-api-go"
)

func TestCache(t *testing.T) {
	req := yahooapi.NewHttpRequestClient(true)

	url := "https://query2.finance.yahoo.com/v10/finance/quoteSummary/AAPL?modules=incomeStatementHistoryQuarterly"
	req.GetResponseBody(url)
	cache := req.GetCache()

	if _, exists := cache[url]; !exists {
		t.Errorf("No cache found for %s", url)
	}
}
