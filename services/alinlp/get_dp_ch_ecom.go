package alinlp

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

// GetDpChEcom invokes the alinlp.GetDpChEcom API synchronously
func (client *Client) GetDpChEcom(request *GetDpChEcomRequest) (response *GetDpChEcomResponse, err error) {
	response = CreateGetDpChEcomResponse()
	err = client.DoAction(request, response)
	return
}

// GetDpChEcomWithChan invokes the alinlp.GetDpChEcom API asynchronously
func (client *Client) GetDpChEcomWithChan(request *GetDpChEcomRequest) (<-chan *GetDpChEcomResponse, <-chan error) {
	responseChan := make(chan *GetDpChEcomResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDpChEcom(request)
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

// GetDpChEcomWithCallback invokes the alinlp.GetDpChEcom API asynchronously
func (client *Client) GetDpChEcomWithCallback(request *GetDpChEcomRequest, callback func(response *GetDpChEcomResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDpChEcomResponse
		var err error
		defer close(result)
		response, err = client.GetDpChEcom(request)
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

// GetDpChEcomRequest is the request struct for api GetDpChEcom
type GetDpChEcomRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Text        string `position:"Body" name:"Text"`
}

// GetDpChEcomResponse is the response struct for api GetDpChEcom
type GetDpChEcomResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetDpChEcomRequest creates a request to invoke GetDpChEcom API
func CreateGetDpChEcomRequest() (request *GetDpChEcomRequest) {
	request = &GetDpChEcomRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetDpChEcom", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDpChEcomResponse creates a response to parse from GetDpChEcom response
func CreateGetDpChEcomResponse() (response *GetDpChEcomResponse) {
	response = &GetDpChEcomResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
