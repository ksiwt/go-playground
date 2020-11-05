package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := serviceError()
	fmt.Printf("%+v", err)
}

func repoError() error {
	e := fmt.Errorf("err")
	return errors.Wrap(e, "err")
}

func domainError() error {
	if err := repoError(); err != nil {
		return err
	}

	return nil
}

func serviceError() error {
	if err := domainError(); err != nil {
		return err
	}

	return nil
}
