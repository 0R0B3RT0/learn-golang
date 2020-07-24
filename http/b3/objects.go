package b3

import (
	"time"
)

type Ticket struct {
	name     string
	business string
	cnpj     string
}

type Cotation struct {
	ticket Ticket
	time   time.Time
}
