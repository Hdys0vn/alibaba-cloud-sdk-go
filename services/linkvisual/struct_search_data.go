package linkvisual

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

// SearchData is a nested struct in linkvisual response
type SearchData struct {
	VectorId       string  `json:"VectorId" xml:"VectorId"`
	EventTime      int64   `json:"EventTime" xml:"EventTime"`
	IotId          string  `json:"IotId" xml:"IotId"`
	NickName       string  `json:"NickName" xml:"NickName"`
	PicUrl         string  `json:"PicUrl" xml:"PicUrl"`
	VectorType     int     `json:"VectorType" xml:"VectorType"`
	GatewayIotId   string  `json:"GatewayIotId" xml:"GatewayIotId"`
	DeviceNickName string  `json:"DeviceNickName" xml:"DeviceNickName"`
	Threshold      float64 `json:"Threshold" xml:"Threshold"`
}