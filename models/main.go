package models

func init() {
	// CoreModels provides a slice of models to migrate when initializing the api.
	CoreModels := make([]interface{}, 4)
	CoreModels[0] = &User{}
	CoreModels[1] = &PlayGroup{}
	CoreModels[2] = &Game{}
	CoreModels[3] = &Character{}
}
