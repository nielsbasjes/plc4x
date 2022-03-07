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
type BACnetConfirmedServiceACK struct {
	Child IBACnetConfirmedServiceACKChild
}

// The corresponding interface
type IBACnetConfirmedServiceACK interface {
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetConfirmedServiceACKParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConfirmedServiceACK, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetConfirmedServiceACKChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetConfirmedServiceACK)
	GetTypeName() string
	IBACnetConfirmedServiceACK
}

// NewBACnetConfirmedServiceACK factory function for BACnetConfirmedServiceACK
func NewBACnetConfirmedServiceACK() *BACnetConfirmedServiceACK {
	return &BACnetConfirmedServiceACK{}
}

func CastBACnetConfirmedServiceACK(structType interface{}) *BACnetConfirmedServiceACK {
	if casted, ok := structType.(BACnetConfirmedServiceACK); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceACK); ok {
		return casted
	}
	return nil
}

func (m *BACnetConfirmedServiceACK) GetTypeName() string {
	return "BACnetConfirmedServiceACK"
}

func (m *BACnetConfirmedServiceACK) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceACK) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetConfirmedServiceACK) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *BACnetConfirmedServiceACK) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceACKParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceACK, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceACK"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetConfirmedServiceACK
	var typeSwitchError error
	switch {
	case serviceChoice == 0x03: // BACnetConfirmedServiceACKGetAlarmSummary
		_parent, typeSwitchError = BACnetConfirmedServiceACKGetAlarmSummaryParse(readBuffer)
	case serviceChoice == 0x04: // BACnetConfirmedServiceACKGetEnrollmentSummary
		_parent, typeSwitchError = BACnetConfirmedServiceACKGetEnrollmentSummaryParse(readBuffer)
	case serviceChoice == 0x1D: // BACnetConfirmedServiceACKGetEventInformation
		_parent, typeSwitchError = BACnetConfirmedServiceACKGetEventInformationParse(readBuffer)
	case serviceChoice == 0x06: // BACnetConfirmedServiceACKAtomicReadFile
		_parent, typeSwitchError = BACnetConfirmedServiceACKAtomicReadFileParse(readBuffer)
	case serviceChoice == 0x07: // BACnetConfirmedServiceACKAtomicWriteFile
		_parent, typeSwitchError = BACnetConfirmedServiceACKAtomicWriteFileParse(readBuffer)
	case serviceChoice == 0x08: // BACnetConfirmedServiceAddListElement
		_parent, typeSwitchError = BACnetConfirmedServiceAddListElementParse(readBuffer)
	case serviceChoice == 0x0A: // BACnetConfirmedServiceACKCreateObject
		_parent, typeSwitchError = BACnetConfirmedServiceACKCreateObjectParse(readBuffer)
	case serviceChoice == 0x0C: // BACnetConfirmedServiceACKReadProperty
		_parent, typeSwitchError = BACnetConfirmedServiceACKReadPropertyParse(readBuffer)
	case serviceChoice == 0x0E: // BACnetConfirmedServiceACKReadPropertyMultiple
		_parent, typeSwitchError = BACnetConfirmedServiceACKReadPropertyMultipleParse(readBuffer)
	case serviceChoice == 0x1A: // BACnetConfirmedServiceACKReadRange
		_parent, typeSwitchError = BACnetConfirmedServiceACKReadRangeParse(readBuffer)
	case serviceChoice == 0x12: // BACnetConfirmedServiceACKConfirmedPrivateTransfer
		_parent, typeSwitchError = BACnetConfirmedServiceACKConfirmedPrivateTransferParse(readBuffer)
	case serviceChoice == 0x15: // BACnetConfirmedServiceACKVTOpen
		_parent, typeSwitchError = BACnetConfirmedServiceACKVTOpenParse(readBuffer)
	case serviceChoice == 0x17: // BACnetConfirmedServiceACKVTData
		_parent, typeSwitchError = BACnetConfirmedServiceACKVTDataParse(readBuffer)
	case serviceChoice == 0x18: // BACnetConfirmedServiceACKRemovedAuthenticate
		_parent, typeSwitchError = BACnetConfirmedServiceACKRemovedAuthenticateParse(readBuffer)
	case serviceChoice == 0x0D: // BACnetConfirmedServiceACKRemovedReadPropertyConditional
		_parent, typeSwitchError = BACnetConfirmedServiceACKRemovedReadPropertyConditionalParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceACK"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *BACnetConfirmedServiceACK) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetConfirmedServiceACK) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetConfirmedServiceACK, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceACK"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := uint8(child.GetServiceChoice())
	_serviceChoiceErr := writeBuffer.WriteUint8("serviceChoice", 8, (serviceChoice))

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceACK"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConfirmedServiceACK) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
