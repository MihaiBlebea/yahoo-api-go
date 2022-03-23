package yahooapi

type IncomeStatement struct {
	EndDate                           int `json:"endDate"`
	TotalRevenue                      int `json:"totalRevenue"`
	CostOfRevenue                     int `json:"costOfRevenue"`
	GrossProfit                       int `json:"grossProfit"`
	ResearchDevelopment               int `json:"researchDevelopment"`
	SellingGeneralAdministrative      int `json:"sellingGeneralAdministrative"`
	NonRecurring                      int `json:"nonRecurring"`
	OtherOperatingExpenses            int `json:"otherOperatingExpenses"`
	TotalOperatingExpenses            int `json:"totalOperatingExpenses"`
	OperatingIncome                   int `json:"operatingIncome"`
	TotalOtherIncomeExpenseNet        int `json:"totalOtherIncomeExpenseNet"`
	Ebit                              int `json:"ebit"`
	InterestExpense                   int `json:"interestExpense"`
	IncomeBeforeTax                   int `json:"incomeBeforeTax"`
	IncomeTaxExpense                  int `json:"incomeTaxExpense"`
	MinorityInterest                  int `json:"minorityInterest"`
	NetIncomeFromContinuingOps        int `json:"netIncomeFromContinuingOps"`
	DiscontinuedOperations            int `json:"discontinuedOperations"`
	ExtraordinaryItems                int `json:"extraordinaryItems"`
	EffectOfAccountingCharges         int `json:"effectOfAccountingCharges"`
	OtherItems                        int `json:"otherItems"`
	NetIncome                         int `json:"netIncome"`
	NetIncomeApplicableToCommonShares int `json:"netIncomeApplicableToCommonShares"`
}

type History struct {
	EpsActual       float64 `json:"epsActual"`
	EpsEstimate     float64 `json:"epsEstimate"`
	EpsDifference   float64 `json:"epsDifference"`
	SurprisePercent float64 `json:"surprisePercent"`
	Quarter         float64 `json:"quarter"`
	Period          string  `json:"period"`
}

type BalanceSheet struct {
	EndDate                      int `json:"endDate"`
	Cash                         int `json:"cash"`
	ShortTermInvestments         int `json:"shortTermInvestments"`
	NetReceivables               int `json:"netReceivables"`
	Inventory                    int `json:"inventory"`
	TotalCurrentAssets           int `json:"totalCurrentAssets"`
	LongTermInvestments          int `json:"longTermInvestments"`
	PropertyPlantEquipment       int `json:"propertyPlantEquipment"`
	GoodWill                     int `json:"goodWill"`
	IntangibleAssets             int `json:"intangibleAssets"`
	OtherAssets                  int `json:"otherAssets"`
	DeferredLongTermAssetCharges int `json:"deferredLongTermAssetCharges"`
	TotalAssets                  int `json:"totalAssets"`
	AccountsPayable              int `json:"accountsPayable"`
	OtherCurrentLiab             int `json:"otherCurrentLiab"`
	LongTermDebt                 int `json:"longTermDebt"`
	OtherLiab                    int `json:"otherLiab"`
	TotalCurrentLiabilities      int `json:"totalCurrentLiabilities"`
	TotalLiab                    int `json:"totalLiab"`
	CommonStock                  int `json:"commonStock"`
	RetainedEarnings             int `json:"retainedEarnings"`
	TreasuryStock                int `json:"treasuryStock"`
	CapitalSurplus               int `json:"capitalSurplus"`
	OtherStockholderEquity       int `json:"otherStockholderEquity"`
	TotalStockholderEquity       int `json:"totalStockholderEquity"`
	NetTangibleAssets            int `json:"netTangibleAssets"`
}
