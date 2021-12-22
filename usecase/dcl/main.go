package dcl

import (
	"context"

	"github.com/sysdevguru/fiskil/models"
)

type TaskRepo interface {
	UpdateLogs(ctx context.Context, msgs *[]models.Log, severities *[]models.Severity) error
}

type UseCase struct {
	taskName string
	repo     TaskRepo
}

func New(taskName string, repo TaskRepo) *UseCase {
	return &UseCase{
		taskName,
		repo,
	}
}
