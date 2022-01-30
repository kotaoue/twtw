package wd

import "github.com/kotaoue/go-wd"

func Get() string {
	dir, err := wd.Get()
	if err != nil {
		panic(err)
	}
	return dir
}
