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

// GetQueryOptimizeDataTop invokes the das.GetQueryOptimizeDataTop API synchronously
func (client *Client) GetQueryOptimizeDataTop(request *GetQueryOptimizeDataTopRequest) (response *GetQueryOptimizeDataTopResponse, err error) {
	response = CreateGetQueryOptimizeDataTopResponse()
	err = client.DoAction(request, response)
	return
}

// GetQueryOptimizeDataTopWithChan invokes the das.GetQueryOptimizeDataTop API asynchronously
func (client *Client) GetQueryOptimizeDataTopWithChan(request *GetQueryOptimizeDataTopRequest) (<-chan *GetQueryOptimizeDataTopResponse, <-chan error) {
	responseChan := make(chan *GetQueryOptimizeDataTopResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQueryOptimizeDataTop(request)
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

// GetQueryOptimizeDataTopWithCallback invokes the das.GetQueryOptimizeDataTop API asynchronously
func (client *Client) GetQueryOptimizeDataTopWithCallback(request *GetQueryOptimizeDataTopRequest, callback func(response *GetQueryOptimizeDataTopResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQueryOptimizeDataTopResponse
		var err error
		defer close(result)
		response, err = client.GetQueryOptimizeDataTop(request)
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

// GetQueryOptimizeDataTopRequest is the request struct for api GetQueryOptimizeDataTop
type GetQueryOptimizeDataTopRequest struct {
	*requests.RpcRequest
	Type           string `position:"Query" name:"Type"`
	TagNames       string `position:"Query" name:"TagNames"`
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	Engine         string `position:"Query" name:"Engine"`
	InstanceIds    string `position:"Query" name:"InstanceIds"`
	Time           string `position:"Query" name:"Time"`
	Region         string `position:"Query" name:"Region"`
}

// GetQueryOptimizeDataTopResponse is the response struct for api GetQueryOptimizeDataTop
type GetQueryOptimizeDataTopResponse struct {
	*responses.BaseResponse
	Code      string                        `json:"Code" xml:"Code"`
	Message   string                        `json:"Message" xml:"Message"`
	RequestId string                        `json:"RequestId" xml:"RequestId"`
	Success   string                        `json:"Success" xml:"Success"`
	Data      DataInGetQueryOptimizeDataTop `json:"Data" xml:"Data"`
}

// CreateGetQueryOptimizeDataTopRequest creates a request to invoke GetQueryOptimizeDataTop API
func CreateGetQueryOptimizeDataTopRequest() (request *GetQueryOptimizeDataTopRequest) {
	request = &GetQueryOptimizeDataTopRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetQueryOptimizeDataTop", "", "")
	request.Method = requests.GET
	return
}

// CreateGetQueryOptimizeDataTopResponse creates a response to parse from GetQueryOptimizeDataTop response
func CreateGetQueryOptimizeDataTopResponse() (response *GetQueryOptimizeDataTopResponse) {
	response = &GetQueryOptimizeDataTopResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
