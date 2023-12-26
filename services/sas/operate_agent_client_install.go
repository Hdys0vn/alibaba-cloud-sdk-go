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

// OperateAgentClientInstall invokes the sas.OperateAgentClientInstall API synchronously
func (client *Client) OperateAgentClientInstall(request *OperateAgentClientInstallRequest) (response *OperateAgentClientInstallResponse, err error) {
	response = CreateOperateAgentClientInstallResponse()
	err = client.DoAction(request, response)
	return
}

// OperateAgentClientInstallWithChan invokes the sas.OperateAgentClientInstall API asynchronously
func (client *Client) OperateAgentClientInstallWithChan(request *OperateAgentClientInstallRequest) (<-chan *OperateAgentClientInstallResponse, <-chan error) {
	responseChan := make(chan *OperateAgentClientInstallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperateAgentClientInstall(request)
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

// OperateAgentClientInstallWithCallback invokes the sas.OperateAgentClientInstall API asynchronously
func (client *Client) OperateAgentClientInstallWithCallback(request *OperateAgentClientInstallRequest, callback func(response *OperateAgentClientInstallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperateAgentClientInstallResponse
		var err error
		defer close(result)
		response, err = client.OperateAgentClientInstall(request)
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

// OperateAgentClientInstallRequest is the request struct for api OperateAgentClientInstall
type OperateAgentClientInstallRequest struct {
	*requests.RpcRequest
	SourceIp    string `position:"Query" name:"SourceIp"`
	InstanceIds string `position:"Query" name:"InstanceIds"`
	Uuids       string `position:"Query" name:"Uuids"`
}

// OperateAgentClientInstallResponse is the response struct for api OperateAgentClientInstall
type OperateAgentClientInstallResponse struct {
	*responses.BaseResponse
	RequestId                     string                      `json:"RequestId" xml:"RequestId"`
	AegisCelintInstallResposeList []AegisCelintInstallRespose `json:"AegisCelintInstallResposeList" xml:"AegisCelintInstallResposeList"`
}

// CreateOperateAgentClientInstallRequest creates a request to invoke OperateAgentClientInstall API
func CreateOperateAgentClientInstallRequest() (request *OperateAgentClientInstallRequest) {
	request = &OperateAgentClientInstallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "OperateAgentClientInstall", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOperateAgentClientInstallResponse creates a response to parse from OperateAgentClientInstall response
func CreateOperateAgentClientInstallResponse() (response *OperateAgentClientInstallResponse) {
	response = &OperateAgentClientInstallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
