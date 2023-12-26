package edas

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

// Conf is a nested struct in edas response
type Conf struct {
	PostStart          string `json:"PostStart" xml:"PostStart"`
	Readiness          string `json:"Readiness" xml:"Readiness"`
	AhasEnabled        bool   `json:"AhasEnabled" xml:"AhasEnabled"`
	K8sCmdArgs         string `json:"K8sCmdArgs" xml:"K8sCmdArgs"`
	Liveness           string `json:"Liveness" xml:"Liveness"`
	DeployAcrossNodes  string `json:"DeployAcrossNodes" xml:"DeployAcrossNodes"`
	K8sCmd             string `json:"K8sCmd" xml:"K8sCmd"`
	PreStop            string `json:"PreStop" xml:"PreStop"`
	DeployAcrossZones  string `json:"DeployAcrossZones" xml:"DeployAcrossZones"`
	JarStartArgs       string `json:"JarStartArgs" xml:"JarStartArgs"`
	K8sNasInfo         string `json:"K8sNasInfo" xml:"K8sNasInfo"`
	JarStartOptions    string `json:"JarStartOptions" xml:"JarStartOptions"`
	RuntimeClassName   string `json:"RuntimeClassName" xml:"RuntimeClassName"`
	K8sLocalvolumeInfo string `json:"K8sLocalvolumeInfo" xml:"K8sLocalvolumeInfo"`
	K8sVolumeInfo      string `json:"K8sVolumeInfo" xml:"K8sVolumeInfo"`
	Affinity           string `json:"Affinity" xml:"Affinity"`
	Tolerations        string `json:"Tolerations" xml:"Tolerations"`
	UserBaseImageUrl   string `json:"UserBaseImageUrl" xml:"UserBaseImageUrl"`
}