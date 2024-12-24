package entities

type Entity[T any] struct {
	id    *UniqueEntityID
	props T
}

// Cria uma nova entidade com base nas propriedades e no ID fornecidos
func NewEntity[T any](props T, id ...*UniqueEntityID) *Entity[T] {
	var entityID *UniqueEntityID
	if len(id) > 0 && id[0] != nil {
		entityID = id[0]
	} else {
		entityID = NewUniqueEntityID()
	}

	return &Entity[T]{
		id:    entityID,
		props: props,
	}
}

func (e *Entity[T]) ID() *UniqueEntityID {
	return e.id
}

// Props retorna as propriedades da entidade
func (e *Entity[T]) Props() *T {
	return &e.props
}

// Equals verifica se duas entidades são equivalentes com base no ID
func (e *Entity[T]) Equals(other *Entity[T]) bool {
	if other == nil {
		return false
	}
	return e.id.Equals(other.id)
}
