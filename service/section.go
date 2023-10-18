package service

import (
	"github.com/TeamTells/portal_documents_service/api"
	"github.com/TeamTells/portal_documents_service/storage"
)

type SectionService struct {
	store *storage.Storage
}

func NewSectionService(store *storage.Storage) SectionService {
	return SectionService{store: store}
}

func (s *SectionService) GetSections() []*api.Sections {
	items, err := s.store.GetSections()
	if err != nil {
		return []*api.Sections{}
	}
	return items
}
