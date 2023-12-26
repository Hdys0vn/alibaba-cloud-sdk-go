package imageseg

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

// Element is a nested struct in imageseg response
type Element struct {
	Name     string                 `json:"Name" xml:"Name"`
	X        int                    `json:"X" xml:"X"`
	ClassUrl map[string]interface{} `json:"ClassUrl" xml:"ClassUrl"`
	ImageURL string                 `json:"ImageURL" xml:"ImageURL"`
	Y        int                    `json:"Y" xml:"Y"`
	Width    int                    `json:"Width" xml:"Width"`
	Height   int                    `json:"Height" xml:"Height"`
}
