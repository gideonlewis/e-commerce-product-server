package datatypes

import "strings"

type SortType string

const (
	SortTypeASC  SortType = "ASC"
	SortTypeDESC SortType = "DESC"
)

func (s *SortType) IsValid() bool {
	if s != nil {
		switch SortType(strings.ToUpper(string(*s))) {
		case SortTypeASC, SortTypeDESC:
			return true
		default:
			return false
		}
	}

	return false
}

func (s *SortType) IsNull() bool {
	return *s == ""
}

func (s *SortType) String() string {
	return string(*s)
}
