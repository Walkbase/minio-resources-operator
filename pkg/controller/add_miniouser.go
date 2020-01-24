package controller

import (
	"github.com/robotinfra/minio-resources-operator/pkg/controller/miniouser"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, miniouser.Add)
}
