package imageprocess

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

// Lesion is a nested struct in imageprocess response
type Lesion struct {
	GCVolume           string             `json:"GCVolume" xml:"GCVolume"`
	Stage2Volume       string             `json:"Stage2Volume" xml:"Stage2Volume"`
	LiverVolume        float64            `json:"LiverVolume" xml:"LiverVolume"`
	Probabilities      string             `json:"Probabilities" xml:"Probabilities"`
	NonGCVolume        string             `json:"NonGCVolume" xml:"NonGCVolume"`
	NonCRCVolumel      string             `json:"NonCRCVolumel" xml:"NonCRCVolumel"`
	ColorectumVolume   string             `json:"ColorectumVolume" xml:"ColorectumVolume"`
	EcVolume           string             `json:"EcVolume" xml:"EcVolume"`
	StomachVolume      string             `json:"StomachVolume" xml:"StomachVolume"`
	Mask               string             `json:"Mask" xml:"Mask"`
	BenignVolume       string             `json:"BenignVolume" xml:"BenignVolume"`
	PdacVol            string             `json:"PdacVol" xml:"PdacVol"`
	CRCVolume          string             `json:"CRCVolume" xml:"CRCVolume"`
	CVDProbability     float64            `json:"CVDProbability" xml:"CVDProbability"`
	EsoVolume          string             `json:"EsoVolume" xml:"EsoVolume"`
	PancVol            string             `json:"PancVol" xml:"PancVol"`
	NonPdacVol         string             `json:"NonPdacVol" xml:"NonPdacVol"`
	Possibilities      []string           `json:"Possibilities" xml:"Possibilities"`
	ResultURL          []string           `json:"ResultURL" xml:"ResultURL"`
	PatientLevelResult PatientLevelResult `json:"PatientLevelResult" xml:"PatientLevelResult"`
	FeatureScore       FeatureScore       `json:"FeatureScore" xml:"FeatureScore"`
	LesionList         []LesionListItem   `json:"LesionList" xml:"LesionList"`
}
