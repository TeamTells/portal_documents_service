package api

import (
	"encoding/json"
	"github.com/TeamTells/portal_documents_service/service"
	"net/http"
)

type Implementation struct {
	section *service.SectionService
}

func NewImplementation(section *service.SectionService) Implementation {
	return Implementation{section: section}
}

func (i *Implementation) returnData(w http.ResponseWriter, r any) {
	bytes, err := json.Marshal(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(bytes) // NOTE that we should handle the error here
	w.WriteHeader(http.StatusOK)
}

func (i *Implementation) GetSections(w http.ResponseWriter, r *http.Request) {
	var response = i.section.GetSections()
	i.returnData(w, response)
}
