package dto

import (
	"github.com/google/uuid"
)

type RatingDto struct {
	Score      int       `json:"score" validate:"required"`
	Difficulty int       `json:"difficulty" validate:"required"`
	ProblemID  uuid.UUID `json:"problem_id" validate:"required"`
	UserID     uuid.UUID `json:"user_id" validate:"required"`
}
