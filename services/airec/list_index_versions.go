package airec

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

// ListIndexVersions invokes the airec.ListIndexVersions API synchronously
func (client *Client) ListIndexVersions(request *ListIndexVersionsRequest) (response *ListIndexVersionsResponse, err error) {
	response = CreateListIndexVersionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListIndexVersionsWithChan invokes the airec.ListIndexVersions API asynchronously
func (client *Client) ListIndexVersionsWithChan(request *ListIndexVersionsRequest) (<-chan *ListIndexVersionsResponse, <-chan error) {
	responseChan := make(chan *ListIndexVersionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIndexVersions(request)
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

// ListIndexVersionsWithCallback invokes the airec.ListIndexVersions API asynchronously
func (client *Client) ListIndexVersionsWithCallback(request *ListIndexVersionsRequest, callback func(response *ListIndexVersionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIndexVersionsResponse
		var err error
		defer close(result)
		response, err = client.ListIndexVersions(request)
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

// ListIndexVersionsRequest is the request struct for api ListIndexVersions
type ListIndexVersionsRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"instanceId"`
	AlgorithmId string `position:"Path" name:"algorithmId"`
}

// ListIndexVersionsResponse is the response struct for api ListIndexVersions
type ListIndexVersionsResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"requestId" xml:"requestId"`
	Result    []IndeVersion `json:"result" xml:"result"`
}

// CreateListIndexVersionsRequest creates a request to invoke ListIndexVersions API
func CreateListIndexVersionsRequest() (request *ListIndexVersionsRequest) {
	request = &ListIndexVersionsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "ListIndexVersions", "/v2/openapi/instances/[instanceId]/filtering-algorithms/[algorithmId]/index-versions", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListIndexVersionsResponse creates a response to parse from ListIndexVersions response
func CreateListIndexVersionsResponse() (response *ListIndexVersionsResponse) {
	response = &ListIndexVersionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
