package graph

import (
	"time"

	"gorm.io/gorm"
)

var layout = "2006-01-02 15:04:05"

func tToS(t time.Time) *string {
	var s string

	if t.IsZero() {
		s = ""
	} else {
		s = t.Format(layout)
	}

	return &s
}

func dtToS(t gorm.DeletedAt) *string {
	var s string

	if t.Valid {
		s = t.Time.Format(layout)
	} else {
		s = ""
	}

	return &s
}

func getString(p *string, alter string) string {
	if p == nil {
		return alter
	}
	return *p
}

func getInt(p *int, alter int) int {
	if p == nil {
		return alter
	}
	return *p
}
