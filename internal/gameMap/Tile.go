package mapModel

type Tile struct {
	Id         int16            `json:"id"`
	Properties []TileProperties `json:"properties"`
}

type TileProperties struct {
	Name string `json:"name"`
	Type string `json:"type"`
	/*
		Поле "Value" имеет тип "any", потому что в tileInfo.json в полях "tiles.properties" поле "value" может быть
		и bool, и string. Для backend нужно только поля с bool, а поля с string не имеют значения.

		TODO: почистить json и убрать тип "any"
	*/
	Value any `json:"value"`
}
