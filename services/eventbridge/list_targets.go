package eventbridge

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

// ListTargets invokes the eventbridge.ListTargets API synchronously
func (client *Client) ListTargets(request *ListTargetsRequest) (response *ListTargetsResponse, err error) {
	response = CreateListTargetsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTargetsWithChan invokes the eventbridge.ListTargets API asynchronously
func (client *Client) ListTargetsWithChan(request *ListTargetsRequest) (<-chan *ListTargetsResponse, <-chan error) {
	responseChan := make(chan *ListTargetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTargets(request)
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

// ListTargetsWithCallback invokes the eventbridge.ListTargets API asynchronously
func (client *Client) ListTargetsWithCallback(request *ListTargetsRequest, callback func(response *ListTargetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTargetsResponse
		var err error
		defer close(result)
		response, err = client.ListTargets(request)
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

// ListTargetsRequest is the request struct for api ListTargets
type ListTargetsRequest struct {
	*requests.RpcRequest
	RuleName     string           `position:"Query" name:"RuleName"`
	EventBusName string           `position:"Query" name:"EventBusName"`
	NextToken    string           `position:"Query" name:"NextToken"`
	Limit        requests.Integer `position:"Query" name:"Limit"`
	Arn          string           `position:"Query" name:"Arn"`
}

// ListTargetsResponse is the response struct for api ListTargets
type ListTargetsResponse struct {
	*responses.BaseResponse
	Message   string            `json:"Message" xml:"Message"`
	RequestId string            `json:"RequestId" xml:"RequestId"`
	Code      string            `json:"Code" xml:"Code"`
	Success   bool              `json:"Success" xml:"Success"`
	Data      DataInListTargets `json:"Data" xml:"Data"`
}

// CreateListTargetsRequest creates a request to invoke ListTargets API
func CreateListTargetsRequest() (request *ListTargetsRequest) {
	request = &ListTargetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eventbridge", "2020-04-01", "ListTargets", "", "")
	request.Method = requests.POST
	return
}

// CreateListTargetsResponse creates a response to parse from ListTargets response
func CreateListTargetsResponse() (response *ListTargetsResponse) {
	response = &ListTargetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}