package cutkey

type CutData struct {

	// ej Modelo Objeto: usuario := Object{Name: "Usuario",Fields: []Field{
	// 		{Name: "name"}, //0
	// 		{Name: "email"},//1
	// 		{Name: "phone"},//2
	// 	},}
	// ej data normal con todos los campos: {"name":"John Doe","email":"johndoe@example.com","phone":"555"}
	// version recortada Data: {"John Doe","johndoe@example.com","555"}
	// Index al codificar = {"0:0","1:1","2:2"}
	// ej no mail: {"marcel", "777"}
	// Index al codificar = {"0:0","1:2"}
	Index map[uint8]uint8 `json:"i"`
	//Data ej en mapa: "Data":[{"id":"222","name":"manzana","valor":"1200"}]
	// ahora: "Data":["222","manzana","1200"]
	Data []string `json:"d"`
}

type CutResponse struct {
	//Type,Object,Module,Message
	//ej: ["read","user","Users","ok"]
	CutOptions []string  `json:"o"`
	CutData    []CutData `json:"d"`
}
