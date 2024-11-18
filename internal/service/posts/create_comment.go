package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/NXRts/fsatcampus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: request.CommentContent,
		CreateAt:       now,
		UpdateAt:       now,
		CreateBy:       strconv.FormatInt(userID, 10),
		UpdateBy:       strconv.FormatInt(userID, 10),
	}
	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create comment to repository")
		return err
	}
	return nil
}
