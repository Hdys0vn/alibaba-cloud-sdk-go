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

// GetWeChGeneral invokes the alinlp.GetWeChGeneral API synchronously
func (client *Client) GetWeChGeneral(request *GetWeChGeneralRequest) (response *GetWeChGeneralResponse, err error) {
	response = CreateGetWeChGeneralResponse()
	err = client.DoAction(request, response)
	return
}

// GetWeChGeneralWithChan invokes the alinlp.GetWeChGeneral API asynchronously
func (client *Client) GetWeChGeneralWithChan(request *GetWeChGeneralRequest) (<-chan *GetWeChGeneralResponse, <-chan error) {
	responseChan := make(chan *GetWeChGeneralResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWeChGeneral(request)
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

// GetWeChGeneralWithCallback invokes the alinlp.GetWeChGeneral API asynchronously
func (client *Client) GetWeChGeneralWithCallback(request *GetWeChGeneralRequest, callback func(response *GetWeChGeneralResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWeChGeneralResponse
		var err error
		defer close(result)
		response, err = client.GetWeChGeneral(request)
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

// GetWeChGeneralRequest is the request struct for api GetWeChGeneral
type GetWeChGeneralRequest struct {
	*requests.RpcRequest
	Type        string `position:"Body" name:"Type"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Size        string `position:"Body" name:"Size"`
	Text        string `position:"Body" name:"Text"`
	Operation   string `position:"Body" name:"Operation"`
}

// GetWeChGeneralResponse is the response struct for api GetWeChGeneral
type GetWeChGeneralResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetWeChGeneralRequest creates a request to invoke GetWeChGeneral API
func CreateGetWeChGeneralRequest() (request *GetWeChGeneralRequest) {
	request = &GetWeChGeneralRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetWeChGeneral", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetWeChGeneralResponse creates a response to parse from GetWeChGeneral response
func CreateGetWeChGeneralResponse() (response *GetWeChGeneralResponse) {
	response = &GetWeChGeneralResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
