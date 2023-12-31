package devops_rdc

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

// ListDevopsProjects invokes the devops_rdc.ListDevopsProjects API synchronously
func (client *Client) ListDevopsProjects(request *ListDevopsProjectsRequest) (response *ListDevopsProjectsResponse, err error) {
	response = CreateListDevopsProjectsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDevopsProjectsWithChan invokes the devops_rdc.ListDevopsProjects API asynchronously
func (client *Client) ListDevopsProjectsWithChan(request *ListDevopsProjectsRequest) (<-chan *ListDevopsProjectsResponse, <-chan error) {
	responseChan := make(chan *ListDevopsProjectsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDevopsProjects(request)
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

// ListDevopsProjectsWithCallback invokes the devops_rdc.ListDevopsProjects API asynchronously
func (client *Client) ListDevopsProjectsWithCallback(request *ListDevopsProjectsRequest, callback func(response *ListDevopsProjectsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDevopsProjectsResponse
		var err error
		defer close(result)
		response, err = client.ListDevopsProjects(request)
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

// ListDevopsProjectsRequest is the request struct for api ListDevopsProjects
type ListDevopsProjectsRequest struct {
	*requests.RpcRequest
	SelectBy  string           `position:"Body" name:"SelectBy"`
	PageSize  requests.Integer `position:"Body" name:"PageSize"`
	OrderBy   string           `position:"Body" name:"OrderBy"`
	OrgId     string           `position:"Body" name:"OrgId"`
	PageToken string           `position:"Body" name:"PageToken"`
}

// ListDevopsProjectsResponse is the response struct for api ListDevopsProjects
type ListDevopsProjectsResponse struct {
	*responses.BaseResponse
	Successful bool                       `json:"Successful" xml:"Successful"`
	ErrorCode  string                     `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string                     `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string                     `json:"RequestId" xml:"RequestId"`
	Object     ObjectInListDevopsProjects `json:"Object" xml:"Object"`
}

// CreateListDevopsProjectsRequest creates a request to invoke ListDevopsProjects API
func CreateListDevopsProjectsRequest() (request *ListDevopsProjectsRequest) {
	request = &ListDevopsProjectsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "ListDevopsProjects", "", "")
	request.Method = requests.POST
	return
}

// CreateListDevopsProjectsResponse creates a response to parse from ListDevopsProjects response
func CreateListDevopsProjectsResponse() (response *ListDevopsProjectsResponse) {
	response = &ListDevopsProjectsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
