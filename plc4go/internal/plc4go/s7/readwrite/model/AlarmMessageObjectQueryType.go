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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const AlarmMessageObjectQueryType_VARIABLESPEC uint8 = 0x12

// The data-structure of this message
type AlarmMessageObjectQueryType struct {
	LengthDataset  uint8
	EventState     *State
	AckStateGoing  *State
	AckStateComing *State
	TimeComing     *DateAndTime
	ValueComing    *AssociatedValueType
	TimeGoing      *DateAndTime
	ValueGoing     *AssociatedValueType
}

// The corresponding interface
type IAlarmMessageObjectQueryType interface {
	// GetLengthDataset returns LengthDataset (property field)
	GetLengthDataset() uint8
	// GetEventState returns EventState (property field)
	GetEventState() *State
	// GetAckStateGoing returns AckStateGoing (property field)
	GetAckStateGoing() *State
	// GetAckStateComing returns AckStateComing (property field)
	GetAckStateComing() *State
	// GetTimeComing returns TimeComing (property field)
	GetTimeComing() *DateAndTime
	// GetValueComing returns ValueComing (property field)
	GetValueComing() *AssociatedValueType
	// GetTimeGoing returns TimeGoing (property field)
	GetTimeGoing() *DateAndTime
	// GetValueGoing returns ValueGoing (property field)
	GetValueGoing() *AssociatedValueType
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *AlarmMessageObjectQueryType) GetLengthDataset() uint8 {
	return m.LengthDataset
}

func (m *AlarmMessageObjectQueryType) GetEventState() *State {
	return m.EventState
}

func (m *AlarmMessageObjectQueryType) GetAckStateGoing() *State {
	return m.AckStateGoing
}

func (m *AlarmMessageObjectQueryType) GetAckStateComing() *State {
	return m.AckStateComing
}

func (m *AlarmMessageObjectQueryType) GetTimeComing() *DateAndTime {
	return m.TimeComing
}

func (m *AlarmMessageObjectQueryType) GetValueComing() *AssociatedValueType {
	return m.ValueComing
}

func (m *AlarmMessageObjectQueryType) GetTimeGoing() *DateAndTime {
	return m.TimeGoing
}

