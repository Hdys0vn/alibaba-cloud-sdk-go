package dms_enterprise

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

// Event is a nested struct in dms_enterprise response
type Event struct {
	EventLength    int64    `json:"EventLength" xml:"EventLength"`
	EventTimestamp string   `json:"EventTimestamp" xml:"EventTimestamp"`
	EventType      string   `json:"EventType" xml:"EventType"`
	EventId        int64    `json:"EventId" xml:"EventId"`
	RollSQL        string   `json:"RollSQL" xml:"RollSQL"`
	DataAfter      []string `json:"DataAfter" xml:"DataAfter"`
	DataBefore     []string `json:"DataBefore" xml:"DataBefore"`
}
