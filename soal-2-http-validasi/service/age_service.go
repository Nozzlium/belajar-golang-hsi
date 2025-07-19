package service

type AgeService struct {
}

func NewAgeService() *AgeService {
	return &AgeService{}
}

func (controller *AgeService) Validate(email string, age int) bool {
	return email != "" && age >= 18
}
