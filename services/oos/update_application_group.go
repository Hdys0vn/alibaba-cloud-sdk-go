package oos

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

// UpdateApplicationGroup invokes the oos.UpdateApplicationGroup API synchronously
func (client *Client) UpdateApplicationGroup(request *UpdateApplicationGroupRequest) (response *UpdateApplicationGroupResponse, err error) {
	response = CreateUpdateApplicationGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateApplicationGroupWithChan invokes the oos.UpdateApplicationGroup API asynchronously
func (client *Client) UpdateApplicationGroupWithChan(request *UpdateApplicationGroupRequest) (<-chan *UpdateApplicationGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateApplicationGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateApplicationGroup(request)
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

// UpdateApplicationGroupWithCallback invokes the oos.UpdateApplicationGroup API asynchronously
func (client *Client) UpdateApplicationGroupWithCallback(request *UpdateApplicationGroupRequest, callback func(response *UpdateApplicationGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateApplicationGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateApplicationGroup(request)
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

// UpdateApplicationGroupRequest is the request struct for api UpdateApplicationGroup
type UpdateApplicationGroupRequest struct {
	*requests.RpcRequest
	NewName         string `position:"Query" name:"NewName"`
	ApplicationName string `position:"Query" name:"ApplicationName"`
	Name            string `position:"Query" name:"Name"`
}

// UpdateApplicationGroupResponse is the response struct for api UpdateApplicationGroup
type UpdateApplicationGroupResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	ApplicationGroup ApplicationGroup `json:"ApplicationGroup" xml:"ApplicationGroup"`
}

// CreateUpdateApplicationGroupRequest creates a request to invoke UpdateApplicationGroup API
func CreateUpdateApplicationGroupRequest() (request *UpdateApplicationGroupRequest) {
	request = &UpdateApplicationGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "UpdateApplicationGroup", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateApplicationGroupResponse creates a response to parse from UpdateApplicationGroup response
func CreateUpdateApplicationGroupResponse() (response *UpdateApplicationGroupResponse) {
	response = &UpdateApplicationGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}