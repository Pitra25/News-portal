package newsportal

type Tags []Tag

func (t Tags) Index() map[int]Tag {
	tagsIndex := make(map[int]Tag, len(t))
	for _, v := range t {
		tagsIndex[v.ID] = v
	}

	return tagsIndex
}

type NewsList []News

func (nl NewsList) UniqueTagIDs() []int {
	var tagIDs []int

	tagIds := make(map[int]struct{})
	for _, v := range nl {
		for _, tag := range v.TagIDs {
			if _, ok := tagIds[tag]; !ok {
				tagIDs = append(tagIDs, tag)
				tagIds[tag] = struct{}{}
			}
		}
	}

	return tagIDs
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
