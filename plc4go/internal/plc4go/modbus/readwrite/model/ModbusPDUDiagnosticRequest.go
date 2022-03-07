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
type ModbusPDUDiagnosticRequest struct {
	*ModbusPDU
	SubFunction uint16
	Data        uint16
}

// The corresponding interface
type IModbusPDUDiagnosticRequest interface {
	IModbusPDU
	// GetSubFunction returns SubFunction (property field)
	GetSubFunction() uint16
	// GetData returns Data (property field)
	GetData() uint16
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
func (m *ModbusPDUDiagnosticRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUDiagnosticRequest) GetFunctionFlag() uint8 {
	return 0x08
}

func (m *ModbusPDUDiagnosticRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUDiagnosticRequest) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *ModbusPDUDiagnosticRequest) GetSubFunction() uint16 {
	return m.SubFunction
}

func (m *ModbusPDUDiagnosticRequest) GetData() uint16 {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUDiagnosticRequest factory function for ModbusPDUDiagnosticRequest
func NewModbusPDUDiagnosticRequest(subFunction uint16, data uint16) *ModbusPDU {
	child := &ModbusPDUDiagnosticRequest{
		SubFunction: subFunction,
		Data:        data,
		ModbusPDU:   NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUDiagnosticRequest(structType interface{}) *ModbusPDUDiagnosticRequest {
	if casted, ok := structType.(ModbusPDUDiagnosticRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUDiagnosticRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUDiagnosticRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUDiagnosticRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUDiagnosticRequest) GetTypeName() string {
	return "ModbusPDUDiagnosticRequest"
}

func (m *ModbusPDUDiagnosticRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUDiagnosticRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subFunction)
	lengthInBits += 16

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *ModbusPDUDiagnosticRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUDiagnosticRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUDiagnosticRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (subFunction)
	_subFunction, _subFunctionErr := readBuffer.ReadUint16("subFunction", 16)
	if _subFunctionErr != nil {
		return nil, errors.Wrap(_subFunctionErr, "Error parsing 'subFunction' field")
	}
	subFunction := _subFunction

	// Simple Field (data)
	_data, _dataErr := readBuffer.ReadUint16("data", 16)
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field")
	}
	data := _data

	if closeErr := readBuffer.CloseContext("ModbusPDUDiagnosticRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUDiagnosticRequest{
		SubFunction: subFunction,
		Data:        data,
		ModbusPDU:   &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUDiagnosticRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUDiagnosticRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (subFunction)
		subFunction := uint16(m.SubFunction)
		_subFunctionErr := writeBuffer.WriteUint16("subFunction", 16, (subFunction))
		if _subFunctionErr != nil {
			return errors.Wrap(_subFunctionErr, "Error serializing 'subFunction' field")
		}

		// Simple Field (data)
		data := uint16(m.Data)
		_dataErr := writeBuffer.WriteUint16("data", 16, (data))
		if _dataErr != nil {
			return errors.Wrap(_dataErr, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUDiagnosticRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUDiagnosticRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
