package get_object_like_api_server

import "runtime"

type pointerRepository struct {
	value entity
}

func (r pointerRepository) getObject() *model {
	return &model{Name: r.value.Name, Value: r.value.Value}
}

type pointerService struct {
	repo pointerRepository
}

func (s pointerService) getObject() *model {
	return s.repo.getObject()
}

type pointerHandler struct {
	service pointerService
}

func (h pointerHandler) getObject() *response {
	m := h.service.getObject()
	return &response{Name: m.Name, Value: m.Value}
}

func mainFunctionByPointer() {
	repo := pointerRepository{value: entity{Name: "test", Value: 1}}
	service := pointerService{repo: repo}
	handler := pointerHandler{service: service}

	runtime.GC()

	_ = handler.getObject()
}
