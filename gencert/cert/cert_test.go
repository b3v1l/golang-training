package cert

import (
	"fmt"
	"testing"
)

func TestValidCertdata(t *testing.T) {
	c, err := New("", "polo", "2019-02-02")
	if err != nil {
		t.Errorf("error %v - Cert data must be valid !\n", err)
	}
	if c == nil {
		t.Errorf("Cert must non empty. Got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course must be GOLANG. Got=%v\n", c.Course)
	}
}

func TestEmptyStringCourse(t *testing.T) {

	_, err := New("", "polo", "2019-02-02")
	if err != nil {
		t.Error("Error on empty course")
	}
}

func TestCoourseTooLong(t *testing.T) {

	course := "dfdskjfdskfsk"
	_, err := New(course, "polopo", "2019-02-02")
	if err != nil {
		t.Errorf("Course name too long ! len(%s)", course)
	}

}

func TestNameLen(t *testing.T) {
	name := "s"
	fmt.Printf("len(%d", len(name))
	_, err := New("test", name, "2019-02-02")
	if err != nil {
		t.Errorf("Name too long: len(%v)", len(name))
	}

}
