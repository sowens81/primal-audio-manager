package models

type PageSettings struct {
	Page    int
	PerPage int
}

const (
	defaultPage    = 1
	defaultPerPage = 100
)

type PageOption func(*PageSettings)

func NewPageSettings(opts ...PageOption) PageSettings {
	ps := PageSettings{
		Page:    defaultPage,
		PerPage: defaultPerPage,
	}

	for _, opt := range opts {
		opt(&ps)
	}

	return ps
}

func WithPage(page int) PageOption {
	return func(p *PageSettings) {
		if page > 0 {
			p.Page = page
		}
	}
}

func WithPerPage(perPage int) PageOption {
	return func(p *PageSettings) {
		if perPage > 0 {
			p.PerPage = perPage
		}
	}
}
