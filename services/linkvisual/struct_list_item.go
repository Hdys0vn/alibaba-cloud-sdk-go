package linkvisual

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

// ListItem is a nested struct in linkvisual response
type ListItem struct {
	Name             string                `json:"Name" xml:"Name"`
	TemplateId       string                `json:"TemplateId" xml:"TemplateId"`
	StreamType       int                   `json:"StreamType" xml:"StreamType"`
	ModifiedTime     string                `json:"ModifiedTime" xml:"ModifiedTime"`
	DeviceGroupId    string                `json:"DeviceGroupId" xml:"DeviceGroupId"`
	ControlType      string                `json:"ControlType" xml:"ControlType"`
	EventType        int                   `json:"EventType" xml:"EventType"`
	FileSize         int                   `json:"FileSize" xml:"FileSize"`
	UserId           string                `json:"UserId" xml:"UserId"`
	ThumbUrl         string                `json:"ThumbUrl" xml:"ThumbUrl"`
	RecordType       int                   `json:"RecordType" xml:"RecordType"`
	FileName         string                `json:"FileName" xml:"FileName"`
	EventData        string                `json:"EventData" xml:"EventData"`
	ControlId        string                `json:"ControlId" xml:"ControlId"`
	Params           string                `json:"Params" xml:"Params"`
	BeginTime        string                `json:"BeginTime" xml:"BeginTime"`
	CustomUserId     string                `json:"CustomUserId" xml:"CustomUserId"`
	EventDesc        string                `json:"EventDesc" xml:"EventDesc"`
	UserGroupId      string                `json:"UserGroupId" xml:"UserGroupId"`
	VideoFrameNumber int                   `json:"VideoFrameNumber" xml:"VideoFrameNumber"`
	IotId            string                `json:"IotId" xml:"IotId"`
	PicId            string                `json:"PicId" xml:"PicId"`
	EndTime          string                `json:"EndTime" xml:"EndTime"`
	AllDay           int                   `json:"AllDay" xml:"AllDay"`
	EventId          string                `json:"EventId" xml:"EventId"`
	EventPicId       string                `json:"EventPicId" xml:"EventPicId"`
	Default          int                   `json:"Default" xml:"Default"`
	PlanId           string                `json:"PlanId" xml:"PlanId"`
	EventTime        string                `json:"EventTime" xml:"EventTime"`
	PicUrl           string                `json:"PicUrl" xml:"PicUrl"`
	PicCreateTime    int64                 `json:"PicCreateTime" xml:"PicCreateTime"`
	SnapshotUrl      string                `json:"SnapshotUrl" xml:"SnapshotUrl"`
	TimeSectionList  []TimeSectionListItem `json:"TimeSectionList" xml:"TimeSectionList"`
}
