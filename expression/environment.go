package expression

type Environment struct {
	Names map[string]Expression
}

func NewEnvironment() Environment {
	return Environment{map[string]Expression{}}
}
