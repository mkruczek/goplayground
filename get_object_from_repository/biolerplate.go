package get_object_from_repository

type Object struct {
	Name  string
	Value int
}

type RepositoryByValue interface {
	GetObject() Object
}

type RepositoryImplByValue struct {
	value Object
}

func (r RepositoryImplByValue) GetObject() Object {
	return r.value
}

type ServiceByValue struct {
	repo RepositoryByValue
}

func (s ServiceByValue) GetObject() Object {
	return s.repo.GetObject()
}

func mainFunctionByValue() {
	repo := RepositoryImplByPointer{value: Object{Name: "test", Value: 1}}
	service := ServiceByPointer{repo: repo}

	_ = service.GetObject()
}

type RepositoryByPointer interface {
	GetObject() *Object
}

type RepositoryImplByPointer struct {
	value Object
}

func (r RepositoryImplByPointer) GetObject() *Object {
	return &r.value
}

type ServiceByPointer struct {
	repo RepositoryByPointer
}

func (s ServiceByPointer) GetObject() *Object {
	return s.repo.GetObject()
}

func mainFunctionByPointer() {
	repo := RepositoryImplByPointer{value: Object{Name: "test", Value: 1}}
	service := ServiceByPointer{repo: repo}

	_ = service.GetObject()
}
