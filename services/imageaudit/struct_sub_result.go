package imageaudit

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

// SubResult is a nested struct in imageaudit response
type SubResult struct {
	Suggestion          string            `json:"Suggestion" xml:"Suggestion"`
	Label               string            `json:"Label" xml:"Label"`
	Scene               string            `json:"Scene" xml:"Scene"`
	Rate                float64           `json:"Rate" xml:"Rate"`
	OCRDataList         []string          `json:"OCRDataList" xml:"OCRDataList"`
	SfaceDataList       []SfaceData       `json:"SfaceDataList" xml:"SfaceDataList"`
	HintWordsInfoList   []HintWordsInfo   `json:"HintWordsInfoList" xml:"HintWordsInfoList"`
	ProgramCodeDataList []ProgramCodeData `json:"ProgramCodeDataList" xml:"ProgramCodeDataList"`
	Frames              []Frame           `json:"Frames" xml:"Frames"`
	LogoDataList        []LogoData        `json:"LogoDataList" xml:"LogoDataList"`
}
