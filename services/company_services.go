package services

import (
	database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models"
	companyModels "github.com/nikola43/ecoapigorm/models/company"
	_ "github.com/nikola43/ecoapigorm/models/employee"
	_ "github.com/nikola43/ecoapigorm/utils"
)

func CreateCompany(employeeID uint, createEmployeeRequest *companyModels.CreateCompanyRequest) (*companyModels.CreateCompanyResponse, error) {
	company := models.Company{
		Name:       createEmployeeRequest.Name,
		//EmployeeID: employeeID,
	}

	if err := database.GormDB.Create(&company).Error; err != nil {
		return nil, err
	}

	createCompanyResponse := &companyModels.CreateCompanyResponse{
		ID:         company.ID,
		Name:       company.Name,
		//EmployeeID: company.EmployeeID,
		CreatedAt:  company.CreatedAt.String(),
	}

	database.GormDB.Model(&models.Employee{}).Where("id = ? ", employeeID).
		Update("company_id", createCompanyResponse.ID)

	database.GormDB.Model(&models.Employee{}).Where("id = ? ", employeeID).
		Update("is_first_login", false)



	return createCompanyResponse, nil
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

func GetClinicsByCompanyID(companyId uint) ([]models.Clinic, error) {
	clinics := make([]models.Clinic, 0)

	 err := database.GormDB.
		Where("company_id = ?", companyId).
		Preload("Clients").
		Preload("Employees").
		Find(&clinics).Error

	if	err != nil {
		return nil, err
	}

	return clinics, nil
}

func GetCompaniesByEmployeeID(employeeID uint) ([]models.Company, error) {
	var list []models.Company

	result := database.GormDB.Where("employee_id = ?", employeeID).
		Preload("Clinics").
		Preload("Employees").
		Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}

	return list, nil
}
