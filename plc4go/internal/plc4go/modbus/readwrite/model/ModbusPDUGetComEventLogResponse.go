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
type ModbusPDUGetComEventLogResponse struct {
	*ModbusPDU
	Status       uint16
	EventCount   uint16
	MessageCount uint16
	Events       []byte
}

// The corresponding interface
type IModbusPDUGetComEventLogResponse interface {
	IModbusPDU
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetEventCount returns EventCount (property field)
	GetEventCount() uint16
	// GetMessageCount returns MessageCount (property field)
	GetMessageCount() uint16
	// GetEvents returns Events (property field)
	GetEvents() []byte
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
func (m *ModbusPDUGetComEventLogResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUGetComEventLogResponse) GetFunctionFlag() uint8 {
	return 0x0C
}

func (m *ModbusPDUGetComEventLogResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUGetComEventLogResponse) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *ModbusPDUGetComEventLogResponse) GetStatus() uint16 {
	return m.Status
}

func (m *ModbusPDUGetComEventLogResponse) GetEventCount() uint16 {
	return m.EventCount
}

func (m *ModbusPDUGetComEventLogResponse) GetMessageCount() uint16 {
	return m.MessageCount
}

func (m *ModbusPDUGetComEventLogResponse) GetEvents() []byte {
	return m.Events
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUGetComEventLogResponse factory function for ModbusPDUGetComEventLogResponse
func NewModbusPDUGetComEventLogResponse(status uint16, eventCount uint16, messageCount uint16, events []byte) *ModbusPDU {
	child := &ModbusPDUGetComEventLogResponse{
		Status:       status,
		EventCount:   eventCount,
		MessageCount: messageCount,
		Events:       events,
		ModbusPDU:    NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUGetComEventLogResponse(structType interface{}) *ModbusPDUGetComEventLogResponse {
	if casted, ok := structType.(ModbusPDUGetComEventLogResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventLogResponse); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUGetComEventLogResponse(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUGetComEventLogResponse(casted.Child)
	}
	return nil
}

func (m *ModbusPDUGetComEventLogResponse) GetTypeName() string {
	return "ModbusPDUGetComEventLogResponse"
}

func (m *ModbusPDUGetComEventLogResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUGetComEventLogResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 16

	// Simple field (eventCount)
	lengthInBits += 16

	// Simple field (messageCount)
	lengthInBits += 16

	// Array field
	if len(m.Events) > 0 {
		lengthInBits += 8 * uint16(len(m.Events))
	}

	return lengthInBits
}

func (m *ModbusPDUGetComEventLogResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUGetComEventLogResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventLogResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint16("status", 16)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field")
	}
	status := _status

	// Simple Field (eventCount)
	_eventCount, _eventCountErr := readBuffer.ReadUint16("eventCount", 16)
	if _eventCountErr != nil {
		return nil, errors.Wrap(_eventCountErr, "Error parsing 'eventCount' field")
	}
	eventCount := _eventCount

	// Simple Field (messageCount)
	_messageCount, _messageCountErr := readBuffer.ReadUint16("messageCount", 16)
	if _messageCountErr != nil {
		return nil, errors.Wrap(_messageCountErr, "Error parsing 'messageCount' field")
	}
	messageCount := _messageCount
	// Byte Array field (events)
	numberOfBytesevents := int(uint16(byteCount) - uint16(uint16(6)))
	events, _readArrayErr := readBuffer.ReadByteArray("events", numberOfBytesevents)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'events' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventLogResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUGetComEventLogResponse{
		Status:       status,
		EventCount:   eventCount,
		MessageCount: messageCount,
		Events:       events,
		ModbusPDU:    &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUGetComEventLogResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventLogResponse"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(uint8(len(m.GetEvents()))) + uint8(uint8(6)))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Simple Field (status)
		status := uint16(m.Status)
		_statusErr := writeBuffer.WriteUint16("status", 16, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (eventCount)
		eventCount := uint16(m.EventCount)
		_eventCountErr := writeBuffer.WriteUint16("eventCount", 16, (eventCount))
		if _eventCountErr != nil {
			return errors.Wrap(_eventCountErr, "Error serializing 'eventCount' field")
		}

		// Simple Field (messageCount)
		messageCount := uint16(m.MessageCount)
		_messageCountErr := writeBuffer.WriteUint16("messageCount", 16, (messageCount))
		if _messageCountErr != nil {
			return errors.Wrap(_messageCountErr, "Error serializing 'messageCount' field")
		}

		// Array Field (events)
		if m.Events != nil {
			// Byte Array field (events)
			_writeArrayErr := writeBuffer.WriteByteArray("events", m.Events)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'events' field")
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventLogResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUGetComEventLogResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
