package storage

import "github.com/TeamTells/portal_documents_service/api"

var GetAllSectionsSQL = "select * from sections"

func (s *Storage) GetSections() ([]*api.Sections, error) {

	var item []*api.Sections

	stmt, err := s.db.Prepare(GetAllSectionsSQL)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow().Scan(&item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
