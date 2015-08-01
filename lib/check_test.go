package lib

import (
	"testing"
)

//
func TestIsEmail(t *testing.T) {
	rigit := []string{"me@ckeyer.com", "ckeyer.com@gmail.com"}
	wrong := []string{"asdf.com", "ckeyer@com"}
	for _, v := range rigit {
		if IsEmail(v) != true {
			t.Errorf("%s is a right email address", v)
		}
	}
	for _, v := range wrong {
		if IsEmail(v) != false {
			t.Errorf("%s is a wrong email address", v)
		}
	}
}

func TestIsMobilePhone(t *testing.T) {
	rigit := []string{"18001231233", "15312345678"}
	wrong := []string{"1331234567", "00012345678"}
	for _, v := range rigit {
		if IsMobildPhone(v) != true {
			t.Errorf("%s is a right phone number", v)
		}
	}
	for _, v := range wrong {
		if IsMobildPhone(v) != false {
			t.Errorf("%s is a wrong phone number", v)
		}
	}
}

func TestIsIDcard(t *testing.T) {
	rigit := []string{"123456123412121234", "12345612341212123x"}
	wrong := []string{"1234561234121212xx", "123213213123123"}
	for _, v := range rigit {
		if IsIDcard(v) != true {
			t.Errorf("%s is a right ID number", v)
		}
	}
	for _, v := range wrong {
		if IsIDcard(v) != false {
			t.Errorf("%s is a wrong ID number", v)
		}
	}
}
