package rd_scrapper

type CompanyInfo struct {
	Nid               string `json:"nid"`
	EntrepreneurName  string `json:"entrepreneurName"`
	CompanyName       string `json:"companyName"`
	VatFlag           string `json:"vatFlag"`
	DocTaxInvoiceFlag string `json:"docTaxInvoiceFlag"`
	DocRecieptFlag    string `json:"docRecieptFlag"`
	RegisDate         string `json:"regisDate"`
	StartDate         string `json:"startDate"`
	EndDate           string `json:"endDate"`
	SourceFlag        string `json:"sourceFlag"`
	ActiveStatus      string `json:"activeStatus"`
	CreateDate        string `json:"createDate"`
	UpdateDate        string `json:"updateDate"`
	Remark            string `json:"remark"`
	IsicName          string `json:"isicName"`
	RegisDateTh       string `json:"regisDateTh"`
	EndDateTh         string `json:"endDateTh"`
	StartDateTh       string `json:"startDateTh"`
	EtaxInvoiceId     string `json:"etaxInvoiceId"`
	IsicCode          string `json:"isicCode"`
	Total             string `json:"total"`
}

type SearchRequest struct {
	Isic    string `json:"isic,omitempty"`
	TaxName string `json:"taxName,omitempty"`
	TaxNo   string `json:"taxNo,omitempty"`
	Index   int    `json:"index"`
	Page    int    `json:"page"`
}
