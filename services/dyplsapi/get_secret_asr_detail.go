package dyplsapi

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

// GetSecretAsrDetail invokes the dyplsapi.GetSecretAsrDetail API synchronously
func (client *Client) GetSecretAsrDetail(request *GetSecretAsrDetailRequest) (response *GetSecretAsrDetailResponse, err error) {
	response = CreateGetSecretAsrDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetSecretAsrDetailWithChan invokes the dyplsapi.GetSecretAsrDetail API asynchronously
func (client *Client) GetSecretAsrDetailWithChan(request *GetSecretAsrDetailRequest) (<-chan *GetSecretAsrDetailResponse, <-chan error) {
	responseChan := make(chan *GetSecretAsrDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSecretAsrDetail(request)
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

// GetSecretAsrDetailWithCallback invokes the dyplsapi.GetSecretAsrDetail API asynchronously
func (client *Client) GetSecretAsrDetailWithCallback(request *GetSecretAsrDetailRequest, callback func(response *GetSecretAsrDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSecretAsrDetailResponse
		var err error
		defer close(result)
		response, err = client.GetSecretAsrDetail(request)
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

// GetSecretAsrDetailRequest is the request struct for api GetSecretAsrDetail
type GetSecretAsrDetailRequest struct {
	*requests.RpcRequest
	CallId               string           `position:"Query" name:"CallId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	CallTime             string           `position:"Query" name:"CallTime"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
}

// GetSecretAsrDetailResponse is the response struct for api GetSecretAsrDetail
type GetSecretAsrDetailResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetSecretAsrDetailRequest creates a request to invoke GetSecretAsrDetail API
func CreateGetSecretAsrDetailRequest() (request *GetSecretAsrDetailRequest) {
	request = &GetSecretAsrDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "GetSecretAsrDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateGetSecretAsrDetailResponse creates a response to parse from GetSecretAsrDetail response
func CreateGetSecretAsrDetailResponse() (response *GetSecretAsrDetailResponse) {
	response = &GetSecretAsrDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
