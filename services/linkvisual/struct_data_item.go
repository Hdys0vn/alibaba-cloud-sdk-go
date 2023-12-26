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

// DataItem is a nested struct in linkvisual response
type DataItem struct {
	Name               string            `json:"Name" xml:"Name"`
	CustomUserId       string            `json:"CustomUserId" xml:"CustomUserId"`
	ModifyTime         int64             `json:"ModifyTime" xml:"ModifyTime"`
	VodUrl             string            `json:"VodUrl" xml:"VodUrl"`
	Day                int               `json:"Day" xml:"Day"`
	BlurScore          float64           `json:"BlurScore" xml:"BlurScore"`
	CreateTime         int64             `json:"CreateTime" xml:"CreateTime"`
	Gender             int               `json:"Gender" xml:"Gender"`
	Age                int               `json:"Age" xml:"Age"`
	PoseScore          float64           `json:"PoseScore" xml:"PoseScore"`
	GoodForRecognition bool              `json:"GoodForRecognition" xml:"GoodForRecognition"`
	IotId              string            `json:"IotId" xml:"IotId"`
	EndTime            string            `json:"EndTime" xml:"EndTime"`
	GoodForLibrary     bool              `json:"GoodForLibrary" xml:"GoodForLibrary"`
	OcclusionScore     float64           `json:"OcclusionScore" xml:"OcclusionScore"`
	UserId             string            `json:"UserId" xml:"UserId"`
	FaceProbability    float64           `json:"FaceProbability" xml:"FaceProbability"`
	FileName           string            `json:"FileName" xml:"FileName"`
	Params             string            `json:"Params" xml:"Params"`
	BeginTime          string            `json:"BeginTime" xml:"BeginTime"`
	FaceRects          []string          `json:"FaceRects" xml:"FaceRects"`
	Landmarks          []string          `json:"Landmarks" xml:"Landmarks"`
	FacePicList        []FacePicListItem `json:"FacePicList" xml:"FacePicList"`
}