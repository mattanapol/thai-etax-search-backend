package main

import (
	"context"
	"github.com/mattanapol/thailand-etax-search-backend/internal/g_scrapper"
	"github.com/mattanapol/thailand-etax-search-backend/internal/mongo"
	"github.com/mattanapol/thailand-etax-search-backend/internal/rd_scrapper"
	"github.com/schollz/progressbar/v3"
	"log"
	"strings"
	"time"
)

// go run cmd/keyword-scrapper/main.go
func main() {
	ctx := context.Background()
	scrapperService := g_scrapper.NewService()
	//rdScrapperService := rd_scrapper.NewService()
	companyInfoRepository := mongo.NewCompanyRepository()

	//pullCompanyInfoFromRd(ctx, rdScrapperService, companyInfoRepository)
	getGoogleSearchResult(ctx, companyInfoRepository, scrapperService)
}

func getGoogleSearchResult(ctx context.Context,
	companyInfoRepository mongo.CompanyRepository,
	scrapperService g_scrapper.Service,
) {
	bar := progressbar.Default(4000)

	page := 1
	pageSize := 10
	bar.Add(page * pageSize)
	for {
		pagedResult, err := companyInfoRepository.GetPaged(ctx, page, pageSize)
		if err != nil {
			log.Fatal(err)
		}
		if len(pagedResult) == 0 {
			break
		}
		for _, companyInfo := range pagedResult {
			scrapperResultList, err := scrapperService.Scrap(companyInfo.EntrepreneurName)
			if err != nil {
				log.Fatal(err)
			}
			_, err = companyInfoRepository.Update(ctx,
				companyInfo.Nid,
				&mongo.CompanyInfo{
					Nid:          companyInfo.Nid,
					SearchResult: mapSearchResult(scrapperResultList),
				},
			)
			if err != nil {
				log.Fatal("error while update company info:", err)
			}
			bar.Add(1)
			time.Sleep(750 * time.Millisecond)
		}
		page++
	}
}

func pullCompanyInfoFromRd(ctx context.Context,
	rdScrapperService rd_scrapper.Service,
	companyInfoRepository mongo.CompanyRepository,
) {
	page := 1
	for {
		pagedResult, err := rdScrapperService.GetPagedCompanyInfo(page)
		if err != nil {
			log.Fatal(err)
		}
		if len(pagedResult) == 0 {
			break
		}
		for _, companyInfo := range pagedResult {
			log.Println("Saving company info: ", companyInfo)
			if exist, err := companyInfoRepository.ExistByTaxId(ctx, companyInfo.Nid); err != nil {
				log.Fatal("error while check exist company info:", err)
			} else if exist {
				_, err := companyInfoRepository.Update(ctx,
					companyInfo.Nid,
					mapCompanyInfo(companyInfo),
				)
				if err != nil {
					log.Fatal("error while update company info:", err)
				}
			} else {
				_, err := companyInfoRepository.
					Add(ctx,
						companyInfo.Nid,
						mapCompanyInfo(companyInfo),
					)
				if err != nil {
					log.Fatal("error while add company info:", err)
				}
			}
		}
		page++
	}
}

func mapCompanyInfo(companyInfo rd_scrapper.CompanyInfo) *mongo.CompanyInfo {
	return &mongo.CompanyInfo{
		Nid:               strings.TrimSpace(companyInfo.Nid),
		EntrepreneurName:  strings.TrimSpace(companyInfo.EntrepreneurName),
		CompanyName:       strings.TrimSpace(companyInfo.CompanyName),
		VatFlag:           strings.TrimSpace(companyInfo.VatFlag),
		DocTaxInvoiceFlag: strings.TrimSpace(companyInfo.DocTaxInvoiceFlag),
		DocReceiptFlag:    strings.TrimSpace(companyInfo.DocRecieptFlag),
		RegisDate:         strings.TrimSpace(companyInfo.RegisDate),
		StartDate:         strings.TrimSpace(companyInfo.StartDate),
		EndDate:           strings.TrimSpace(companyInfo.EndDate),
		SourceFlag:        strings.TrimSpace(companyInfo.SourceFlag),
		ActiveStatus:      strings.TrimSpace(companyInfo.ActiveStatus),
		CreateDate:        strings.TrimSpace(companyInfo.CreateDate),
		UpdateDate:        strings.TrimSpace(companyInfo.UpdateDate),
		Remark:            strings.TrimSpace(companyInfo.Remark),
		IsicName:          strings.TrimSpace(companyInfo.IsicName),
		RegisDateTh:       strings.TrimSpace(companyInfo.RegisDateTh),
		EndDateTh:         strings.TrimSpace(companyInfo.EndDateTh),
		StartDateTh:       strings.TrimSpace(companyInfo.StartDateTh),
		EtaxInvoiceId:     strings.TrimSpace(companyInfo.EtaxInvoiceId),
		IsicCode:          strings.TrimSpace(companyInfo.IsicCode),
		Total:             strings.TrimSpace(companyInfo.Total),
	}
}

func mapSearchResult(searchResult []g_scrapper.ScrapperResult) []mongo.GoogleSearch {
	var result []mongo.GoogleSearch
	for _, search := range searchResult {
		result = append(result, mongo.GoogleSearch{
			Title:       strings.TrimSpace(search.Title),
			Link:        strings.TrimSpace(search.Link),
			Description: strings.TrimSpace(search.Description),
			Rank:        search.Position,
		},
		)
	}
	return result
}
