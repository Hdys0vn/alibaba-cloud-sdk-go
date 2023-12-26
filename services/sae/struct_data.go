package sae

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

// Data is a nested struct in sae response
type Data struct {
	PythonModules                 string                   `json:"PythonModules" xml:"PythonModules"`
	LoadBalanceType               string                   `json:"LoadBalanceType" xml:"LoadBalanceType"`
	LastDisableTime               int64                    `json:"LastDisableTime" xml:"LastDisableTime"`
	VpcName                       string                   `json:"VpcName" xml:"VpcName"`
	Logo                          string                   `json:"Logo" xml:"Logo"`
	Python                        string                   `json:"Python" xml:"Python"`
	NasId                         string                   `json:"NasId" xml:"NasId"`
	CompletionTime                int64                    `json:"CompletionTime" xml:"CompletionTime"`
	PreStop                       string                   `json:"PreStop" xml:"PreStop"`
	NamespaceDescription          string                   `json:"NamespaceDescription" xml:"NamespaceDescription"`
	Message                       string                   `json:"Message" xml:"Message"`
	CertIds                       string                   `json:"CertIds" xml:"CertIds"`
	JarStartArgs                  string                   `json:"JarStartArgs" xml:"JarStartArgs"`
	InternetSlbId                 string                   `json:"InternetSlbId" xml:"InternetSlbId"`
	InternetSlbExpired            bool                     `json:"InternetSlbExpired" xml:"InternetSlbExpired"`
	RepoTag                       string                   `json:"RepoTag" xml:"RepoTag"`
	UpdateStrategy                string                   `json:"UpdateStrategy" xml:"UpdateStrategy"`
	Name                          string                   `json:"Name" xml:"Name"`
	IntranetSlbId                 string                   `json:"IntranetSlbId" xml:"IntranetSlbId"`
	Timeout                       int64                    `json:"Timeout" xml:"Timeout"`
	SecurityGroupId               string                   `json:"SecurityGroupId" xml:"SecurityGroupId"`
	Group                         string                   `json:"Group" xml:"Group"`
	PackageVersion                string                   `json:"PackageVersion" xml:"PackageVersion"`
	NameSpaceShortId              string                   `json:"NameSpaceShortId" xml:"NameSpaceShortId"`
	EnableGreyTagRoute            bool                     `json:"EnableGreyTagRoute" xml:"EnableGreyTagRoute"`
	Replicas                      int                      `json:"Replicas" xml:"Replicas"`
	PhpConfigLocation             string                   `json:"PhpConfigLocation" xml:"PhpConfigLocation"`
	CoStatus                      string                   `json:"CoStatus" xml:"CoStatus"`
	SecretData                    map[string]interface{}   `json:"SecretData" xml:"SecretData"`
	TenantId                      string                   `json:"TenantId" xml:"TenantId"`
	NamespaceName                 string                   `json:"NamespaceName" xml:"NamespaceName"`
	JobId                         string                   `json:"JobId" xml:"JobId"`
	VSwitchId                     string                   `json:"VSwitchId" xml:"VSwitchId"`
	PackageRuntimeCustomBuild     string                   `json:"PackageRuntimeCustomBuild" xml:"PackageRuntimeCustomBuild"`
	PostStart                     string                   `json:"PostStart" xml:"PostStart"`
	ListenerProtocol              string                   `json:"ListenerProtocol" xml:"ListenerProtocol"`
	RepoName                      string                   `json:"RepoName" xml:"RepoName"`
	EnableAhas                    string                   `json:"EnableAhas" xml:"EnableAhas"`
	BackoffLimit                  int64                    `json:"BackoffLimit" xml:"BackoffLimit"`
	State                         string                   `json:"State" xml:"State"`
	NextToken                     string                   `json:"NextToken" xml:"NextToken"`
	SpringApplicationName         string                   `json:"SpringApplicationName" xml:"SpringApplicationName"`
	PackageUrl                    string                   `json:"PackageUrl" xml:"PackageUrl"`
	PackageType                   string                   `json:"PackageType" xml:"PackageType"`
	LastChangeOrderRunning        bool                     `json:"LastChangeOrderRunning" xml:"LastChangeOrderRunning"`
	Failed                        int64                    `json:"Failed" xml:"Failed"`
	KafkaConfigs                  string                   `json:"KafkaConfigs" xml:"KafkaConfigs"`
	SlbType                       string                   `json:"SlbType" xml:"SlbType"`
	EdasAppName                   string                   `json:"EdasAppName" xml:"EdasAppName"`
	OssAkSecret                   string                   `json:"OssAkSecret" xml:"OssAkSecret"`
	LastChangeOrderStatus         string                   `json:"LastChangeOrderStatus" xml:"LastChangeOrderStatus"`
	VpcId                         string                   `json:"VpcId" xml:"VpcId"`
	CertId                        string                   `json:"CertId" xml:"CertId"`
	SecretType                    string                   `json:"SecretType" xml:"SecretType"`
	CrUrl                         string                   `json:"CrUrl" xml:"CrUrl"`
	EnableMicroRegistration       bool                     `json:"EnableMicroRegistration" xml:"EnableMicroRegistration"`
	Jdk                           string                   `json:"Jdk" xml:"Jdk"`
	Metadata                      map[string]interface{}   `json:"Metadata" xml:"Metadata"`
	Php                           string                   `json:"Php" xml:"Php"`
	MicroRegistration             string                   `json:"MicroRegistration" xml:"MicroRegistration"`
	JumpServerAppId               string                   `json:"JumpServerAppId" xml:"JumpServerAppId"`
	Cpu                           int                      `json:"Cpu" xml:"Cpu"`
	PipelineStatus                int                      `json:"PipelineStatus" xml:"PipelineStatus"`
	CreateTime                    int64                    `json:"CreateTime" xml:"CreateTime"`
	AppDescription                string                   `json:"AppDescription" xml:"AppDescription"`
	PhpPECLExtensions             string                   `json:"PhpPECLExtensions" xml:"PhpPECLExtensions"`
	TriggerConfig                 string                   `json:"TriggerConfig" xml:"TriggerConfig"`
	WarStartOptions               string                   `json:"WarStartOptions" xml:"WarStartOptions"`
	Workload                      string                   `json:"Workload" xml:"Workload"`
	ProgrammingLanguage           string                   `json:"ProgrammingLanguage" xml:"ProgrammingLanguage"`
	Description                   string                   `json:"Description" xml:"Description"`
	IsNeedApproval                bool                     `json:"IsNeedApproval" xml:"IsNeedApproval"`
	IngressId                     int64                    `json:"IngressId" xml:"IngressId"`
	NamespaceId                   string                   `json:"NamespaceId" xml:"NamespaceId"`
	Data                          map[string]interface{}   `json:"Data" xml:"Data"`
	SecretName                    string                   `json:"SecretName" xml:"SecretName"`
	IntranetSlbExpired            bool                     `json:"IntranetSlbExpired" xml:"IntranetSlbExpired"`
	TomcatConfig                  string                   `json:"TomcatConfig" xml:"TomcatConfig"`
	VSwitchName                   string                   `json:"VSwitchName" xml:"VSwitchName"`
	LastChangeOrderId             string                   `json:"LastChangeOrderId" xml:"LastChangeOrderId"`
	EdasContainerVersion          string                   `json:"EdasContainerVersion" xml:"EdasContainerVersion"`
	SlsConfigs                    string                   `json:"SlsConfigs" xml:"SlsConfigs"`
	Suspend                       bool                     `json:"Suspend" xml:"Suspend"`
	Slice                         bool                     `json:"Slice" xml:"Slice"`
	ListenerPort                  int                      `json:"ListenerPort" xml:"ListenerPort"`
	PipelineName                  string                   `json:"PipelineName" xml:"PipelineName"`
	MseApplicationId              string                   `json:"MseApplicationId" xml:"MseApplicationId"`
	UpdateTime                    int64                    `json:"UpdateTime" xml:"UpdateTime"`
	Id                            int64                    `json:"Id" xml:"Id"`
	GreyTagRouteId                int64                    `json:"GreyTagRouteId" xml:"GreyTagRouteId"`
	Command                       string                   `json:"Command" xml:"Command"`
	NasConfigs                    string                   `json:"NasConfigs" xml:"NasConfigs"`
	JarStartOptions               string                   `json:"JarStartOptions" xml:"JarStartOptions"`
	PhpExtensions                 string                   `json:"PhpExtensions" xml:"PhpExtensions"`
	PhpConfig                     string                   `json:"PhpConfig" xml:"PhpConfig"`
	Envs                          string                   `json:"Envs" xml:"Envs"`
	RepoType                      string                   `json:"RepoType" xml:"RepoType"`
	SlbId                         string                   `json:"SlbId" xml:"SlbId"`
	TotalSize                     int                      `json:"TotalSize" xml:"TotalSize"`
	NotificationExpired           bool                     `json:"NotificationExpired" xml:"NotificationExpired"`
	ChangeOrderId                 string                   `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Timezone                      string                   `json:"Timezone" xml:"Timezone"`
	OssAkId                       string                   `json:"OssAkId" xml:"OssAkId"`
	PageSize                      int                      `json:"PageSize" xml:"PageSize"`
	MinReadyInstanceRatio         int                      `json:"MinReadyInstanceRatio" xml:"MinReadyInstanceRatio"`
	PvtzDiscovery                 string                   `json:"PvtzDiscovery" xml:"PvtzDiscovery"`
	CustomHostAlias               string                   `json:"CustomHostAlias" xml:"CustomHostAlias"`
	EnableImageAccl               bool                     `json:"EnableImageAccl" xml:"EnableImageAccl"`
	NextPipelineId                string                   `json:"NextPipelineId" xml:"NextPipelineId"`
	Memory                        int                      `json:"Memory" xml:"Memory"`
	PhpArmsConfigLocation         string                   `json:"PhpArmsConfigLocation" xml:"PhpArmsConfigLocation"`
	ImagePullSecrets              string                   `json:"ImagePullSecrets" xml:"ImagePullSecrets"`
	Version                       string                   `json:"Version" xml:"Version"`
	RepoNamespace                 string                   `json:"RepoNamespace" xml:"RepoNamespace"`
	Liveness                      string                   `json:"Liveness" xml:"Liveness"`
	CurrentPage                   int                      `json:"CurrentPage" xml:"CurrentPage"`
	AcrInstanceId                 string                   `json:"AcrInstanceId" xml:"AcrInstanceId"`
	AppName                       string                   `json:"AppName" xml:"AppName"`
	DubboApplicationName          string                   `json:"DubboApplicationName" xml:"DubboApplicationName"`
	JumpServerIp                  string                   `json:"JumpServerIp" xml:"JumpServerIp"`
	RefAppId                      string                   `json:"RefAppId" xml:"RefAppId"`
	MountHost                     string                   `json:"MountHost" xml:"MountHost"`
	StartTime                     int64                    `json:"StartTime" xml:"StartTime"`
	CurrentPoint                  int                      `json:"CurrentPoint" xml:"CurrentPoint"`
	ImageUrl                      string                   `json:"ImageUrl" xml:"ImageUrl"`
	Active                        int64                    `json:"Active" xml:"Active"`
	IntranetIp                    string                   `json:"IntranetIp" xml:"IntranetIp"`
	ServiceType                   string                   `json:"ServiceType" xml:"ServiceType"`
	ServiceName                   string                   `json:"ServiceName" xml:"ServiceName"`
	ShowBatch                     bool                     `json:"ShowBatch" xml:"ShowBatch"`
	BatchWaitTime                 int                      `json:"BatchWaitTime" xml:"BatchWaitTime"`
	RepoId                        int                      `json:"RepoId" xml:"RepoId"`
	SliceEnvs                     string                   `json:"SliceEnvs" xml:"SliceEnvs"`
	BelongRegion                  string                   `json:"BelongRegion" xml:"BelongRegion"`
	AppId                         string                   `json:"AppId" xml:"AppId"`
	ScaleRuleName                 string                   `json:"ScaleRuleName" xml:"ScaleRuleName"`
	AcrAssumeRoleArn              string                   `json:"AcrAssumeRoleArn" xml:"AcrAssumeRoleArn"`
	Succeeded                     int64                    `json:"Succeeded" xml:"Succeeded"`
	TerminationGracePeriodSeconds int                      `json:"TerminationGracePeriodSeconds" xml:"TerminationGracePeriodSeconds"`
	WebContainer                  string                   `json:"WebContainer" xml:"WebContainer"`
	AssociateEip                  bool                     `json:"AssociateEip" xml:"AssociateEip"`
	RepoOriginType                string                   `json:"RepoOriginType" xml:"RepoOriginType"`
	MinReadyInstances             int                      `json:"MinReadyInstances" xml:"MinReadyInstances"`
	CommandArgs                   string                   `json:"CommandArgs" xml:"CommandArgs"`
	CurrentStageId                string                   `json:"CurrentStageId" xml:"CurrentStageId"`
	RegionId                      string                   `json:"RegionId" xml:"RegionId"`
	PipelineId                    string                   `json:"PipelineId" xml:"PipelineId"`
	UserId                        string                   `json:"UserId" xml:"UserId"`
	ScaleRuleEnabled              bool                     `json:"ScaleRuleEnabled" xml:"ScaleRuleEnabled"`
	InternetIp                    string                   `json:"InternetIp" xml:"InternetIp"`
	AppCount                      int64                    `json:"AppCount" xml:"AppCount"`
	ConfigMapId                   int64                    `json:"ConfigMapId" xml:"ConfigMapId"`
	ConcurrencyPolicy             string                   `json:"ConcurrencyPolicy" xml:"ConcurrencyPolicy"`
	Readiness                     string                   `json:"Readiness" xml:"Readiness"`
	ScaleRuleType                 string                   `json:"ScaleRuleType" xml:"ScaleRuleType"`
	SecretId                      int64                    `json:"SecretId" xml:"SecretId"`
	VpcWebHookUrls                []string                 `json:"VpcWebHookUrls" xml:"VpcWebHookUrls"`
	RefedAppIds                   []string                 `json:"RefedAppIds" xml:"RefedAppIds"`
	PublicWebHookUrls             []string                 `json:"PublicWebHookUrls" xml:"PublicWebHookUrls"`
	Order                         Order                    `json:"Order" xml:"Order"`
	RealTimeRes                   RealTimeRes              `json:"RealTimeRes" xml:"RealTimeRes"`
	DefaultRule                   DefaultRule              `json:"DefaultRule" xml:"DefaultRule"`
	Metric                        Metric                   `json:"Metric" xml:"Metric"`
	BagUsage                      BagUsage                 `json:"BagUsage" xml:"BagUsage"`
	Timer                         Timer                    `json:"Timer" xml:"Timer"`
	Summary                       Summary                  `json:"Summary" xml:"Summary"`
	OssMountDescs                 []OssMountDesc           `json:"OssMountDescs" xml:"OssMountDescs"`
	TagResources                  []TagResource            `json:"TagResources" xml:"TagResources"`
	Rules                         []Rule                   `json:"Rules" xml:"Rules"`
	Applications                  []Application            `json:"Applications" xml:"Applications"`
	Intranet                      []IntranetItem           `json:"Intranet" xml:"Intranet"`
	DubboRules                    []DubboRule              `json:"DubboRules" xml:"DubboRules"`
	Tags                          []Tag                    `json:"Tags" xml:"Tags"`
	Instances                     []Instance               `json:"Instances" xml:"Instances"`
	Methods                       []Method                 `json:"Methods" xml:"Methods"`
	Internet                      []InternetItem           `json:"Internet" xml:"Internet"`
	MountDesc                     []MountDescItem          `json:"MountDesc" xml:"MountDesc"`
	ConfigMaps                    []ConfigMap              `json:"ConfigMaps" xml:"ConfigMaps"`
	RelateApps                    []RelateApp              `json:"RelateApps" xml:"RelateApps"`
	LogConfigs                    []LogConfig              `json:"LogConfigs" xml:"LogConfigs"`
	ScRules                       []ScRule                 `json:"ScRules" xml:"ScRules"`
	ChangeOrderList               []ChangeOrder            `json:"ChangeOrderList" xml:"ChangeOrderList"`
	IngressList                   []Ingress                `json:"IngressList" xml:"IngressList"`
	Svcs                          []Svc                    `json:"Svcs" xml:"Svcs"`
	Secrets                       []Secret                 `json:"Secrets" xml:"Secrets"`
	ConfigMapMountDesc            []ConfigMapMountDescItem `json:"ConfigMapMountDesc" xml:"ConfigMapMountDesc"`
	AlbRules                      []AlbRule                `json:"AlbRules" xml:"AlbRules"`
	ApplicationScalingRules       []ApplicationScalingRule `json:"ApplicationScalingRules" xml:"ApplicationScalingRules"`
	Namespaces                    []Namespace              `json:"Namespaces" xml:"Namespaces"`
	AppEventEntity                []AppEventEntityItem     `json:"AppEventEntity" xml:"AppEventEntity"`
	StageList                     []Stage                  `json:"StageList" xml:"StageList"`
}