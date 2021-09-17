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
type BACnetTagApplicationBoolean struct {
	Parent *BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationBoolean interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTagApplicationBoolean) TagClass() TagClass {
	return TagClass_APPLICATION_TAGS
}

func (m *BACnetTagApplicationBoolean) InitializeParent(parent *BACnetTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetTagApplicationBoolean(tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetTag {
	child := &BACnetTagApplicationBoolean{
		Parent: NewBACnetTag(tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetTagApplicationBoolean(structType interface{}) *BACnetTagApplicationBoolean {
	castFunc := func(typ interface{}) *BACnetTagApplicationBoolean {
		if casted, ok := typ.(BACnetTagApplicationBoolean); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagApplicationBoolean); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTag); ok {
			return CastBACnetTagApplicationBoolean(casted.Child)
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return CastBACnetTagApplicationBoolean(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagApplicationBoolean) GetTypeName() string {
	return "BACnetTagApplicationBoolean"
}

func (m *BACnetTagApplicationBoolean) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTagApplicationBoolean) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetTagApplicationBoolean) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagApplicationBooleanParse(readBuffer utils.ReadBuffer) (*BACnetTag, error) {
	if pullErr := readBuffer.PullContext("BACnetTagApplicationBoolean"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTagApplicationBoolean"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetTagApplicationBoolean{
		Parent: &BACnetTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetTagApplicationBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTagApplicationBoolean"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetTagApplicationBoolean"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetTagApplicationBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
