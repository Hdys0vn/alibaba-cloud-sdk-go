package trademark

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

// TmProduces is a nested struct in trademark response
type TmProduces struct {
	BizId               string                                          `json:"BizId" xml:"BizId"`
	OrderId             string                                          `json:"OrderId" xml:"OrderId"`
	MaterialName        string                                          `json:"MaterialName" xml:"MaterialName"`
	TmIcon              string                                          `json:"TmIcon" xml:"TmIcon"`
	TmName              string                                          `json:"TmName" xml:"TmName"`
	TmNumber            string                                          `json:"TmNumber" xml:"TmNumber"`
	CreateTime          int64                                           `json:"CreateTime" xml:"CreateTime"`
	Type                int                                             `json:"Type" xml:"Type"`
	Status              int                                             `json:"Status" xml:"Status"`
	UserId              string                                          `json:"UserId" xml:"UserId"`
	OrderPrice          float64                                         `json:"OrderPrice" xml:"OrderPrice"`
	MaterialId          int64                                           `json:"MaterialId" xml:"MaterialId"`
	LoaUrl              string                                          `json:"LoaUrl" xml:"LoaUrl"`
	Note                string                                          `json:"Note" xml:"Note"`
	UpdateTime          int64                                           `json:"UpdateTime" xml:"UpdateTime"`
	SupplementStatus    int                                             `json:"SupplementStatus" xml:"SupplementStatus"`
	SupplementId        int64                                           `json:"SupplementId" xml:"SupplementId"`
	Flags               Flags                                           `json:"Flags" xml:"Flags"`
	FirstClassification FirstClassification                             `json:"FirstClassification" xml:"FirstClassification"`
	RenewResponse       RenewResponse                                   `json:"RenewResponse" xml:"RenewResponse"`
	ThirdClassification ThirdClassificationInQueryTradeMarkApplications `json:"ThirdClassification" xml:"ThirdClassification"`
}
