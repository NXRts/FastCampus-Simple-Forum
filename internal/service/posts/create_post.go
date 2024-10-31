// package posts

// import (
// 	"context"
// 	"log"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/NXRts/fsatcampus/internal/model/posts"
// )

// func (s *service) CreatePost(ctx context.Context, UserId int64, req posts.CreatePostRequest) error {
// 	postHashtags := strings.Join(req.PostHashtags, ",")

// 	now := time.Now()

//		model := posts.PostModel{
//			UserId:       UserId,
//			PostTitle:    req.PostTitle,
//			PostContent:  req.PostContent,
//			PostHashtags: postHashtags,
//			CreateAt:     now,
//			UpdateAt:     now,
//			CreateBy:     strconv.FormatInt(UserId, 10),
//			UpdateBy:     strconv.FormatInt(UserId, 10),
//		}
//		err := s.postRepo.CreatePost(ctx, model)
//		if err != nil {
//			log.Error().Err(err).Msg("error create post to repository")
//			return err
//		}
//		return nil
//	}
package posts

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/NXRts/fsatcampus/internal/model/posts"
)

func (s *service) CreatePost(ctx context.Context, UserId int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHashtags, ",")
	now := time.Now()

	model := posts.PostModel{
		UserId:       UserId,
		PostTitle:    req.PostTitle,
		PostContent:  req.PostContent,
		PostHashtags: postHashtags,
		CreateAt:     now,
		UpdateAt:     now,
		CreateBy:     strconv.FormatInt(UserId, 10),
		UpdateBy:     strconv.FormatInt(UserId, 10),
	}
	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Println("error create post to repository:", err)
		return err
	}
	return nil
}
