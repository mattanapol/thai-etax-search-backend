package mongo

func NewCompanyRepository() CompanyRepository {
	configuration, err := newMongoDbConfiguration()
	if err != nil {
		return nil
	}
	db := SetupMongoDb(*configuration.Database)
	return newCompanyRepository(db)
}
