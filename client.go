package yahooapi

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

var instance *Client

type Client struct {
	baseUrl string
	request Request
}

func NewClient(request Request) *Client {
	if instance != nil {
		fmt.Println("Getting from instance")
		return instance
	}

	client := &Client{
		request: request,
		baseUrl: "https://query2.finance.yahoo.com/v10/finance/quoteSummary",
	}

	instance = client
	fmt.Println("Getting a new instance")

	return client
}

func (c *Client) IncomeStatementHistoryQuarterly(symbol string) ([]IncomeStatement, error) {
	url := fmt.Sprintf(
		"%s/%s?modules=%s",
		c.baseUrl,
		strings.ToUpper(symbol),
		"incomeStatementHistoryQuarterly",
	)

	body, err := c.request.GetResponseBody(url)
	if err != nil {
		return []IncomeStatement{}, err
	}

	statements := []IncomeStatement{}
	data := gjson.Get(string(body), "quoteSummary.result.0.incomeStatementHistoryQuarterly.incomeStatementHistory")

	data.ForEach(func(key, value gjson.Result) bool {
		statements = append(statements, IncomeStatement{
			EndDate:                           int(value.Get("endDate.raw").Int()),
			TotalRevenue:                      int(value.Get("totalRevenue.raw").Int()),
			CostOfRevenue:                     int(value.Get("costOfRevenue.raw").Int()),
			GrossProfit:                       int(value.Get("grossProfit.raw").Int()),
			ResearchDevelopment:               int(value.Get("researchDevelopment.raw").Int()),
			SellingGeneralAdministrative:      int(value.Get("sellingGeneralAdministrative.raw").Int()),
			NonRecurring:                      int(value.Get("nonRecurring.raw").Int()),
			OtherOperatingExpenses:            int(value.Get("otherOperatingExpenses.raw").Int()),
			TotalOperatingExpenses:            int(value.Get("totalOperatingExpenses.raw").Int()),
			OperatingIncome:                   int(value.Get("operatingIncome.raw").Int()),
			TotalOtherIncomeExpenseNet:        int(value.Get("totalOtherIncomeExpenseNet.raw").Int()),
			Ebit:                              int(value.Get("ebit.raw").Int()),
			InterestExpense:                   int(value.Get("interestExpense.raw").Int()),
			IncomeBeforeTax:                   int(value.Get("incomeBeforeTax.raw").Int()),
			IncomeTaxExpense:                  int(value.Get("incomeTaxExpense.raw").Int()),
			MinorityInterest:                  int(value.Get("minorityInterest.raw").Int()),
			NetIncomeFromContinuingOps:        int(value.Get("netIncomeFromContinuingOps.raw").Int()),
			DiscontinuedOperations:            int(value.Get("discontinuedOperations.raw").Int()),
			ExtraordinaryItems:                int(value.Get("extraordinaryItems.raw").Int()),
			EffectOfAccountingCharges:         int(value.Get("effectOfAccountingCharges.raw").Int()),
			OtherItems:                        int(value.Get("otherItems.raw").Int()),
			NetIncome:                         int(value.Get("netIncome.raw").Int()),
			NetIncomeApplicableToCommonShares: int(value.Get("netIncomeApplicableToCommonShares.raw").Int()),
		})

		return true
	})

	return statements, nil
}

func (c *Client) EarningsHistory(symbol string) ([]History, error) {
	url := fmt.Sprintf(
		"%s/%s?modules=%s",
		c.baseUrl,
		strings.ToUpper(symbol),
		"earningsHistory",
	)

	body, err := c.request.GetResponseBody(url)
	if err != nil {
		return []History{}, err
	}

	histories := []History{}
	history := gjson.Get(string(body), "quoteSummary.result.0.earningsHistory.history")

	history.ForEach(func(key, value gjson.Result) bool {
		var history History
		history.EpsActual = value.Get("epsActual.raw").Float()
		history.EpsEstimate = value.Get("epsEstimate.raw").Float()
		history.EpsDifference = value.Get("epsDifference.raw").Float()
		history.SurprisePercent = value.Get("surprisePercent.raw").Float()
		history.Quarter = value.Get("quarter.raw").Float()
		history.Period = value.Get("period").String()

		histories = append(histories, history)

		return true
	})

	return histories, nil
}

func (c *Client) BalanceSheetHistoryQuarterly(symbol string) ([]BalanceSheet, error) {
	url := fmt.Sprintf(
		"%s/%s?modules=%s",
		c.baseUrl,
		strings.ToUpper(symbol),
		"balanceSheetHistoryQuarterly",
	)

	body, err := c.request.GetResponseBody(url)
	if err != nil {
		return []BalanceSheet{}, err
	}

	balanceSheets := []BalanceSheet{}
	data := gjson.Get(string(body), "quoteSummary.result.0.balanceSheetHistoryQuarterly.balanceSheetStatements")

	data.ForEach(func(key, value gjson.Result) bool {
		balanceSheets = append(balanceSheets, BalanceSheet{
			EndDate:                      int(value.Get("endDate.raw").Int()),
			Cash:                         int(value.Get("cash.raw").Int()),
			ShortTermInvestments:         int(value.Get("shortTermInvestments.raw").Int()),
			NetReceivables:               int(value.Get("netReceivables.raw").Int()),
			Inventory:                    int(value.Get("inventory.raw").Int()),
			TotalCurrentAssets:           int(value.Get("totalCurrentAssets.raw").Int()),
			LongTermInvestments:          int(value.Get("longTermInvestments.raw").Int()),
			PropertyPlantEquipment:       int(value.Get("propertyPlantEquipment.raw").Int()),
			GoodWill:                     int(value.Get("goodWill.raw").Int()),
			IntangibleAssets:             int(value.Get("intangibleAssets.raw").Int()),
			OtherAssets:                  int(value.Get("otherAssets.raw").Int()),
			DeferredLongTermAssetCharges: int(value.Get("deferredLongTermAssetCharges.raw").Int()),
			TotalAssets:                  int(value.Get("totalAssets.raw").Int()),
			AccountsPayable:              int(value.Get("accountsPayable.raw").Int()),
			OtherCurrentLiab:             int(value.Get("otherCurrentLiab.raw").Int()),
			LongTermDebt:                 int(value.Get("longTermDebt.raw").Int()),
			OtherLiab:                    int(value.Get("otherLiab.raw").Int()),
			TotalCurrentLiabilities:      int(value.Get("totalCurrentLiabilities.raw").Int()),
			TotalLiab:                    int(value.Get("totalLiab.raw").Int()),
			CommonStock:                  int(value.Get("commonStock.raw").Int()),
			RetainedEarnings:             int(value.Get("retainedEarnings.raw").Int()),
			TreasuryStock:                int(value.Get("treasuryStock.raw").Int()),
			CapitalSurplus:               int(value.Get("capitalSurplus.raw").Int()),
			OtherStockholderEquity:       int(value.Get("otherStockholderEquity.raw").Int()),
			TotalStockholderEquity:       int(value.Get("totalStockholderEquity.raw").Int()),
			NetTangibleAssets:            int(value.Get("netTangibleAssets.raw").Int()),
		})

		return true
	})

	return balanceSheets, nil
}
