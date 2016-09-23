package gamemod

import (
	"fmt"
	"reflect"
)

// Allowed types
//  map[string]BaiscModel
//  map[int64]BasicModel
//  *ArrayList
//  int64
//  string
//  nil
type Model interface{}

type ModelEntry interface{}

type IndexType int64

const (
	INT64  IndexType = 0
	STRING           = 1
)

type Index struct {
	indexType IndexType

	intValue    int64
	stringValue string
}

func indexTypeToReflectType(indexType IndexType) reflect.Type {
	return map[IndexType]reflect.Type{
		STRING: reflect.TypeOf(""),
		INT64:  reflect.TypeOf(int64(0)),
	}[indexType]
}

type Path []Index

type MutationType int64

const (
	INSERT MutationType = 0
	DELETE              = 5
	CLEAR               = 1
	SET                 = 2
	MOVE                = 3
	SWAP                = 4
)

type ModelMutation struct {
	mutationType MutationType
	path         Path
	index        Index
	param        Model // Only for INSERT and SET

	sourcePath  Path  // Only for MOVE or SWAP
	sourceIndex Index // Only for MOVE or SWAP
}

type TypeNotMatchError struct {
	path Path

	actualType   IndexType
	expectedType IndexType
}

func (e TypeNotMatchError) Error() string {
	return fmt.Sprintf("Invalid type: %v, expected: %v, path: %v",
		e.actualType, e.expectedType, e.path)
}

type EntryNotExistError struct {
	path Path
}

func (e EntryNotExistError) Error() string {
	return fmt.Sprintf("Entry not exist: %v ", e.path)
}

type InvalidModelTypeError struct {
	modelType reflect.Type
}

func (e InvalidModelTypeError) Error() string {
	return fmt.Sprintf("Invalid model type: %v ", e.modelType)
}

func MergePath(path Path, index Index) Path {
	return Path(
		append(path, index),
	)
}

func FindInstance(model Model, path Path) (result Model, err error) {
	var ok bool
	m := model
	for _, index := range path {
		if m == nil {
			return nil, EntryNotExistError{path: path}
		}
		switch m := model.(type) {
		case map[string]interface{}:
			if index.indexType == STRING {
				if model, ok = m[index.stringValue]; !ok {
					return nil, EntryNotExistError{path: path}
				}
			} else {
				return nil, TypeNotMatchError{
					path:         path,
					actualType:   index.indexType,
					expectedType: STRING,
				}
			}
		case map[int64]interface{}:
			if index.indexType == INT64 {
				if model, ok = m[index.intValue]; !ok {
					return nil, EntryNotExistError{path: path}
				}
			} else {
				return nil, TypeNotMatchError{
					path:         path,
					actualType:   index.indexType,
					expectedType: INT64,
				}
			}
		case *ArrayList:
			if index.indexType == INT64 {
				model, err = m.At(int(index.intValue))
				if err != nil {
					return
				}
			} else {
				return nil, TypeNotMatchError{
					path:         path,
					actualType:   index.indexType,
					expectedType: INT64,
				}
			}
		default:
			return nil, InvalidModelTypeError{modelType: reflect.TypeOf(model)}
		}
	}
	return model, nil
}

func ApplyMutations(model Model, mutations []*ModelMutation) (err error) {
	for _, mu := range mutations {
		err = ApplyMutation(model, mu)
		if err != nil {
			return
		}
	}
	return nil
}

func ApplyInsert(model Model, index Index, param Model) (err error) {
	switch m := model.(type) {
	case map[string]interface{}:
		if index.indexType != STRING {
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: STRING,
			}
		}
		if _, ok := m[index.stringValue]; ok {
			return EntryNotExistError{path: Path([]Index{})}
		}
		m[index.stringValue] = param
	case map[int64]interface{}:
		if index.indexType != INT64 {
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: INT64,
			}
		}
		if _, ok := m[index.intValue]; ok {
			return EntryNotExistError{path: Path([]Index{})}
		}
		m[index.intValue] = param
	case *ArrayList:
		if index.indexType != INT64 {
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: INT64,
			}
		}
		err = m.PushAt(int(index.intValue), param)
		if err != nil {
			return
		}
	default:
		return InvalidModelTypeError{modelType: reflect.TypeOf(model)}
	}
	return nil
}

