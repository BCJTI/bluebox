package bluebox

const (
	MovementStatusBooked    = "BOOKED"
	MovementStatusInTransit = "IN_TRANSIT"
	MovementStatusArrived   = "ARRIVED"
)

const (
	ShipmentSubscriptionModeInsert = 1
	ShipmentSubscriptionModeUpdate = 2
	ShipmentSubscriptionModeUpsert = 3
	ShipmentSubscriptionModeDelete = 4
)

const (
	SubscriptionStatusOK                                    = "100"
	SubscriptionStatusUnsupportedCarrier                    = "201"
	SubscriptionStatusDuplicateSubscription                 = "202"
	SubscriptionStatusAccountNotAllowed                     = "301"
	SubscriptionStatusInvalidJSON                           = "302"
	SubscriptionStatusInvalidMAWB                           = "303"
	SubscriptionStatusNotFound                              = "404"
	SubscriptionStatusFoundSubscriptionInconsistentShipment = "405"
	SubscriptionStatusSubDeleteNotFound                     = "406"
	SubscriptionStatusSubUpdateNotFound                     = "407"
)

const (
	VehicleTypePlain = "P1"
	VehicleTypeTruck = "T1"
)
