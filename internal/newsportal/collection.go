package newsportal

//go:generate colgen -imports=News-portal/internal/db
//colgen:News,Category,Tag
//colgen:News:UniqueTagIDs,MapP(db)
//colgen:Tag:MapP(db)
//colgen:Category:MapP(db)

// MapP converts slice of type T to slice of type M with given converter with pointers.
func MapP[T, M any](a []T, f func(*T) *M) []M {
	n := make([]M, len(a))
	for i := range a {
		n[i] = *f(&a[i])
	}
	return n
}

// Map converts slice of type T to slice of type M with given converter.
func Map[T, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i := range a {
		n[i] = f(a[i])
	}
	return n
}

func (nl NewsList) SetTags(tags Tags) {
	tagIndex := tags.Index()
	for i, v := range nl {
		for _, tag := range v.TagIDs {
			if t, ok := tagIndex[tag]; ok {
				nl[i].Tags = append(nl[i].Tags, t)
			}
		}
	}
}
