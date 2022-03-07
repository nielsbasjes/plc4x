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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CEMIAdditionalInformationBusmonitorInfo_LEN uint8 = uint8(1)

// The data-structure of this message
type CEMIAdditionalInformationBusmonitorInfo struct {
	*CEMIAdditionalInformation
	FrameErrorFlag  bool
	BitErrorFlag    bool
	ParityErrorFlag bool
	UnknownFlag     bool
	LostFlag        bool
	SequenceNumber  uint8
}

// The corresponding interface
type ICEMIAdditionalInformationBusmonitorInfo interface {
	ICEMIAdditionalInformation
	// GetFrameErrorFlag returns FrameErrorFlag (property field)
	GetFrameErrorFlag() bool
	// GetBitErrorFlag returns BitErrorFlag (property field)
	GetBitErrorFlag() bool
	// GetParityErrorFlag returns ParityErrorFlag (property field)
	GetParityErrorFlag() bool
	// GetUnknownFlag returns UnknownFlag (property field)
	GetUnknownFlag() bool
	// GetLostFlag returns LostFlag (property field)
	GetLostFlag() bool
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
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
func (m *CEMIAdditionalInformationBusmonitorInfo) GetAdditionalInformationType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *CEMIAdditionalInformationBusmonitorInfo) InitializeParent(parent *CEMIAdditionalInformation) {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *CEMIAdditionalInformationBusmonitorInfo) GetFrameErrorFlag() bool {
	return m.FrameErrorFlag
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetBitErrorFlag() bool {
	return m.BitErrorFlag
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetParityErrorFlag() bool {
	return m.ParityErrorFlag
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetUnknownFlag() bool {
	return m.UnknownFlag
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetLostFlag() bool {
	return m.LostFlag
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCEMIAdditionalInformationBusmonitorInfo factory function for CEMIAdditionalInformationBusmonitorInfo
func NewCEMIAdditionalInformationBusmonitorInfo(frameErrorFlag bool, bitErrorFlag bool, parityErrorFlag bool, unknownFlag bool, lostFlag bool, sequenceNumber uint8) *CEMIAdditionalInformation {
	child := &CEMIAdditionalInformationBusmonitorInfo{
		FrameErrorFlag:            frameErrorFlag,
		BitErrorFlag:              bitErrorFlag,
		ParityErrorFlag:           parityErrorFlag,
		UnknownFlag:               unknownFlag,
		LostFlag:                  lostFlag,
		SequenceNumber:            sequenceNumber,
		CEMIAdditionalInformation: NewCEMIAdditionalInformation(),
	}
	child.Child = child
	return child.CEMIAdditionalInformation
}

func CastCEMIAdditionalInformationBusmonitorInfo(structType interface{}) *CEMIAdditionalInformationBusmonitorInfo {
	if casted, ok := structType.(CEMIAdditionalInformationBusmonitorInfo); ok {
		return &casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformationBusmonitorInfo); ok {
		return casted
	}
	if casted, ok := structType.(CEMIAdditionalInformation); ok {
		return CastCEMIAdditionalInformationBusmonitorInfo(casted.Child)
	}
	if casted, ok := structType.(*CEMIAdditionalInformation); ok {
		return CastCEMIAdditionalInformationBusmonitorInfo(casted.Child)
	}
	return nil
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetTypeName() string {
	return "CEMIAdditionalInformationBusmonitorInfo"
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (len)
	lengthInBits += 8

	// Simple field (frameErrorFlag)
	lengthInBits += 1

	// Simple field (bitErrorFlag)
	lengthInBits += 1

	// Simple field (parityErrorFlag)
	lengthInBits += 1

	// Simple field (unknownFlag)
	lengthInBits += 1

	// Simple field (lostFlag)
	lengthInBits += 1

	// Simple field (sequenceNumber)
	lengthInBits += 3

	return lengthInBits
}

func (m *CEMIAdditionalInformationBusmonitorInfo) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CEMIAdditionalInformationBusmonitorInfoParse(readBuffer utils.ReadBuffer) (*CEMIAdditionalInformation, error) {
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformationBusmonitorInfo"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Const Field (len)
	len, _lenErr := readBuffer.ReadUint8("len", 8)
	if _lenErr != nil {
		return nil, errors.Wrap(_lenErr, "Error parsing 'len' field")
	}
	if len != CEMIAdditionalInformationBusmonitorInfo_LEN {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CEMIAdditionalInformationBusmonitorInfo_LEN) + " but got " + fmt.Sprintf("%d", len))
	}

	// Simple Field (frameErrorFlag)
	_frameErrorFlag, _frameErrorFlagErr := readBuffer.ReadBit("frameErrorFlag")
	if _frameErrorFlagErr != nil {
		return nil, errors.Wrap(_frameErrorFlagErr, "Error parsing 'frameErrorFlag' field")
	}
	frameErrorFlag := _frameErrorFlag

	// Simple Field (bitErrorFlag)
	_bitErrorFlag, _bitErrorFlagErr := readBuffer.ReadBit("bitErrorFlag")
	if _bitErrorFlagErr != nil {
		return nil, errors.Wrap(_bitErrorFlagErr, "Error parsing 'bitErrorFlag' field")
	}
	bitErrorFlag := _bitErrorFlag

	// Simple Field (parityErrorFlag)
	_parityErrorFlag, _parityErrorFlagErr := readBuffer.ReadBit("parityErrorFlag")
	if _parityErrorFlagErr != nil {
		return nil, errors.Wrap(_parityErrorFlagErr, "Error parsing 'parityErrorFlag' field")
	}
	parityErrorFlag := _parityErrorFlag

	// Simple Field (unknownFlag)
	_unknownFlag, _unknownFlagErr := readBuffer.ReadBit("unknownFlag")
	if _unknownFlagErr != nil {
		return nil, errors.Wrap(_unknownFlagErr, "Error parsing 'unknownFlag' field")
	}
	unknownFlag := _unknownFlag

	// Simple Field (lostFlag)
	_lostFlag, _lostFlagErr := readBuffer.ReadBit("lostFlag")
	if _lostFlagErr != nil {
		return nil, errors.Wrap(_lostFlagErr, "Error parsing 'lostFlag' field")
	}
	lostFlag := _lostFlag

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadUint8("sequenceNumber", 3)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field")
	}
	sequenceNumber := _sequenceNumber

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformationBusmonitorInfo"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CEMIAdditionalInformationBusmonitorInfo{
		FrameErrorFlag:            frameErrorFlag,
		BitErrorFlag:              bitErrorFlag,
		ParityErrorFlag:           parityErrorFlag,
		UnknownFlag:               unknownFlag,
		LostFlag:                  lostFlag,
		SequenceNumber:            sequenceNumber,
		CEMIAdditionalInformation: &CEMIAdditionalInformation{},
	}
	_child.CEMIAdditionalInformation.Child = _child
	return _child.CEMIAdditionalInformation, nil
}

func (m *CEMIAdditionalInformationBusmonitorInfo) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CEMIAdditionalInformationBusmonitorInfo"); pushErr != nil {
			return pushErr
		}

		// Const Field (len)
		_lenErr := writeBuffer.WriteUint8("len", 8, 1)
		if _lenErr != nil {
			return errors.Wrap(_lenErr, "Error serializing 'len' field")
		}

		// Simple Field (frameErrorFlag)
		frameErrorFlag := bool(m.FrameErrorFlag)
		_frameErrorFlagErr := writeBuffer.WriteBit("frameErrorFlag", (frameErrorFlag))
		if _frameErrorFlagErr != nil {
			return errors.Wrap(_frameErrorFlagErr, "Error serializing 'frameErrorFlag' field")
		}

		// Simple Field (bitErrorFlag)
		bitErrorFlag := bool(m.BitErrorFlag)
		_bitErrorFlagErr := writeBuffer.WriteBit("bitErrorFlag", (bitErrorFlag))
		if _bitErrorFlagErr != nil {
			return errors.Wrap(_bitErrorFlagErr, "Error serializing 'bitErrorFlag' field")
		}

		// Simple Field (parityErrorFlag)
		parityErrorFlag := bool(m.ParityErrorFlag)
		_parityErrorFlagErr := writeBuffer.WriteBit("parityErrorFlag", (parityErrorFlag))
		if _parityErrorFlagErr != nil {
			return errors.Wrap(_parityErrorFlagErr, "Error serializing 'parityErrorFlag' field")
		}

		// Simple Field (unknownFlag)
		unknownFlag := bool(m.UnknownFlag)
		_unknownFlagErr := writeBuffer.WriteBit("unknownFlag", (unknownFlag))
		if _unknownFlagErr != nil {
			return errors.Wrap(_unknownFlagErr, "Error serializing 'unknownFlag' field")
		}

		// Simple Field (lostFlag)
		lostFlag := bool(m.LostFlag)
		_lostFlagErr := writeBuffer.WriteBit("lostFlag", (lostFlag))
		if _lostFlagErr != nil {
			return errors.Wrap(_lostFlagErr, "Error serializing 'lostFlag' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.SequenceNumber)
		_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 3, (sequenceNumber))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("CEMIAdditionalInformationBusmonitorInfo"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CEMIAdditionalInformationBusmonitorInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
