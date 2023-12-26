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

// GetTcChGeneral invokes the alinlp.GetTcChGeneral API synchronously
func (client *Client) GetTcChGeneral(request *GetTcChGeneralRequest) (response *GetTcChGeneralResponse, err error) {
	response = CreateGetTcChGeneralResponse()
	err = client.DoAction(request, response)
	return
}

// GetTcChGeneralWithChan invokes the alinlp.GetTcChGeneral API asynchronously
func (client *Client) GetTcChGeneralWithChan(request *GetTcChGeneralRequest) (<-chan *GetTcChGeneralResponse, <-chan error) {
	responseChan := make(chan *GetTcChGeneralResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTcChGeneral(request)
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

// GetTcChGeneralWithCallback invokes the alinlp.GetTcChGeneral API asynchronously
func (client *Client) GetTcChGeneralWithCallback(request *GetTcChGeneralRequest, callback func(response *GetTcChGeneralResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTcChGeneralResponse
		var err error
		defer close(result)
		response, err = client.GetTcChGeneral(request)
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

// GetTcChGeneralRequest is the request struct for api GetTcChGeneral
type GetTcChGeneralRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Text        string `position:"Body" name:"Text"`
}

// GetTcChGeneralResponse is the response struct for api GetTcChGeneral
type GetTcChGeneralResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetTcChGeneralRequest creates a request to invoke GetTcChGeneral API
func CreateGetTcChGeneralRequest() (request *GetTcChGeneralRequest) {
	request = &GetTcChGeneralRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetTcChGeneral", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTcChGeneralResponse creates a response to parse from GetTcChGeneral response
func CreateGetTcChGeneralResponse() (response *GetTcChGeneralResponse) {
	response = &GetTcChGeneralResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}