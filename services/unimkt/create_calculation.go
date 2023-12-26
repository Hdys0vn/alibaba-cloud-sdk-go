package unimkt

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

// CreateCalculation invokes the unimkt.CreateCalculation API synchronously
func (client *Client) CreateCalculation(request *CreateCalculationRequest) (response *CreateCalculationResponse, err error) {
	response = CreateCreateCalculationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCalculationWithChan invokes the unimkt.CreateCalculation API asynchronously
func (client *Client) CreateCalculationWithChan(request *CreateCalculationRequest) (<-chan *CreateCalculationResponse, <-chan error) {
	responseChan := make(chan *CreateCalculationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCalculation(request)
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

// CreateCalculationWithCallback invokes the unimkt.CreateCalculation API asynchronously
func (client *Client) CreateCalculationWithCallback(request *CreateCalculationRequest, callback func(response *CreateCalculationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCalculationResponse
		var err error
		defer close(result)
		response, err = client.CreateCalculation(request)
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

// CreateCalculationRequest is the request struct for api CreateCalculation
type CreateCalculationRequest struct {
	*requests.RpcRequest
	CloudCodeUserId string `position:"Query" name:"CloudCodeUserId"`
	ClientToken     string `position:"Query" name:"ClientToken"`
	QueryString     string `position:"Query" name:"QueryString"`
}

// CreateCalculationResponse is the response struct for api CreateCalculation
type CreateCalculationResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Result       string `json:"Result" xml:"Result"`
}

// CreateCreateCalculationRequest creates a request to invoke CreateCalculation API
func CreateCreateCalculationRequest() (request *CreateCalculationRequest) {
	request = &CreateCalculationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "CreateCalculation", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCalculationResponse creates a response to parse from CreateCalculation response
func CreateCreateCalculationResponse() (response *CreateCalculationResponse) {
	response = &CreateCalculationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
