package retailadvqa_public

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

// ListDigitalTemplates invokes the retailadvqa_public.ListDigitalTemplates API synchronously
func (client *Client) ListDigitalTemplates(request *ListDigitalTemplatesRequest) (response *ListDigitalTemplatesResponse, err error) {
	response = CreateListDigitalTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDigitalTemplatesWithChan invokes the retailadvqa_public.ListDigitalTemplates API asynchronously
func (client *Client) ListDigitalTemplatesWithChan(request *ListDigitalTemplatesRequest) (<-chan *ListDigitalTemplatesResponse, <-chan error) {
	responseChan := make(chan *ListDigitalTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDigitalTemplates(request)
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

// ListDigitalTemplatesWithCallback invokes the retailadvqa_public.ListDigitalTemplates API asynchronously
func (client *Client) ListDigitalTemplatesWithCallback(request *ListDigitalTemplatesRequest, callback func(response *ListDigitalTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDigitalTemplatesResponse
		var err error
		defer close(result)
		response, err = client.ListDigitalTemplates(request)
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

// ListDigitalTemplatesRequest is the request struct for api ListDigitalTemplates
type ListDigitalTemplatesRequest struct {
	*requests.RpcRequest
	AccessId    string           `position:"Query" name:"AccessId"`
	SmsSign     string           `position:"Query" name:"SmsSign"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	TenantId    string           `position:"Query" name:"TenantId"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	Keyword     string           `position:"Query" name:"Keyword"`
	ColumnName  string           `position:"Query" name:"ColumnName"`
	Order       string           `position:"Query" name:"Order"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// ListDigitalTemplatesResponse is the response struct for api ListDigitalTemplates
type ListDigitalTemplatesResponse struct {
	*responses.BaseResponse
	ErrorCode string                     `json:"ErrorCode" xml:"ErrorCode"`
	ErrorDesc string                     `json:"ErrorDesc" xml:"ErrorDesc"`
	Success   bool                       `json:"Success" xml:"Success"`
	TraceId   string                     `json:"TraceId" xml:"TraceId"`
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Data      DataInListDigitalTemplates `json:"Data" xml:"Data"`
}

// CreateListDigitalTemplatesRequest creates a request to invoke ListDigitalTemplates API
func CreateListDigitalTemplatesRequest() (request *ListDigitalTemplatesRequest) {
	request = &ListDigitalTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "ListDigitalTemplates", "", "")
	request.Method = requests.GET
	return
}

// CreateListDigitalTemplatesResponse creates a response to parse from ListDigitalTemplates response
func CreateListDigitalTemplatesResponse() (response *ListDigitalTemplatesResponse) {
	response = &ListDigitalTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
