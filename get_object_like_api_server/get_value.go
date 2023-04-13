package get_object_like_api_server

type valueRepository struct {
	value entity
}

func (r valueRepository) getObject() model {
	return model{Name: r.value.Name, Value: r.value.Value}
}

type valueService struct {
	repo valueRepository
}

func (s valueService) getObject() model {
	return s.repo.getObject()
}

type valueHandler struct {
	service valueService
}

func (h valueHandler) getObject() response {
	m := h.service.getObject()
	return response{Name: m.Name, Value: m.Value}
}

func mainFunctionByValue(n int) {

	for i := 0; i < n; i++ {
		repo := valueRepository{value: entity{Name: "test", Value: 1}}
		service := valueService{repo: repo}
		handler := valueHandler{service: service}

		_ = handler.getObject()
	}
}
