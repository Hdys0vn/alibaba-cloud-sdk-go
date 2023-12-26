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

// GetEcChGeneral invokes the alinlp.GetEcChGeneral API synchronously
func (client *Client) GetEcChGeneral(request *GetEcChGeneralRequest) (response *GetEcChGeneralResponse, err error) {
	response = CreateGetEcChGeneralResponse()
	err = client.DoAction(request, response)
	return
}

// GetEcChGeneralWithChan invokes the alinlp.GetEcChGeneral API asynchronously
func (client *Client) GetEcChGeneralWithChan(request *GetEcChGeneralRequest) (<-chan *GetEcChGeneralResponse, <-chan error) {
	responseChan := make(chan *GetEcChGeneralResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEcChGeneral(request)
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

// GetEcChGeneralWithCallback invokes the alinlp.GetEcChGeneral API asynchronously
func (client *Client) GetEcChGeneralWithCallback(request *GetEcChGeneralRequest, callback func(response *GetEcChGeneralResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEcChGeneralResponse
		var err error
		defer close(result)
		response, err = client.GetEcChGeneral(request)
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

// GetEcChGeneralRequest is the request struct for api GetEcChGeneral
type GetEcChGeneralRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Text        string `position:"Body" name:"Text"`
}

// GetEcChGeneralResponse is the response struct for api GetEcChGeneral
type GetEcChGeneralResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetEcChGeneralRequest creates a request to invoke GetEcChGeneral API
func CreateGetEcChGeneralRequest() (request *GetEcChGeneralRequest) {
	request = &GetEcChGeneralRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetEcChGeneral", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetEcChGeneralResponse creates a response to parse from GetEcChGeneral response
func CreateGetEcChGeneralResponse() (response *GetEcChGeneralResponse) {
	response = &GetEcChGeneralResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
