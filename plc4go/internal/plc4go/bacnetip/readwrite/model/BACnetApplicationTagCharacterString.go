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
type BACnetApplicationTagCharacterString struct {
	*BACnetApplicationTag
	Payload *BACnetTagPayloadCharacterString
}

// The corresponding interface
type IBACnetApplicationTagCharacterString interface {
	IBACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() *BACnetTagPayloadCharacterString
	// GetValue returns Value (virtual field)
	GetValue() string
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
///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetApplicationTagCharacterString) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader) {
	m.BACnetApplicationTag.Header = header
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetApplicationTagCharacterString) GetPayload() *BACnetTagPayloadCharacterString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////
func (m *BACnetApplicationTagCharacterString) GetValue() string {
	return m.GetPayload().GetValue()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagCharacterString factory function for BACnetApplicationTagCharacterString
func NewBACnetApplicationTagCharacterString(payload *BACnetTagPayloadCharacterString, header *BACnetTagHeader) *BACnetApplicationTag {
	child := &BACnetApplicationTagCharacterString{
		Payload:              payload,
		BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagCharacterString(structType interface{}) *BACnetApplicationTagCharacterString {
	if casted, ok := structType.(BACnetApplicationTagCharacterString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetApplicationTagCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetApplicationTag); ok {
		return CastBACnetApplicationTagCharacterString(casted.Child)
	}
	if casted, ok := structType.(*BACnetApplicationTag); ok {
		return CastBACnetApplicationTagCharacterString(casted.Child)
	}
	return nil
}

func (m *BACnetApplicationTagCharacterString) GetTypeName() string {
	return "BACnetApplicationTagCharacterString"
}

func (m *BACnetApplicationTagCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetApplicationTagCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetApplicationTagCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetApplicationTagCharacterStringParse(readBuffer utils.ReadBuffer, header *BACnetTagHeader) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagCharacterString"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, pullErr
	}
	_payload, _payloadErr := BACnetTagPayloadCharacterStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := CastBACnetTagPayloadCharacterString(_payload)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_value := payload.GetValue()
	value := string(_value)
	_ = value

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagCharacterString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagCharacterString{
		Payload:              CastBACnetTagPayloadCharacterString(payload),
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagCharacterString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return pushErr
		}
		_payloadErr := m.Payload.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return popErr
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
		// Virtual field
		if _valueErr := writeBuffer.WriteVirtual("value", m.GetValue()); _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagCharacterString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
