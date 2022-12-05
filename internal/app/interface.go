package app

import "github.com/kormiltsev/filereaders/readers"

type DataTypeSwitcher interface {
	Convert(*readers.InterfaceCSV)
}
