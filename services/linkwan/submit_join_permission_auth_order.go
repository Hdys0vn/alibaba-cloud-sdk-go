package linkwan

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

// SubmitJoinPermissionAuthOrder invokes the linkwan.SubmitJoinPermissionAuthOrder API synchronously
func (client *Client) SubmitJoinPermissionAuthOrder(request *SubmitJoinPermissionAuthOrderRequest) (response *SubmitJoinPermissionAuthOrderResponse, err error) {
	response = CreateSubmitJoinPermissionAuthOrderResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitJoinPermissionAuthOrderWithChan invokes the linkwan.SubmitJoinPermissionAuthOrder API asynchronously
func (client *Client) SubmitJoinPermissionAuthOrderWithChan(request *SubmitJoinPermissionAuthOrderRequest) (<-chan *SubmitJoinPermissionAuthOrderResponse, <-chan error) {
	responseChan := make(chan *SubmitJoinPermissionAuthOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitJoinPermissionAuthOrder(request)
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

// SubmitJoinPermissionAuthOrderWithCallback invokes the linkwan.SubmitJoinPermissionAuthOrder API asynchronously
func (client *Client) SubmitJoinPermissionAuthOrderWithCallback(request *SubmitJoinPermissionAuthOrderRequest, callback func(response *SubmitJoinPermissionAuthOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitJoinPermissionAuthOrderResponse
		var err error
		defer close(result)
		response, err = client.SubmitJoinPermissionAuthOrder(request)
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

// SubmitJoinPermissionAuthOrderRequest is the request struct for api SubmitJoinPermissionAuthOrder
type SubmitJoinPermissionAuthOrderRequest struct {
	*requests.RpcRequest
	JoinPermissionId string `position:"Query" name:"JoinPermissionId"`
	RenterAliyunId   string `position:"Query" name:"RenterAliyunId"`
	ApiProduct       string `position:"Body" name:"ApiProduct"`
	ApiRevision      string `position:"Body" name:"ApiRevision"`
}

// SubmitJoinPermissionAuthOrderResponse is the response struct for api SubmitJoinPermissionAuthOrder
type SubmitJoinPermissionAuthOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateSubmitJoinPermissionAuthOrderRequest creates a request to invoke SubmitJoinPermissionAuthOrder API
func CreateSubmitJoinPermissionAuthOrderRequest() (request *SubmitJoinPermissionAuthOrderRequest) {
	request = &SubmitJoinPermissionAuthOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "SubmitJoinPermissionAuthOrder", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitJoinPermissionAuthOrderResponse creates a response to parse from SubmitJoinPermissionAuthOrder response
func CreateSubmitJoinPermissionAuthOrderResponse() (response *SubmitJoinPermissionAuthOrderResponse) {
	response = &SubmitJoinPermissionAuthOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
