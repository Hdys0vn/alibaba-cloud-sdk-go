package ehpc

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

// CreateCluster invokes the ehpc.CreateCluster API synchronously
func (client *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	response = CreateCreateClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateClusterWithChan invokes the ehpc.CreateCluster API asynchronously
func (client *Client) CreateClusterWithChan(request *CreateClusterRequest) (<-chan *CreateClusterResponse, <-chan error) {
	responseChan := make(chan *CreateClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCluster(request)
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

// CreateClusterWithCallback invokes the ehpc.CreateCluster API asynchronously
func (client *Client) CreateClusterWithCallback(request *CreateClusterRequest, callback func(response *CreateClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateCluster(request)
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

// CreateClusterRequest is the request struct for api CreateCluster
type CreateClusterRequest struct {
	*requests.RpcRequest
	AdditionalVolumes           *[]CreateClusterAdditionalVolumes `position:"Query" name:"AdditionalVolumes"  type:"Repeated"`
	AddOns                      *[]CreateClusterAddOns            `position:"Query" name:"AddOns"  type:"Repeated"`
	EcsOrderManagerInstanceType string                            `position:"Query" name:"EcsOrder.Manager.InstanceType"`
	KeyPairName                 string                            `position:"Query" name:"KeyPairName"`
	SecurityGroupName           string                            `position:"Query" name:"SecurityGroupName"`
	WithoutNas                  requests.Boolean                  `position:"Query" name:"WithoutNas"`
	ImageOwnerAlias             string                            `position:"Query" name:"ImageOwnerAlias"`
	DeployMode                  string                            `position:"Query" name:"DeployMode"`
	EcsOrderManagerCount        requests.Integer                  `position:"Query" name:"EcsOrder.Manager.Count"`
	ResourceGroupId             string                            `position:"Query" name:"ResourceGroupId"`
	Password                    string                            `position:"Query" name:"Password"`
	EcsOrderLoginCount          requests.Integer                  `position:"Query" name:"EcsOrder.Login.Count"`
	WithoutElasticIp            requests.Boolean                  `position:"Query" name:"WithoutElasticIp"`
	RemoteVisEnable             string                            `position:"Query" name:"RemoteVisEnable"`
	SystemDiskSize              requests.Integer                  `position:"Query" name:"SystemDiskSize"`
	Tag                         *[]CreateClusterTag               `position:"Query" name:"Tag"  type:"Repeated"`
	ComputeSpotPriceLimit       string                            `position:"Query" name:"ComputeSpotPriceLimit"`
	AutoRenewPeriod             requests.Integer                  `position:"Query" name:"AutoRenewPeriod"`
	Period                      requests.Integer                  `position:"Query" name:"Period"`
	RemoteDirectory             string                            `position:"Query" name:"RemoteDirectory"`
	EcsOrderComputeCount        requests.Integer                  `position:"Query" name:"EcsOrder.Compute.Count"`
	ComputeSpotStrategy         string                            `position:"Query" name:"ComputeSpotStrategy"`
	PostInstallScript           *[]CreateClusterPostInstallScript `position:"Query" name:"PostInstallScript"  type:"Repeated"`
	RamNodeTypes                *[]string                         `position:"Query" name:"RamNodeTypes"  type:"Repeated"`
	VSwitchId                   string                            `position:"Query" name:"VSwitchId"`
	PeriodUnit                  string                            `position:"Query" name:"PeriodUnit"`
	ComputeEnableHt             requests.Boolean                  `position:"Query" name:"ComputeEnableHt"`
	AutoRenew                   string                            `position:"Query" name:"AutoRenew"`
	Domain                      string                            `position:"Query" name:"Domain"`
	Name                        string                            `position:"Query" name:"Name"`
	VolumeId                    string                            `position:"Query" name:"VolumeId"`
	ZoneId                      string                            `position:"Query" name:"ZoneId"`
	SccClusterId                string                            `position:"Query" name:"SccClusterId"`
	VolumeMountOption           string                            `position:"Query" name:"VolumeMountOption"`
	ImageId                     string                            `position:"Query" name:"ImageId"`
	SystemDiskLevel             string                            `position:"Query" name:"SystemDiskLevel"`
	ClientToken                 string                            `position:"Query" name:"ClientToken"`
	EhpcVersion                 string                            `position:"Query" name:"EhpcVersion"`
	AccountType                 string                            `position:"Query" name:"AccountType"`
	SecurityGroupId             string                            `position:"Query" name:"SecurityGroupId"`
	Description                 string                            `position:"Query" name:"Description"`
	EcsOrderComputeInstanceType string                            `position:"Query" name:"EcsOrder.Compute.InstanceType"`
	JobQueue                    string                            `position:"Query" name:"JobQueue"`
	VolumeType                  string                            `position:"Query" name:"VolumeType"`
	SystemDiskType              string                            `position:"Query" name:"SystemDiskType"`
	DeploymentSetId             string                            `position:"Query" name:"DeploymentSetId"`
	VolumeProtocol              string                            `position:"Query" name:"VolumeProtocol"`
	ClientVersion               string                            `position:"Query" name:"ClientVersion"`
	OsTag                       string                            `position:"Query" name:"OsTag"`
	ClusterVersion              string                            `position:"Query" name:"ClusterVersion"`
	IsComputeEss                requests.Boolean                  `position:"Query" name:"IsComputeEss"`
	RamRoleName                 string                            `position:"Query" name:"RamRoleName"`
	NetworkInterfaceTrafficMode string                            `position:"Query" name:"NetworkInterfaceTrafficMode"`
	Plugin                      string                            `position:"Query" name:"Plugin"`
	Application                 *[]CreateClusterApplication       `position:"Query" name:"Application"  type:"Repeated"`
	EcsChargeType               string                            `position:"Query" name:"EcsChargeType"`
	InputFileUrl                string                            `position:"Query" name:"InputFileUrl"`
	VpcId                       string                            `position:"Query" name:"VpcId"`
	HaEnable                    requests.Boolean                  `position:"Query" name:"HaEnable"`
	WithoutAgent                requests.Boolean                  `position:"Query" name:"WithoutAgent"`
	SchedulerType               string                            `position:"Query" name:"SchedulerType"`
	VolumeMountpoint            string                            `position:"Query" name:"VolumeMountpoint"`
	EcsOrderLoginInstanceType   string                            `position:"Query" name:"EcsOrder.Login.InstanceType"`
}

// CreateClusterAdditionalVolumes is a repeated param struct in CreateClusterRequest
type CreateClusterAdditionalVolumes struct {
	VolumeType        string                                 `name:"VolumeType"`
	VolumeMountOption string                                 `name:"VolumeMountOption"`
	VolumeProtocol    string                                 `name:"VolumeProtocol"`
	LocalDirectory    string                                 `name:"LocalDirectory"`
	RemoteDirectory   string                                 `name:"RemoteDirectory"`
	Roles             *[]CreateClusterAdditionalVolumesRoles `name:"Roles" type:"Repeated"`
	VolumeId          string                                 `name:"VolumeId"`
	VolumeMountpoint  string                                 `name:"VolumeMountpoint"`
	Location          string                                 `name:"Location"`
	JobQueue          string                                 `name:"JobQueue"`
}

// CreateClusterAddOns is a repeated param struct in CreateClusterRequest
type CreateClusterAddOns struct {
	DeployMode   string `name:"DeployMode"`
	Port         string `name:"Port"`
	ConfigFile   string `name:"ConfigFile"`
	DefaultStart string `name:"DefaultStart"`
	Name         string `name:"Name"`
	DBType       string `name:"DBType"`
	Version      string `name:"Version"`
}

// CreateClusterTag is a repeated param struct in CreateClusterRequest
type CreateClusterTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateClusterPostInstallScript is a repeated param struct in CreateClusterRequest
type CreateClusterPostInstallScript struct {
	Args string `name:"Args"`
	Url  string `name:"Url"`
}

// CreateClusterApplication is a repeated param struct in CreateClusterRequest
type CreateClusterApplication struct {
	Tag string `name:"Tag"`
}

// CreateClusterAdditionalVolumesRoles is a repeated param struct in CreateClusterRequest
type CreateClusterAdditionalVolumesRoles struct {
	Name string `name:"Name"`
}

// CreateClusterResponse is the response struct for api CreateCluster
type CreateClusterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
}

// CreateCreateClusterRequest creates a request to invoke CreateCluster API
func CreateCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "CreateCluster", "", "")
	request.Method = requests.GET
	return
}

// CreateCreateClusterResponse creates a response to parse from CreateCluster response
func CreateCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
