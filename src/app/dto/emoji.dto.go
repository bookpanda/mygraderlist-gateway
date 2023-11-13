package dto

import (
	"github.com/google/uuid"
)

type EmojiDto struct {
	Emoji     string    `json:"emoji" validate:"required"`
	ProblemID uuid.UUID `json:"problem_id" validate:"required"`
	UserID    uuid.UUID `json:"user_id" validate:"required"`
}
