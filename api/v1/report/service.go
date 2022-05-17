package report

import "context"
import "github.com/harisaginting/tech-test-kredivo/pkg/tracer"

type Service struct {
	repo Repository
}

func ProviderService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (service *Service) List(ctx context.Context, res *ResponseList) {
	trace := tracer.Span(ctx,"ListUser")
	defer trace.End()
	users := service.repo.FindAll(ctx)
	res.Items = users
	res.Total = len(users)

	tracer.SetAttributeInt(trace,"total User",res.Total)
	return
}