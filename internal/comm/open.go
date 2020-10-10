package comm

import (
	"os"

	"github.com/jorgefuertes/mister-modemu/internal/console"
	"github.com/tarm/serial"
)

// Open - Open the serial port
func Open(port *string, baud *int) {
	prefix := "SER/OPEN"
	if _, err := os.Stat(*port); os.IsNotExist(err) {
		console.Error(prefix, "Cannot find port ", *port)
		os.Exit(4)
	}
	console.Debug(prefix, "Opening serial port")
	var err error
	m.port, err = serial.OpenPort(&serial.Config{Name: *port, Baud: *baud})
	if err != nil {
		console.Error(prefix, err.Error())
		os.Exit(1)
	}
	console.Debug(prefix, "Serial port open")
	resetStatus()
}

// Close - Closes the port
func Close() {
	m.port.Flush()
	m.port.Close()
}