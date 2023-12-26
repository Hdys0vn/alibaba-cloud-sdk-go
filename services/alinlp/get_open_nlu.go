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

// GetOpenNLU invokes the alinlp.GetOpenNLU API synchronously
func (client *Client) GetOpenNLU(request *GetOpenNLURequest) (response *GetOpenNLUResponse, err error) {
	response = CreateGetOpenNLUResponse()
	err = client.DoAction(request, response)
	return
}

// GetOpenNLUWithChan invokes the alinlp.GetOpenNLU API asynchronously
func (client *Client) GetOpenNLUWithChan(request *GetOpenNLURequest) (<-chan *GetOpenNLUResponse, <-chan error) {
	responseChan := make(chan *GetOpenNLUResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOpenNLU(request)
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

// GetOpenNLUWithCallback invokes the alinlp.GetOpenNLU API asynchronously
func (client *Client) GetOpenNLUWithCallback(request *GetOpenNLURequest, callback func(response *GetOpenNLUResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOpenNLUResponse
		var err error
		defer close(result)
		response, err = client.GetOpenNLU(request)
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

// GetOpenNLURequest is the request struct for api GetOpenNLU
type GetOpenNLURequest struct {
	*requests.RpcRequest
	Sentence    string `position:"Body" name:"Sentence"`
	Business    string `position:"Query" name:"Business"`
	Labels      string `position:"Body" name:"Labels"`
	Task        string `position:"Body" name:"Task"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Examples    string `position:"Body" name:"Examples"`
}

// GetOpenNLUResponse is the response struct for api GetOpenNLU
type GetOpenNLUResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetOpenNLURequest creates a request to invoke GetOpenNLU API
func CreateGetOpenNLURequest() (request *GetOpenNLURequest) {
	request = &GetOpenNLURequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetOpenNLU", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOpenNLUResponse creates a response to parse from GetOpenNLU response
func CreateGetOpenNLUResponse() (response *GetOpenNLUResponse) {
	response = &GetOpenNLUResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
