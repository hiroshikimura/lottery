package helper_test

import (
	. "github.com/r7kamura/gospel"
	"lottery/helper"
	"regexp"
	"testing"
)

func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

// go test with Gospel
// https://github.com/r7kamura/gospel
func TestHelper2(t *testing.T) {
	Describe(t, "test Helper2()", func() {
		uuid := helper.UUIDGen()
		Context("Context description 1", func() {
			Before(func() {
				// Called before each examples in this Context.
			})

			After(func() {
				// Called after each examples in this Context.
			})
			It("It description 1", func() {
				Expect(uuid).To(NotEqual, nil)
				Expect(check_regexp(`[A-Za-z0-9]{8}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{12}`, uuid.String())).To(Equal, true)
			})
			It("It description 2", func() {
				Expect(uuid).To(NotEqual, nil)
				Expect(check_regexp(`[A-Za-z0-9]{8}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{12}`, uuid.String())).To(Equal, true)
			})
		})
		Context("Context description 2", func() {
			Before(func() {
				// Called before each examples in this Context.
			})

			After(func() {
				// Called after each examples in this Context.
			})
			It("It description 1", func() {
				Expect(uuid).To(NotEqual, nil)
				Expect(check_regexp(`[A-Za-z0-9]{8}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{12}`, uuid.String())).To(Equal, true)
			})
			It("It description 2", func() {
				Expect(uuid).To(NotEqual, nil)
				Expect(check_regexp(`[A-Za-z0-9]{8}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{12}`, uuid.String())).To(Equal, true)
			})
		})
	})
}
