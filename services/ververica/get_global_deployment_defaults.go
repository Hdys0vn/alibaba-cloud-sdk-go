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

// GetGlobalDeploymentDefaults invokes the ververica.GetGlobalDeploymentDefaults API synchronously
func (client *Client) GetGlobalDeploymentDefaults(request *GetGlobalDeploymentDefaultsRequest) (response *GetGlobalDeploymentDefaultsResponse, err error) {
	response = CreateGetGlobalDeploymentDefaultsResponse()
	err = client.DoAction(request, response)
	return
}

// GetGlobalDeploymentDefaultsWithChan invokes the ververica.GetGlobalDeploymentDefaults API asynchronously
func (client *Client) GetGlobalDeploymentDefaultsWithChan(request *GetGlobalDeploymentDefaultsRequest) (<-chan *GetGlobalDeploymentDefaultsResponse, <-chan error) {
	responseChan := make(chan *GetGlobalDeploymentDefaultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGlobalDeploymentDefaults(request)
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

// GetGlobalDeploymentDefaultsWithCallback invokes the ververica.GetGlobalDeploymentDefaults API asynchronously
func (client *Client) GetGlobalDeploymentDefaultsWithCallback(request *GetGlobalDeploymentDefaultsRequest, callback func(response *GetGlobalDeploymentDefaultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGlobalDeploymentDefaultsResponse
		var err error
		defer close(result)
		response, err = client.GetGlobalDeploymentDefaults(request)
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

// GetGlobalDeploymentDefaultsRequest is the request struct for api GetGlobalDeploymentDefaults
type GetGlobalDeploymentDefaultsRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Namespace string `position:"Header" name:"namespace"`
}

// GetGlobalDeploymentDefaultsResponse is the response struct for api GetGlobalDeploymentDefaults
type GetGlobalDeploymentDefaultsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
	Data      string `json:"data" xml:"data"`
}

// CreateGetGlobalDeploymentDefaultsRequest creates a request to invoke GetGlobalDeploymentDefaults API
func CreateGetGlobalDeploymentDefaultsRequest() (request *GetGlobalDeploymentDefaultsRequest) {
	request = &GetGlobalDeploymentDefaultsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "GetGlobalDeploymentDefaults", "/pop/workspaces/[workspace]/api/v1/global-deployment-defaults", "", "")
	request.Method = requests.GET
	return
}

// CreateGetGlobalDeploymentDefaultsResponse creates a response to parse from GetGlobalDeploymentDefaults response
func CreateGetGlobalDeploymentDefaultsResponse() (response *GetGlobalDeploymentDefaultsResponse) {
	response = &GetGlobalDeploymentDefaultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
