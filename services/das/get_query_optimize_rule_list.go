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

// GetQueryOptimizeRuleList invokes the das.GetQueryOptimizeRuleList API synchronously
func (client *Client) GetQueryOptimizeRuleList(request *GetQueryOptimizeRuleListRequest) (response *GetQueryOptimizeRuleListResponse, err error) {
	response = CreateGetQueryOptimizeRuleListResponse()
	err = client.DoAction(request, response)
	return
}

// GetQueryOptimizeRuleListWithChan invokes the das.GetQueryOptimizeRuleList API asynchronously
func (client *Client) GetQueryOptimizeRuleListWithChan(request *GetQueryOptimizeRuleListRequest) (<-chan *GetQueryOptimizeRuleListResponse, <-chan error) {
	responseChan := make(chan *GetQueryOptimizeRuleListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQueryOptimizeRuleList(request)
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

// GetQueryOptimizeRuleListWithCallback invokes the das.GetQueryOptimizeRuleList API asynchronously
func (client *Client) GetQueryOptimizeRuleListWithCallback(request *GetQueryOptimizeRuleListRequest, callback func(response *GetQueryOptimizeRuleListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQueryOptimizeRuleListResponse
		var err error
		defer close(result)
		response, err = client.GetQueryOptimizeRuleList(request)
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

// GetQueryOptimizeRuleListRequest is the request struct for api GetQueryOptimizeRuleList
type GetQueryOptimizeRuleListRequest struct {
	*requests.RpcRequest
	TagNames       string `position:"Query" name:"TagNames"`
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	Engine         string `position:"Query" name:"Engine"`
	InstanceIds    string `position:"Query" name:"InstanceIds"`
	Region         string `position:"Query" name:"Region"`
}

// GetQueryOptimizeRuleListResponse is the response struct for api GetQueryOptimizeRuleList
type GetQueryOptimizeRuleListResponse struct {
	*responses.BaseResponse
	Code      string                         `json:"Code" xml:"Code"`
	Message   string                         `json:"Message" xml:"Message"`
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Success   string                         `json:"Success" xml:"Success"`
	Data      DataInGetQueryOptimizeRuleList `json:"Data" xml:"Data"`
}

// CreateGetQueryOptimizeRuleListRequest creates a request to invoke GetQueryOptimizeRuleList API
func CreateGetQueryOptimizeRuleListRequest() (request *GetQueryOptimizeRuleListRequest) {
	request = &GetQueryOptimizeRuleListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetQueryOptimizeRuleList", "", "")
	request.Method = requests.GET
	return
}

// CreateGetQueryOptimizeRuleListResponse creates a response to parse from GetQueryOptimizeRuleList response
func CreateGetQueryOptimizeRuleListResponse() (response *GetQueryOptimizeRuleListResponse) {
	response = &GetQueryOptimizeRuleListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
