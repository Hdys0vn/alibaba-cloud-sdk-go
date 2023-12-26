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

// Download is a nested struct in arms response
type Download struct {
	DownloadTransmissionSize    int64  `json:"DownloadTransmissionSize" xml:"DownloadTransmissionSize"`
	DownloadCustomHost          int64  `json:"DownloadCustomHost" xml:"DownloadCustomHost"`
	DownloadCustomHostIp        string `json:"DownloadCustomHostIp" xml:"DownloadCustomHostIp"`
	WhiteList                   string `json:"WhiteList" xml:"WhiteList"`
	DownloadKernel              int64  `json:"DownloadKernel" xml:"DownloadKernel"`
	QuickProtocol               string `json:"QuickProtocol" xml:"QuickProtocol"`
	MonitorTimeout              int64  `json:"MonitorTimeout" xml:"MonitorTimeout"`
	ConnectionTimeout           int64  `json:"ConnectionTimeout" xml:"ConnectionTimeout"`
	VerifyWay                   int64  `json:"VerifyWay" xml:"VerifyWay"`
	ValidateKeywords            string `json:"ValidateKeywords" xml:"ValidateKeywords"`
	DownloadRedirect            int64  `json:"DownloadRedirect" xml:"DownloadRedirect"`
	DownloadCustomHeaderContent string `json:"DownloadCustomHeaderContent" xml:"DownloadCustomHeaderContent"`
}