package mse

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

// ListClusterTypes invokes the mse.ListClusterTypes API synchronously
func (client *Client) ListClusterTypes(request *ListClusterTypesRequest) (response *ListClusterTypesResponse, err error) {
	response = CreateListClusterTypesResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterTypesWithChan invokes the mse.ListClusterTypes API asynchronously
func (client *Client) ListClusterTypesWithChan(request *ListClusterTypesRequest) (<-chan *ListClusterTypesResponse, <-chan error) {
	responseChan := make(chan *ListClusterTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterTypes(request)
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

// ListClusterTypesWithCallback invokes the mse.ListClusterTypes API asynchronously
func (client *Client) ListClusterTypesWithCallback(request *ListClusterTypesRequest, callback func(response *ListClusterTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterTypesResponse
		var err error
		defer close(result)
		response, err = client.ListClusterTypes(request)
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

// ListClusterTypesRequest is the request struct for api ListClusterTypes
type ListClusterTypesRequest struct {
	*requests.RpcRequest
	MseSessionId   string `position:"Query" name:"MseSessionId"`
	ConnectType    string `position:"Query" name:"ConnectType"`
	MseVersion     string `position:"Query" name:"MseVersion"`
	AcceptLanguage string `position:"Query" name:"AcceptLanguage"`
}

// ListClusterTypesResponse is the response struct for api ListClusterTypes
type ListClusterTypesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	ErrorCode      string     `json:"ErrorCode" xml:"ErrorCode"`
	Code           int        `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	DynamicMessage string     `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           []DataItem `json:"Data" xml:"Data"`
}

// CreateListClusterTypesRequest creates a request to invoke ListClusterTypes API
func CreateListClusterTypesRequest() (request *ListClusterTypesRequest) {
	request = &ListClusterTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListClusterTypes", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListClusterTypesResponse creates a response to parse from ListClusterTypes response
func CreateListClusterTypesResponse() (response *ListClusterTypesResponse) {
	response = &ListClusterTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
