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

// FeatureScore is a nested struct in imageprocess response
type FeatureScore struct {
	MyoEpiRatio          []string `json:"MyoEpiRatio" xml:"MyoEpiRatio"`
	AscAoMaxDiam         []string `json:"AscAoMaxDiam" xml:"AscAoMaxDiam"`
	CoronaryCalciumVol   []string `json:"CoronaryCalciumVol" xml:"CoronaryCalciumVol"`
	EatVolume            []string `json:"EatVolume" xml:"EatVolume"`
	AortaCalciumScore    []string `json:"AortaCalciumScore" xml:"AortaCalciumScore"`
	CardioThoracicRatio  []string `json:"CardioThoracicRatio" xml:"CardioThoracicRatio"`
	EatHUMean            []string `json:"EatHUMean" xml:"EatHUMean"`
	EatHUSTD             []string `json:"EatHUSTD" xml:"EatHUSTD"`
	RightLungLowattRatio []string `json:"RightLungLowattRatio" xml:"RightLungLowattRatio"`
	AscendAortaLength    []string `json:"AscendAortaLength" xml:"AscendAortaLength"`
	LeftLungLowattRatio  []string `json:"LeftLungLowattRatio" xml:"LeftLungLowattRatio"`
	DeepFeature          []string `json:"DeepFeature" xml:"DeepFeature"`
	AortaCalciumVolume   []string `json:"AortaCalciumVolume" xml:"AortaCalciumVolume"`
	CoronaryCalciumScore []string `json:"CoronaryCalciumScore" xml:"CoronaryCalciumScore"`
}
