package opensearch

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

// ListFunctionInstances invokes the opensearch.ListFunctionInstances API synchronously
func (client *Client) ListFunctionInstances(request *ListFunctionInstancesRequest) (response *ListFunctionInstancesResponse, err error) {
	response = CreateListFunctionInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListFunctionInstancesWithChan invokes the opensearch.ListFunctionInstances API asynchronously
func (client *Client) ListFunctionInstancesWithChan(request *ListFunctionInstancesRequest) (<-chan *ListFunctionInstancesResponse, <-chan error) {
	responseChan := make(chan *ListFunctionInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFunctionInstances(request)
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

// ListFunctionInstancesWithCallback invokes the opensearch.ListFunctionInstances API asynchronously
func (client *Client) ListFunctionInstancesWithCallback(request *ListFunctionInstancesRequest, callback func(response *ListFunctionInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFunctionInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListFunctionInstances(request)
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

// ListFunctionInstancesRequest is the request struct for api ListFunctionInstances
type ListFunctionInstancesRequest struct {
	*requests.RoaRequest
	Output           string           `position:"Query" name:"output"`
	ModelType        string           `position:"Query" name:"modelType"`
	FunctionName     string           `position:"Path" name:"functionName"`
	PageSize         requests.Integer `position:"Query" name:"pageSize"`
	FunctionType     string           `position:"Query" name:"functionType"`
	Source           string           `position:"Query" name:"source"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
	PageNumber       requests.Integer `position:"Query" name:"pageNumber"`
}

// ListFunctionInstancesResponse is the response struct for api ListFunctionInstances
type ListFunctionInstancesResponse struct {
	*responses.BaseResponse
	Status     string                              `json:"Status" xml:"Status"`
	HttpCode   int64                               `json:"HttpCode" xml:"HttpCode"`
	TotalCount int64                               `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                              `json:"RequestId" xml:"RequestId"`
	Message    string                              `json:"Message" xml:"Message"`
	Code       string                              `json:"Code" xml:"Code"`
	Latency    int64                               `json:"Latency" xml:"Latency"`
	Result     []ResultItemInListFunctionInstances `json:"Result" xml:"Result"`
}

// CreateListFunctionInstancesRequest creates a request to invoke ListFunctionInstances API
func CreateListFunctionInstancesRequest() (request *ListFunctionInstancesRequest) {
	request = &ListFunctionInstancesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListFunctionInstances", "/v4/openapi/app-groups/[appGroupIdentity]/functions/[functionName]/instances", "", "")
	request.Method = requests.GET
	return
}

// CreateListFunctionInstancesResponse creates a response to parse from ListFunctionInstances response
func CreateListFunctionInstancesResponse() (response *ListFunctionInstancesResponse) {
	response = &ListFunctionInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
