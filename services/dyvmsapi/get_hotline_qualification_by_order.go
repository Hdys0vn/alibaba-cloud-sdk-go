package dyvmsapi

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

// GetHotlineQualificationByOrder invokes the dyvmsapi.GetHotlineQualificationByOrder API synchronously
func (client *Client) GetHotlineQualificationByOrder(request *GetHotlineQualificationByOrderRequest) (response *GetHotlineQualificationByOrderResponse, err error) {
	response = CreateGetHotlineQualificationByOrderResponse()
	err = client.DoAction(request, response)
	return
}

// GetHotlineQualificationByOrderWithChan invokes the dyvmsapi.GetHotlineQualificationByOrder API asynchronously
func (client *Client) GetHotlineQualificationByOrderWithChan(request *GetHotlineQualificationByOrderRequest) (<-chan *GetHotlineQualificationByOrderResponse, <-chan error) {
	responseChan := make(chan *GetHotlineQualificationByOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetHotlineQualificationByOrder(request)
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

// GetHotlineQualificationByOrderWithCallback invokes the dyvmsapi.GetHotlineQualificationByOrder API asynchronously
func (client *Client) GetHotlineQualificationByOrderWithCallback(request *GetHotlineQualificationByOrderRequest, callback func(response *GetHotlineQualificationByOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetHotlineQualificationByOrderResponse
		var err error
		defer close(result)
		response, err = client.GetHotlineQualificationByOrder(request)
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

// GetHotlineQualificationByOrderRequest is the request struct for api GetHotlineQualificationByOrder
type GetHotlineQualificationByOrderRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OrderId              string           `position:"Query" name:"OrderId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// GetHotlineQualificationByOrderResponse is the response struct for api GetHotlineQualificationByOrder
type GetHotlineQualificationByOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetHotlineQualificationByOrderRequest creates a request to invoke GetHotlineQualificationByOrder API
func CreateGetHotlineQualificationByOrderRequest() (request *GetHotlineQualificationByOrderRequest) {
	request = &GetHotlineQualificationByOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "GetHotlineQualificationByOrder", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetHotlineQualificationByOrderResponse creates a response to parse from GetHotlineQualificationByOrder response
func CreateGetHotlineQualificationByOrderResponse() (response *GetHotlineQualificationByOrderResponse) {
	response = &GetHotlineQualificationByOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}