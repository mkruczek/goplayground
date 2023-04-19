package caller

import (
	"some-benchmark/structVsPassRepoAsParam/repo"
	"some-benchmark/structVsPassRepoAsParam/service"
)

func callerStructWithRepo() {

	r := &repo.Repo{}
	s := service.Service{Repo: r}

	for i := 0; i < 1000; i++ {
		s.Get()
		s.Create()
		s.Delete()
		s.Update()
	}
}

func callerPassRepoAsParam() {

	r := &repo.Repo{}

	for i := 0; i < 1000; i++ {
		service.Get(r)
		service.Create(r)
		service.Delete(r)
		service.Update(r)
	}
}
