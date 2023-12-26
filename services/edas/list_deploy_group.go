package edas

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

// ListDeployGroup invokes the edas.ListDeployGroup API synchronously
func (client *Client) ListDeployGroup(request *ListDeployGroupRequest) (response *ListDeployGroupResponse, err error) {
	response = CreateListDeployGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListDeployGroupWithChan invokes the edas.ListDeployGroup API asynchronously
func (client *Client) ListDeployGroupWithChan(request *ListDeployGroupRequest) (<-chan *ListDeployGroupResponse, <-chan error) {
	responseChan := make(chan *ListDeployGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDeployGroup(request)
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

// ListDeployGroupWithCallback invokes the edas.ListDeployGroup API asynchronously
func (client *Client) ListDeployGroupWithCallback(request *ListDeployGroupRequest, callback func(response *ListDeployGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDeployGroupResponse
		var err error
		defer close(result)
		response, err = client.ListDeployGroup(request)
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

// ListDeployGroupRequest is the request struct for api ListDeployGroup
type ListDeployGroupRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// ListDeployGroupResponse is the response struct for api ListDeployGroup
type ListDeployGroupResponse struct {
	*responses.BaseResponse
	Code            int             `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	DeployGroupList DeployGroupList `json:"DeployGroupList" xml:"DeployGroupList"`
}

// CreateListDeployGroupRequest creates a request to invoke ListDeployGroup API
func CreateListDeployGroupRequest() (request *ListDeployGroupRequest) {
	request = &ListDeployGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListDeployGroup", "/pop/v5/app/deploy_group_list", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDeployGroupResponse creates a response to parse from ListDeployGroup response
func CreateListDeployGroupResponse() (response *ListDeployGroupResponse) {
	response = &ListDeployGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}