package log

func DebugF(msg string, fields ...any) {
	exLogger.loggerF.Debug(msg, fields)
}

func InfoF(msg string, fields ...any) {
	exLogger.loggerF.Info(msg, fields)
}

func WarnF(msg string, fields ...any) {
	exLogger.loggerF.Warn(msg, fields)
}

func ErrorF(msg string, fields ...any) {
	exLogger.loggerF.Error(msg, fields)
}

func DPanicF(msg string, fields ...any) {
	exLogger.loggerF.DPanic(msg, fields)
}

func PanicF(msg string, fields ...any) {
	exLogger.loggerF.Panic(msg, fields)
}

func FatalF(msg string, fields ...any) {
	exLogger.loggerF.Fatal(msg, fields)
}
