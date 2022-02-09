package main

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

//go:generate mockgen -source=main.go -package=main -destination=mock_iface.go
type Iface interface {
	Foo(a, b, c, d int)
}

type TestReporter interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Reporter struct{}

func (r *Reporter) Errorf(format string, args ...any) {
	fmt.Printf(format, args...)
}
func (r *Reporter) Fatalf(format string, args ...any) {
	fmt.Printf(format, args...)
}

func main() {
	mockController := gomock.NewController(&Reporter{})

	good := NewMockIface(mockController)
	good.EXPECT().Foo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Do(func(a, b, c, d int) {
		fmt.Println("Good mock called!")
	}).AnyTimes()
	good.Foo(1, 2, 3, 4)

	bad := NewMockIface(mockController)
	bad.EXPECT().Foo(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Do(func(args ...int) {
		fmt.Println("Bad mock called!")
	}).AnyTimes()
	bad.Foo(1, 2, 3, 4)
}
