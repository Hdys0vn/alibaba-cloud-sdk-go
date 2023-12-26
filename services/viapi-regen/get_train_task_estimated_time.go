package viapi_regen

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

// GetTrainTaskEstimatedTime invokes the viapi_regen.GetTrainTaskEstimatedTime API synchronously
func (client *Client) GetTrainTaskEstimatedTime(request *GetTrainTaskEstimatedTimeRequest) (response *GetTrainTaskEstimatedTimeResponse, err error) {
	response = CreateGetTrainTaskEstimatedTimeResponse()
	err = client.DoAction(request, response)
	return
}

// GetTrainTaskEstimatedTimeWithChan invokes the viapi_regen.GetTrainTaskEstimatedTime API asynchronously
func (client *Client) GetTrainTaskEstimatedTimeWithChan(request *GetTrainTaskEstimatedTimeRequest) (<-chan *GetTrainTaskEstimatedTimeResponse, <-chan error) {
	responseChan := make(chan *GetTrainTaskEstimatedTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTrainTaskEstimatedTime(request)
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

// GetTrainTaskEstimatedTimeWithCallback invokes the viapi_regen.GetTrainTaskEstimatedTime API asynchronously
func (client *Client) GetTrainTaskEstimatedTimeWithCallback(request *GetTrainTaskEstimatedTimeRequest, callback func(response *GetTrainTaskEstimatedTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTrainTaskEstimatedTimeResponse
		var err error
		defer close(result)
		response, err = client.GetTrainTaskEstimatedTime(request)
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

// GetTrainTaskEstimatedTimeRequest is the request struct for api GetTrainTaskEstimatedTime
type GetTrainTaskEstimatedTimeRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Body" name:"Id"`
}

// GetTrainTaskEstimatedTimeResponse is the response struct for api GetTrainTaskEstimatedTime
type GetTrainTaskEstimatedTimeResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetTrainTaskEstimatedTimeRequest creates a request to invoke GetTrainTaskEstimatedTime API
func CreateGetTrainTaskEstimatedTimeRequest() (request *GetTrainTaskEstimatedTimeRequest) {
	request = &GetTrainTaskEstimatedTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "GetTrainTaskEstimatedTime", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTrainTaskEstimatedTimeResponse creates a response to parse from GetTrainTaskEstimatedTime response
func CreateGetTrainTaskEstimatedTimeResponse() (response *GetTrainTaskEstimatedTimeResponse) {
	response = &GetTrainTaskEstimatedTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
