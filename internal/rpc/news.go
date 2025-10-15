package rpc

import (
	"context"
)

/*** News ***/

// News return list news.
//
//zenrpc:404 not found
func (ns *NewsService) News(ctx context.Context, params *queryParams) ([]News, error) {
	list, err := ns.m.GetNewsByFilters(ctx, params.NewFilter())
	if err != nil {
		return nil, newInternalError(err)
	} else if list == nil {
		return nil, noContentError
	}

	return NewNewsList(list), nil
}

// GetById returns news by ID.
//
//zenrpc:id news id
//zenrpc:404 not found
func (ns *NewsService) GetById(ctx context.Context, id int) (*News, error) {
	news, err := ns.m.GetNewsById(ctx, id)
	if err != nil {
		return nil, newInternalError(err)
	} else if news == nil {
		return nil, noContentError
	}

	return NewNews(news), nil
}

// Count returns count news by filters.
//
//zenrpc:404 not found
func (ns *NewsService) Count(ctx context.Context, params *queryParams) (int, error) {
	count, err := ns.m.GetNewsCount(ctx, params.NewFilter())
	if err != nil {
		return 0, newInternalError(err)
	}

	return count, nil
}

// NewsSummaries returns news summary by filters.
//
//zenrpc:404 not found
func (ns *NewsService) NewsSummaries(ctx context.Context, params *queryParams) ([]NewsSummary, error) {
	list, err := ns.m.GetNewsByFilters(ctx, params.NewFilter())
	if err != nil {
		return nil, newInternalError(err)
	} else if list == nil {
		return nil, noContentError
	}

	return NewNewsSummaries(list), nil
}

// AddNews return new news.
//
//zenrpc:404 not found
func (ns *NewsService) AddNews(ctx context.Context, in NewsInput) (*News, error) {
	res, err := ns.m.AddNews(ctx, newsToManager(&in))
	if err != nil {
		return nil, newInternalError(err)
	}

	return NewNews(res), nil
}

/*** Category ***/

// Categories return list category
//
//zenrpc:404 not found
func (ns *NewsService) Categories(ctx context.Context) ([]Category, error) {
	list, err := ns.m.GetAllCategory(ctx)
	if err != nil {
		return nil, newInternalError(err)
	} else if list == nil {
		return nil, noContentError
	}

	return NewCategories(list), nil
}

// AddCategory return new category.
//
//zenrpc:404 not found
func (ns *NewsService) AddCategory(ctx context.Context, in CategoryInput) (*Category, error) {
	res, err := ns.m.AddCategory(ctx, categoryToManager(&in))
	if err != nil {
		return nil, newInternalError(err)
	}

	return NewCategory(res), nil
}

/*** Tag ***/

// Tags return list tag.
//
//zenrpc:404 not found
func (ns *NewsService) Tags(ctx context.Context) ([]Tag, error) {
	list, err := ns.m.GetAllTag(ctx)
	if err != nil {
		return nil, newInternalError(err)
	} else if list == nil {
		return nil, noContentError
	}

	return NewTags(list), nil
}

// AddTag return new tag.
//
//zenrpc:404 not found
func (ns *NewsService) AddTag(ctx context.Context, in TagInput) (*Tag, error) {
	res, err := ns.m.AddTag(ctx, tagToManager(&in))
	if err != nil {
		return nil, newInternalError(err)
	}

	return NewTag(res), nil
}
