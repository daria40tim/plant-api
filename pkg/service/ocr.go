package service

import (
	plantapi "github.com/daria40tim/plant-api"
	"github.com/daria40tim/plant-api/pkg/repository"
)

type OCRService struct {
	repo repository.OCR
}

func NewOCRService(repo repository.OCR) *OCRService {
	return &OCRService{repo: repo}
}

func (s *OCRService) GetDirs() ([]plantapi.Dir, error) {
	return s.repo.GetDirs()
}

func (s *OCRService) GetImg(index int, dir string) (plantapi.CardOCR, error) {
	return s.repo.GetImg(index, dir)
}
