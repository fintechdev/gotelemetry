package gotelemetry

type Widget struct {
	Variant        string  `json:"variant"`
	Board_id       string  `json:"board_id"`
	Column         float64 `json:"column"`
	Row            float64 `json:"row"`
	Width          float64 `json:"width"`
	Height         float64 `json:"height"`
	In_board_index float64 `json:"in_board_index"`
	Background     string  `json:"background"`
	Flow_ids       []string
	Updated_at     string
	Created_at     string
	Id             string
	Account_id     string
}

func (w *Widget) CreateWidget(credentials Credentials) error {
	request, err := buildRequest("POST", credentials, "/widgets", w)

	if err != nil {
		return err
	}

	var responseBody interface{}
	responseBody, err = sendJSONRequest(request)

	if err.(*Error).StatusCode == 400 {
		return err
	}

	w.mapToWidget(responseBody.(map[string]interface{}))

	return err
}

func (w *Widget) mapToWidget(m map[string]interface{}) {

	w.Height = m["height"].(float64)
	w.In_board_index = m["in_board_index"].(float64)
	w.Background = m["background"].(string)
	w.Board_id = m["board_id"].(string)

	w.Flow_ids = make([]string, 1)
	w.Flow_ids[0] = m["flow_ids"].([]interface{})[0].(string)

	w.Updated_at = m["updated_at"].(string)
	w.Id = m["id"].(string)
	w.Variant = m["variant"].(string)
	w.Column = m["column"].(float64)
	w.Row = m["row"].(float64)
	w.Width = m["width"].(float64)
	w.Account_id = m["account_id"].(string)
	w.Created_at = m["created_at"].(string)

}
