package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("nil request - error", func(t *testing.T) {
		t.Parallel()

		rsl := base.Validate(nil)

		assert.Equal(t, []string{"request is nil"}, rsl)
	})

	t.Run("all fields filled - no errors", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "u1",
			Title:       "title",
			Description: "desc",
		}
		rsl := base.Validate(in)

		assert.Empty(t, rsl)
	})

	t.Run("empty user id - error", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "",
			Title:       "title",
			Description: "desc",
		}
		rsl := base.Validate(in)

		assert.Equal(t, []string{"user id is empty"}, rsl)
	})

	t.Run("empty title - error", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "u1",
			Title:       "",
			Description: "desc",
		}
		rsl := base.Validate(in)

		assert.Equal(t, []string{"title is empty"}, rsl)
	})

	t.Run("empty description - error", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "u1",
			Title:       "title",
			Description: "",
		}
		rsl := base.Validate(in)

		assert.Equal(t, []string{"description is empty"}, rsl)
	})

	t.Run("all fields empty - 3 errors", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "",
			Title:       "",
			Description: "",
		}
		rsl := base.Validate(in)

		assert.Equal(t, []string{
			"user id is empty",
			"title is empty",
			"description is empty",
		}, rsl)
	})

	t.Run("two fields empty - 2 errors (user id + title)", func(t *testing.T) {
		t.Parallel()

		in := &base.ValidateRequest{
			UserID:      "",
			Title:       "",
			Description: "desc",
		}
		rsl := base.Validate(in)

		assert.Equal(t, []string{
			"user id is empty",
			"title is empty",
		}, rsl)
	})
}