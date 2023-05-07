# cutkey

protocolo escrito en go para intercambio de data json sin enviar el nombre de los campos solo funciona con mapa de string 
(map[string]string)

la idea principal es enviar la menor data posible que no es util.

esta pensado para trabajar con go del lado del cliente compilado a WebAssembly.

se eligió la librería cbor para codificar según la chatGpt es compatible con
tinyGo ..no lo he probado aun..
