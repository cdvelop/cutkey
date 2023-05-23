package cutkey

import (
	"log"

	"github.com/cdvelop/model"
)

func Add(models *map[string]model.Object) *Cut {

	if models == nil || len(*models) == 0 {
		log.Fatalln("Modelos de Objetos No encontrado en cutkey")
	}

	return &Cut{models: models}
}
