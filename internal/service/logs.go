// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	ILogs interface {
		Info(args ...interface{})
		Infof(template string, args ...interface{})
		Warn(args ...interface{})
		Warnf(template string, args ...interface{})
		Error(args ...interface{})
		Errorf(template string, args ...interface{})
		Fatal(args ...interface{})
		Fatalf(template string, args ...interface{})
	}
)

var (
	localLogs ILogs
)

func Logs() ILogs {
	if localLogs == nil {
		panic("implement not found for interface ILogs, forgot register?")
	}
	return localLogs
}

func RegisterLogs(i ILogs) {
	localLogs = i
}
