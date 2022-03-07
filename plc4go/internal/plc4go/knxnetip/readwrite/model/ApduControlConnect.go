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
type ApduControlConnect struct {
	*ApduControl
}

// The corresponding interface
type IApduControlConnect interface {
	IApduControl
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
func (m *ApduControlConnect) GetControlType() uint8 {
	return 0x0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduControlConnect) InitializeParent(parent *ApduControl) {}

// NewApduControlConnect factory function for ApduControlConnect
func NewApduControlConnect() *ApduControl {
	child := &ApduControlConnect{
		ApduControl: NewApduControl(),
	}
	child.Child = child
	return child.ApduControl
}

func CastApduControlConnect(structType interface{}) *ApduControlConnect {
	if casted, ok := structType.(ApduControlConnect); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduControlConnect); ok {
		return casted
	}
	if casted, ok := structType.(ApduControl); ok {
		return CastApduControlConnect(casted.Child)
	}
	if casted, ok := structType.(*ApduControl); ok {
		return CastApduControlConnect(casted.Child)
	}
	return nil
}

func (m *ApduControlConnect) GetTypeName() string {
	return "ApduControlConnect"
}

func (m *ApduControlConnect) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduControlConnect) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduControlConnect) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduControlConnectParse(readBuffer utils.ReadBuffer) (*ApduControl, error) {
	if pullErr := readBuffer.PullContext("ApduControlConnect"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduControlConnect"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduControlConnect{
		ApduControl: &ApduControl{},
	}
	_child.ApduControl.Child = _child
	return _child.ApduControl, nil
}

func (m *ApduControlConnect) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlConnect"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduControlConnect"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduControlConnect) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
