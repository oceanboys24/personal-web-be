package services

import (
	"personal-web-be/config"
	model "personal-web-be/models"

	"github.com/google/uuid"
)




func GetWorkService() ([]model.WorkModel, error) {
	var result []model.WorkModel

	err := config.SupaClient.DB.From("work_experience").Select("*").OrderBy("created_at", "asc").Execute(&result)
	if err != nil {
		return nil ,err
	}

	return result , nil
	
}
func GetWorkServiceById(id string) ([]model.WorkModel, error) {
	var result []model.WorkModel

	err := config.SupaClient.DB.From("work_experience").Select("*").Eq("id", id).Execute(&result)
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

func UpdateWorkService(id uuid.UUID,work map[string]interface{}) error {
	finalId := id.String()
	var result map[string]interface{}
	err := config.SupaClient.DB.From("work_experience").Update(work).Eq("id", finalId).Execute(&result) 
	if err != nil {
		return nil
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