package rpc

import (
	"github.com/vmkteam/zenrpc/v2"

	"context"
)

//zenrpc:params=queryParams{}
func (rpc *Service) GetAllNews(ctx context.Context, params queryParams) ([]News, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	list, err := rpc.m.GetNewsByFilters(r.Context(), params.NewFilter())
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	if len(list) == 0 {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, "no list found")
	}

	return NewNewsList(list), nil
}

func (rpc *Service) GetNewsById(ctx context.Context, id int) (*News, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	news, err := rpc.m.GetNewsById(r.Context(), id)
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	return NewNews(news), nil
}

//zenrpc:params=queryParams{}
func (rpc *Service) GetNewsCount(ctx context.Context, params queryParams) (int, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	count, err := rpc.m.GetNewsCount(r.Context(), params.NewFilter())
	if err != nil {
		return 0, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	return count, nil
}

//zenrpc:params=queryParams{}
func (rpc *Service) GetAllShortNews(ctx context.Context, params queryParams) ([]NewsSummary, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	list, err := rpc.m.GetNewsByFilters(r.Context(), params.NewFilter())
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	if len(list) == 0 {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, "no news found")
	}

	return NewNewsSummaries(list), nil
}

func (rpc *Service) GetAllCategories(ctx context.Context) ([]Category, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	categories, err := rpc.m.GetAllCategory(r.Context())
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	return NewCategories(categories), nil
}

func (rpc *Service) GetAllTags(ctx context.Context) ([]Tag, *zenrpc.Error) {
	r, _ := zenrpc.RequestFromContext(ctx)

	tags, err := rpc.m.GetAllTag(r.Context())
	if err != nil {
		return nil, zenrpc.NewStringError(zenrpc.InternalError, err.Error())
	}

	return NewTags(tags), nil
}
