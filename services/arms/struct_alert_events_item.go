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

// AlertEventsItem is a nested struct in arms response
type AlertEventsItem struct {
	AlertName       string `json:"AlertName" xml:"AlertName"`
	Severity        string `json:"Severity" xml:"Severity"`
	State           string `json:"State" xml:"State"`
	StartTime       string `json:"StartTime" xml:"StartTime"`
	EndTime         string `json:"EndTime" xml:"EndTime"`
	ReceiveTime     string `json:"ReceiveTime" xml:"ReceiveTime"`
	IntegrationName string `json:"IntegrationName" xml:"IntegrationName"`
	IntegrationType string `json:"IntegrationType" xml:"IntegrationType"`
	GeneratorURL    string `json:"GeneratorURL" xml:"GeneratorURL"`
	Description     string `json:"Description" xml:"Description"`
	Annotations     string `json:"Annotations" xml:"Annotations"`
	Labels          string `json:"Labels" xml:"Labels"`
}
