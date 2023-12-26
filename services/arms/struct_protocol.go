package arms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Protocol is a nested struct in arms response
type Protocol struct {
	CharacterEncoding         int64          `json:"CharacterEncoding" xml:"CharacterEncoding"`
	VerifyWay                 int64          `json:"VerifyWay" xml:"VerifyWay"`
	VerifyContent             string         `json:"VerifyContent" xml:"VerifyContent"`
	ReceivedDataSize          int64          `json:"ReceivedDataSize" xml:"ReceivedDataSize"`
	ProtocolMonitorTimeout    int64          `json:"ProtocolMonitorTimeout" xml:"ProtocolMonitorTimeout"`
	ProtocolConnectionTimeout int64          `json:"ProtocolConnectionTimeout" xml:"ProtocolConnectionTimeout"`
	CustomHost                int64          `json:"CustomHost" xml:"CustomHost"`
	CustomHostIp              string         `json:"CustomHostIp" xml:"CustomHostIp"`
	RequestContent            RequestContent `json:"RequestContent" xml:"RequestContent"`
}