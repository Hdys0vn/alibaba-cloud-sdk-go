package emr

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

// ModifyClusterTemplate invokes the emr.ModifyClusterTemplate API synchronously
func (client *Client) ModifyClusterTemplate(request *ModifyClusterTemplateRequest) (response *ModifyClusterTemplateResponse, err error) {
	response = CreateModifyClusterTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyClusterTemplateWithChan invokes the emr.ModifyClusterTemplate API asynchronously
func (client *Client) ModifyClusterTemplateWithChan(request *ModifyClusterTemplateRequest) (<-chan *ModifyClusterTemplateResponse, <-chan error) {
	responseChan := make(chan *ModifyClusterTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyClusterTemplate(request)
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

// ModifyClusterTemplateWithCallback invokes the emr.ModifyClusterTemplate API asynchronously
func (client *Client) ModifyClusterTemplateWithCallback(request *ModifyClusterTemplateRequest, callback func(response *ModifyClusterTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyClusterTemplateResponse
		var err error
		defer close(result)
		response, err = client.ModifyClusterTemplate(request)
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

// ModifyClusterTemplateRequest is the request struct for api ModifyClusterTemplate
type ModifyClusterTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer                        `position:"Query" name:"ResourceOwnerId"`
	LogPath                string                                  `position:"Query" name:"LogPath"`
	MasterPwd              string                                  `position:"Query" name:"MasterPwd"`
	Configurations         string                                  `position:"Query" name:"Configurations"`
	SshEnable              requests.Boolean                        `position:"Query" name:"SshEnable"`
	KeyPairName            string                                  `position:"Query" name:"KeyPairName"`
	MetaStoreType          string                                  `position:"Query" name:"MetaStoreType"`
	SecurityGroupName      string                                  `position:"Query" name:"SecurityGroupName"`
	MachineType            string                                  `position:"Query" name:"MachineType"`
	ResourceGroupId        string                                  `position:"Query" name:"ResourceGroupId"`
	BootstrapAction        *[]ModifyClusterTemplateBootstrapAction `position:"Query" name:"BootstrapAction"  type:"Repeated"`
	MetaStoreConf          string                                  `position:"Query" name:"MetaStoreConf"`
	EmrVer                 string                                  `position:"Query" name:"EmrVer"`
	Tag                    *[]ModifyClusterTemplateTag             `position:"Query" name:"Tag"  type:"Repeated"`
	IsOpenPublicIp         requests.Boolean                        `position:"Query" name:"IsOpenPublicIp"`
	Period                 requests.Integer                        `position:"Query" name:"Period"`
	InstanceGeneration     string                                  `position:"Query" name:"InstanceGeneration"`
	VSwitchId              string                                  `position:"Query" name:"VSwitchId"`
	ClusterType            string                                  `position:"Query" name:"ClusterType"`
	AutoRenew              requests.Boolean                        `position:"Query" name:"AutoRenew"`
	OptionSoftWareList     *[]string                               `position:"Query" name:"OptionSoftWareList"  type:"Repeated"`
	NetType                string                                  `position:"Query" name:"NetType"`
	ZoneId                 string                                  `position:"Query" name:"ZoneId"`
	UseCustomHiveMetaDb    requests.Boolean                        `position:"Query" name:"UseCustomHiveMetaDb"`
	InitCustomHiveMetaDb   requests.Boolean                        `position:"Query" name:"InitCustomHiveMetaDb"`
	IoOptimized            requests.Boolean                        `position:"Query" name:"IoOptimized"`
	SecurityGroupId        string                                  `position:"Query" name:"SecurityGroupId"`
	EasEnable              requests.Boolean                        `position:"Query" name:"EasEnable"`
	DepositType            string                                  `position:"Query" name:"DepositType"`
	DataDiskKMSKeyId       string                                  `position:"Query" name:"DataDiskKMSKeyId"`
	UseLocalMetaDb         requests.Boolean                        `position:"Query" name:"UseLocalMetaDb"`
	TemplateName           string                                  `position:"Query" name:"TemplateName"`
	UserDefinedEmrEcsRole  string                                  `position:"Query" name:"UserDefinedEmrEcsRole"`
	DataDiskEncrypted      requests.Boolean                        `position:"Query" name:"DataDiskEncrypted"`
	VpcId                  string                                  `position:"Query" name:"VpcId"`
	BizId                  string                                  `position:"Query" name:"BizId"`
	HostGroup              *[]ModifyClusterTemplateHostGroup       `position:"Query" name:"HostGroup"  type:"Repeated"`
	ChargeType             string                                  `position:"Query" name:"ChargeType"`
	Config                 *[]ModifyClusterTemplateConfig          `position:"Query" name:"Config"  type:"Repeated"`
	HighAvailabilityEnable requests.Boolean                        `position:"Query" name:"HighAvailabilityEnable"`
}

// ModifyClusterTemplateBootstrapAction is a repeated param struct in ModifyClusterTemplateRequest
type ModifyClusterTemplateBootstrapAction struct {
	Path                  string `name:"Path"`
	ExecutionTarget       string `name:"ExecutionTarget"`
	ExecutionMoment       string `name:"ExecutionMoment"`
	Arg                   string `name:"Arg"`
	Name                  string `name:"Name"`
	ExecutionFailStrategy string `name:"ExecutionFailStrategy"`
}

// ModifyClusterTemplateTag is a repeated param struct in ModifyClusterTemplateRequest
type ModifyClusterTemplateTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ModifyClusterTemplateHostGroup is a repeated param struct in ModifyClusterTemplateRequest
type ModifyClusterTemplateHostGroup struct {
	Period             string `name:"Period"`
	SysDiskCapacity    string `name:"SysDiskCapacity"`
	DiskCapacity       string `name:"DiskCapacity"`
	SysDiskType        string `name:"SysDiskType"`
	ClusterId          string `name:"ClusterId"`
	DiskType           string `name:"DiskType"`
	HostGroupName      string `name:"HostGroupName"`
	VSwitchId          string `name:"VSwitchId"`
	DiskCount          string `name:"DiskCount"`
	AutoRenew          string `name:"AutoRenew"`
	HostGroupId        string `name:"HostGroupId"`
	NodeCount          string `name:"NodeCount"`
	InstanceType       string `name:"InstanceType"`
	Comment            string `name:"Comment"`
	ChargeType         string `name:"ChargeType"`
	MultiInstanceTypes string `name:"MultiInstanceTypes"`
	CreateType         string `name:"CreateType"`
	HostGroupType      string `name:"HostGroupType"`
}

// ModifyClusterTemplateConfig is a repeated param struct in ModifyClusterTemplateRequest
type ModifyClusterTemplateConfig struct {
	ConfigKey   string `name:"ConfigKey"`
	FileName    string `name:"FileName"`
	Encrypt     string `name:"Encrypt"`
	Replace     string `name:"Replace"`
	ConfigValue string `name:"ConfigValue"`
	ServiceName string `name:"ServiceName"`
}

// ModifyClusterTemplateResponse is the response struct for api ModifyClusterTemplate
type ModifyClusterTemplateResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ClusterTemplateId string `json:"ClusterTemplateId" xml:"ClusterTemplateId"`
}

// CreateModifyClusterTemplateRequest creates a request to invoke ModifyClusterTemplate API
func CreateModifyClusterTemplateRequest() (request *ModifyClusterTemplateRequest) {
	request = &ModifyClusterTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ModifyClusterTemplate", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyClusterTemplateResponse creates a response to parse from ModifyClusterTemplate response
func CreateModifyClusterTemplateResponse() (response *ModifyClusterTemplateResponse) {
	response = &ModifyClusterTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
