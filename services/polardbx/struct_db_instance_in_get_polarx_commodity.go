package polardbx

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

// DBInstanceInGetPolarxCommodity is a nested struct in polardbx response
type DBInstanceInGetPolarxCommodity struct {
	Status             string     `json:"Status" xml:"Status"`
	Description        string     `json:"Description" xml:"Description"`
	ZoneId             string     `json:"ZoneId" xml:"ZoneId"`
	VPCId              string     `json:"VPCId" xml:"VPCId"`
	CreateTime         string     `json:"CreateTime" xml:"CreateTime"`
	Expired            string     `json:"Expired" xml:"Expired"`
	PayType            string     `json:"PayType" xml:"PayType"`
	DBType             string     `json:"DBType" xml:"DBType"`
	LockMode           string     `json:"LockMode" xml:"LockMode"`
	StorageUsed        int64      `json:"StorageUsed" xml:"StorageUsed"`
	DBVersion          string     `json:"DBVersion" xml:"DBVersion"`
	Network            string     `json:"Network" xml:"Network"`
	RegionId           string     `json:"RegionId" xml:"RegionId"`
	Engine             string     `json:"Engine" xml:"Engine"`
	Id                 string     `json:"Id" xml:"Id"`
	ConnectionString   string     `json:"ConnectionString" xml:"ConnectionString"`
	Port               string     `json:"Port" xml:"Port"`
	MinorVersion       string     `json:"MinorVersion" xml:"MinorVersion"`
	LatestMinorVersion string     `json:"LatestMinorVersion" xml:"LatestMinorVersion"`
	DBNodeCount        int        `json:"DBNodeCount" xml:"DBNodeCount"`
	DBInstanceType     string     `json:"DBInstanceType" xml:"DBInstanceType"`
	MaintainStartTime  string     `json:"MaintainStartTime" xml:"MaintainStartTime"`
	MaintainEndTime    string     `json:"MaintainEndTime" xml:"MaintainEndTime"`
	VSwitchId          string     `json:"VSwitchId" xml:"VSwitchId"`
	CommodityCode      string     `json:"CommodityCode" xml:"CommodityCode"`
	ExpireDate         string     `json:"ExpireDate" xml:"ExpireDate"`
	Type               string     `json:"Type" xml:"Type"`
	DBNodeClass        string     `json:"DBNodeClass" xml:"DBNodeClass"`
	ReadDBInstances    []string   `json:"ReadDBInstances" xml:"ReadDBInstances"`
	DBNodes            []DBNode   `json:"DBNodes" xml:"DBNodes"`
	ConnAddrs          []ConnAddr `json:"ConnAddrs" xml:"ConnAddrs"`
}
