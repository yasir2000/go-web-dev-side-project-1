package cms

import (
	"reflect"
	"strconv"
	"testing"
)

var p *Page

func Test_CreatePage(t *testing.T) {
	p = &Page{
		Title:   "test",
		Content: "test",
	}
	id, err := CreatePage(p)
	if err != nil {
		t.Errorf("Failed to create page: %s\n", err.Error())
	}
	p.ID = id
}

func Test_GetPage(t *testing.T) {
	page, err := GetPage(strconv.Itoa(p.ID))
	if err != nil {
		t.Errorf("Failed to get page: %s\n", err.Error())
	}

	if page.ID != p.ID {
		t.Errorf("Page IDs do not match: %d\n vs %d\n", page.ID, p.ID)
	}

	if reflect.DeepEqual(page, p) != true {
		t.Errorf("Pages do not match: %+v\n vs %+v\n", page, p)
	}
}
