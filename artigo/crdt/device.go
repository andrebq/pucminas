package crdt

type (
	// Device holds information about a device in the network
	Device struct {
		// ID holds the xid ID associated with a given device
		ID string

		// PublicName used to identify the device
		PublicName string

		// PublicKey used to identify this device
		PublicKey PublicKey

		// LastClock holds the last clock value for this device
		LastClock Clock
	}

	// Sensor is a CRDT type representing a sensor. Since this is a CRDT
	// instead of just one device performing a reading, this sensor allows for
	// multiple devices to generate a reading.
	Sensor struct {
		// ID holds the unique ID of the sensor
		ID string

		// PublicName holds the public name for a sensor
		PublicName string

		// Readings contains the various readings from the sensors
		Readings []Reading
	}

	// Reading is a value that a given Device generated for a given region
	Reading struct {
		Stats

		// DeviceID indicates which device performed the reading
		DeviceID string

		// Clock of when the reading was performed
		Clock Clock
	}

	// Clock is a epoch based clock. It allows clock drift to be fixed by adjusting
	// the epoch
	Clock struct {
		// Epoch of the clock
		Epoch uint32
		// Unix milliseconds
		Unix int64
	}

	// Stats holds a series of statistics for a given sensor
	Stats struct {
		Current float64
		Min     float64
		Max     float64
		Count   float64
		Sum     float64
	}

	// Sensor holds information about a given sensor

	// PublicKey holds the hex string of a public key
	PublicKey string
)
