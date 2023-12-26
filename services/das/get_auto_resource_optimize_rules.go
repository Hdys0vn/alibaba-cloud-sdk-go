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

// GetAutoResourceOptimizeRules invokes the das.GetAutoResourceOptimizeRules API synchronously
func (client *Client) GetAutoResourceOptimizeRules(request *GetAutoResourceOptimizeRulesRequest) (response *GetAutoResourceOptimizeRulesResponse, err error) {
	response = CreateGetAutoResourceOptimizeRulesResponse()
	err = client.DoAction(request, response)
	return
}

// GetAutoResourceOptimizeRulesWithChan invokes the das.GetAutoResourceOptimizeRules API asynchronously
func (client *Client) GetAutoResourceOptimizeRulesWithChan(request *GetAutoResourceOptimizeRulesRequest) (<-chan *GetAutoResourceOptimizeRulesResponse, <-chan error) {
	responseChan := make(chan *GetAutoResourceOptimizeRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAutoResourceOptimizeRules(request)
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

// GetAutoResourceOptimizeRulesWithCallback invokes the das.GetAutoResourceOptimizeRules API asynchronously
func (client *Client) GetAutoResourceOptimizeRulesWithCallback(request *GetAutoResourceOptimizeRulesRequest, callback func(response *GetAutoResourceOptimizeRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAutoResourceOptimizeRulesResponse
		var err error
		defer close(result)
		response, err = client.GetAutoResourceOptimizeRules(request)
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

// GetAutoResourceOptimizeRulesRequest is the request struct for api GetAutoResourceOptimizeRules
type GetAutoResourceOptimizeRulesRequest struct {
	*requests.RpcRequest
	ConsoleContext string `position:"Query" name:"ConsoleContext"`
	InstanceIds    string `position:"Query" name:"InstanceIds"`
}

// GetAutoResourceOptimizeRulesResponse is the response struct for api GetAutoResourceOptimizeRules
type GetAutoResourceOptimizeRulesResponse struct {
	*responses.BaseResponse
	Code      int64  `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAutoResourceOptimizeRulesRequest creates a request to invoke GetAutoResourceOptimizeRules API
func CreateGetAutoResourceOptimizeRulesRequest() (request *GetAutoResourceOptimizeRulesRequest) {
	request = &GetAutoResourceOptimizeRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetAutoResourceOptimizeRules", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAutoResourceOptimizeRulesResponse creates a response to parse from GetAutoResourceOptimizeRules response
func CreateGetAutoResourceOptimizeRulesResponse() (response *GetAutoResourceOptimizeRulesResponse) {
	response = &GetAutoResourceOptimizeRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}