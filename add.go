package cutkey

import (
	"log"

	model "github.com/cdvelop/go_model"
)

func Add(models *map[string]model.Object) *Cut {

	if models == nil || len(*models) == 0 {
		log.Fatalln("Modelos de Objetos No encontrado en cutkey")
	}

	return &Cut{models: models}
}
