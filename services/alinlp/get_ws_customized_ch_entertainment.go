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

// GetWsCustomizedChEntertainment invokes the alinlp.GetWsCustomizedChEntertainment API synchronously
func (client *Client) GetWsCustomizedChEntertainment(request *GetWsCustomizedChEntertainmentRequest) (response *GetWsCustomizedChEntertainmentResponse, err error) {
	response = CreateGetWsCustomizedChEntertainmentResponse()
	err = client.DoAction(request, response)
	return
}

// GetWsCustomizedChEntertainmentWithChan invokes the alinlp.GetWsCustomizedChEntertainment API asynchronously
func (client *Client) GetWsCustomizedChEntertainmentWithChan(request *GetWsCustomizedChEntertainmentRequest) (<-chan *GetWsCustomizedChEntertainmentResponse, <-chan error) {
	responseChan := make(chan *GetWsCustomizedChEntertainmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWsCustomizedChEntertainment(request)
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

// GetWsCustomizedChEntertainmentWithCallback invokes the alinlp.GetWsCustomizedChEntertainment API asynchronously
func (client *Client) GetWsCustomizedChEntertainmentWithCallback(request *GetWsCustomizedChEntertainmentRequest, callback func(response *GetWsCustomizedChEntertainmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWsCustomizedChEntertainmentResponse
		var err error
		defer close(result)
		response, err = client.GetWsCustomizedChEntertainment(request)
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

// GetWsCustomizedChEntertainmentRequest is the request struct for api GetWsCustomizedChEntertainment
type GetWsCustomizedChEntertainmentRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Body" name:"ServiceCode"`
	TokenizerId string `position:"Body" name:"TokenizerId"`
	Text        string `position:"Body" name:"Text"`
	OutType     string `position:"Body" name:"OutType"`
}

// GetWsCustomizedChEntertainmentResponse is the response struct for api GetWsCustomizedChEntertainment
type GetWsCustomizedChEntertainmentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetWsCustomizedChEntertainmentRequest creates a request to invoke GetWsCustomizedChEntertainment API
func CreateGetWsCustomizedChEntertainmentRequest() (request *GetWsCustomizedChEntertainmentRequest) {
	request = &GetWsCustomizedChEntertainmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetWsCustomizedChEntertainment", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetWsCustomizedChEntertainmentResponse creates a response to parse from GetWsCustomizedChEntertainment response
func CreateGetWsCustomizedChEntertainmentResponse() (response *GetWsCustomizedChEntertainmentResponse) {
	response = &GetWsCustomizedChEntertainmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