func ApplyDelete(model Model, index Index) (err error) {
	switch m := model.(type) {
	case map[string]interface{}:
		if index.indexType != STRING {
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: STRING,
			}
		}
		if _, ok := m[index.stringValue]; !ok {
			return EntryNotExistError{path: Path([]Index{})}
		}
		delete(m, index.stringValue)
	case map[int64]interface{}:
		if index.indexType != INT64 {
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: INT64,
			}
		}
		if _, ok := m[index.intValue]; !ok {
			return EntryNotExistError{path: Path([]Index{})}
		}
		delete(m, index.intValue)
	case *ArrayList:
		if index.indexType != INT64 {
			return
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: INT64,
			}
		}
		_, err = m.PopAt(int(index.intValue))
		if err != nil {
			return
		}
	default:
		return InvalidModelTypeError{modelType: reflect.TypeOf(model)}
	}
	return nil
}

func ApplySet(model Model, index Index, param Model) (err error) {
	switch m := model.(type) {
	case map[string]interface{}:
		if index.indexType != STRING {
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: STRING,
			}
		}
		m[index.stringValue] = param
	case map[int64]interface{}:
		if index.indexType != INT64 {
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: INT64,
			}
		}
		m[index.intValue] = param
	case *ArrayList:
		if index.indexType != INT64 {
			return TypeNotMatchError{
				path:         Path([]Index{}),
				actualType:   index.indexType,
				expectedType: INT64,
			}
		}
		err = m.Set(int(index.intValue), param)
		if err != nil {
			return
		}
	default:
		return InvalidModelTypeError{modelType: reflect.TypeOf(model)}
	}
	return nil
}

func ApplyMutation(model Model, mutation *ModelMutation) (err error) {
	var m Model
	switch mutation.mutationType {
	case INSERT:
		m, err = FindInstance(model, mutation.path)
		if err != nil {
			return err
		}
		return ApplyInsert(m, mutation.index, mutation.param)
	case DELETE:
		m, err = FindInstance(model, mutation.path)
		if err != nil {
			return err
		}
		return ApplyDelete(m, mutation.index)
	case SET:
		m, err = FindInstance(model, mutation.path)
		if err != nil {
			return err
		}
		return ApplySet(m, mutation.index, mutation.param)
	case MOVE:
		// Find source model and its parent
		var sourceModel, sourceParentModel, destParentModel Model
		sourceModel, err = FindInstance(model, MergePath(mutation.sourcePath, mutation.sourceIndex))
		if err != nil {
			return err
		}
		sourceParentModel, err = FindInstance(model, mutation.sourcePath)
		if err != nil {
			return err
		}

		// Destination path
		destParentModel, err = FindInstance(model, mutation.path)
		if err != nil {
			return err
		}
		err = ApplyInsert(destParentModel, mutation.index, sourceModel)
		if err != nil {
			return err
		}
		return ApplyDelete(sourceParentModel, mutation.sourceIndex)
	case SWAP:
		var sourceModel, sourceParentModel, destModel, destParentModel Model
		// Find source model and its parent
		sourceModel, err = FindInstance(model, MergePath(mutation.sourcePath, mutation.sourceIndex))
		if err != nil {
			return err
		}
		sourceParentModel, err = FindInstance(model, mutation.sourcePath)
		if err != nil {
			return err
		}

		// Destination model and its parent
		destModel, err = FindInstance(model, MergePath(mutation.path, mutation.index))
		if err != nil {
			return err
		}
		destParentModel, err = FindInstance(model, mutation.path)
		if err != nil {
			return err
		}
		err = ApplySet(destParentModel, mutation.index, sourceModel)
		if err != nil {
			return err
		}
		return ApplySet(sourceParentModel, mutation.sourceIndex, destModel)
	}
	return nil
}

func main() {
	m1 := Model(1)
	m2 := Model(map[string]interface{}{
		"ABC": m1,
	})
	m3 := Model(map[int64]interface{}{
		1: m2,
		2: m1,
	})
	m4 := Model(NewArrayList([]interface{}{
		Model(2),
		Model(3),
		m3,
	}))

	fmt.Println(m4)
	fmt.Println(ApplyMutation(m4, &ModelMutation{
		mutationType: SWAP,
		path: Path(
			[]Index{Index{indexType: INT64, intValue: 2}},
		),
		index: Index{
			indexType: INT64,
			intValue:  1,
		},
		sourcePath: Path([]Index{}),
		sourceIndex: Index{
			indexType: INT64,
			intValue:  1,
		},
	}))
	fmt.Println(m4)
}
