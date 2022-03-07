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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ApduDataIndividualAddressWrite struct {
	*ApduData

	// Arguments.
	DataLength uint8
}

// The corresponding interface
type IApduDataIndividualAddressWrite interface {
	IApduData
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
func (m *ApduDataIndividualAddressWrite) GetApciType() uint8 {
	return 0x3
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataIndividualAddressWrite) InitializeParent(parent *ApduData) {}

// NewApduDataIndividualAddressWrite factory function for ApduDataIndividualAddressWrite
func NewApduDataIndividualAddressWrite(dataLength uint8) *ApduData {
	child := &ApduDataIndividualAddressWrite{
		ApduData: NewApduData(dataLength),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataIndividualAddressWrite(structType interface{}) *ApduDataIndividualAddressWrite {
	if casted, ok := structType.(ApduDataIndividualAddressWrite); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataIndividualAddressWrite); ok {
		return casted
	}
	if casted, ok := structType.(ApduData); ok {
		return CastApduDataIndividualAddressWrite(casted.Child)
	}
	if casted, ok := structType.(*ApduData); ok {
		return CastApduDataIndividualAddressWrite(casted.Child)
	}
	return nil
}

func (m *ApduDataIndividualAddressWrite) GetTypeName() string {
	return "ApduDataIndividualAddressWrite"
}

func (m *ApduDataIndividualAddressWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataIndividualAddressWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataIndividualAddressWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataIndividualAddressWriteParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataIndividualAddressWrite"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataIndividualAddressWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataIndividualAddressWrite{
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataIndividualAddressWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataIndividualAddressWrite"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataIndividualAddressWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataIndividualAddressWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
