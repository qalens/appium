package appium

import "fmt"

type someDriver struct{}

var FirefoxDriver = someDriver{}

func (s someDriver) Log(msg string) {
	fmt.Println("Log: " + msg)
}
