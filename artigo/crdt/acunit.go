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

// Update creates a new ACUnit object which is the result of combining
// the current state with the next ControlMessage
func (a *ACUnit) Update(m ControlMessage) ACUnit {
	var cp ACUnit
	cp.ID = a.ID
	cp.Messages = append(cp.Messages, a.Messages...)
	cp.updateInPlace(m)
	return cp
}

func (a *ACUnit) updateInPlace(m ControlMessage) {
	var updated bool
	for i, v := range a.Messages {
		if v.ControlUnitID == m.ControlUnitID {
			if v.Clock.Less(m.Clock) {
				a.Messages[i] = v
				updated = true
			}
		} else {
			a.Messages[i] = v
		}
	}
	if !updated {
		a.Messages = append(a.Messages, m)
	}

}

// Merge takes another ACUnit and returns the combination of two ACUnits
// the orignal values remain intact, to have sucessful merge, they MUST
// have the same ID
func (a *ACUnit) Merge(o *ACUnit) ACUnit {
	if a.ID != o.ID {
		return ACUnit{}
	}
	var cp ACUnit
	cp.ID = a.ID
	cp.Messages = append(cp.Messages, a.Messages...)
	for _, m := range o.Messages {
		cp.updateInPlace(m)
	}
	return cp
}

// Query returns the status of the ACUnit, the state is the number of ons
// minus the number of offs. Positive values are translated to true
// negative ones are translated to false. Zero means off
func (a *ACUnit) Query() bool {
	var ons, offs int
	for _, m := range a.Messages {
		if m.Status {
			ons++
		} else {
			offs++
		}
	}
	return (ons - offs) > 0
}
