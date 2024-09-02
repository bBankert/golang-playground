package outputtable

import (
	"example.com/playground/interfaces/displayer"
	"example.com/playground/interfaces/saver"
)

type Outputtable interface {
	saver.Saver
	displayer.Displayer
}

func OutputData(data Outputtable) {
	data.Display()
}
