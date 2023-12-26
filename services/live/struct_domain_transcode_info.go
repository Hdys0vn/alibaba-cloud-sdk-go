package live

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

// DomainTranscodeInfo is a nested struct in live response
type DomainTranscodeInfo struct {
	TranscodeApp              string                    `json:"TranscodeApp" xml:"TranscodeApp"`
	TranscodeTemplate         string                    `json:"TranscodeTemplate" xml:"TranscodeTemplate"`
	IsLazy                    bool                      `json:"IsLazy" xml:"IsLazy"`
	TranscodeName             string                    `json:"TranscodeName" xml:"TranscodeName"`
	CustomTranscodeParameters CustomTranscodeParameters `json:"CustomTranscodeParameters" xml:"CustomTranscodeParameters"`
	EncryptParameters         EncryptParameters         `json:"EncryptParameters" xml:"EncryptParameters"`
}