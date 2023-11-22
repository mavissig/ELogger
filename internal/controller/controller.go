package controller

import "fmt"

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Info(s1 string, s2 string) error {
	fmt.Printf("s1: %s | s2: %s\n", s1, s2)
	return nil
}
