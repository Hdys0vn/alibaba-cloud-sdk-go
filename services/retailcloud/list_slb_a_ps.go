package retailcloud

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

// ListSlbAPs invokes the retailcloud.ListSlbAPs API synchronously
func (client *Client) ListSlbAPs(request *ListSlbAPsRequest) (response *ListSlbAPsResponse, err error) {
	response = CreateListSlbAPsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSlbAPsWithChan invokes the retailcloud.ListSlbAPs API asynchronously
func (client *Client) ListSlbAPsWithChan(request *ListSlbAPsRequest) (<-chan *ListSlbAPsResponse, <-chan error) {
	responseChan := make(chan *ListSlbAPsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSlbAPs(request)
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

// ListSlbAPsWithCallback invokes the retailcloud.ListSlbAPs API asynchronously
func (client *Client) ListSlbAPsWithCallback(request *ListSlbAPsRequest, callback func(response *ListSlbAPsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSlbAPsResponse
		var err error
		defer close(result)
		response, err = client.ListSlbAPs(request)
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

// ListSlbAPsRequest is the request struct for api ListSlbAPs
type ListSlbAPsRequest struct {
	*requests.RpcRequest
	ProtocolList *[]string        `position:"Body" name:"ProtocolList"  type:"Repeated"`
	SlbId        string           `position:"Query" name:"SlbId"`
	AppId        requests.Integer `position:"Query" name:"AppId"`
	Name         string           `position:"Query" name:"Name"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	EnvId        requests.Integer `position:"Query" name:"EnvId"`
	NetworkMode  string           `position:"Query" name:"NetworkMode"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
}

// ListSlbAPsResponse is the response struct for api ListSlbAPs
type ListSlbAPsResponse struct {
	*responses.BaseResponse
	Code       int             `json:"Code" xml:"Code"`
	ErrorMsg   string          `json:"ErrorMsg" xml:"ErrorMsg"`
	PageNumber int             `json:"PageNumber" xml:"PageNumber"`
	PageSize   int             `json:"PageSize" xml:"PageSize"`
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	TotalCount int64           `json:"TotalCount" xml:"TotalCount"`
	Data       []SlbAPInstance `json:"Data" xml:"Data"`
}

// CreateListSlbAPsRequest creates a request to invoke ListSlbAPs API
func CreateListSlbAPsRequest() (request *ListSlbAPsRequest) {
	request = &ListSlbAPsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListSlbAPs", "", "")
	request.Method = requests.POST
	return
}

// CreateListSlbAPsResponse creates a response to parse from ListSlbAPs response
func CreateListSlbAPsResponse() (response *ListSlbAPsResponse) {
	response = &ListSlbAPsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
