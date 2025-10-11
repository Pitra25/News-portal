package rpc

import (
	"github.com/vmkteam/zenrpc/v2"

	"context"
)

func (ns *NewsService) GetAllNews(ctx context.Context, params queryParams) ([]News, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	list, err := ns.m.GetNewsByFilters(r.Context(), params.NewFilter())
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	if len(list) == 0 {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, "no list found")
	}

	return NewNewsList(list), nil
	//return nil, nil
}

func (ns *NewsService) GetNewsById(ctx context.Context, id int) (*News, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	news, err := ns.m.GetNewsById(r.Context(), id)
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	return NewNews(news), nil
	//return nil, nil
}

func (ns *NewsService) GetNewsCount(ctx context.Context, params queryParams) (int, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	count, err := ns.m.GetNewsCount(r.Context(), params.NewFilter())
	if err != nil {
		return 0, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	return count, nil
	//return 0, nil
}

func (s *ShortNewsService) GetAllShortNews(ctx context.Context, params queryParams) ([]NewsSummary, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	list, err := s.m.GetNewsByFilters(r.Context(), params.NewFilter())
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	if len(list) == 0 {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, "no news found")
	}

	return NewNewsSummaries(list), nil
	//return nil, nil
}

func (s *CategoriesService) GetAllCategories(ctx context.Context) ([]Category, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	categories, err := s.m.GetAllCategory(r.Context())
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	return NewCategories(categories), nil
	//return nil, nil
}

func (s *TagsService) GetAllTags(ctx context.Context) ([]Tag, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	tags, err := s.m.GetAllTag(r.Context())
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	return NewTags(tags), nil
	//return nil, nil
}
