package aliyuncvc

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

// CustomLayout invokes the aliyuncvc.CustomLayout API synchronously
func (client *Client) CustomLayout(request *CustomLayoutRequest) (response *CustomLayoutResponse, err error) {
	response = CreateCustomLayoutResponse()
	err = client.DoAction(request, response)
	return
}

// CustomLayoutWithChan invokes the aliyuncvc.CustomLayout API asynchronously
func (client *Client) CustomLayoutWithChan(request *CustomLayoutRequest) (<-chan *CustomLayoutResponse, <-chan error) {
	responseChan := make(chan *CustomLayoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CustomLayout(request)
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

// CustomLayoutWithCallback invokes the aliyuncvc.CustomLayout API asynchronously
func (client *Client) CustomLayoutWithCallback(request *CustomLayoutRequest, callback func(response *CustomLayoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CustomLayoutResponse
		var err error
		defer close(result)
		response, err = client.CustomLayout(request)
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

// CustomLayoutRequest is the request struct for api CustomLayout
type CustomLayoutRequest struct {
	*requests.RpcRequest
	LiveUUID   string `position:"Body" name:"LiveUUID"`
	LayoutInfo string `position:"Body" name:"LayoutInfo"`
}

// CustomLayoutResponse is the response struct for api CustomLayout
type CustomLayoutResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCustomLayoutRequest creates a request to invoke CustomLayout API
func CreateCustomLayoutRequest() (request *CustomLayoutRequest) {
	request = &CustomLayoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "CustomLayout", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCustomLayoutResponse creates a response to parse from CustomLayout response
func CreateCustomLayoutResponse() (response *CustomLayoutResponse) {
	response = &CustomLayoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
