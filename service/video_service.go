package service

import "github.com/khilmi-aminudin/videos-app/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	Videos []entity.Video
}

func NewVideoService() VideoService {
	return &videoService{}
}

// FindAll implements VideoService
func (s *videoService) FindAll() []entity.Video {
	return s.Videos
}

// Save implements VideoService
func (s *videoService) Save(video entity.Video) entity.Video {
	s.Videos = append(s.Videos, video)
	return video
}
