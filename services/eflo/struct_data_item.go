package eflo

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

// DataItem is a nested struct in eflo response
type DataItem struct {
	ErId                      string                         `json:"ErId" xml:"ErId"`
	RouteMaps                 int64                          `json:"RouteMaps" xml:"RouteMaps"`
	NetworkInterfaceCount     int                            `json:"NetworkInterfaceCount" xml:"NetworkInterfaceCount"`
	AccessPointId             string                         `json:"AccessPointId" xml:"AccessPointId"`
	ReceptionInstanceId       string                         `json:"ReceptionInstanceId" xml:"ReceptionInstanceId"`
	NcType                    string                         `json:"NcType" xml:"NcType"`
	Mac                       string                         `json:"Mac" xml:"Mac"`
	Across                    bool                           `json:"Across" xml:"Across"`
	MasterZoneId              string                         `json:"MasterZoneId" xml:"MasterZoneId"`
	Used                      bool                           `json:"Used" xml:"Used"`
	InstanceType              string                         `json:"InstanceType" xml:"InstanceType"`
	NetworkInterfaceName      string                         `json:"NetworkInterfaceName" xml:"NetworkInterfaceName"`
	NextHopId                 string                         `json:"NextHopId" xml:"NextHopId"`
	VccName                   string                         `json:"VccName" xml:"VccName"`
	TaskId                    string                         `json:"TaskId" xml:"TaskId"`
	ServiceCidr               string                         `json:"ServiceCidr" xml:"ServiceCidr"`
	ElasticNetworkInterfaceId string                         `json:"ElasticNetworkInterfaceId" xml:"ElasticNetworkInterfaceId"`
	Gateway                   string                         `json:"Gateway" xml:"Gateway"`
	CommodityCode             string                         `json:"CommodityCode" xml:"CommodityCode"`
	VpcId                     string                         `json:"VpcId" xml:"VpcId"`
	Type                      string                         `json:"Type" xml:"Type"`
	ReceptionInstanceName     string                         `json:"ReceptionInstanceName" xml:"ReceptionInstanceName"`
	Message                   string                         `json:"Message" xml:"Message"`
	Cidr                      string                         `json:"Cidr" xml:"Cidr"`
	IpAddressMac              string                         `json:"IpAddressMac" xml:"IpAddressMac"`
	VpdId                     string                         `json:"VpdId" xml:"VpdId"`
	ErAttachmentId            string                         `json:"ErAttachmentId" xml:"ErAttachmentId"`
	Action                    string                         `json:"Action" xml:"Action"`
	InstanceName              string                         `json:"InstanceName" xml:"InstanceName"`
	ReceptionInstanceOwner    string                         `json:"ReceptionInstanceOwner" xml:"ReceptionInstanceOwner"`
	VccRouteEntryId           string                         `json:"VccRouteEntryId" xml:"VccRouteEntryId"`
	CenOwnerId                string                         `json:"CenOwnerId" xml:"CenOwnerId"`
	IpName                    string                         `json:"IpName" xml:"IpName"`
	LineOperator              string                         `json:"LineOperator" xml:"LineOperator"`
	Spec                      string                         `json:"Spec" xml:"Spec"`
	Rate                      string                         `json:"Rate" xml:"Rate"`
	ErRouteEntryId            string                         `json:"ErRouteEntryId" xml:"ErRouteEntryId"`
	ExpirationDate            string                         `json:"ExpirationDate" xml:"ExpirationDate"`
	ZoneId                    string                         `json:"ZoneId" xml:"ZoneId"`
	TransmissionInstanceType  string                         `json:"TransmissionInstanceType" xml:"TransmissionInstanceType"`
	ErName                    string                         `json:"ErName" xml:"ErName"`
	AutoReceiveAllRoute       bool                           `json:"AutoReceiveAllRoute" xml:"AutoReceiveAllRoute"`
	Status                    string                         `json:"Status" xml:"Status"`
	ErRouteMapId              string                         `json:"ErRouteMapId" xml:"ErRouteMapId"`
	PrivateIpAddress          string                         `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
	NextHopType               string                         `json:"NextHopType" xml:"NextHopType"`
	ReceptionInstanceType     string                         `json:"ReceptionInstanceType" xml:"ReceptionInstanceType"`
	GrantRuleId               string                         `json:"GrantRuleId" xml:"GrantRuleId"`
	TransmissionInstanceOwner string                         `json:"TransmissionInstanceOwner" xml:"TransmissionInstanceOwner"`
	CreateTime                string                         `json:"CreateTime" xml:"CreateTime"`
	SubnetCount               int                            `json:"SubnetCount" xml:"SubnetCount"`
	PortType                  string                         `json:"PortType" xml:"PortType"`
	GmtCreate                 string                         `json:"GmtCreate" xml:"GmtCreate"`
	NodeId                    string                         `json:"NodeId" xml:"NodeId"`
	RegionId                  string                         `json:"RegionId" xml:"RegionId"`
	SubnetName                string                         `json:"SubnetName" xml:"SubnetName"`
	ResourceGroupId           string                         `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Mask                      string                         `json:"Mask" xml:"Mask"`
	InstanceId                string                         `json:"InstanceId" xml:"InstanceId"`
	SubnetId                  string                         `json:"SubnetId" xml:"SubnetId"`
	RouteMapNum               int                            `json:"RouteMapNum" xml:"RouteMapNum"`
	ResourceTenantId          string                         `json:"ResourceTenantId" xml:"ResourceTenantId"`
	Description               string                         `json:"Description" xml:"Description"`
	TenantId                  string                         `json:"TenantId" xml:"TenantId"`
	ErAttachmentName          string                         `json:"ErAttachmentName" xml:"ErAttachmentName"`
	Connections               int64                          `json:"Connections" xml:"Connections"`
	VpdName                   string                         `json:"VpdName" xml:"VpdName"`
	Product                   string                         `json:"Product" xml:"Product"`
	VpdRouteEntryId           string                         `json:"VpdRouteEntryId" xml:"VpdRouteEntryId"`
	Dependence                map[string]interface{}         `json:"Dependence" xml:"Dependence"`
	Ip                        string                         `json:"Ip" xml:"Ip"`
	VSwitchId                 string                         `json:"VSwitchId" xml:"VSwitchId"`
	GmtModified               string                         `json:"GmtModified" xml:"GmtModified"`
	Quota                     int                            `json:"Quota" xml:"Quota"`
	BgpCidr                   string                         `json:"BgpCidr" xml:"BgpCidr"`
	BandwidthStr              string                         `json:"BandwidthStr" xml:"BandwidthStr"`
	CenId                     string                         `json:"CenId" xml:"CenId"`
	ConnectionType            string                         `json:"ConnectionType" xml:"ConnectionType"`
	NcCount                   int                            `json:"NcCount" xml:"NcCount"`
	CurrentNode               string                         `json:"CurrentNode" xml:"CurrentNode"`
	TransmissionInstanceName  string                         `json:"TransmissionInstanceName" xml:"TransmissionInstanceName"`
	VccId                     string                         `json:"VccId" xml:"VccId"`
	RouteType                 string                         `json:"RouteType" xml:"RouteType"`
	NetworkInterfaceId        string                         `json:"NetworkInterfaceId" xml:"NetworkInterfaceId"`
	GrantTenantId             string                         `json:"GrantTenantId" xml:"GrantTenantId"`
	TransmissionInstanceId    string                         `json:"TransmissionInstanceId" xml:"TransmissionInstanceId"`
	ServiceMac                string                         `json:"ServiceMac" xml:"ServiceMac"`
	DestinationCidrBlock      string                         `json:"DestinationCidrBlock" xml:"DestinationCidrBlock"`
	SecondaryCidrBlocks       []string                       `json:"SecondaryCidrBlocks" xml:"SecondaryCidrBlocks"`
	Ethernet                  []string                       `json:"Ethernet" xml:"Ethernet"`
	VpdBaseInfo               VpdBaseInfo                    `json:"VpdBaseInfo" xml:"VpdBaseInfo"`
	SubnetBaseInfo            SubnetBaseInfo                 `json:"SubnetBaseInfo" xml:"SubnetBaseInfo"`
	ErInfos                   []ErInfo                       `json:"ErInfos" xml:"ErInfos"`
	Tags                      []Tag                          `json:"Tags" xml:"Tags"`
	PrivateIpAddressMacGroup  []PrivateIpAddressMacGroupItem `json:"PrivateIpAddressMacGroup" xml:"PrivateIpAddressMacGroup"`
}