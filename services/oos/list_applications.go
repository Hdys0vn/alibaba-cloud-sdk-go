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

// ListApplications invokes the oos.ListApplications API synchronously
func (client *Client) ListApplications(request *ListApplicationsRequest) (response *ListApplicationsResponse, err error) {
	response = CreateListApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListApplicationsWithChan invokes the oos.ListApplications API asynchronously
func (client *Client) ListApplicationsWithChan(request *ListApplicationsRequest) (<-chan *ListApplicationsResponse, <-chan error) {
	responseChan := make(chan *ListApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApplications(request)
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

// ListApplicationsWithCallback invokes the oos.ListApplications API asynchronously
func (client *Client) ListApplicationsWithCallback(request *ListApplicationsRequest, callback func(response *ListApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApplicationsResponse
		var err error
		defer close(result)
		response, err = client.ListApplications(request)
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

// ListApplicationsRequest is the request struct for api ListApplications
type ListApplicationsRequest struct {
	*requests.RpcRequest
	NextToken       string           `position:"Query" name:"NextToken"`
	Tags            string           `position:"Query" name:"Tags"`
	Names           string           `position:"Query" name:"Names"`
	Name            string           `position:"Query" name:"Name"`
	MaxResults      requests.Integer `position:"Query" name:"MaxResults"`
	ApplicationType string           `position:"Query" name:"ApplicationType"`
}

// ListApplicationsResponse is the response struct for api ListApplications
type ListApplicationsResponse struct {
	*responses.BaseResponse
	NextToken    string        `json:"NextToken" xml:"NextToken"`
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	MaxResults   int           `json:"MaxResults" xml:"MaxResults"`
	Applications []Application `json:"Applications" xml:"Applications"`
}

// CreateListApplicationsRequest creates a request to invoke ListApplications API
func CreateListApplicationsRequest() (request *ListApplicationsRequest) {
	request = &ListApplicationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ListApplications", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListApplicationsResponse creates a response to parse from ListApplications response
func CreateListApplicationsResponse() (response *ListApplicationsResponse) {
	response = &ListApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
