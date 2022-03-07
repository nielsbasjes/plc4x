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
type ApduDataExtKeyWrite struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// The corresponding interface
type IApduDataExtKeyWrite interface {
	IApduDataExt
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
func (m *ApduDataExtKeyWrite) GetExtApciType() uint8 {
	return 0x13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtKeyWrite) InitializeParent(parent *ApduDataExt) {}

// NewApduDataExtKeyWrite factory function for ApduDataExtKeyWrite
func NewApduDataExtKeyWrite(length uint8) *ApduDataExt {
	child := &ApduDataExtKeyWrite{
		ApduDataExt: NewApduDataExt(length),
	}
	child.Child = child
	return child.ApduDataExt
}

func CastApduDataExtKeyWrite(structType interface{}) *ApduDataExtKeyWrite {
	if casted, ok := structType.(ApduDataExtKeyWrite); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtKeyWrite); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtKeyWrite(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtKeyWrite(casted.Child)
	}
	return nil
}

func (m *ApduDataExtKeyWrite) GetTypeName() string {
	return "ApduDataExtKeyWrite"
}

func (m *ApduDataExtKeyWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtKeyWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtKeyWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtKeyWriteParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtKeyWrite"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtKeyWrite"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtKeyWrite{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child.ApduDataExt, nil
}

func (m *ApduDataExtKeyWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtKeyWrite"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtKeyWrite"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtKeyWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
