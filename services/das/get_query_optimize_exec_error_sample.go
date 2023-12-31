package das

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

// GetQueryOptimizeExecErrorSample invokes the das.GetQueryOptimizeExecErrorSample API synchronously
func (client *Client) GetQueryOptimizeExecErrorSample(request *GetQueryOptimizeExecErrorSampleRequest) (response *GetQueryOptimizeExecErrorSampleResponse, err error) {
	response = CreateGetQueryOptimizeExecErrorSampleResponse()
	err = client.DoAction(request, response)
	return
}

// GetQueryOptimizeExecErrorSampleWithChan invokes the das.GetQueryOptimizeExecErrorSample API asynchronously
func (client *Client) GetQueryOptimizeExecErrorSampleWithChan(request *GetQueryOptimizeExecErrorSampleRequest) (<-chan *GetQueryOptimizeExecErrorSampleResponse, <-chan error) {
	responseChan := make(chan *GetQueryOptimizeExecErrorSampleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQueryOptimizeExecErrorSample(request)
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

// GetQueryOptimizeExecErrorSampleWithCallback invokes the das.GetQueryOptimizeExecErrorSample API asynchronously
func (client *Client) GetQueryOptimizeExecErrorSampleWithCallback(request *GetQueryOptimizeExecErrorSampleRequest, callback func(response *GetQueryOptimizeExecErrorSampleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQueryOptimizeExecErrorSampleResponse
		var err error
		defer close(result)
		response, err = client.GetQueryOptimizeExecErrorSample(request)
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

// GetQueryOptimizeExecErrorSampleRequest is the request struct for api GetQueryOptimizeExecErrorSample
type GetQueryOptimizeExecErrorSampleRequest struct {
	*requests.RpcRequest
	SqlId          string `position:"Query" name:"SqlId"`
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	Engine         string `position:"Query" name:"Engine"`
	Time           string `position:"Query" name:"Time"`
}

// GetQueryOptimizeExecErrorSampleResponse is the response struct for api GetQueryOptimizeExecErrorSample
type GetQueryOptimizeExecErrorSampleResponse struct {
	*responses.BaseResponse
	Code      string                                `json:"Code" xml:"Code"`
	Message   string                                `json:"Message" xml:"Message"`
	RequestId string                                `json:"RequestId" xml:"RequestId"`
	Success   string                                `json:"Success" xml:"Success"`
	Data      DataInGetQueryOptimizeExecErrorSample `json:"Data" xml:"Data"`
}

// CreateGetQueryOptimizeExecErrorSampleRequest creates a request to invoke GetQueryOptimizeExecErrorSample API
func CreateGetQueryOptimizeExecErrorSampleRequest() (request *GetQueryOptimizeExecErrorSampleRequest) {
	request = &GetQueryOptimizeExecErrorSampleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetQueryOptimizeExecErrorSample", "", "")
	request.Method = requests.GET
	return
}

// CreateGetQueryOptimizeExecErrorSampleResponse creates a response to parse from GetQueryOptimizeExecErrorSample response
func CreateGetQueryOptimizeExecErrorSampleResponse() (response *GetQueryOptimizeExecErrorSampleResponse) {
	response = &GetQueryOptimizeExecErrorSampleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
