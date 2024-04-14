package types

// ValidateBasic is used for validating the packet
func (p IbcEncodedDataPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcEncodedDataPacketData) GetBytes() ([]byte, error) {
	var modulePacket CodenetPacketData

	modulePacket.Packet = &CodenetPacketData_IbcEncodedDataPacket{&p}

	return modulePacket.Marshal()
}
