package cutkey

import (
	"log"

	model "github.com/cdvelop/go_model"
)

type cut struct {
	models *map[string]model.Object
}

func Add(models *map[string]model.Object) *cut {

	if models == nil || len(*models) == 0 {
		log.Fatalln("Modelos de Objetos No encontrado en cutkey")
	}

	return &cut{models: models}
}
