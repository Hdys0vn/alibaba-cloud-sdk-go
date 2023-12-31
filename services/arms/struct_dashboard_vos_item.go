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

// DashboardVosItem is a nested struct in arms response
type DashboardVosItem struct {
	Type           string    `json:"Type" xml:"Type"`
	Time           string    `json:"Time" xml:"Time"`
	NeedUpdate     bool      `json:"NeedUpdate" xml:"NeedUpdate"`
	Kind           string    `json:"Kind" xml:"Kind"`
	Language       string    `json:"Language" xml:"Language"`
	Url            string    `json:"Url" xml:"Url"`
	HttpsUrl       string    `json:"HttpsUrl" xml:"HttpsUrl"`
	DashboardType  string    `json:"DashboardType" xml:"DashboardType"`
	Exporter       string    `json:"Exporter" xml:"Exporter"`
	Version        string    `json:"Version" xml:"Version"`
	IsArmsExporter bool      `json:"IsArmsExporter" xml:"IsArmsExporter"`
	HttpUrl        string    `json:"HttpUrl" xml:"HttpUrl"`
	Title          string    `json:"Title" xml:"Title"`
	Name           string    `json:"Name" xml:"Name"`
	Id             string    `json:"Id" xml:"Id"`
	Uid            string    `json:"Uid" xml:"Uid"`
	Tags           []string  `json:"Tags" xml:"Tags"`
	I18nChild      I18nChild `json:"I18nChild" xml:"I18nChild"`
}
