package dto

import (
	"github.com/google/uuid"
)

type LikeDto struct {
	ProblemID uuid.UUID `json:"problem_id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
}
