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
type KnxGroupAddress3Level struct {
	*KnxGroupAddress
	MainGroup   uint8
	MiddleGroup uint8
	SubGroup    uint8
}

// The corresponding interface
type IKnxGroupAddress3Level interface {
	IKnxGroupAddress
	// GetMainGroup returns MainGroup (property field)
	GetMainGroup() uint8
	// GetMiddleGroup returns MiddleGroup (property field)
	GetMiddleGroup() uint8
	// GetSubGroup returns SubGroup (property field)
	GetSubGroup() uint8
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
func (m *KnxGroupAddress3Level) GetNumLevels() uint8 {
	return uint8(3)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *KnxGroupAddress3Level) InitializeParent(parent *KnxGroupAddress) {}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *KnxGroupAddress3Level) GetMainGroup() uint8 {
	return m.MainGroup
}

func (m *KnxGroupAddress3Level) GetMiddleGroup() uint8 {
	return m.MiddleGroup
}

func (m *KnxGroupAddress3Level) GetSubGroup() uint8 {
	return m.SubGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxGroupAddress3Level factory function for KnxGroupAddress3Level
func NewKnxGroupAddress3Level(mainGroup uint8, middleGroup uint8, subGroup uint8) *KnxGroupAddress {
	child := &KnxGroupAddress3Level{
		MainGroup:       mainGroup,
		MiddleGroup:     middleGroup,
		SubGroup:        subGroup,
		KnxGroupAddress: NewKnxGroupAddress(),
	}
	child.Child = child
	return child.KnxGroupAddress
}

func CastKnxGroupAddress3Level(structType interface{}) *KnxGroupAddress3Level {
	if casted, ok := structType.(KnxGroupAddress3Level); ok {
		return &casted
	}
	if casted, ok := structType.(*KnxGroupAddress3Level); ok {
		return casted
	}
	if casted, ok := structType.(KnxGroupAddress); ok {
		return CastKnxGroupAddress3Level(casted.Child)
	}
	if casted, ok := structType.(*KnxGroupAddress); ok {
		return CastKnxGroupAddress3Level(casted.Child)
	}
	return nil
}

func (m *KnxGroupAddress3Level) GetTypeName() string {
	return "KnxGroupAddress3Level"
}

func (m *KnxGroupAddress3Level) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *KnxGroupAddress3Level) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (mainGroup)
	lengthInBits += 5

	// Simple field (middleGroup)
	lengthInBits += 3

	// Simple field (subGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxGroupAddress3Level) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxGroupAddress3LevelParse(readBuffer utils.ReadBuffer, numLevels uint8) (*KnxGroupAddress, error) {
	if pullErr := readBuffer.PullContext("KnxGroupAddress3Level"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (mainGroup)
	_mainGroup, _mainGroupErr := readBuffer.ReadUint8("mainGroup", 5)
	if _mainGroupErr != nil {
		return nil, errors.Wrap(_mainGroupErr, "Error parsing 'mainGroup' field")
	}
	mainGroup := _mainGroup

	// Simple Field (middleGroup)
	_middleGroup, _middleGroupErr := readBuffer.ReadUint8("middleGroup", 3)
	if _middleGroupErr != nil {
		return nil, errors.Wrap(_middleGroupErr, "Error parsing 'middleGroup' field")
	}
	middleGroup := _middleGroup

	// Simple Field (subGroup)
	_subGroup, _subGroupErr := readBuffer.ReadUint8("subGroup", 8)
	if _subGroupErr != nil {
		return nil, errors.Wrap(_subGroupErr, "Error parsing 'subGroup' field")
	}
	subGroup := _subGroup

	if closeErr := readBuffer.CloseContext("KnxGroupAddress3Level"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxGroupAddress3Level{
		MainGroup:       mainGroup,
		MiddleGroup:     middleGroup,
		SubGroup:        subGroup,
		KnxGroupAddress: &KnxGroupAddress{},
	}
	_child.KnxGroupAddress.Child = _child
	return _child.KnxGroupAddress, nil
}

func (m *KnxGroupAddress3Level) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxGroupAddress3Level"); pushErr != nil {
			return pushErr
		}

		// Simple Field (mainGroup)
		mainGroup := uint8(m.MainGroup)
		_mainGroupErr := writeBuffer.WriteUint8("mainGroup", 5, (mainGroup))
		if _mainGroupErr != nil {
			return errors.Wrap(_mainGroupErr, "Error serializing 'mainGroup' field")
		}

		// Simple Field (middleGroup)
		middleGroup := uint8(m.MiddleGroup)
		_middleGroupErr := writeBuffer.WriteUint8("middleGroup", 3, (middleGroup))
		if _middleGroupErr != nil {
			return errors.Wrap(_middleGroupErr, "Error serializing 'middleGroup' field")
		}

		// Simple Field (subGroup)
		subGroup := uint8(m.SubGroup)
		_subGroupErr := writeBuffer.WriteUint8("subGroup", 8, (subGroup))
		if _subGroupErr != nil {
			return errors.Wrap(_subGroupErr, "Error serializing 'subGroup' field")
		}

		if popErr := writeBuffer.PopContext("KnxGroupAddress3Level"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxGroupAddress3Level) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
