//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"errors"
	"plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"
)

// The data-structure of this message
type COTPPacketDisconnectRequest struct {
	destinationReference uint16
	sourceReference      uint16
	protocolClass        ICOTPProtocolClass
	COTPPacket
}

// The corresponding interface
type ICOTPPacketDisconnectRequest interface {
	ICOTPPacket
	Serialize(io spi.WriteBuffer)
}

// Accessors for discriminator values.
func (m COTPPacketDisconnectRequest) TpduCode() uint8 {
	return 0x80
}

func (m COTPPacketDisconnectRequest) initialize(parameters []ICOTPParameter, payload *IS7Message) spi.Message {
	m.parameters = parameters
	m.payload = payload
	return m
}

func NewCOTPPacketDisconnectRequest(destinationReference uint16, sourceReference uint16, protocolClass ICOTPProtocolClass) COTPPacketInitializer {
	return &COTPPacketDisconnectRequest{destinationReference: destinationReference, sourceReference: sourceReference, protocolClass: protocolClass}
}

func CastICOTPPacketDisconnectRequest(structType interface{}) ICOTPPacketDisconnectRequest {
	castFunc := func(typ interface{}) ICOTPPacketDisconnectRequest {
		if iCOTPPacketDisconnectRequest, ok := typ.(ICOTPPacketDisconnectRequest); ok {
			return iCOTPPacketDisconnectRequest
		}
		return nil
	}
	return castFunc(structType)
}

func CastCOTPPacketDisconnectRequest(structType interface{}) COTPPacketDisconnectRequest {
	castFunc := func(typ interface{}) COTPPacketDisconnectRequest {
		if sCOTPPacketDisconnectRequest, ok := typ.(COTPPacketDisconnectRequest); ok {
			return sCOTPPacketDisconnectRequest
		}
		return COTPPacketDisconnectRequest{}
	}
	return castFunc(structType)
}

func (m COTPPacketDisconnectRequest) LengthInBits() uint16 {
	var lengthInBits uint16 = m.COTPPacket.LengthInBits()

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	// Enum Field (protocolClass)
	lengthInBits += 8

	return lengthInBits
}

func (m COTPPacketDisconnectRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPPacketDisconnectRequestParse(io *spi.ReadBuffer) (COTPPacketInitializer, error) {

	// Simple Field (destinationReference)
	destinationReference, _destinationReferenceErr := io.ReadUint16(16)
	if _destinationReferenceErr != nil {
		return nil, errors.New("Error parsing 'destinationReference' field " + _destinationReferenceErr.Error())
	}

	// Simple Field (sourceReference)
	sourceReference, _sourceReferenceErr := io.ReadUint16(16)
	if _sourceReferenceErr != nil {
		return nil, errors.New("Error parsing 'sourceReference' field " + _sourceReferenceErr.Error())
	}

	// Enum field (protocolClass)
	protocolClass, _protocolClassErr := COTPProtocolClassParse(io)
	if _protocolClassErr != nil {
		return nil, errors.New("Error parsing 'protocolClass' field " + _protocolClassErr.Error())
	}

	// Create the instance
	return NewCOTPPacketDisconnectRequest(destinationReference, sourceReference, protocolClass), nil
}

func (m COTPPacketDisconnectRequest) Serialize(io spi.WriteBuffer) {
	ser := func() {

		// Simple Field (destinationReference)
		destinationReference := uint16(m.destinationReference)
		io.WriteUint16(16, (destinationReference))

		// Simple Field (sourceReference)
		sourceReference := uint16(m.sourceReference)
		io.WriteUint16(16, (sourceReference))

		// Enum field (protocolClass)
		protocolClass := CastCOTPProtocolClass(m.protocolClass)
		protocolClass.Serialize(io)

	}
	COTPPacketSerialize(io, m.COTPPacket, CastICOTPPacket(m), ser)
}
