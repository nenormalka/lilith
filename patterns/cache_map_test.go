package patterns

import (
	"testing"
	"time"
)

func TestCacheMap_FullReload(t *testing.T) {
	enableReload := false
	cm := NewCacheMap(func(oldCache map[string]string) map[string]string {
		if !enableReload {
			return oldCache
		}

		return map[string]string{
			"key1": "value3",
			"key2": "value2",
		}
	}, time.Millisecond*100)

	cm.Set("key1", "value1")

	value, ok := cm.Get("key1")
	if !ok {
		t.Errorf("expected key1 to be in cache")
	}

	if value != "value1" {
		t.Errorf("expected key1 to be value1")
	}

	enableReload = true

	time.Sleep(time.Millisecond * 200)

	value, ok = cm.Get("key1")
	if !ok {
		t.Errorf("expected key1 to be in cache")
	}

	if value != "value3" {
		t.Errorf("expected key1 to be value3")
	}

	value, ok = cm.Get("key2")
	if !ok {
		t.Errorf("expected key2 to be in cache")
	}

	if value != "value2" {
		t.Errorf("expected key2 to be value2")
	}
}

func TestCacheMap_ReloadWithExpiredAt(t *testing.T) {
	type test struct {
		value     string
		expiredAt time.Time
	}

	cm := NewCacheMap(func(oldCache map[string]test) map[string]test {
		now := time.Now()

		for key := range oldCache {
			if now.After(oldCache[key].expiredAt) {
				delete(oldCache, key)
			}
		}

		return oldCache
	}, time.Millisecond*100)

	cm.Set("key1", test{value: "value1", expiredAt: time.Now().Add(time.Millisecond * 200)})
	cm.Set("key2", test{value: "value2", expiredAt: time.Now().Add(time.Second * 2)})

	value, ok := cm.Get("key1")
	if !ok {
		t.Errorf("expected key1 to be in cache")
	}

	if value.value != "value1" {
		t.Errorf("expected key1 to be value1")
	}

	value, ok = cm.Get("key2")
	if !ok {
		t.Errorf("expected key2 to be in cache")
	}

	if value.value != "value2" {
		t.Errorf("expected key2 to be value2")
	}

	time.Sleep(time.Millisecond * 200)

	value, ok = cm.Get("key1")
	if ok {
		t.Errorf("expected key1 to be expired")
	}

	value, ok = cm.Get("key2")
	if !ok {
		t.Errorf("expected key2 to be in cache")
	}
}
