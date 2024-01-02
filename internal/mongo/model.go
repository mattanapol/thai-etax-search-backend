package mongo

import "time"

type CompanyInfo struct {
	CreatedOn  time.Time `bson:"created_on"`
	ModifiedOn time.Time `bson:"modified_on"`

	Nid               string         `bson:"nid,omitempty"`
	EntrepreneurName  string         `bson:"entrepreneurName,omitempty"`
	CompanyName       string         `bson:"companyName,omitempty"`
	VatFlag           string         `bson:"vatFlag,omitempty"`
	DocTaxInvoiceFlag string         `bson:"docTaxInvoiceFlag,omitempty"`
	DocReceiptFlag    string         `bson:"docReceiptFlag,omitempty"`
	RegisDate         string         `bson:"regisDate,omitempty"`
	StartDate         string         `bson:"startDate,omitempty"`
	EndDate           string         `bson:"endDate,omitempty"`
	SourceFlag        string         `bson:"sourceFlag,omitempty"`
	ActiveStatus      string         `bson:"activeStatus,omitempty"`
	CreateDate        string         `bson:"createDate,omitempty"`
	UpdateDate        string         `bson:"updateDate,omitempty"`
	Remark            string         `bson:"remark,omitempty"`
	IsicName          string         `bson:"isicName,omitempty"`
	RegisDateTh       string         `bson:"regisDateTh,omitempty"`
	EndDateTh         string         `bson:"endDateTh,omitempty"`
	StartDateTh       string         `bson:"startDateTh,omitempty"`
	EtaxInvoiceId     string         `bson:"etaxInvoiceId,omitempty"`
	IsicCode          string         `bson:"isicCode,omitempty"`
	Total             string         `bson:"total,omitempty"`
	SearchResult      []GoogleSearch `bson:"searchResult,omitempty"`
}

type GoogleSearch struct {
	Title       string `bson:"title,omitempty"`
	Link        string `bson:"link,omitempty"`
	Description string `bson:"description,omitempty"`
	Rank        int    `bson:"rank"`
}
