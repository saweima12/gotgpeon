package ants

import "github.com/panjf2000/ants/v2"

var cPool *ants.Pool

func Init() error {
	var err error

	cPool, err = ants.NewPool(20000)
	if err != nil {
		return err
	}

	return nil
}

func Submit(task func()) error {
	return cPool.Submit(task)
}

func Release() {
	cPool.Release()
}
