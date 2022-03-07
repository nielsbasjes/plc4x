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
type EipDisconnectRequest struct {
	*EipPacket
}

// The corresponding interface
type IEipDisconnectRequest interface {
	IEipPacket
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
func (m *EipDisconnectRequest) GetCommand() uint16 {
	return 0x0066
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *EipDisconnectRequest) InitializeParent(parent *EipPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.EipPacket.SessionHandle = sessionHandle
	m.EipPacket.Status = status
	m.EipPacket.SenderContext = senderContext
	m.EipPacket.Options = options
}

// NewEipDisconnectRequest factory function for EipDisconnectRequest
func NewEipDisconnectRequest(sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *EipPacket {
	child := &EipDisconnectRequest{
		EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	child.Child = child
	return child.EipPacket
}

func CastEipDisconnectRequest(structType interface{}) *EipDisconnectRequest {
	if casted, ok := structType.(EipDisconnectRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*EipDisconnectRequest); ok {
		return casted
	}
	if casted, ok := structType.(EipPacket); ok {
		return CastEipDisconnectRequest(casted.Child)
	}
	if casted, ok := structType.(*EipPacket); ok {
		return CastEipDisconnectRequest(casted.Child)
	}
	return nil
}

func (m *EipDisconnectRequest) GetTypeName() string {
	return "EipDisconnectRequest"
}

func (m *EipDisconnectRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *EipDisconnectRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *EipDisconnectRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func EipDisconnectRequestParse(readBuffer utils.ReadBuffer) (*EipPacket, error) {
	if pullErr := readBuffer.PullContext("EipDisconnectRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("EipDisconnectRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &EipDisconnectRequest{
		EipPacket: &EipPacket{},
	}
	_child.EipPacket.Child = _child
	return _child.EipPacket, nil
}

func (m *EipDisconnectRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipDisconnectRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("EipDisconnectRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *EipDisconnectRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
