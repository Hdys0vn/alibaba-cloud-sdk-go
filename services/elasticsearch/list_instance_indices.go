package elasticsearch

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

// ListInstanceIndices invokes the elasticsearch.ListInstanceIndices API synchronously
func (client *Client) ListInstanceIndices(request *ListInstanceIndicesRequest) (response *ListInstanceIndicesResponse, err error) {
	response = CreateListInstanceIndicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstanceIndicesWithChan invokes the elasticsearch.ListInstanceIndices API asynchronously
func (client *Client) ListInstanceIndicesWithChan(request *ListInstanceIndicesRequest) (<-chan *ListInstanceIndicesResponse, <-chan error) {
	responseChan := make(chan *ListInstanceIndicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstanceIndices(request)
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

// ListInstanceIndicesWithCallback invokes the elasticsearch.ListInstanceIndices API asynchronously
func (client *Client) ListInstanceIndicesWithCallback(request *ListInstanceIndicesRequest, callback func(response *ListInstanceIndicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstanceIndicesResponse
		var err error
		defer close(result)
		response, err = client.ListInstanceIndices(request)
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

// ListInstanceIndicesRequest is the request struct for api ListInstanceIndices
type ListInstanceIndicesRequest struct {
	*requests.RoaRequest
	All         requests.Boolean `position:"Query" name:"all"`
	InstanceId  string           `position:"Path" name:"InstanceId"`
	IsManaged   requests.Boolean `position:"Query" name:"isManaged"`
	Size        requests.Integer `position:"Query" name:"size"`
	Name        string           `position:"Query" name:"name"`
	Page        requests.Integer `position:"Query" name:"page"`
	IsOpenstore requests.Boolean `position:"Query" name:"isOpenstore"`
}

// ListInstanceIndicesResponse is the response struct for api ListInstanceIndices
type ListInstanceIndicesResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Headers   Headers      `json:"Headers" xml:"Headers"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateListInstanceIndicesRequest creates a request to invoke ListInstanceIndices API
func CreateListInstanceIndicesRequest() (request *ListInstanceIndicesRequest) {
	request = &ListInstanceIndicesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListInstanceIndices", "/openapi/instances/[InstanceId]/indices", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListInstanceIndicesResponse creates a response to parse from ListInstanceIndices response
func CreateListInstanceIndicesResponse() (response *ListInstanceIndicesResponse) {
	response = &ListInstanceIndicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
