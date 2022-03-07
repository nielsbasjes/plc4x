/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type COTPPacketTpduError struct {
	*COTPPacket
	DestinationReference uint16
	RejectCause          uint8

	// Arguments.
	CotpLen uint16
}

// The corresponding interface
type ICOTPPacketTpduError interface {
	ICOTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetRejectCause returns RejectCause (property field)
	GetRejectCause() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////
func (m *COTPPacketTpduError) GetTpduCode() uint8 {
	return 0x70
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *COTPPacketTpduError) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
	m.COTPPacket.Parameters = parameters
	m.COTPPacket.Payload = payload
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *COTPPacketTpduError) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *COTPPacketTpduError) GetRejectCause() uint8 {
	return m.RejectCause
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPPacketTpduError factory function for COTPPacketTpduError
func NewCOTPPacketTpduError(destinationReference uint16, rejectCause uint8, parameters []*COTPParameter, payload *S7Message, cotpLen uint16) *COTPPacket {
	child := &COTPPacketTpduError{
		DestinationReference: destinationReference,
		RejectCause:          rejectCause,
		COTPPacket:           NewCOTPPacket(parameters, payload, cotpLen),
	}
	child.Child = child
	return child.COTPPacket
}

func CastCOTPPacketTpduError(structType interface{}) *COTPPacketTpduError {
	if casted, ok := structType.(COTPPacketTpduError); ok {
		return &casted
	}
	if casted, ok := structType.(*COTPPacketTpduError); ok {
		return casted
	}
	if casted, ok := structType.(COTPPacket); ok {
		return CastCOTPPacketTpduError(casted.Child)
	}
	if casted, ok := structType.(*COTPPacket); ok {
		return CastCOTPPacketTpduError(casted.Child)
	}
	return nil
}

func (m *COTPPacketTpduError) GetTypeName() string {
	return "COTPPacketTpduError"
}

func (m *COTPPacketTpduError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *COTPPacketTpduError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (rejectCause)
	lengthInBits += 8

	return lengthInBits
}

func (m *COTPPacketTpduError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketTpduErrorParse(readBuffer utils.ReadBuffer, cotpLen uint16) (*COTPPacket, error) {
	if pullErr := readBuffer.PullContext("COTPPacketTpduError"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (destinationReference)
	_destinationReference, _destinationReferenceErr := readBuffer.ReadUint16("destinationReference", 16)
	if _destinationReferenceErr != nil {
		return nil, errors.Wrap(_destinationReferenceErr, "Error parsing 'destinationReference' field")
	}
	destinationReference := _destinationReference

	// Simple Field (rejectCause)
	_rejectCause, _rejectCauseErr := readBuffer.ReadUint8("rejectCause", 8)
	if _rejectCauseErr != nil {
		return nil, errors.Wrap(_rejectCauseErr, "Error parsing 'rejectCause' field")
	}
	rejectCause := _rejectCause

	if closeErr := readBuffer.CloseContext("COTPPacketTpduError"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &COTPPacketTpduError{
		DestinationReference: destinationReference,
		RejectCause:          rejectCause,
		COTPPacket:           &COTPPacket{},
	}
	_child.COTPPacket.Child = _child
	return _child.COTPPacket, nil
}

func (m *COTPPacketTpduError) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketTpduError"); pushErr != nil {
			return pushErr
		}

		// Simple Field (destinationReference)
		destinationReference := uint16(m.DestinationReference)
		_destinationReferenceErr := writeBuffer.WriteUint16("destinationReference", 16, (destinationReference))
		if _destinationReferenceErr != nil {
			return errors.Wrap(_destinationReferenceErr, "Error serializing 'destinationReference' field")
		}

		// Simple Field (rejectCause)
		rejectCause := uint8(m.RejectCause)
		_rejectCauseErr := writeBuffer.WriteUint8("rejectCause", 8, (rejectCause))
		if _rejectCauseErr != nil {
			return errors.Wrap(_rejectCauseErr, "Error serializing 'rejectCause' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketTpduError"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *COTPPacketTpduError) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
