package helper_test

import (
	. "github.com/pranavraja/zen"
	"lottery/helper"
	"regexp"
	"testing"
)

// go test with zen
// https://github.com/pranavraja/zen
func TestUUIDGen(t *testing.T) {
	Desc(t, "test descriptions...", func(it It) {
		check_regexp := func(reg, str string) bool {
			return regexp.MustCompile(reg).Match([]byte(str))
		}
		uuid := helper.UUIDGen()
		before_func := func() {
			// before
		}
		after_func := func() {
			// after
		}
		setup := Setup(before_func, after_func)
		it("test description 1", setup(func(expect Expect) {
			expect(check_regexp(`[A-Za-z0-9]{8}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{4}-[A-Za-z0-9]{12}`, uuid.String())).ToEqual(true)
		}))
	})
}
