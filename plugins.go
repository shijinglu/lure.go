package lure

import "sync"

var (
	functions map[string]IFunction
	extTypes  map[string]IDataFactory
	once      sync.Once
)

// IFunction defines the interface for customized LURE functions
type IFunction interface {
	// name identifies what function this is
	name() string
	// derive calculates the output from the input
	derive(params []IData) IData
}

// // IDataFactory defines the interface for extended data types
// type IDataFactory interface {
// 	create(raw string, extKey string) IData
// }

// IDataFactory is a function that can create IData
type IDataFactory = func(raw string, extKey string) IData

// getFunctions gets the plugin for functions
func getFunctions() map[string]IFunction {
	once.Do(func() {
		functions = make(map[string]IFunction)
		functions["md5mod"] = nil
	})
	return functions
}

// getExtDataTypes gets the plugin for extended data types
func getExtDataTypes() map[string]IDataFactory {
	once.Do(func() {
		extTypes = make(map[string]IDataFactory)
		extTypes["semver"] = VersionDataConstructor
	})
	return extTypes
}
