package mse

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

// ListApplicationsWithTagRules invokes the mse.ListApplicationsWithTagRules API synchronously
func (client *Client) ListApplicationsWithTagRules(request *ListApplicationsWithTagRulesRequest) (response *ListApplicationsWithTagRulesResponse, err error) {
	response = CreateListApplicationsWithTagRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListApplicationsWithTagRulesWithChan invokes the mse.ListApplicationsWithTagRules API asynchronously
func (client *Client) ListApplicationsWithTagRulesWithChan(request *ListApplicationsWithTagRulesRequest) (<-chan *ListApplicationsWithTagRulesResponse, <-chan error) {
	responseChan := make(chan *ListApplicationsWithTagRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApplicationsWithTagRules(request)
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

// ListApplicationsWithTagRulesWithCallback invokes the mse.ListApplicationsWithTagRules API asynchronously
func (client *Client) ListApplicationsWithTagRulesWithCallback(request *ListApplicationsWithTagRulesRequest, callback func(response *ListApplicationsWithTagRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApplicationsWithTagRulesResponse
		var err error
		defer close(result)
		response, err = client.ListApplicationsWithTagRules(request)
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

// ListApplicationsWithTagRulesRequest is the request struct for api ListApplicationsWithTagRules
type ListApplicationsWithTagRulesRequest struct {
	*requests.RpcRequest
	MseSessionId   string           `position:"Query" name:"MseSessionId"`
	Source         string           `position:"Query" name:"Source"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	AppName        string           `position:"Query" name:"AppName"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	AppId          string           `position:"Query" name:"AppId"`
	Namespace      string           `position:"Query" name:"Namespace"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
	Region         string           `position:"Query" name:"Region"`
}

// ListApplicationsWithTagRulesResponse is the response struct for api ListApplicationsWithTagRules
type ListApplicationsWithTagRulesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListApplicationsWithTagRulesRequest creates a request to invoke ListApplicationsWithTagRules API
func CreateListApplicationsWithTagRulesRequest() (request *ListApplicationsWithTagRulesRequest) {
	request = &ListApplicationsWithTagRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListApplicationsWithTagRules", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListApplicationsWithTagRulesResponse creates a response to parse from ListApplicationsWithTagRules response
func CreateListApplicationsWithTagRulesResponse() (response *ListApplicationsWithTagRulesResponse) {
	response = &ListApplicationsWithTagRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}