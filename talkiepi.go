package talkiepi

import (
	"crypto/tls"

	"github.com/alexmotz/gpio"
	"github.com/alexmotz/gumble/gumble"
	"github.com/alexmotz/gumble/gumbleopenal"
)

// Raspberry Pi GPIO pin assignments (CPU pin definitions)
const (
	OnlineLEDPin       uint = 24
	ParticipantsLEDPin uint = 23
	TransmitLEDPin     uint = 12
	ButtonPin          uint = 25
)

type Talkiepi struct {
	Config *gumble.Config
	Client *gumble.Client

	Address   string
	TLSConfig tls.Config

	ConnectAttempts uint

	Stream *gumbleopenal.Stream

	ChannelName    string
	IsConnected    bool
	IsTransmitting bool

	GPIOEnabled     bool
	OnlineLED       gpio.Pin
	ParticipantsLED gpio.Pin
	TransmitLED     gpio.Pin
	Button          gpio.Pin
	ButtonState     uint
}
