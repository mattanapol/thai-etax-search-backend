package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CompanyRepository interface {
	ExistByTaxId(ctx context.Context, taxId string) (bool, error)
	Add(
		context context.Context,
		taxId string,
		product *CompanyInfo,
	) (*CompanyInfo, error)
	Update(
		context context.Context,
		taxId string,
		product *CompanyInfo,
	) (*CompanyInfo, error)
	GetPaged(context context.Context,
		pageNum int,
		pageSize int,
	) ([]CompanyInfo, error)
}

type repository struct {
	client *mongo.Collection
}

func newCompanyRepository(client *mongo.Database) CompanyRepository {
	return &repository{
		client: client.Collection("companyInfo"),
	}
}

func (r *repository) ExistByTaxId(ctx context.Context, taxId string) (bool, error) {
	count, err := r.client.CountDocuments(ctx,
		bson.D{{Key: "nid", Value: taxId}},
	)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return false, err
	}

	if int(count) <= 0 {
		return false, nil
	}

	return true, nil
}

func (r *repository) Add(
	context context.Context,
	taxId string,
	companyInfo *CompanyInfo,
) (*CompanyInfo, error) {
	if taxId == "" {
		return nil, fmt.Errorf("taxId is empty")
	}

	companyInfo.CreatedOn = time.Now()
	companyInfo.ModifiedOn = time.Now()

	_, err := r.client.InsertOne(context, companyInfo)
	if err != nil {
		return nil, err
	}

	return companyInfo, nil
}

func (r *repository) Update(
	context context.Context,
	taxId string,
	companyInfo *CompanyInfo,
) (*CompanyInfo, error) {
	if taxId == "" {
		return nil, fmt.Errorf("taxId is empty")
	}

	companyInfo.ModifiedOn = time.Now()

	_, err := r.client.UpdateOne(context,
		bson.D{{Key: "nid", Value: taxId}},
		bson.D{{Key: "$set", Value: companyInfo}},
	)
	if err != nil {
		return nil, err
	}

	return companyInfo, nil
}

func (r *repository) GetPaged(context context.Context,
	pageNum int,
	pageSize int,
) ([]CompanyInfo, error) {
	var (
		companyInfos []CompanyInfo
		err          error
		opt          options.FindOptions
		skip         = (pageNum - 1) * pageSize
	)

	opt.SetSkip(int64(skip))
	opt.SetLimit(int64(pageSize))

	cursor, err := r.client.Find(context, primitive.D{}, &opt)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if err = cursor.All(context, &companyInfos); err != nil {
		return nil, err
	}

	return companyInfos, nil
}

func getNoSearchResultFilter() bson.M {
	return bson.M{
		"$or": []bson.M{
			{"searchResult": bson.M{"$exists": false}},
			{"searchResult": bson.M{"$size": 0}},
		},
	}
}
