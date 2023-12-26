package iot

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

// Data is a nested struct in iot response
type Data struct {
	LatestDeploymentStatus    int                                              `json:"LatestDeploymentStatus" xml:"LatestDeploymentStatus"`
	SourceData                string                                           `json:"SourceData" xml:"SourceData"`
	ContainerConfig           string                                           `json:"ContainerConfig" xml:"ContainerConfig"`
	RoleName                  string                                           `json:"RoleName" xml:"RoleName"`
	RequestMethod             string                                           `json:"RequestMethod" xml:"RequestMethod"`
	DevEui                    string                                           `json:"DevEui" xml:"DevEui"`
	CoordinateSystem          int                                              `json:"CoordinateSystem" xml:"CoordinateSystem"`
	OssPreSignedAddress       string                                           `json:"OssPreSignedAddress" xml:"OssPreSignedAddress"`
	DeviceConnState           string                                           `json:"DeviceConnState" xml:"DeviceConnState"`
	FileId                    string                                           `json:"FileId" xml:"FileId"`
	TslUri                    string                                           `json:"TslUri" xml:"TslUri"`
	Code                      string                                           `json:"Code" xml:"Code"`
	CsvUrl                    string                                           `json:"CsvUrl" xml:"CsvUrl"`
	OssAddress                string                                           `json:"OssAddress" xml:"OssAddress"`
	FirmwareId                string                                           `json:"FirmwareId" xml:"FirmwareId"`
	ApplyId                   int64                                            `json:"ApplyId" xml:"ApplyId"`
	LongJobId                 string                                           `json:"LongJobId" xml:"LongJobId"`
	Host                      string                                           `json:"Host" xml:"Host"`
	PageCount                 int64                                            `json:"PageCount" xml:"PageCount"`
	FailSum                   int                                              `json:"FailSum" xml:"FailSum"`
	Count                     int64                                            `json:"Count" xml:"Count"`
	Size                      string                                           `json:"Size" xml:"Size"`
	Udi                       string                                           `json:"Udi" xml:"Udi"`
	ProductName               string                                           `json:"ProductName" xml:"ProductName"`
	Name                      string                                           `json:"Name" xml:"Name"`
	SuccessSum                int                                              `json:"SuccessSum" xml:"SuccessSum"`
	GmtCreateTimestamp        int64                                            `json:"GmtCreateTimestamp" xml:"GmtCreateTimestamp"`
	SpeechCode                string                                           `json:"SpeechCode" xml:"SpeechCode"`
	Key                       string                                           `json:"Key" xml:"Key"`
	GmtCreate                 string                                           `json:"GmtCreate" xml:"GmtCreate"`
	EnableSoundCode           bool                                             `json:"EnableSoundCode" xml:"EnableSoundCode"`
	SourceConnState           string                                           `json:"SourceConnState" xml:"SourceConnState"`
	Voice                     string                                           `json:"Voice" xml:"Voice"`
	DateFormat                string                                           `json:"DateFormat" xml:"DateFormat"`
	Volume                    int                                              `json:"Volume" xml:"Volume"`
	TenantId                  string                                           `json:"TenantId" xml:"TenantId"`
	JobId                     string                                           `json:"JobId" xml:"JobId"`
	Id                        int                                              `json:"Id" xml:"Id"`
	Ip                        string                                           `json:"Ip" xml:"Ip"`
	ScriptType                string                                           `json:"ScriptType" xml:"ScriptType"`
	PageNum                   int                                              `json:"PageNum" xml:"PageNum"`
	AvailableQuota            int                                              `json:"AvailableQuota" xml:"AvailableQuota"`
	EdgeVersion               string                                           `json:"EdgeVersion" xml:"EdgeVersion"`
	DriverId                  string                                           `json:"DriverId" xml:"DriverId"`
	Signature                 string                                           `json:"Signature" xml:"Signature"`
	Adcode                    int64                                            `json:"Adcode" xml:"Adcode"`
	DeviceCount               int                                              `json:"DeviceCount" xml:"DeviceCount"`
	ProtocolType              string                                           `json:"ProtocolType" xml:"ProtocolType"`
	AuthType                  string                                           `json:"AuthType" xml:"AuthType"`
	SourceAccessToken         string                                           `json:"SourceAccessToken" xml:"SourceAccessToken"`
	GmtCompleted              string                                           `json:"GmtCompleted" xml:"GmtCompleted"`
	ProgressId                string                                           `json:"ProgressId" xml:"ProgressId"`
	UtcCreate                 string                                           `json:"UtcCreate" xml:"UtcCreate"`
	TargetType                string                                           `json:"TargetType" xml:"TargetType"`
	BizType                   string                                           `json:"BizType" xml:"BizType"`
	Password                  string                                           `json:"Password" xml:"Password"`
	GmtCompletedTimestamp     int64                                            `json:"GmtCompletedTimestamp" xml:"GmtCompletedTimestamp"`
	PreviewSize               int                                              `json:"PreviewSize" xml:"PreviewSize"`
	RoleArn                   string                                           `json:"RoleArn" xml:"RoleArn"`
	Tags                      string                                           `json:"Tags" xml:"Tags"`
	Configuration             string                                           `json:"Configuration" xml:"Configuration"`
	ExpiredQuota              int                                              `json:"ExpiredQuota" xml:"ExpiredQuota"`
	Longitude                 float64                                          `json:"Longitude" xml:"Longitude"`
	FailedResultCsvFile       string                                           `json:"FailedResultCsvFile" xml:"FailedResultCsvFile"`
	Province                  string                                           `json:"Province" xml:"Province"`
	BeginTime                 int64                                            `json:"BeginTime" xml:"BeginTime"`
	Id2                       bool                                             `json:"Id2" xml:"Id2"`
	NodeType                  int                                              `json:"NodeType" xml:"NodeType"`
	ConfigCheckRule           string                                           `json:"ConfigCheckRule" xml:"ConfigCheckRule"`
	TslStr                    string                                           `json:"TslStr" xml:"TslStr"`
	ApiSrn                    string                                           `json:"ApiSrn" xml:"ApiSrn"`
	ScriptUrl                 string                                           `json:"ScriptUrl" xml:"ScriptUrl"`
	OSSAccessKeyId            string                                           `json:"OSSAccessKeyId" xml:"OSSAccessKeyId"`
	Text                      string                                           `json:"Text" xml:"Text"`
	DynamicGroupExpression    string                                           `json:"DynamicGroupExpression" xml:"DynamicGroupExpression"`
	GroupName                 string                                           `json:"GroupName" xml:"GroupName"`
	CreateTime                int64                                            `json:"CreateTime" xml:"CreateTime"`
	FirmwareUrl               string                                           `json:"FirmwareUrl" xml:"FirmwareUrl"`
	RoleAttachTimestamp       int64                                            `json:"RoleAttachTimestamp" xml:"RoleAttachTimestamp"`
	GmtOpened                 int64                                            `json:"GmtOpened" xml:"GmtOpened"`
	Description               string                                           `json:"Description" xml:"Description"`
	Sn                        string                                           `json:"Sn" xml:"Sn"`
	ApiPath                   string                                           `json:"ApiPath" xml:"ApiPath"`
	Country                   string                                           `json:"Country" xml:"Country"`
	Protocol                  string                                           `json:"Protocol" xml:"Protocol"`
	TopicFilter               string                                           `json:"TopicFilter" xml:"TopicFilter"`
	ProductKey                string                                           `json:"ProductKey" xml:"ProductKey"`
	City                      string                                           `json:"City" xml:"City"`
	DisplayName               string                                           `json:"DisplayName" xml:"DisplayName"`
	IotId                     string                                           `json:"IotId" xml:"IotId"`
	GroupDesc                 string                                           `json:"GroupDesc" xml:"GroupDesc"`
	SpeechRate                int                                              `json:"SpeechRate" xml:"SpeechRate"`
	IsEnable                  string                                           `json:"IsEnable" xml:"IsEnable"`
	RoleAttachTime            string                                           `json:"RoleAttachTime" xml:"RoleAttachTime"`
	RequestProtocol           string                                           `json:"RequestProtocol" xml:"RequestProtocol"`
	Spec                      int                                              `json:"Spec" xml:"Spec"`
	Nickname                  string                                           `json:"Nickname" xml:"Nickname"`
	ProjectId                 string                                           `json:"ProjectId" xml:"ProjectId"`
	IsBeian                   string                                           `json:"IsBeian" xml:"IsBeian"`
	GroupId                   string                                           `json:"GroupId" xml:"GroupId"`
	TunnelId                  string                                           `json:"TunnelId" xml:"TunnelId"`
	LatestDeploymentType      string                                           `json:"LatestDeploymentType" xml:"LatestDeploymentType"`
	Type                      string                                           `json:"Type" xml:"Type"`
	ThingModelJson            string                                           `json:"ThingModelJson" xml:"ThingModelJson"`
	Versions                  string                                           `json:"Versions" xml:"Versions"`
	LastUpdateTime            int64                                            `json:"LastUpdateTime" xml:"LastUpdateTime"`
	TotalCount                int64                                            `json:"TotalCount" xml:"TotalCount"`
	AliyunCommodityCode       string                                           `json:"AliyunCommodityCode" xml:"AliyunCommodityCode"`
	OssAccessKeyId            string                                           `json:"OssAccessKeyId" xml:"OssAccessKeyId"`
	UtcCreated                string                                           `json:"UtcCreated" xml:"UtcCreated"`
	DeviceName                string                                           `json:"DeviceName" xml:"DeviceName"`
	CsvFileName               string                                           `json:"CsvFileName" xml:"CsvFileName"`
	IsOpen                    bool                                             `json:"IsOpen" xml:"IsOpen"`
	SourceURI                 string                                           `json:"SourceURI" xml:"SourceURI"`
	TotalSize                 int                                              `json:"TotalSize" xml:"TotalSize"`
	DatasetId                 string                                           `json:"DatasetId" xml:"DatasetId"`
	SpeechType                string                                           `json:"SpeechType" xml:"SpeechType"`
	DownloadUrl               string                                           `json:"DownloadUrl" xml:"DownloadUrl"`
	PageSize                  int                                              `json:"PageSize" xml:"PageSize"`
	ResultDataInString        string                                           `json:"ResultDataInString" xml:"ResultDataInString"`
	SourceType                string                                           `json:"SourceType" xml:"SourceType"`
	ResultJson                string                                           `json:"ResultJson" xml:"ResultJson"`
	InstanceId                string                                           `json:"InstanceId" xml:"InstanceId"`
	UtcClosed                 string                                           `json:"UtcClosed" xml:"UtcClosed"`
	Policy                    string                                           `json:"Policy" xml:"Policy"`
	DeviceOnline              int                                              `json:"DeviceOnline" xml:"DeviceOnline"`
	ProductSecret             string                                           `json:"ProductSecret" xml:"ProductSecret"`
	AccessKeyId               string                                           `json:"AccessKeyId" xml:"AccessKeyId"`
	JoinEui                   string                                           `json:"JoinEui" xml:"JoinEui"`
	CurrentPage               int                                              `json:"CurrentPage" xml:"CurrentPage"`
	ObjectStorage             string                                           `json:"ObjectStorage" xml:"ObjectStorage"`
	Total                     int64                                            `json:"Total" xml:"Total"`
	DataFormat                int                                              `json:"DataFormat" xml:"DataFormat"`
	BizEnable                 bool                                             `json:"BizEnable" xml:"BizEnable"`
	Latitude                  float64                                          `json:"Latitude" xml:"Latitude"`
	EndTime                   int64                                            `json:"EndTime" xml:"EndTime"`
	HasNext                   bool                                             `json:"HasNext" xml:"HasNext"`
	DeviceActive              int                                              `json:"DeviceActive" xml:"DeviceActive"`
	DriverVersion             string                                           `json:"DriverVersion" xml:"DriverVersion"`
	PageNo                    int                                              `json:"PageNo" xml:"PageNo"`
	Progress                  int                                              `json:"Progress" xml:"Progress"`
	DriverConfig              string                                           `json:"DriverConfig" xml:"DriverConfig"`
	BizCode                   string                                           `json:"BizCode" xml:"BizCode"`
	Token                     string                                           `json:"Token" xml:"Token"`
	AuthMode                  int                                              `json:"AuthMode" xml:"AuthMode"`
	UtcCreatedOn              string                                           `json:"UtcCreatedOn" xml:"UtcCreatedOn"`
	AppId                     string                                           `json:"AppId" xml:"AppId"`
	MessageId                 string                                           `json:"MessageId" xml:"MessageId"`
	FileUrl                   string                                           `json:"FileUrl" xml:"FileUrl"`
	ShareId                   string                                           `json:"ShareId" xml:"ShareId"`
	AudioFormat               string                                           `json:"AudioFormat" xml:"AudioFormat"`
	TargetData                string                                           `json:"TargetData" xml:"TargetData"`
	SourceConfig              string                                           `json:"SourceConfig" xml:"SourceConfig"`
	DeploymentId              string                                           `json:"DeploymentId" xml:"DeploymentId"`
	BizId                     string                                           `json:"BizId" xml:"BizId"`
	GmtModifiedTimestamp      int64                                            `json:"GmtModifiedTimestamp" xml:"GmtModifiedTimestamp"`
	Status                    int                                              `json:"Status" xml:"Status"`
	DeviceSecret              string                                           `json:"DeviceSecret" xml:"DeviceSecret"`
	AsyncExecute              bool                                             `json:"AsyncExecute" xml:"AsyncExecute"`
	GmtModified               string                                           `json:"GmtModified" xml:"GmtModified"`
	AuditResult               int                                              `json:"AuditResult" xml:"AuditResult"`
	ModifiedTime              int64                                            `json:"ModifiedTime" xml:"ModifiedTime"`
	VersionState              string                                           `json:"VersionState" xml:"VersionState"`
	Argument                  string                                           `json:"Argument" xml:"Argument"`
	ExpiringQuota             int                                              `json:"ExpiringQuota" xml:"ExpiringQuota"`
	TunnelState               string                                           `json:"TunnelState" xml:"TunnelState"`
	InvalidDeviceNameList     InvalidDeviceNameListInBatchCheckDeviceNames     `json:"InvalidDeviceNameList" xml:"InvalidDeviceNameList"`
	InvalidManufacturerList   InvalidManufacturerListInBatchCheckVehicleDevice `json:"InvalidManufacturerList" xml:"InvalidManufacturerList"`
	RepeatedDeviceNameList    RepeatedDeviceNameListInBatchCheckDeviceNames    `json:"RepeatedDeviceNameList" xml:"RepeatedDeviceNameList"`
	ResultList                ResultList                                       `json:"ResultList" xml:"ResultList"`
	InvalidDeviceModelList    InvalidDeviceModelListInBatchCheckVehicleDevice  `json:"InvalidDeviceModelList" xml:"InvalidDeviceModelList"`
	FieldNameList             FieldNameList                                    `json:"FieldNameList" xml:"FieldNameList"`
	InvalidDeviceSecretList   InvalidDeviceSecretListInBatchCheckImportDevice  `json:"InvalidDeviceSecretList" xml:"InvalidDeviceSecretList"`
	InvalidSnList             InvalidSnListInBatchCheckImportDevice            `json:"InvalidSnList" xml:"InvalidSnList"`
	RepeatedDeviceIdList      RepeatedDeviceIdListInBatchCheckVehicleDevice    `json:"RepeatedDeviceIdList" xml:"RepeatedDeviceIdList"`
	Result                    Result                                           `json:"Result" xml:"Result"`
	InvalidDeviceNicknameList InvalidDeviceNicknameList                        `json:"InvalidDeviceNicknameList" xml:"InvalidDeviceNicknameList"`
	InvalidDeviceIdList       InvalidDeviceIdListInBatchCheckVehicleDevice     `json:"InvalidDeviceIdList" xml:"InvalidDeviceIdList"`
	RouteContext              RouteContext                                     `json:"RouteContext" xml:"RouteContext"`
	QuerySetting              QuerySetting                                     `json:"QuerySetting" xml:"QuerySetting"`
	TokenInfo                 TokenInfo                                        `json:"TokenInfo" xml:"TokenInfo"`
	SqlTemplateDTO            SqlTemplateDTO                                   `json:"SqlTemplateDTO" xml:"SqlTemplateDTO"`
	JtProtocolDeviceData      JtProtocolDeviceData                             `json:"JtProtocolDeviceData" xml:"JtProtocolDeviceData"`
	SoundCodeConfig           SoundCodeConfig                                  `json:"SoundCodeConfig" xml:"SoundCodeConfig"`
	Header                    []HeaderItem                                     `json:"Header" xml:"Header"`
	TaskList                  []Task                                           `json:"TaskList" xml:"TaskList"`
	Points                    []PointsItem                                     `json:"Points" xml:"Points"`
	InvalidDetailList         []InvalidDetailListItem                          `json:"InvalidDetailList" xml:"InvalidDetailList"`
	List                      ListInGetThingTopo                               `json:"List" xml:"List"`
	DynamicRegClientIds       []DynamicRegClientId                             `json:"DynamicRegClientIds" xml:"DynamicRegClientIds"`
	FailDeviceSimpleInfoList  FailDeviceSimpleInfoList                         `json:"FailDeviceSimpleInfoList" xml:"FailDeviceSimpleInfoList"`
	ModelVersions             []ModelVersion                                   `json:"ModelVersions" xml:"ModelVersions"`
	Details                   Details                                          `json:"Details" xml:"Details"`
}