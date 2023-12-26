package rtc

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

// Rule is a nested struct in rtc response
type Rule struct {
	Status            string   `json:"Status" xml:"Status"`
	MediaEncode       int      `json:"MediaEncode" xml:"MediaEncode"`
	PlayDomain        string   `json:"PlayDomain" xml:"PlayDomain"`
	CallBack          string   `json:"CallBack" xml:"CallBack"`
	CreateTime        string   `json:"CreateTime" xml:"CreateTime"`
	RuleId            int64    `json:"RuleId" xml:"RuleId"`
	RuleName          string   `json:"RuleName" xml:"RuleName"`
	ChannelIds        []string `json:"ChannelIds" xml:"ChannelIds"`
	ChannelIdPrefixes []string `json:"ChannelIdPrefixes" xml:"ChannelIdPrefixes"`
}
