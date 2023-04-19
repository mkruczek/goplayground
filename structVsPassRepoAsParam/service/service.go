package service

type Repo interface {
	Get()
	Create()
	Delete()
	Update()
}

type Service struct {
	Repo Repo
}

func (s *Service) Get() {
	s.Repo.Get()
}

func (s *Service) Create() {
	s.Repo.Create()
}

func (s *Service) Delete() {
	s.Repo.Delete()
}

func (s *Service) Update() {
	s.Repo.Update()
}

func Get(r Repo) {
	r.Get()
}

func Create(r Repo) {
	r.Create()
}

func Delete(r Repo) {
	r.Delete()
}

func Update(r Repo) {
	r.Update()
}
