package ecs

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

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// RunInstances invokes the ecs.RunInstances API synchronously
func (client *Client) RunInstances(request *RunInstancesRequest) (response *RunInstancesResponse, err error) {
	response = CreateRunInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// RunInstancesWithChan invokes the ecs.RunInstances API asynchronously
func (client *Client) RunInstancesWithChan(request *RunInstancesRequest) (<-chan *RunInstancesResponse, <-chan error) {
	responseChan := make(chan *RunInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunInstances(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// RunInstancesWithCallback invokes the ecs.RunInstances API asynchronously
func (client *Client) RunInstancesWithCallback(request *RunInstancesRequest, callback func(response *RunInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunInstancesResponse
		var err error
		defer close(result)
		response, err = client.RunInstances(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// RunInstancesRequest is the request struct for api RunInstances
type RunInstancesRequest struct {
	*requests.RpcRequest
	UniqueSuffix                             requests.Boolean                `position:"Query" name:"UniqueSuffix"`
	SecurityEnhancementStrategy              string                          `position:"Query" name:"SecurityEnhancementStrategy"`
	MinAmount                                requests.Integer                `position:"Query" name:"MinAmount"`
	DeletionProtection                       requests.Boolean                `position:"Query" name:"DeletionProtection"`
	ResourceGroupId                          string                          `position:"Query" name:"ResourceGroupId"`
	PrivatePoolOptionsMatchCriteria          string                          `position:"Query" name:"PrivatePoolOptions.MatchCriteria"`
	HostName                                 string                          `position:"Query" name:"HostName"`
	Password                                 string                          `position:"Query" name:"Password"`
	SystemDisk                               RunInstancesSystemDisk          `position:"Query" name:"SystemDisk"  type:"Struct"`
	ImageOptions                             RunInstancesImageOptions        `position:"Query" name:"ImageOptions"  type:"Struct"`
	DeploymentSetGroupNo                     requests.Integer                `position:"Query" name:"DeploymentSetGroupNo"`
	SystemDiskAutoSnapshotPolicyId           string                          `position:"Query" name:"SystemDisk.AutoSnapshotPolicyId"`
	CpuOptionsCore                           requests.Integer                `position:"Query" name:"CpuOptions.Core"`
	Period                                   requests.Integer                `position:"Query" name:"Period"`
	DryRun                                   requests.Boolean                `position:"Query" name:"DryRun"`
	CpuOptionsNuma                           string                          `position:"Query" name:"CpuOptions.Numa"`
	OwnerId                                  requests.Integer                `position:"Query" name:"OwnerId"`
	SpotStrategy                             string                          `position:"Query" name:"SpotStrategy"`
	PrivateIpAddress                         string                          `position:"Query" name:"PrivateIpAddress"`
	PeriodUnit                               string                          `position:"Query" name:"PeriodUnit"`
	AutoRenew                                requests.Boolean                `position:"Query" name:"AutoRenew"`
	InternetChargeType                       string                          `position:"Query" name:"InternetChargeType"`
	InternetMaxBandwidthIn                   requests.Integer                `position:"Query" name:"InternetMaxBandwidthIn"`
	Affinity                                 string                          `position:"Query" name:"Affinity"`
	ImageId                                  string                          `position:"Query" name:"ImageId"`
	SpotInterruptionBehavior                 string                          `position:"Query" name:"SpotInterruptionBehavior"`
	NetworkInterfaceQueueNumber              requests.Integer                `position:"Query" name:"NetworkInterfaceQueueNumber"`
	SystemTag                                *[]RunInstancesSystemTag        `position:"Query" name:"SystemTag"  type:"Repeated"`
	IoOptimized                              string                          `position:"Query" name:"IoOptimized"`
	SecurityGroupId                          string                          `position:"Query" name:"SecurityGroupId"`
	HibernationOptionsConfigured             requests.Boolean                `position:"Query" name:"HibernationOptions.Configured"`
	SystemDiskPerformanceLevel               string                          `position:"Query" name:"SystemDisk.PerformanceLevel"`
	PasswordInherit                          requests.Boolean                `position:"Query" name:"PasswordInherit"`
	InstanceType                             string                          `position:"Query" name:"InstanceType"`
	Arn                                      *[]RunInstancesArn              `position:"Query" name:"Arn"  type:"Repeated"`
	ResourceOwnerAccount                     string                          `position:"Query" name:"ResourceOwnerAccount"`
	SchedulerOptionsDedicatedHostClusterId   string                          `position:"Query" name:"SchedulerOptions.DedicatedHostClusterId"`
	SystemDiskDiskName                       string                          `position:"Query" name:"SystemDisk.DiskName"`
	DedicatedHostId                          string                          `position:"Query" name:"DedicatedHostId"`
	SpotDuration                             requests.Integer                `position:"Query" name:"SpotDuration"`
	SecurityGroupIds                         *[]string                       `position:"Query" name:"SecurityGroupIds"  type:"Repeated"`
	NetworkOptions                           RunInstancesNetworkOptions      `position:"Query" name:"NetworkOptions"  type:"Struct"`
	SystemDiskSize                           string                          `position:"Query" name:"SystemDisk.Size"`
	ImageFamily                              string                          `position:"Query" name:"ImageFamily"`
	LaunchTemplateName                       string                          `position:"Query" name:"LaunchTemplateName"`
	ResourceOwnerId                          requests.Integer                `position:"Query" name:"ResourceOwnerId"`
	HpcClusterId                             string                          `position:"Query" name:"HpcClusterId"`
	HttpPutResponseHopLimit                  requests.Integer                `position:"Query" name:"HttpPutResponseHopLimit"`
	Isp                                      string                          `position:"Query" name:"Isp"`
	KeyPairName                              string                          `position:"Query" name:"KeyPairName"`
	SpotPriceLimit                           requests.Float                  `position:"Query" name:"SpotPriceLimit"`
	StorageSetPartitionNumber                requests.Integer                `position:"Query" name:"StorageSetPartitionNumber"`
	Tag                                      *[]RunInstancesTag              `position:"Query" name:"Tag"  type:"Repeated"`
	PrivatePoolOptionsId                     string                          `position:"Query" name:"PrivatePoolOptions.Id"`
	AutoRenewPeriod                          requests.Integer                `position:"Query" name:"AutoRenewPeriod"`
	LaunchTemplateId                         string                          `position:"Query" name:"LaunchTemplateId"`
	Ipv6AddressCount                         requests.Integer                `position:"Query" name:"Ipv6AddressCount"`
	HostNames                                *[]string                       `position:"Query" name:"HostNames"  type:"Repeated"`
	CapacityReservationPreference            string                          `position:"Query" name:"CapacityReservationPreference"`
	VSwitchId                                string                          `position:"Query" name:"VSwitchId"`
	InstanceName                             string                          `position:"Query" name:"InstanceName"`
	ZoneId                                   string                          `position:"Query" name:"ZoneId"`
	Ipv6Address                              *[]string                       `position:"Query" name:"Ipv6Address"  type:"Repeated"`
	SecurityOptionsConfidentialComputingMode string                          `position:"Query" name:"SecurityOptions.ConfidentialComputingMode"`
	ClientToken                              string                          `position:"Query" name:"ClientToken"`
	InternetMaxBandwidthOut                  requests.Integer                `position:"Query" name:"InternetMaxBandwidthOut"`
	Description                              string                          `position:"Query" name:"Description"`
	CpuOptionsThreadsPerCore                 requests.Integer                `position:"Query" name:"CpuOptions.ThreadsPerCore"`
	SystemDiskCategory                       string                          `position:"Query" name:"SystemDisk.Category"`
	SecurityOptionsTrustedSystemMode         string                          `position:"Query" name:"SecurityOptions.TrustedSystemMode"`
	CapacityReservationId                    string                          `position:"Query" name:"CapacityReservationId"`
	UserData                                 string                          `position:"Query" name:"UserData"`
	HttpEndpoint                             string                          `position:"Query" name:"HttpEndpoint"`
	InstanceChargeType                       string                          `position:"Query" name:"InstanceChargeType"`
	DeploymentSetId                          string                          `position:"Query" name:"DeploymentSetId"`
	NetworkInterface                         *[]RunInstancesNetworkInterface `position:"Query" name:"NetworkInterface"  type:"Repeated"`
	Amount                                   requests.Integer                `position:"Query" name:"Amount"`
	AutoPay                                  requests.Boolean                `position:"Query" name:"AutoPay"`
	OwnerAccount                             string                          `position:"Query" name:"OwnerAccount"`
	Tenancy                                  string                          `position:"Query" name:"Tenancy"`
	RamRoleName                              string                          `position:"Query" name:"RamRoleName"`
	AutoReleaseTime                          string                          `position:"Query" name:"AutoReleaseTime"`
	CreditSpecification                      string                          `position:"Query" name:"CreditSpecification"`
	LaunchTemplateVersion                    requests.Integer                `position:"Query" name:"LaunchTemplateVersion"`
	SchedulerOptionsManagedPrivateSpaceId    string                          `position:"Query" name:"SchedulerOptions.ManagedPrivateSpaceId"`
	DataDisk                                 *[]RunInstancesDataDisk         `position:"Query" name:"DataDisk"  type:"Repeated"`
	StorageSetId                             string                          `position:"Query" name:"StorageSetId"`
	HttpTokens                               string                          `position:"Query" name:"HttpTokens"`
	SystemDiskDescription                    string                          `position:"Query" name:"SystemDisk.Description"`
}

// RunInstancesSystemDisk is a repeated param struct in RunInstancesRequest
type RunInstancesSystemDisk struct {
	StorageClusterId string `name:"StorageClusterId"`
	ProvisionedIops  string `name:"ProvisionedIops"`
	BurstingEnabled  string `name:"BurstingEnabled"`
	Encrypted        string `name:"Encrypted"`
	KMSKeyId         string `name:"KMSKeyId"`
	EncryptAlgorithm string `name:"EncryptAlgorithm"`
}

// RunInstancesImageOptions is a repeated param struct in RunInstancesRequest
type RunInstancesImageOptions struct {
	LoginAsNonRoot string `name:"LoginAsNonRoot"`
}

// RunInstancesSystemTag is a repeated param struct in RunInstancesRequest
type RunInstancesSystemTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Scope string `name:"Scope"`
}

// RunInstancesArn is a repeated param struct in RunInstancesRequest
type RunInstancesArn struct {
	RoleType      string `name:"RoleType"`
	Rolearn       string `name:"Rolearn"`
	AssumeRoleFor string `name:"AssumeRoleFor"`
}

// RunInstancesNetworkOptions is a repeated param struct in RunInstancesRequest
type RunInstancesNetworkOptions struct {
	EnableJumboFrame string `name:"EnableJumboFrame"`
}

// RunInstancesTag is a repeated param struct in RunInstancesRequest
type RunInstancesTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// RunInstancesNetworkInterface is a repeated param struct in RunInstancesRequest
type RunInstancesNetworkInterface struct {
	VSwitchId                   string    `name:"VSwitchId"`
	NetworkInterfaceName        string    `name:"NetworkInterfaceName"`
	Description                 string    `name:"Description"`
	SecurityGroupId             string    `name:"SecurityGroupId"`
	PrimaryIpAddress            string    `name:"PrimaryIpAddress"`
	QueueNumber                 string    `name:"QueueNumber"`
	SecurityGroupIds            *[]string `name:"SecurityGroupIds" type:"Repeated"`
	NetworkInterfaceTrafficMode string    `name:"NetworkInterfaceTrafficMode"`
	QueuePairNumber             string    `name:"QueuePairNumber"`
	InstanceType                string    `name:"InstanceType"`
	Ipv6AddressCount            string    `name:"Ipv6AddressCount"`
	Ipv6Address                 *[]string `name:"Ipv6Address" type:"Repeated"`
	NetworkCardIndex            string    `name:"NetworkCardIndex"`
	DeleteOnRelease             string    `name:"DeleteOnRelease"`
	NetworkInterfaceId          string    `name:"NetworkInterfaceId"`
}

// RunInstancesDataDisk is a repeated param struct in RunInstancesRequest
type RunInstancesDataDisk struct {
	PerformanceLevel     string `name:"PerformanceLevel"`
	AutoSnapshotPolicyId string `name:"AutoSnapshotPolicyId"`
	Encrypted            string `name:"Encrypted"`
	Description          string `name:"Description"`
	SnapshotId           string `name:"SnapshotId"`
	Device               string `name:"Device"`
	Size                 string `name:"Size"`
	DiskName             string `name:"DiskName"`
	Category             string `name:"Category"`
	EncryptAlgorithm     string `name:"EncryptAlgorithm"`
	DeleteWithInstance   string `name:"DeleteWithInstance"`
	KMSKeyId             string `name:"KMSKeyId"`
	StorageClusterId     string `name:"StorageClusterId"`
	ProvisionedIops      string `name:"ProvisionedIops"`
	BurstingEnabled      string `name:"BurstingEnabled"`
}

// RunInstancesResponse is the response struct for api RunInstances
type RunInstancesResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	OrderId        string         `json:"OrderId" xml:"OrderId"`
	TradePrice     float64        `json:"TradePrice" xml:"TradePrice"`
	InstanceIdSets InstanceIdSets `json:"InstanceIdSets" xml:"InstanceIdSets"`
}

// CreateRunInstancesRequest creates a request to invoke RunInstances API
func CreateRunInstancesRequest() (request *RunInstancesRequest) {
	request = &RunInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "RunInstances", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRunInstancesResponse creates a response to parse from RunInstances response
func CreateRunInstancesResponse() (response *RunInstancesResponse) {
	response = &RunInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
