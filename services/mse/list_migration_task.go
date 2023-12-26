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

// ListMigrationTask invokes the mse.ListMigrationTask API synchronously
func (client *Client) ListMigrationTask(request *ListMigrationTaskRequest) (response *ListMigrationTaskResponse, err error) {
	response = CreateListMigrationTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ListMigrationTaskWithChan invokes the mse.ListMigrationTask API asynchronously
func (client *Client) ListMigrationTaskWithChan(request *ListMigrationTaskRequest) (<-chan *ListMigrationTaskResponse, <-chan error) {
	responseChan := make(chan *ListMigrationTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMigrationTask(request)
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

// ListMigrationTaskWithCallback invokes the mse.ListMigrationTask API asynchronously
func (client *Client) ListMigrationTaskWithCallback(request *ListMigrationTaskRequest, callback func(response *ListMigrationTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMigrationTaskResponse
		var err error
		defer close(result)
		response, err = client.ListMigrationTask(request)
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

// ListMigrationTaskRequest is the request struct for api ListMigrationTask
type ListMigrationTaskRequest struct {
	*requests.RpcRequest
	MseSessionId       string           `position:"Query" name:"MseSessionId"`
	PageNum            requests.Integer `position:"Query" name:"PageNum"`
	RequestPars        string           `position:"Query" name:"RequestPars"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
	OriginInstanceName string           `position:"Query" name:"OriginInstanceName"`
	AcceptLanguage     string           `position:"Query" name:"AcceptLanguage"`
}

// ListMigrationTaskResponse is the response struct for api ListMigrationTask
type ListMigrationTaskResponse struct {
	*responses.BaseResponse
	HttpCode   string     `json:"HttpCode" xml:"HttpCode"`
	PageSize   int64      `json:"PageSize" xml:"PageSize"`
	PageNumber int64      `json:"PageNumber" xml:"PageNumber"`
	TotalCount int64      `json:"TotalCount" xml:"TotalCount"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Message    string     `json:"Message" xml:"Message"`
	ErrorCode  string     `json:"ErrorCode" xml:"ErrorCode"`
	Success    bool       `json:"Success" xml:"Success"`
	Data       []DataItem `json:"Data" xml:"Data"`
}

// CreateListMigrationTaskRequest creates a request to invoke ListMigrationTask API
func CreateListMigrationTaskRequest() (request *ListMigrationTaskRequest) {
	request = &ListMigrationTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListMigrationTask", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListMigrationTaskResponse creates a response to parse from ListMigrationTask response
func CreateListMigrationTaskResponse() (response *ListMigrationTaskResponse) {
	response = &ListMigrationTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}