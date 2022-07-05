/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ParameterChangeReply is the corresponding interface of ParameterChangeReply
type ParameterChangeReply interface {
	utils.LengthAware
	utils.Serializable
	NormalReply
	// GetIsA returns IsA (property field)
	GetIsA() ParameterChange
}

// ParameterChangeReplyExactly can be used when we want exactly this type and not a type which fulfills ParameterChangeReply.
// This is useful for switch cases.
type ParameterChangeReplyExactly interface {
	ParameterChangeReply
	isParameterChangeReply() bool
}

// _ParameterChangeReply is the data-structure of this message
type _ParameterChangeReply struct {
	*_NormalReply
	IsA ParameterChange
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterChangeReply) InitializeParent(parent NormalReply, peekedByte byte) {
	m.PeekedByte = peekedByte
}

func (m *_ParameterChangeReply) GetParent() NormalReply {
	return m._NormalReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterChangeReply) GetIsA() ParameterChange {
	return m.IsA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterChangeReply factory function for _ParameterChangeReply
func NewParameterChangeReply(isA ParameterChange, peekedByte byte, replyLength uint16) *_ParameterChangeReply {
	_result := &_ParameterChangeReply{
		IsA:          isA,
		_NormalReply: NewNormalReply(peekedByte, replyLength),
	}
	_result._NormalReply._NormalReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterChangeReply(structType interface{}) ParameterChangeReply {
	if casted, ok := structType.(ParameterChangeReply); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterChangeReply); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterChangeReply) GetTypeName() string {
	return "ParameterChangeReply"
}

func (m *_ParameterChangeReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ParameterChangeReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (isA)
	lengthInBits += m.IsA.GetLengthInBits()

	return lengthInBits
}

func (m *_ParameterChangeReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ParameterChangeReplyParse(readBuffer utils.ReadBuffer, replyLength uint16) (ParameterChangeReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterChangeReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterChangeReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (isA)
	if pullErr := readBuffer.PullContext("isA"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for isA")
	}
	_isA, _isAErr := ParameterChangeParse(readBuffer)
	if _isAErr != nil {
		return nil, errors.Wrap(_isAErr, "Error parsing 'isA' field of ParameterChangeReply")
	}
	isA := _isA.(ParameterChange)
	if closeErr := readBuffer.CloseContext("isA"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for isA")
	}

	if closeErr := readBuffer.CloseContext("ParameterChangeReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterChangeReply")
	}

	// Create a partially initialized instance
	_child := &_ParameterChangeReply{
		IsA: isA,
		_NormalReply: &_NormalReply{
			ReplyLength: replyLength,
		},
	}
	_child._NormalReply._NormalReplyChildRequirements = _child
	return _child, nil
}

func (m *_ParameterChangeReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterChangeReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterChangeReply")
		}

		// Simple Field (isA)
		if pushErr := writeBuffer.PushContext("isA"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for isA")
		}
		_isAErr := writeBuffer.WriteSerializable(m.GetIsA())
		if popErr := writeBuffer.PopContext("isA"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for isA")
		}
		if _isAErr != nil {
			return errors.Wrap(_isAErr, "Error serializing 'isA' field")
		}

		if popErr := writeBuffer.PopContext("ParameterChangeReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterChangeReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ParameterChangeReply) isParameterChangeReply() bool {
	return true
}

func (m *_ParameterChangeReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
