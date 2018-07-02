package crdt

type (
	// ACUnit represents a given AC unit
	ACUnit struct {
		// ID is the ID for the given ACUnit
		ID string

		// Controls contains the list of control messages
		// for each control unit, this is another CRDT which
		// is queried to get the final status of the AC
		//
		// The latest entry is used to set the current value
		// for the AC, if the same number of control units set
		// the AC to on/off, the default status is used (which is on).
		Messages []ControlMessage
	}

	// ControlMessage is a unit of operation which is sent
	// by control units to toggle an AC on/off
	ControlMessage struct {
		// Status indicates if the AC should be on/off
		Status bool

		// ACUnitID holds the Id of the ac unit associated with this
		// control message
		ACUnitID string

		// ControlUnitID contains the control unit responsible for
		// the operation
		ControlUnitID string

		// Clock contains the timestamp of when the message was dispatched
		Clock Clock
	}

	// ControlUnit is a controller for all the system, it dispatches
	// for ACUnits
	ControlUnit struct {
		// ID contains the control unit ID
		ID string

		// PublicName for the given control unit
		PublicName string

		// PublicKey of this control unit
		PublicKey PublicKey
	}
)
