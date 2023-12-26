package ltl

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

// RelatedData is a nested struct in ltl response
type RelatedData struct {
	RelatedDataKey       string `json:"RelatedDataKey" xml:"RelatedDataKey"`
	RelatedPhaseDataHash string `json:"RelatedPhaseDataHash" xml:"RelatedPhaseDataHash"`
	RelatedDataSeq       string `json:"RelatedDataSeq" xml:"RelatedDataSeq"`
	RelatedPhaseName     string `json:"RelatedPhaseName" xml:"RelatedPhaseName"`
	RelatedPhaseId       string `json:"RelatedPhaseId" xml:"RelatedPhaseId"`
}