package service

import (
	"context"
	"fmt"
	"github.com/xoesae/cid-api/internal/domain/entity"
	"github.com/xoesae/cid-api/internal/domain/repository"
	"github.com/xoesae/cid-api/internal/importer"
	"io"
	"os"
	"strings"
)

type ImportService interface {
	RunImport(xmlPath string) error
}

type importService struct {
	chapterRepository     repository.ChapterRepository
	groupRepository       repository.GroupRepository
	categoryRepository    repository.CategoryRepository
	subcategoryRepository repository.SubcategoryRepository
}

func NewImportService(chRepo repository.ChapterRepository, gpRepo repository.GroupRepository, catRepo repository.CategoryRepository, subcatRepo repository.SubcategoryRepository) ImportService {
	return &importService{
		chapterRepository:     chRepo,
		groupRepository:       gpRepo,
		categoryRepository:    catRepo,
		subcategoryRepository: subcatRepo,
	}
}

func (s importService) RunImport(xmlPath string) error {
	ctx := context.Background()

	file, err := os.Open(xmlPath)

	xmlData, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo XML: %w", err)
	}

	tempFileName := s.saveTmpFile(string(xmlData))
	defer os.Remove(tempFileName)

	return importer.StreamChapters(tempFileName, func(ch importer.Chapter) error {

		chapter := entity.Chapter{
			ID:        0,
			CodeStart: ch.Initial,
			CodeEnd:   ch.Final,
			Roman:     strings.Trim(ch.Roman, " "),
			Name:      ch.Name,
		}

		chapterID, err := s.chapterRepository.Insert(ctx, chapter)
		if err != nil {
			return err
		}

		for _, gp := range ch.Groups {

			group := entity.Group{
				ID:        0,
				ChapterID: chapterID,
				CodeStart: gp.Initial,
				CodeEnd:   gp.Final,
				Name:      gp.Name,
			}

			groupID, err := s.groupRepository.Insert(ctx, group)
			if err != nil {
				return err
			}

			for _, cat := range gp.Categories {

				category := entity.Category{
					ID:      0,
					GroupID: groupID,
					Code:    cat.Code,
					Name:    cat.Name,
				}

				categoryID, err := s.categoryRepository.Insert(ctx, category)
				if err != nil {
					return err
				}

				for _, sub := range cat.Subcategories {

					subcategory := entity.Subcategory{
						ID:         0,
						CategoryID: categoryID,
						Code:       sub.Code,
						Name:       sub.Name,
					}

					subcatID, err := s.subcategoryRepository.Insert(ctx, subcategory)
					if err != nil {
						return err
					}

					fmt.Printf("Imported subcategory: [%d] %s\n", subcatID, sub.Name)
				}

				fmt.Printf("Imported category: [%d] %s\n", categoryID, cat.Name)
			}

			fmt.Printf("Imported group: [%d] %s\n", groupID, gp.Name)
		}

		fmt.Printf("Imported chapter: [%d] %s\n", chapterID, ch.Name)

		return nil
	})
}

func (s importService) saveTmpFile(xmlData string) string {
	// replace unknown chars
	xmlString := strings.ReplaceAll(xmlData, "&cruz;", "")

	tempFile, err := os.CreateTemp("", "parsed_xml_*.xml")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	_, err = tempFile.Write([]byte(xmlString))
	if err != nil {
		panic(err)
	}

	return tempFile.Name()
}
