package gotelemetry

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Constants
const board_id string = "542ef7b37765627da4090000"
const api_key string = "2fcec2bd1c112dad755fbe1ad54db878"

func TestWidgets(t *testing.T) {

	Convey("Widgets", t, func() {

		widget := Widget{Variant: "log", Board_id: board_id, Column: 1, Row: 1, Width: 10, Height: 10, In_board_index: 0, Background: "default"}
		credentials, _ := NewCredentials(api_key)

		Convey("Should return status 201 when a Widget is created", func() {
			results := widget.CreateWidget(credentials)
			So(results.Error(), ShouldContainSubstring, "201")
		})

		Convey("Should return an invalid board_id error when the board does not exist", func() {
			testWidget := widget
			testWidget.Board_id = "I am not a valid board"
			results := testWidget.CreateWidget(credentials)
			So(results.Error(), ShouldContainSubstring, "400")
		})

		Convey("Should return a 400 when a invalid variant is passed", func() {
			testWidget := widget
			testWidget.Variant = "I am not a valid variant"
			results := testWidget.CreateWidget(credentials)
			So(results.Error(), ShouldContainSubstring, "400")
		})
	})

}
