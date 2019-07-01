package go_try

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	withoutError := func() (v int, e error) {
		return 1, nil
	}
	withError := func() (v int, e error) {
		return 1, errors.New("test")
	}

	t.Run("without error", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			assert.Fail(t, "test failed")
		}()
		res := Try(withoutError())
		assert.Equal(t, 1, res)
	})
	t.Run("with error", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			if TryError == nil {
				assert.Fail(t, "test failed")
			}
			assert.EqualError(t, TryError, "test")
		}()
		_ = Try(withError())
		assert.Fail(t, "test failed")
	})
}
