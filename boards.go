package gotelemetry

type Board struct {
	Id                 string    `json:"id,omitempty"`
	Name               string    `json:"name,omitempty"`
	Theme              string    `json:"theme,omitempty"`
	Columns            float64   `json:"columns,omitempty"`
	Rows               float64   `json:"rows,omitempty"`
	Size               []float64 `json:"size,omitempty"`
	Aspect_ratio       string    `json:"aspect_ratio,omitempty"`
	Display_board_name bool      `json:"display_board_name,omitempty"`
	Widget_margins     float64   `json:"widget_margins,omitempty"`
	Widget_padding     float64   `json:"widget_padding,omitempty"`
	Font_size          string
	Font_family        string
	Account_id         string
	Is_virtualized     bool
	Created_at         string
	Channel_ids        []string
	Is_affiliated      bool
	Updated_at         string
}

func (b *Board) CreateBoard(credentials Credentials) error {
	request, err := buildRequest("POST", credentials, "/boards", b)

	if err != nil {
		return err
	}

	var responseBody interface{}
	responseBody, err = sendJSONRequest(request)

	b.mapToBoard(responseBody.(map[string]interface{}))
	return err
}

func (b *Board) DeleteBoard(credentials Credentials) error {
	request, err := buildRequest("DELETE", credentials, "/boards/"+b.Id, nil)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)
	return err
}

func (b *Board) mapToBoard(m map[string]interface{}) {

	b.Account_id = m["account_id"].(string)
	b.Aspect_ratio = m["aspect_ratio"].(string)
	b.Is_virtualized = m["is_virtualized"].(bool)

	b.Size = make([]float64, 2)
	b.Size[0] = m["size"].([]interface{})[0].(float64)
	b.Size[1] = m["size"].([]interface{})[1].(float64)

	b.Created_at = m["created_at"].(string)
	b.Name = m["name"].(string)
	b.Theme = m["theme"].(string)

	b.Channel_ids = make([]string, 1)
	b.Channel_ids[0] = m["channel_ids"].([]interface{})[0].(string)

	b.Display_board_name = m["display_board_name"].(bool)
	b.Font_size = m["font_size"].(string)
	b.Id = m["id"].(string)
	b.Is_affiliated = m["is_affiliated"].(bool)
	b.Rows = m["rows"].(float64)
	b.Updated_at = m["updated_at"].(string)
	b.Columns = m["columns"].(float64)
	b.Font_family = m["font_family"].(string)
	b.Widget_margins = m["widget_margins"].(float64)
	b.Widget_padding = m["widget_padding"].(float64)

}
