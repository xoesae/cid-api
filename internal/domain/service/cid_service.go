package service

import (
	"context"
	"github.com/xoesae/cid-api/internal/domain/entity"
	"github.com/xoesae/cid-api/internal/domain/repository"
	"github.com/xoesae/cid-api/pkg/fault"
	"github.com/xoesae/cid-api/pkg/logger"
	"regexp"
	"strings"
)

type CidService interface {
	GetAllPaginated(ctx context.Context, pageSize int, page int, search string) ([]entity.Subcategory, error)
	GetByCode(ctx context.Context, code string) (entity.Cid, error)
}

type cidService struct {
	categoryRepository    repository.CategoryRepository
	subcategoryRepository repository.SubcategoryRepository
}

func NewCidService(catRepository repository.CategoryRepository, subcatRepo repository.SubcategoryRepository) CidService {
	return &cidService{
		categoryRepository:    catRepository,
		subcategoryRepository: subcatRepo,
	}
}

func (s *cidService) GetAllPaginated(ctx context.Context, pageSize int, page int, search string) ([]entity.Subcategory, error) {
	offset := (page - 1) * pageSize

	reg := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	search = reg.ReplaceAllString(search, "")
	search = strings.TrimSpace(search)

	result, err := s.subcategoryRepository.GetPaginated(pageSize, offset, search)
	if err != nil {
		logger.Get().Error(err.Error())
		return nil, err
	}

	return result, nil
}

func (s *cidService) GetByCode(ctx context.Context, code string) (entity.Cid, error) {
	codeLen := len(code)

	if codeLen < 3 || codeLen > 4 {
		return entity.Cid{}, fault.NewNotFoundFault("could not find cid")
	}

	code = strings.ToUpper(code)

	if codeLen > 3 {
		result, err := s.subcategoryRepository.GetByCode(code)
		if err != nil {
			logger.Get().Debug(err.Error())

			return entity.Cid{}, err
		}

		return entity.Cid{
			Code: result.Code,
			Name: result.Name,
		}, nil
	}

	result, err := s.categoryRepository.GetByCode(code)
	if err != nil {
		logger.Get().Debug(err.Error())

		return entity.Cid{}, err
	}

	return entity.Cid{
		Code: result.Code,
		Name: result.Name,
	}, nil
}
