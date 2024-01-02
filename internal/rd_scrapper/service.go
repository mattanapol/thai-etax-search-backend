package rd_scrapper

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Service interface {
	GetAllCompanyInfo() ([]CompanyInfo, error)
	GetPagedCompanyInfo(page int) ([]CompanyInfo, error)
}

type service struct {
}

func (s *service) GetAllCompanyInfo() ([]CompanyInfo, error) {
	var result []CompanyInfo
	page := 1
	for {
		pagedResult, err := s.GetPagedCompanyInfo(page)
		if err != nil {
			return nil, err
		}
		if len(pagedResult) == 0 {
			break
		}
		result = append(result, pagedResult...)
		page++
	}
	return result, nil
}

func (s *service) GetPagedCompanyInfo(page int) ([]CompanyInfo, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(SearchRequest{
			Index: page,
			Page:  10,
		},
		).
		Post("https://efiling.rd.go.th/rd-questionnaire-service/etax/search")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("error while get company info: %s", resp.Status())
	}
	var result []CompanyInfo
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
