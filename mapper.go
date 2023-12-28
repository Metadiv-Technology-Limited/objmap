package objmap

import (
	"reflect"

	"github.com/Metadiv-Technology-Limited/objmap/internal/util"
)

/*
Map2Model converts the model the generic-type model.
*/
func Map2Model[T any](from any) *T {
	from = util.NeverBePtr(from)
	to := reflect.ValueOf(new(T)).Elem()

	if from == nil {
		return nil
	}

	fields := util.ParseField(from)
	for _, f := range fields {
		to = util.SetField(to, f)
	}

	return to.Addr().Interface().(*T)
}

/*
Map2Models converts the models to the generic-type models.
*/
func Map2Models[T any](from []any) []T {
	var to []T
	for _, f := range from {
		to = append(to, *Map2Model[T](f))
	}
	return to
}

/*
MapModel2Model converts the model to the generic-type target model.
*/
func MapModel2Model[T any](from any, to *T) *T {
	from = util.NeverBePtr(from)

	if from == nil || to == nil {
		return nil
	}

	new := reflect.ValueOf(new(T)).Elem()

	fields := util.ParseField(to)
	for _, f := range fields {
		new = util.SetField(new, f)
	}

	fields = util.ParseField(from)
	for _, f := range fields {
		new = util.SetField(new, f)
	}

	return new.Addr().Interface().(*T)
}

/*
MapModel2Models converts the models to the generic-type target models.
*/
func MapModels2Models[T any](from []any, to []T) []T {
	for i, f := range from {
		to[i] = *MapModel2Model[T](f, &to[i])
	}
	return to
}
