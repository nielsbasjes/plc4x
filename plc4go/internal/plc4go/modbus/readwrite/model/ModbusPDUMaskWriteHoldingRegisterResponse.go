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
type ModbusPDUMaskWriteHoldingRegisterResponse struct {
	*ModbusPDU
	ReferenceAddress uint16
	AndMask          uint16
	OrMask           uint16
}

// The corresponding interface
type IModbusPDUMaskWriteHoldingRegisterResponse interface {
	IModbusPDU
	// GetReferenceAddress returns ReferenceAddress (property field)
	GetReferenceAddress() uint16
	// GetAndMask returns AndMask (property field)
	GetAndMask() uint16
	// GetOrMask returns OrMask (property field)
	GetOrMask() uint16
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
func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetFunctionFlag() uint8 {
	return 0x16
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetReferenceAddress() uint16 {
	return m.ReferenceAddress
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetAndMask() uint16 {
	return m.AndMask
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetOrMask() uint16 {
	return m.OrMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUMaskWriteHoldingRegisterResponse factory function for ModbusPDUMaskWriteHoldingRegisterResponse
func NewModbusPDUMaskWriteHoldingRegisterResponse(referenceAddress uint16, andMask uint16, orMask uint16) *ModbusPDU {
	child := &ModbusPDUMaskWriteHoldingRegisterResponse{
		ReferenceAddress: referenceAddress,
		AndMask:          andMask,
		OrMask:           orMask,
		ModbusPDU:        NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUMaskWriteHoldingRegisterResponse(structType interface{}) *ModbusPDUMaskWriteHoldingRegisterResponse {
	if casted, ok := structType.(ModbusPDUMaskWriteHoldingRegisterResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUMaskWriteHoldingRegisterResponse); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUMaskWriteHoldingRegisterResponse(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUMaskWriteHoldingRegisterResponse(casted.Child)
	}
	return nil
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetTypeName() string {
	return "ModbusPDUMaskWriteHoldingRegisterResponse"
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (referenceAddress)
	lengthInBits += 16

	// Simple field (andMask)
	lengthInBits += 16

	// Simple field (orMask)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUMaskWriteHoldingRegisterResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUMaskWriteHoldingRegisterResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (referenceAddress)
	_referenceAddress, _referenceAddressErr := readBuffer.ReadUint16("referenceAddress", 16)
	if _referenceAddressErr != nil {
		return nil, errors.Wrap(_referenceAddressErr, "Error parsing 'referenceAddress' field")
	}
	referenceAddress := _referenceAddress

	// Simple Field (andMask)
	_andMask, _andMaskErr := readBuffer.ReadUint16("andMask", 16)
	if _andMaskErr != nil {
		return nil, errors.Wrap(_andMaskErr, "Error parsing 'andMask' field")
	}
	andMask := _andMask

	// Simple Field (orMask)
	_orMask, _orMaskErr := readBuffer.ReadUint16("orMask", 16)
	if _orMaskErr != nil {
		return nil, errors.Wrap(_orMaskErr, "Error parsing 'orMask' field")
	}
	orMask := _orMask

	if closeErr := readBuffer.CloseContext("ModbusPDUMaskWriteHoldingRegisterResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUMaskWriteHoldingRegisterResponse{
		ReferenceAddress: referenceAddress,
		AndMask:          andMask,
		OrMask:           orMask,
		ModbusPDU:        &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUMaskWriteHoldingRegisterResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (referenceAddress)
		referenceAddress := uint16(m.ReferenceAddress)
		_referenceAddressErr := writeBuffer.WriteUint16("referenceAddress", 16, (referenceAddress))
		if _referenceAddressErr != nil {
			return errors.Wrap(_referenceAddressErr, "Error serializing 'referenceAddress' field")
		}

		// Simple Field (andMask)
		andMask := uint16(m.AndMask)
		_andMaskErr := writeBuffer.WriteUint16("andMask", 16, (andMask))
		if _andMaskErr != nil {
			return errors.Wrap(_andMaskErr, "Error serializing 'andMask' field")
		}

		// Simple Field (orMask)
		orMask := uint16(m.OrMask)
		_orMaskErr := writeBuffer.WriteUint16("orMask", 16, (orMask))
		if _orMaskErr != nil {
			return errors.Wrap(_orMaskErr, "Error serializing 'orMask' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUMaskWriteHoldingRegisterResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUMaskWriteHoldingRegisterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
