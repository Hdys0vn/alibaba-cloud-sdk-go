package das

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

// ModifyAutoScalingConfig invokes the das.ModifyAutoScalingConfig API synchronously
func (client *Client) ModifyAutoScalingConfig(request *ModifyAutoScalingConfigRequest) (response *ModifyAutoScalingConfigResponse, err error) {
	response = CreateModifyAutoScalingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAutoScalingConfigWithChan invokes the das.ModifyAutoScalingConfig API asynchronously
func (client *Client) ModifyAutoScalingConfigWithChan(request *ModifyAutoScalingConfigRequest) (<-chan *ModifyAutoScalingConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyAutoScalingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAutoScalingConfig(request)
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

// ModifyAutoScalingConfigWithCallback invokes the das.ModifyAutoScalingConfig API asynchronously
func (client *Client) ModifyAutoScalingConfigWithCallback(request *ModifyAutoScalingConfigRequest, callback func(response *ModifyAutoScalingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAutoScalingConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyAutoScalingConfig(request)
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

// ModifyAutoScalingConfigRequest is the request struct for api ModifyAutoScalingConfig
type ModifyAutoScalingConfigRequest struct {
	*requests.RpcRequest
	Bandwidth  ModifyAutoScalingConfigBandwidth `position:"Query" name:"Bandwidth"  type:"Struct"`
	Resource   ModifyAutoScalingConfigResource  `position:"Query" name:"Resource"  type:"Struct"`
	Storage    ModifyAutoScalingConfigStorage   `position:"Query" name:"Storage"  type:"Struct"`
	Spec       ModifyAutoScalingConfigSpec      `position:"Query" name:"Spec"  type:"Struct"`
	InstanceId string                           `position:"Query" name:"InstanceId"`
	Shard      ModifyAutoScalingConfigShard     `position:"Query" name:"Shard"  type:"Struct"`
}

// ModifyAutoScalingConfigBandwidth is a repeated param struct in ModifyAutoScalingConfigRequest
type ModifyAutoScalingConfigBandwidth struct {
	ObservationWindowSize        string `name:"ObservationWindowSize"`
	Upgrade                      string `name:"Upgrade"`
	Apply                        string `name:"Apply"`
	BandwidthUsageLowerThreshold string `name:"BandwidthUsageLowerThreshold"`
	Downgrade                    string `name:"Downgrade"`
	BandwidthUsageUpperThreshold string `name:"BandwidthUsageUpperThreshold"`
}

// ModifyAutoScalingConfigResource is a repeated param struct in ModifyAutoScalingConfigRequest
type ModifyAutoScalingConfigResource struct {
	Apply                          string `name:"Apply"`
	Enable                         string `name:"Enable"`
	UpgradeObservationWindowSize   string `name:"UpgradeObservationWindowSize"`
	DowngradeObservationWindowSize string `name:"DowngradeObservationWindowSize"`
	CpuUsageUpperThreshold         string `name:"CpuUsageUpperThreshold"`
}

// ModifyAutoScalingConfigStorage is a repeated param struct in ModifyAutoScalingConfigRequest
type ModifyAutoScalingConfigStorage struct {
	Upgrade                 string `name:"Upgrade"`
	Apply                   string `name:"Apply"`
	MaxStorage              string `name:"MaxStorage"`
	DiskUsageUpperThreshold string `name:"DiskUsageUpperThreshold"`
}

// ModifyAutoScalingConfigSpec is a repeated param struct in ModifyAutoScalingConfigRequest
type ModifyAutoScalingConfigSpec struct {
	ObservationWindowSize  string `name:"ObservationWindowSize"`
	MaxSpec                string `name:"MaxSpec"`
	Upgrade                string `name:"Upgrade"`
	Apply                  string `name:"Apply"`
	MemUsageUpperThreshold string `name:"MemUsageUpperThreshold"`
	CoolDownTime           string `name:"CoolDownTime"`
	CpuUsageUpperThreshold string `name:"CpuUsageUpperThreshold"`
	MaxReadOnlyNodes       string `name:"MaxReadOnlyNodes"`
	Downgrade              string `name:"Downgrade"`
}

// ModifyAutoScalingConfigShard is a repeated param struct in ModifyAutoScalingConfigRequest
type ModifyAutoScalingConfigShard struct {
	Upgrade                        string `name:"Upgrade"`
	Apply                          string `name:"Apply"`
	MemUsageUpperThreshold         string `name:"MemUsageUpperThreshold"`
	MinShards                      string `name:"MinShards"`
	UpgradeObservationWindowSize   string `name:"UpgradeObservationWindowSize"`
	DowngradeObservationWindowSize string `name:"DowngradeObservationWindowSize"`
	MemUsageLowerThreshold         string `name:"MemUsageLowerThreshold"`
	MaxShards                      string `name:"MaxShards"`
	Downgrade                      string `name:"Downgrade"`
}

// ModifyAutoScalingConfigResponse is the response struct for api ModifyAutoScalingConfig
type ModifyAutoScalingConfigResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateModifyAutoScalingConfigRequest creates a request to invoke ModifyAutoScalingConfig API
func CreateModifyAutoScalingConfigRequest() (request *ModifyAutoScalingConfigRequest) {
	request = &ModifyAutoScalingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "ModifyAutoScalingConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyAutoScalingConfigResponse creates a response to parse from ModifyAutoScalingConfig response
func CreateModifyAutoScalingConfigResponse() (response *ModifyAutoScalingConfigResponse) {
	response = &ModifyAutoScalingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}