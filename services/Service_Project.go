package services

import (
	"personal-web-be/config"
	model "personal-web-be/models"

	"github.com/google/uuid"
)


func GetProjectService() ([]model.ProjectModel, error) {
	var result []model.ProjectModel

	err := config.SupaClient.DB.From("project").Select("*").OrderBy("created_at", "asc").Execute(&result)
	if err != nil {
		return nil ,err
	}

	return result , nil
	
}

func GetProjectServiceById(id string) (*model.ProjectModel, error) {
	var result model.ProjectModel

	err := config.SupaClient.DB.From("project").Select("*").Execute(&result)
	if err != nil {
		return nil ,err
	}

	return &result , nil
	
}


func CreateProjectService(project model.ProjectModel)  error {
	var result []model.ProjectModel

	err := config.SupaClient.DB.From("project").Insert(project).Execute(&result)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProjectService(id uuid.UUID,project map[string]interface{}) error {
	finalId := id.String()
	var result map[string]interface{}
	err := config.SupaClient.DB.From("project").Update(project).Eq("id", finalId).Execute(&result) 
	if err != nil {
		return nil
	}

	return nil
}




func DeleteProjectService(id uuid.UUID) error {
	finalId := id.String()
	err := config.SupaClient.DB.From("project").Delete().Eq("id",finalId).Execute("")
	if err != nil {
		return err
	}

	return nil
}