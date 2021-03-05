package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models"
	companyModels "github.com/nikola43/ecoapigorm/models/company"
	_ "github.com/nikola43/ecoapigorm/models/employee"
	_ "github.com/nikola43/ecoapigorm/utils"
)

func CreateCompany(createEmployeeRequest *companyModels.CreateCompanyRequest) (*companyModels.CreateCompanyResponse, error) {
	company := models.Company{
		Name:       createEmployeeRequest.Name,
		EmployeeID: createEmployeeRequest.EmployeeID,
	}

	if err := database.GormDB.Create(&company).Error; err != nil {
		return nil, err
	}

	createEmployeeResponse := &companyModels.CreateCompanyResponse{
		ID:         company.ID,
		Name:       company.Name,
		EmployeeID: company.EmployeeID,
	}

	return createEmployeeResponse, nil
}

func GetCompanies() ([]models.Company, error) {
	var list []models.Company

	GormDBResult := database.GormDB.Find(list)
	if GormDBResult.Error != nil {
		return nil, GormDBResult.Error
	}

	return list, nil
}

func GetCompanyByID(id uint) (*models.Company, error) {
	Company := models.Company{}

	if err := database.GormDB.First(&Company, id).Error; err != nil {
		return nil, err
	}

	return &Company, nil
}


func GetEmployeesByCompanyID(id uint) ([]models.Employee, error) {
	list := make([]models.Employee, 0)

	if err := database.GormDB.Where("company_id = ?", id).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
