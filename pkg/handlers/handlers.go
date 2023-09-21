package handlers

type Handlers struct {
	*InstanceHandler
}

func New() *Handlers {
	return &Handlers{
		InstanceHandler: newInstanceHandler(),
	}
}
