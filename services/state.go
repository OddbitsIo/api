package services

import (
	"github.com/oddbitsio/api/mdb-repos"
)

func Init() error {
	return mdbrepos.Init()
}

func TearDown() {
	mdbrepos.TearDown()
}