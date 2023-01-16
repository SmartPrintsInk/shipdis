package shipdis

import (
	"os"
)

var authorization string

const ordersEndpoint string = "https://ssapi.shipstation.com/orders"
const holdEndpoint string = "https://ssapi.shipstation.com/orders/holduntil"
const createEndpoint string = "https://ssapi.shipstation.com/orders/createorder"
const shipmentEndpoint string = "https://ssapi.shipstation.com/shipments"
const markAsShipped string = "https://ssapi.shipstation.com/orders/markasshipped"

func setup() {
	apiKey := os.Getenv("ShipstationAPIKey")
	authorization = "Basic " + apiKey
}
