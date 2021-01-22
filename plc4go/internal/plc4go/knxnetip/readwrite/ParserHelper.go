//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package readwrite

import (
	"errors"
	"github.com/apache/plc4x/plc4go/internal/plc4go/knxnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type KnxnetipParserHelper struct {
}

func (m KnxnetipParserHelper) Parse(typeName string, arguments []string, io *utils.ReadBuffer) (spi.Message, error) {
	switch typeName {
	case "HPAIControlEndpoint":
		return model.HPAIControlEndpointParse(io)
	case "TunnelingResponseDataBlock":
		return model.TunnelingResponseDataBlockParse(io)
	case "DeviceConfigurationAckDataBlock":
		return model.DeviceConfigurationAckDataBlockParse(io)
	case "ConnectionRequestInformation":
		return model.ConnectionRequestInformationParse(io)
	case "HPAIDiscoveryEndpoint":
		return model.HPAIDiscoveryEndpointParse(io)
	case "ProjectInstallationIdentifier":
		return model.ProjectInstallationIdentifierParse(io)
	case "ServiceId":
		return model.ServiceIdParse(io)
	case "HPAIDataEndpoint":
		return model.HPAIDataEndpointParse(io)
	case "RelativeTimestamp":
		return model.RelativeTimestampParse(io)
	case "CEMI":
		size, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, err
		}
		return model.CEMIParse(io, size)
	case "KnxNetIpMessage":
		return model.KnxNetIpMessageParse(io)
	case "DeviceStatus":
		return model.DeviceStatusParse(io)
	case "IPAddress":
		return model.IPAddressParse(io)
	case "CEMIAdditionalInformation":
		return model.CEMIAdditionalInformationParse(io)
	case "KnxAddress":
		return model.KnxAddressParse(io)
	case "ConnectionResponseDataBlock":
		return model.ConnectionResponseDataBlockParse(io)
	case "TunnelingRequestDataBlock":
		return model.TunnelingRequestDataBlockParse(io)
	case "DIBDeviceInfo":
		return model.DIBDeviceInfoParse(io)
	case "DeviceConfigurationRequestDataBlock":
		return model.DeviceConfigurationRequestDataBlockParse(io)
	case "DIBSuppSvcFamilies":
		return model.DIBSuppSvcFamiliesParse(io)
	case "LDataFrame":
		return model.LDataFrameParse(io)
	case "KnxGroupAddress":
		numLevels, err := utils.StrToUint8(arguments[0])
		if err != nil {
			return nil, err
		}
		return model.KnxGroupAddressParse(io, numLevels)
	case "MACAddress":
		return model.MACAddressParse(io)
	}
	return nil, errors.New("Unsupported type " + typeName)
}
