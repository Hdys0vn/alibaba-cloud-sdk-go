package videorecog

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

// VideoQualityInfo is a nested struct in videorecog response
type VideoQualityInfo struct {
	CompressiveStrength float64 `json:"CompressiveStrength" xml:"CompressiveStrength"`
	NoiseIntensity      float64 `json:"NoiseIntensity" xml:"NoiseIntensity"`
	Blurriness          float64 `json:"Blurriness" xml:"Blurriness"`
	ColorContrast       float64 `json:"ColorContrast" xml:"ColorContrast"`
	ColorSaturation     float64 `json:"ColorSaturation" xml:"ColorSaturation"`
	Luminance           float64 `json:"Luminance" xml:"Luminance"`
	Colorfulness        float64 `json:"Colorfulness" xml:"Colorfulness"`
	MosScore            float64 `json:"MosScore" xml:"MosScore"`
}
