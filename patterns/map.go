package patterns

type (
	SafeMap[T comparable, M any] interface {
		Load(key T) (value M, ok bool)
		Store(key T, value M)
		LoadOrStore(key T, value M) (actual M, loaded bool)
		LoadAndDelete(key T) (value M, loaded bool)
		Delete(key T)
		Swap(key T, value M) (previous M, loaded bool)
		CompareAndSwap(key T, old, new M) (swapped bool)
		CompareAndDelete(key T, old M) (deleted bool)
		Range(func(key T, value M) (shouldContinue bool))
		Clear()
		Len() int
		Map() map[T]M
	}
)
