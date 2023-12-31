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

// CreateApplicationGroup invokes the oos.CreateApplicationGroup API synchronously
func (client *Client) CreateApplicationGroup(request *CreateApplicationGroupRequest) (response *CreateApplicationGroupResponse, err error) {
	response = CreateCreateApplicationGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateApplicationGroupWithChan invokes the oos.CreateApplicationGroup API asynchronously
func (client *Client) CreateApplicationGroupWithChan(request *CreateApplicationGroupRequest) (<-chan *CreateApplicationGroupResponse, <-chan error) {
	responseChan := make(chan *CreateApplicationGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApplicationGroup(request)
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

// CreateApplicationGroupWithCallback invokes the oos.CreateApplicationGroup API asynchronously
func (client *Client) CreateApplicationGroupWithCallback(request *CreateApplicationGroupRequest, callback func(response *CreateApplicationGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateApplicationGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateApplicationGroup(request)
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

// CreateApplicationGroupRequest is the request struct for api CreateApplicationGroup
type CreateApplicationGroupRequest struct {
	*requests.RpcRequest
	ClientToken     string `position:"Query" name:"ClientToken"`
	Description     string `position:"Query" name:"Description"`
	CmsGroupId      string `position:"Query" name:"CmsGroupId"`
	DeployRegionId  string `position:"Query" name:"DeployRegionId"`
	ApplicationName string `position:"Query" name:"ApplicationName"`
	ImportTagValue  string `position:"Query" name:"ImportTagValue"`
	ImportTagKey    string `position:"Query" name:"ImportTagKey"`
	Name            string `position:"Query" name:"Name"`
}

// CreateApplicationGroupResponse is the response struct for api CreateApplicationGroup
type CreateApplicationGroupResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	ApplicationGroup ApplicationGroup `json:"ApplicationGroup" xml:"ApplicationGroup"`
}

// CreateCreateApplicationGroupRequest creates a request to invoke CreateApplicationGroup API
func CreateCreateApplicationGroupRequest() (request *CreateApplicationGroupRequest) {
	request = &CreateApplicationGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "CreateApplicationGroup", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateApplicationGroupResponse creates a response to parse from CreateApplicationGroup response
func CreateCreateApplicationGroupResponse() (response *CreateApplicationGroupResponse) {
	response = &CreateApplicationGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
