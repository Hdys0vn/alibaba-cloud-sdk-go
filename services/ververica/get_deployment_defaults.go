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

// GetDeploymentDefaults invokes the ververica.GetDeploymentDefaults API synchronously
func (client *Client) GetDeploymentDefaults(request *GetDeploymentDefaultsRequest) (response *GetDeploymentDefaultsResponse, err error) {
	response = CreateGetDeploymentDefaultsResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeploymentDefaultsWithChan invokes the ververica.GetDeploymentDefaults API asynchronously
func (client *Client) GetDeploymentDefaultsWithChan(request *GetDeploymentDefaultsRequest) (<-chan *GetDeploymentDefaultsResponse, <-chan error) {
	responseChan := make(chan *GetDeploymentDefaultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeploymentDefaults(request)
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

// GetDeploymentDefaultsWithCallback invokes the ververica.GetDeploymentDefaults API asynchronously
func (client *Client) GetDeploymentDefaultsWithCallback(request *GetDeploymentDefaultsRequest, callback func(response *GetDeploymentDefaultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeploymentDefaultsResponse
		var err error
		defer close(result)
		response, err = client.GetDeploymentDefaults(request)
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

// GetDeploymentDefaultsRequest is the request struct for api GetDeploymentDefaults
type GetDeploymentDefaultsRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Namespace string `position:"Path" name:"namespace"`
}

// GetDeploymentDefaultsResponse is the response struct for api GetDeploymentDefaults
type GetDeploymentDefaultsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	RequestId string `json:"requestId" xml:"requestId"`
	Data      string `json:"data" xml:"data"`
}

// CreateGetDeploymentDefaultsRequest creates a request to invoke GetDeploymentDefaults API
func CreateGetDeploymentDefaultsRequest() (request *GetDeploymentDefaultsRequest) {
	request = &GetDeploymentDefaultsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "GetDeploymentDefaults", "/pop/workspaces/[workspace]/api/v1/namespaces/[namespace]/deployment-defaults", "", "")
	request.Method = requests.GET
	return
}

// CreateGetDeploymentDefaultsResponse creates a response to parse from GetDeploymentDefaults response
func CreateGetDeploymentDefaultsResponse() (response *GetDeploymentDefaultsResponse) {
	response = &GetDeploymentDefaultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
