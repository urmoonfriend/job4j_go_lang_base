package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_LruCache(t *testing.T) {
	t.Parallel()

	t.Run("put new item", func(t *testing.T) {
		t.Parallel()
		lru := base.NewLruCache(2)
		lru.Put("key1", "value1")
		lru.Put("key2", "value2")

		
		v1 := lru.Get("key1")
		v2 := lru.Get("key2")

		if assert.NotNil(t, v1) {
			assert.Equal(t, "value1", *v1)
		}
		if assert.NotNil(t, v2) {
			assert.Equal(t, "value2", *v2)
		}
	})


	t.Run("put new item, but cache is full", func(t *testing.T) {
		t.Parallel()
		lru := base.NewLruCache(2)
		lru.Put("key1", "value1")
		lru.Put("key2", "value2")
		lru.Put("key3", "value3")

		
		v1 := lru.Get("key1")
		v2 := lru.Get("key2")
		v3 := lru.Get("key3")

		assert.Nil(t, v1)
			
		if assert.NotNil(t, v2) {
			assert.Equal(t, "value2", *v2)
		}
		if assert.NotNil(t, v3) {
			assert.Equal(t, "value3", *v3)
		}
	})

	t.Run("put existing item", func(t *testing.T) {
		t.Parallel()
		lru := base.NewLruCache(2)
		lru.Put("key1", "value1")
		lru.Put("key2", "value2")
		lru.Put("key1", "value1_updated")

		
		v1 := lru.Get("key1")
		v2 := lru.Get("key2")

		if assert.NotNil(t, v1) {
			assert.Equal(t, "value1_updated", *v1)
		}
		if assert.NotNil(t, v2) {
			assert.Equal(t, "value2", *v2)
		}
	})

	t.Run("get non-existing item", func(t *testing.T) {
		t.Parallel()
		lru := base.NewLruCache(2)
		lru.Put("key1", "value1")
		lru.Put("key2", "value2")

		
		v3 := lru.Get("key3")

		assert.Nil(t, v3)
	})

	t.Run("get item updates its usage", func(t *testing.T) {
		t.Parallel()
		lru := base.NewLruCache(2)
		lru.Put("key1", "value1")
		lru.Put("key2", "value2")

		
		v1 := lru.Get("key1")
		assert.NotNil(t, v1)

		lru.Put("key3", "value3")

		v1 = lru.Get("key1")
		v2 := lru.Get("key2")
		v3 := lru.Get("key3")

		assert.Nil(t, v2)
		
		if assert.NotNil(t, v1) {
			assert.Equal(t, "value1", *v1)
		}
		if assert.NotNil(t, v3) {
			assert.Equal(t, "value3", *v3)
		}
	})

	t.Run("get item that was updated", func(t *testing.T) {
		t.Parallel()
		lru := base.NewLruCache(2)
		lru.Put("key1", "value1")
		lru.Put("key2", "value2")
		lru.Put("key1", "value1_updated")

		
		v1 := lru.Get("key1")
		v2 := lru.Get("key2")

		if assert.NotNil(t, v1) {
			assert.Equal(t, "value1_updated", *v1)
		}
		if assert.NotNil(t, v2) {
			assert.Equal(t, "value2", *v2)
		}
	})

	t.Run("get item that was evicted", func(t *testing.T) {
		t.Parallel()
		lru := base.NewLruCache(2)
		lru.Put("key1", "value1")
		lru.Put("key2", "value2")
		lru.Put("key3", "value3")

		
		v1 := lru.Get("key1")
		v2 := lru.Get("key2")
		v3 := lru.Get("key3")

		assert.Nil(t, v1)
		
		if assert.NotNil(t, v2) {
			assert.Equal(t, "value2", *v2)
		}
		if assert.NotNil(t, v3) {
			assert.Equal(t, "value3", *v3)
		}
	})	

}