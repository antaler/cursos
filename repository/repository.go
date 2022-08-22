package repository

import (
	"context"

	"gitlab.com/antaler/cursos/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo

}

func SetStudent(ctx context.Context, student *models.Student) error {

	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}
