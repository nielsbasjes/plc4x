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
type KnxNetIpTunneling struct {
	*ServiceId
	Version uint8
}

// The corresponding interface
type IKnxNetIpTunneling interface {
	IServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
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
func (m *KnxNetIpTunneling) GetServiceType() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *KnxNetIpTunneling) InitializeParent(parent *ServiceId) {}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *KnxNetIpTunneling) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetIpTunneling factory function for KnxNetIpTunneling
func NewKnxNetIpTunneling(version uint8) *ServiceId {
	child := &KnxNetIpTunneling{
		Version:   version,
		ServiceId: NewServiceId(),
	}
	child.Child = child
	return child.ServiceId
}

func CastKnxNetIpTunneling(structType interface{}) *KnxNetIpTunneling {
	if casted, ok := structType.(KnxNetIpTunneling); ok {
		return &casted
	}
	if casted, ok := structType.(*KnxNetIpTunneling); ok {
		return casted
	}
	if casted, ok := structType.(ServiceId); ok {
		return CastKnxNetIpTunneling(casted.Child)
	}
	if casted, ok := structType.(*ServiceId); ok {
		return CastKnxNetIpTunneling(casted.Child)
	}
	return nil
}

func (m *KnxNetIpTunneling) GetTypeName() string {
	return "KnxNetIpTunneling"
}

func (m *KnxNetIpTunneling) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *KnxNetIpTunneling) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxNetIpTunneling) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetIpTunnelingParse(readBuffer utils.ReadBuffer) (*ServiceId, error) {
	if pullErr := readBuffer.PullContext("KnxNetIpTunneling"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetIpTunneling"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxNetIpTunneling{
		Version:   version,
		ServiceId: &ServiceId{},
	}
	_child.ServiceId.Child = _child
	return _child.ServiceId, nil
}

func (m *KnxNetIpTunneling) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpTunneling"); pushErr != nil {
			return pushErr
		}

		// Simple Field (version)
		version := uint8(m.Version)
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpTunneling"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxNetIpTunneling) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
