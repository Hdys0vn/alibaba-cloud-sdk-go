package aliyuncvc

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

// ListEvaluations invokes the aliyuncvc.ListEvaluations API synchronously
func (client *Client) ListEvaluations(request *ListEvaluationsRequest) (response *ListEvaluationsResponse, err error) {
	response = CreateListEvaluationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListEvaluationsWithChan invokes the aliyuncvc.ListEvaluations API asynchronously
func (client *Client) ListEvaluationsWithChan(request *ListEvaluationsRequest) (<-chan *ListEvaluationsResponse, <-chan error) {
	responseChan := make(chan *ListEvaluationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEvaluations(request)
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

// ListEvaluationsWithCallback invokes the aliyuncvc.ListEvaluations API asynchronously
func (client *Client) ListEvaluationsWithCallback(request *ListEvaluationsRequest, callback func(response *ListEvaluationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEvaluationsResponse
		var err error
		defer close(result)
		response, err = client.ListEvaluations(request)
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

// ListEvaluationsRequest is the request struct for api ListEvaluations
type ListEvaluationsRequest struct {
	*requests.RpcRequest
}

// ListEvaluationsResponse is the response struct for api ListEvaluations
type ListEvaluationsResponse struct {
	*responses.BaseResponse
	UserEvaluation string `json:"UserEvaluation" xml:"UserEvaluation"`
	ErrorCode      int    `json:"ErrorCode" xml:"ErrorCode"`
	Message        string `json:"Message" xml:"Message"`
	Success        bool   `json:"Success" xml:"Success"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateListEvaluationsRequest creates a request to invoke ListEvaluations API
func CreateListEvaluationsRequest() (request *ListEvaluationsRequest) {
	request = &ListEvaluationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ListEvaluations", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListEvaluationsResponse creates a response to parse from ListEvaluations response
func CreateListEvaluationsResponse() (response *ListEvaluationsResponse) {
	response = &ListEvaluationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
