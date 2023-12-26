package dataworks_public

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

// ListNodeIO invokes the dataworks_public.ListNodeIO API synchronously
func (client *Client) ListNodeIO(request *ListNodeIORequest) (response *ListNodeIOResponse, err error) {
	response = CreateListNodeIOResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodeIOWithChan invokes the dataworks_public.ListNodeIO API asynchronously
func (client *Client) ListNodeIOWithChan(request *ListNodeIORequest) (<-chan *ListNodeIOResponse, <-chan error) {
	responseChan := make(chan *ListNodeIOResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodeIO(request)
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

// ListNodeIOWithCallback invokes the dataworks_public.ListNodeIO API asynchronously
func (client *Client) ListNodeIOWithCallback(request *ListNodeIORequest, callback func(response *ListNodeIOResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodeIOResponse
		var err error
		defer close(result)
		response, err = client.ListNodeIO(request)
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

// ListNodeIORequest is the request struct for api ListNodeIO
type ListNodeIORequest struct {
	*requests.RpcRequest
	ProjectEnv string           `position:"Body" name:"ProjectEnv"`
	NodeId     requests.Integer `position:"Body" name:"NodeId"`
	IoType     string           `position:"Body" name:"IoType"`
}

// ListNodeIOResponse is the response struct for api ListNodeIO
type ListNodeIOResponse struct {
	*responses.BaseResponse
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string     `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string     `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool       `json:"Success" xml:"Success"`
	Data           []DataItem `json:"Data" xml:"Data"`
}

// CreateListNodeIORequest creates a request to invoke ListNodeIO API
func CreateListNodeIORequest() (request *ListNodeIORequest) {
	request = &ListNodeIORequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListNodeIO", "", "")
	request.Method = requests.POST
	return
}

// CreateListNodeIOResponse creates a response to parse from ListNodeIO response
func CreateListNodeIOResponse() (response *ListNodeIOResponse) {
	response = &ListNodeIOResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}