func (m *AlarmMessageObjectQueryType) GetValueGoing() *AssociatedValueType {
	return m.ValueGoing
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////
func (m *AlarmMessageObjectQueryType) GetVariableSpec() uint8 {
	return AlarmMessageObjectQueryType_VARIABLESPEC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageObjectQueryType factory function for AlarmMessageObjectQueryType
func NewAlarmMessageObjectQueryType(lengthDataset uint8, eventState *State, ackStateGoing *State, ackStateComing *State, timeComing *DateAndTime, valueComing *AssociatedValueType, timeGoing *DateAndTime, valueGoing *AssociatedValueType) *AlarmMessageObjectQueryType {
	return &AlarmMessageObjectQueryType{LengthDataset: lengthDataset, EventState: eventState, AckStateGoing: ackStateGoing, AckStateComing: ackStateComing, TimeComing: timeComing, ValueComing: valueComing, TimeGoing: timeGoing, ValueGoing: valueGoing}
}

func CastAlarmMessageObjectQueryType(structType interface{}) *AlarmMessageObjectQueryType {
	if casted, ok := structType.(AlarmMessageObjectQueryType); ok {
		return &casted
	}
	if casted, ok := structType.(*AlarmMessageObjectQueryType); ok {
		return casted
	}
	return nil
}

func (m *AlarmMessageObjectQueryType) GetTypeName() string {
	return "AlarmMessageObjectQueryType"
}

func (m *AlarmMessageObjectQueryType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AlarmMessageObjectQueryType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (lengthDataset)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 16

	// Const Field (variableSpec)
	lengthInBits += 8

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits()

	// Simple field (ackStateGoing)
	lengthInBits += m.AckStateGoing.GetLengthInBits()

	// Simple field (ackStateComing)
	lengthInBits += m.AckStateComing.GetLengthInBits()

	// Simple field (timeComing)
	lengthInBits += m.TimeComing.GetLengthInBits()

	// Simple field (valueComing)
	lengthInBits += m.ValueComing.GetLengthInBits()

	// Simple field (timeGoing)
	lengthInBits += m.TimeGoing.GetLengthInBits()

	// Simple field (valueGoing)
	lengthInBits += m.ValueGoing.GetLengthInBits()

	return lengthInBits
}

func (m *AlarmMessageObjectQueryType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AlarmMessageObjectQueryTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessageObjectQueryType, error) {
	if pullErr := readBuffer.PullContext("AlarmMessageObjectQueryType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (lengthDataset)
	_lengthDataset, _lengthDatasetErr := readBuffer.ReadUint8("lengthDataset", 8)
	if _lengthDatasetErr != nil {
		return nil, errors.Wrap(_lengthDatasetErr, "Error parsing 'lengthDataset' field")
	}
	lengthDataset := _lengthDataset

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Const Field (variableSpec)
	variableSpec, _variableSpecErr := readBuffer.ReadUint8("variableSpec", 8)
	if _variableSpecErr != nil {
		return nil, errors.Wrap(_variableSpecErr, "Error parsing 'variableSpec' field")
	}
	if variableSpec != AlarmMessageObjectQueryType_VARIABLESPEC {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AlarmMessageObjectQueryType_VARIABLESPEC) + " but got " + fmt.Sprintf("%d", variableSpec))
	}

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, pullErr
	}
	_eventState, _eventStateErr := StateParse(readBuffer)
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field")
	}
	eventState := CastState(_eventState)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (ackStateGoing)
	if pullErr := readBuffer.PullContext("ackStateGoing"); pullErr != nil {
		return nil, pullErr
	}
	_ackStateGoing, _ackStateGoingErr := StateParse(readBuffer)
	if _ackStateGoingErr != nil {
		return nil, errors.Wrap(_ackStateGoingErr, "Error parsing 'ackStateGoing' field")
	}
	ackStateGoing := CastState(_ackStateGoing)
	if closeErr := readBuffer.CloseContext("ackStateGoing"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (ackStateComing)
	if pullErr := readBuffer.PullContext("ackStateComing"); pullErr != nil {
		return nil, pullErr
	}
	_ackStateComing, _ackStateComingErr := StateParse(readBuffer)
	if _ackStateComingErr != nil {
		return nil, errors.Wrap(_ackStateComingErr, "Error parsing 'ackStateComing' field")
	}
	ackStateComing := CastState(_ackStateComing)
	if closeErr := readBuffer.CloseContext("ackStateComing"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timeComing)
	if pullErr := readBuffer.PullContext("timeComing"); pullErr != nil {
		return nil, pullErr
	}
	_timeComing, _timeComingErr := DateAndTimeParse(readBuffer)
	if _timeComingErr != nil {
		return nil, errors.Wrap(_timeComingErr, "Error parsing 'timeComing' field")
	}
	timeComing := CastDateAndTime(_timeComing)
	if closeErr := readBuffer.CloseContext("timeComing"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (valueComing)
	if pullErr := readBuffer.PullContext("valueComing"); pullErr != nil {
		return nil, pullErr
	}
	_valueComing, _valueComingErr := AssociatedValueTypeParse(readBuffer)
	if _valueComingErr != nil {
		return nil, errors.Wrap(_valueComingErr, "Error parsing 'valueComing' field")
	}
	valueComing := CastAssociatedValueType(_valueComing)
	if closeErr := readBuffer.CloseContext("valueComing"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timeGoing)
	if pullErr := readBuffer.PullContext("timeGoing"); pullErr != nil {
		return nil, pullErr
	}
	_timeGoing, _timeGoingErr := DateAndTimeParse(readBuffer)
	if _timeGoingErr != nil {
		return nil, errors.Wrap(_timeGoingErr, "Error parsing 'timeGoing' field")
	}
	timeGoing := CastDateAndTime(_timeGoing)
	if closeErr := readBuffer.CloseContext("timeGoing"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (valueGoing)
	if pullErr := readBuffer.PullContext("valueGoing"); pullErr != nil {
		return nil, pullErr
	}
	_valueGoing, _valueGoingErr := AssociatedValueTypeParse(readBuffer)
	if _valueGoingErr != nil {
		return nil, errors.Wrap(_valueGoingErr, "Error parsing 'valueGoing' field")
	}
	valueGoing := CastAssociatedValueType(_valueGoing)
	if closeErr := readBuffer.CloseContext("valueGoing"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageObjectQueryType"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAlarmMessageObjectQueryType(lengthDataset, eventState, ackStateGoing, ackStateComing, timeComing, valueComing, timeGoing, valueGoing), nil
}

func (m *AlarmMessageObjectQueryType) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AlarmMessageObjectQueryType"); pushErr != nil {
		return pushErr
	}

	// Simple Field (lengthDataset)
	lengthDataset := uint8(m.LengthDataset)
	_lengthDatasetErr := writeBuffer.WriteUint8("lengthDataset", 8, (lengthDataset))
	if _lengthDatasetErr != nil {
		return errors.Wrap(_lengthDatasetErr, "Error serializing 'lengthDataset' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Const Field (variableSpec)
	_variableSpecErr := writeBuffer.WriteUint8("variableSpec", 8, 0x12)
	if _variableSpecErr != nil {
		return errors.Wrap(_variableSpecErr, "Error serializing 'variableSpec' field")
	}

	// Simple Field (eventState)
	if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
		return pushErr
	}
	_eventStateErr := m.EventState.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
		return popErr
	}
	if _eventStateErr != nil {
		return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
	}

	// Simple Field (ackStateGoing)
	if pushErr := writeBuffer.PushContext("ackStateGoing"); pushErr != nil {
		return pushErr
	}
	_ackStateGoingErr := m.AckStateGoing.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("ackStateGoing"); popErr != nil {
		return popErr
	}
	if _ackStateGoingErr != nil {
		return errors.Wrap(_ackStateGoingErr, "Error serializing 'ackStateGoing' field")
	}

	// Simple Field (ackStateComing)
	if pushErr := writeBuffer.PushContext("ackStateComing"); pushErr != nil {
		return pushErr
	}
	_ackStateComingErr := m.AckStateComing.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("ackStateComing"); popErr != nil {
		return popErr
	}
	if _ackStateComingErr != nil {
		return errors.Wrap(_ackStateComingErr, "Error serializing 'ackStateComing' field")
	}

	// Simple Field (timeComing)
	if pushErr := writeBuffer.PushContext("timeComing"); pushErr != nil {
		return pushErr
	}
	_timeComingErr := m.TimeComing.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("timeComing"); popErr != nil {
		return popErr
	}
	if _timeComingErr != nil {
		return errors.Wrap(_timeComingErr, "Error serializing 'timeComing' field")
	}

	// Simple Field (valueComing)
	if pushErr := writeBuffer.PushContext("valueComing"); pushErr != nil {
		return pushErr
	}
	_valueComingErr := m.ValueComing.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("valueComing"); popErr != nil {
		return popErr
	}
	if _valueComingErr != nil {
		return errors.Wrap(_valueComingErr, "Error serializing 'valueComing' field")
	}

	// Simple Field (timeGoing)
	if pushErr := writeBuffer.PushContext("timeGoing"); pushErr != nil {
		return pushErr
	}
	_timeGoingErr := m.TimeGoing.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("timeGoing"); popErr != nil {
		return popErr
	}
	if _timeGoingErr != nil {
		return errors.Wrap(_timeGoingErr, "Error serializing 'timeGoing' field")
	}

	// Simple Field (valueGoing)
	if pushErr := writeBuffer.PushContext("valueGoing"); pushErr != nil {
		return pushErr
	}
	_valueGoingErr := m.ValueGoing.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("valueGoing"); popErr != nil {
		return popErr
	}
	if _valueGoingErr != nil {
		return errors.Wrap(_valueGoingErr, "Error serializing 'valueGoing' field")
	}

	if popErr := writeBuffer.PopContext("AlarmMessageObjectQueryType"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AlarmMessageObjectQueryType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
