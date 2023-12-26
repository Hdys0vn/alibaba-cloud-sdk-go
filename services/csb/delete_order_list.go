package csb

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

// DeleteOrderList invokes the csb.DeleteOrderList API synchronously
// api document: https://help.aliyun.com/api/csb/deleteorderlist.html
func (client *Client) DeleteOrderList(request *DeleteOrderListRequest) (response *DeleteOrderListResponse, err error) {
	response = CreateDeleteOrderListResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteOrderListWithChan invokes the csb.DeleteOrderList API asynchronously
// api document: https://help.aliyun.com/api/csb/deleteorderlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteOrderListWithChan(request *DeleteOrderListRequest) (<-chan *DeleteOrderListResponse, <-chan error) {
	responseChan := make(chan *DeleteOrderListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteOrderList(request)
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

// DeleteOrderListWithCallback invokes the csb.DeleteOrderList API asynchronously
// api document: https://help.aliyun.com/api/csb/deleteorderlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteOrderListWithCallback(request *DeleteOrderListRequest, callback func(response *DeleteOrderListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteOrderListResponse
		var err error
		defer close(result)
		response, err = client.DeleteOrderList(request)
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

// DeleteOrderListRequest is the request struct for api DeleteOrderList
type DeleteOrderListRequest struct {
	*requests.RpcRequest
	Data string `position:"Body" name:"Data"`
}

// DeleteOrderListResponse is the response struct for api DeleteOrderList
type DeleteOrderListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteOrderListRequest creates a request to invoke DeleteOrderList API
func CreateDeleteOrderListRequest() (request *DeleteOrderListRequest) {
	request = &DeleteOrderListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "DeleteOrderList", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteOrderListResponse creates a response to parse from DeleteOrderList response
func CreateDeleteOrderListResponse() (response *DeleteOrderListResponse) {
	response = &DeleteOrderListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}