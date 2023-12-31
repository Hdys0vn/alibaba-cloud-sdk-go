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

// GetSaChGeneral invokes the alinlp.GetSaChGeneral API synchronously
func (client *Client) GetSaChGeneral(request *GetSaChGeneralRequest) (response *GetSaChGeneralResponse, err error) {
	response = CreateGetSaChGeneralResponse()
	err = client.DoAction(request, response)
	return
}

// GetSaChGeneralWithChan invokes the alinlp.GetSaChGeneral API asynchronously
func (client *Client) GetSaChGeneralWithChan(request *GetSaChGeneralRequest) (<-chan *GetSaChGeneralResponse, <-chan error) {
	responseChan := make(chan *GetSaChGeneralResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSaChGeneral(request)
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

// GetSaChGeneralWithCallback invokes the alinlp.GetSaChGeneral API asynchronously
func (client *Client) GetSaChGeneralWithCallback(request *GetSaChGeneralRequest, callback func(response *GetSaChGeneralResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSaChGeneralResponse
		var err error
		defer close(result)
		response, err = client.GetSaChGeneral(request)
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

// GetSaChGeneralRequest is the request struct for api GetSaChGeneral
type GetSaChGeneralRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Text        string `position:"Body" name:"Text"`
}

// GetSaChGeneralResponse is the response struct for api GetSaChGeneral
type GetSaChGeneralResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetSaChGeneralRequest creates a request to invoke GetSaChGeneral API
func CreateGetSaChGeneralRequest() (request *GetSaChGeneralRequest) {
	request = &GetSaChGeneralRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetSaChGeneral", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSaChGeneralResponse creates a response to parse from GetSaChGeneral response
func CreateGetSaChGeneralResponse() (response *GetSaChGeneralResponse) {
	response = &GetSaChGeneralResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
