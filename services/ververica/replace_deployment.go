package ververica

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

// ReplaceDeployment invokes the ververica.ReplaceDeployment API synchronously
func (client *Client) ReplaceDeployment(request *ReplaceDeploymentRequest) (response *ReplaceDeploymentResponse, err error) {
	response = CreateReplaceDeploymentResponse()
	err = client.DoAction(request, response)
	return
}

// ReplaceDeploymentWithChan invokes the ververica.ReplaceDeployment API asynchronously
func (client *Client) ReplaceDeploymentWithChan(request *ReplaceDeploymentRequest) (<-chan *ReplaceDeploymentResponse, <-chan error) {
	responseChan := make(chan *ReplaceDeploymentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReplaceDeployment(request)
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

// ReplaceDeploymentWithCallback invokes the ververica.ReplaceDeployment API asynchronously
func (client *Client) ReplaceDeploymentWithCallback(request *ReplaceDeploymentRequest, callback func(response *ReplaceDeploymentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReplaceDeploymentResponse
		var err error
		defer close(result)
		response, err = client.ReplaceDeployment(request)
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

// ReplaceDeploymentRequest is the request struct for api ReplaceDeployment
type ReplaceDeploymentRequest struct {
	*requests.RoaRequest
	Workspace    string `position:"Path" name:"workspace"`
	ParamsJson   string `position:"Body" name:"paramsJson"`
	DeploymentId string `position:"Path" name:"deploymentId"`
	Namespace    string `position:"Path" name:"namespace"`
}

// ReplaceDeploymentResponse is the response struct for api ReplaceDeployment
type ReplaceDeploymentResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
	Data      int    `json:"data" xml:"data"`
}

// CreateReplaceDeploymentRequest creates a request to invoke ReplaceDeployment API
func CreateReplaceDeploymentRequest() (request *ReplaceDeploymentRequest) {
	request = &ReplaceDeploymentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ReplaceDeployment", "/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployments/[deploymentId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateReplaceDeploymentResponse creates a response to parse from ReplaceDeployment response
func CreateReplaceDeploymentResponse() (response *ReplaceDeploymentResponse) {
	response = &ReplaceDeploymentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}