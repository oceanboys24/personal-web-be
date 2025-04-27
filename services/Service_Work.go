package services

import (
	"personal-web-be/config"
	model "personal-web-be/models"

	"github.com/google/uuid"
)




func GetWorkService() ([]model.WorkModel, error) {
	var result []model.WorkModel

	err := config.SupaClient.DB.From("work_experience").Select("*").Execute(&result)
	if err != nil {
		return nil ,err
	}

	return result , nil
	
}
func CreateWorkService(work model.WorkModel)  error {
	var result []model.WorkModel

	err := config.SupaClient.DB.From("work_experience").Insert(work).Execute(&result)
	if err != nil {
		return err
	}

	return nil
}



func DeleteWorkService(id uuid.UUID) error {
	finalId := id.String()
	err := config.SupaClient.DB.From("work_experience").Delete().Eq("id",finalId).Execute("")
	if err != nil {
		return err
	}

	return nil
}