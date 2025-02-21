package todoist

import (
	"regexp"
	"strings"
)

type Section struct {
	HaveID
	HaveProjectID
	Collapsed    bool   `json:"collapsed"`
	Name         string `json:"name"`
	IsArchived   bool   `json:"is_archived"`
	IsDeleted    bool   `json:"is_deleted"`
	SectionOrder int    `json:"section_order"`
}

type Sections []Section

func (a Sections) GetIDsByName(name string) []string {
	ids := []string{}
	name = strings.ToLower(name)
	for _, section := range a {
		match, _ := regexp.MatchString(name, strings.ToLower(section.Name))
		if match {
			ids = append(ids, section.ID)
		}
	}
	return ids
}

