package simple

type simpleRepository struct {
}

func newSimpleRepository() *simpleRepository {
	return &simpleRepository{}
}

type simpleServices struct {
	*simpleRepository
}

func newSimpleServices(simpleRepository *simpleRepository) *simpleServices {
	return &simpleServices{simpleRepository: simpleRepository}
}
