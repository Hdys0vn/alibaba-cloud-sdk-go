package sas

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

// ModifyVulTargetConfig invokes the sas.ModifyVulTargetConfig API synchronously
func (client *Client) ModifyVulTargetConfig(request *ModifyVulTargetConfigRequest) (response *ModifyVulTargetConfigResponse, err error) {
	response = CreateModifyVulTargetConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVulTargetConfigWithChan invokes the sas.ModifyVulTargetConfig API asynchronously
func (client *Client) ModifyVulTargetConfigWithChan(request *ModifyVulTargetConfigRequest) (<-chan *ModifyVulTargetConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyVulTargetConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVulTargetConfig(request)
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

// ModifyVulTargetConfigWithCallback invokes the sas.ModifyVulTargetConfig API asynchronously
func (client *Client) ModifyVulTargetConfigWithCallback(request *ModifyVulTargetConfigRequest, callback func(response *ModifyVulTargetConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVulTargetConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyVulTargetConfig(request)
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

// ModifyVulTargetConfigRequest is the request struct for api ModifyVulTargetConfig
type ModifyVulTargetConfigRequest struct {
	*requests.RpcRequest
	Type     string `position:"Query" name:"Type"`
	Uuid     string `position:"Query" name:"Uuid"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Config   string `position:"Query" name:"Config"`
}

// ModifyVulTargetConfigResponse is the response struct for api ModifyVulTargetConfig
type ModifyVulTargetConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVulTargetConfigRequest creates a request to invoke ModifyVulTargetConfig API
func CreateModifyVulTargetConfigRequest() (request *ModifyVulTargetConfigRequest) {
	request = &ModifyVulTargetConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyVulTargetConfig", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVulTargetConfigResponse creates a response to parse from ModifyVulTargetConfig response
func CreateModifyVulTargetConfigResponse() (response *ModifyVulTargetConfigResponse) {
	response = &ModifyVulTargetConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
