package model

import "go.uber.org/fx"

var options []fx.Option

// Provide 提供
func Provide(constructors ...interface{}) {
	options = append(options, fx.Provide(constructors...))
}

// Invoke 调用
func Invoke(funcEs ...interface{}) {
	options = append(options, fx.Invoke(funcEs...))
}

func GetOptions() []fx.Option {
	return options
}
