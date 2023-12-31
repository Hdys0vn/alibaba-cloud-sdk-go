package smartag

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

// DescribeGrantSagRules invokes the smartag.DescribeGrantSagRules API synchronously
func (client *Client) DescribeGrantSagRules(request *DescribeGrantSagRulesRequest) (response *DescribeGrantSagRulesResponse, err error) {
	response = CreateDescribeGrantSagRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGrantSagRulesWithChan invokes the smartag.DescribeGrantSagRules API asynchronously
func (client *Client) DescribeGrantSagRulesWithChan(request *DescribeGrantSagRulesRequest) (<-chan *DescribeGrantSagRulesResponse, <-chan error) {
	responseChan := make(chan *DescribeGrantSagRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGrantSagRules(request)
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

// DescribeGrantSagRulesWithCallback invokes the smartag.DescribeGrantSagRules API asynchronously
func (client *Client) DescribeGrantSagRulesWithCallback(request *DescribeGrantSagRulesRequest, callback func(response *DescribeGrantSagRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGrantSagRulesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGrantSagRules(request)
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

// DescribeGrantSagRulesRequest is the request struct for api DescribeGrantSagRules
type DescribeGrantSagRulesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// DescribeGrantSagRulesResponse is the response struct for api DescribeGrantSagRules
type DescribeGrantSagRulesResponse struct {
	*responses.BaseResponse
	TotalCount int                               `json:"TotalCount" xml:"TotalCount"`
	PageSize   int                               `json:"PageSize" xml:"PageSize"`
	RequestId  string                            `json:"RequestId" xml:"RequestId"`
	PageNumber int                               `json:"PageNumber" xml:"PageNumber"`
	GrantRules GrantRulesInDescribeGrantSagRules `json:"GrantRules" xml:"GrantRules"`
}

// CreateDescribeGrantSagRulesRequest creates a request to invoke DescribeGrantSagRules API
func CreateDescribeGrantSagRulesRequest() (request *DescribeGrantSagRulesRequest) {
	request = &DescribeGrantSagRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeGrantSagRules", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGrantSagRulesResponse creates a response to parse from DescribeGrantSagRules response
func CreateDescribeGrantSagRulesResponse() (response *DescribeGrantSagRulesResponse) {
	response = &DescribeGrantSagRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
