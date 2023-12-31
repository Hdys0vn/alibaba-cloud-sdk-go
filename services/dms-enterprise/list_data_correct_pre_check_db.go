package dms_enterprise

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

// ListDataCorrectPreCheckDB invokes the dms_enterprise.ListDataCorrectPreCheckDB API synchronously
func (client *Client) ListDataCorrectPreCheckDB(request *ListDataCorrectPreCheckDBRequest) (response *ListDataCorrectPreCheckDBResponse, err error) {
	response = CreateListDataCorrectPreCheckDBResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataCorrectPreCheckDBWithChan invokes the dms_enterprise.ListDataCorrectPreCheckDB API asynchronously
func (client *Client) ListDataCorrectPreCheckDBWithChan(request *ListDataCorrectPreCheckDBRequest) (<-chan *ListDataCorrectPreCheckDBResponse, <-chan error) {
	responseChan := make(chan *ListDataCorrectPreCheckDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataCorrectPreCheckDB(request)
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

// ListDataCorrectPreCheckDBWithCallback invokes the dms_enterprise.ListDataCorrectPreCheckDB API asynchronously
func (client *Client) ListDataCorrectPreCheckDBWithCallback(request *ListDataCorrectPreCheckDBRequest, callback func(response *ListDataCorrectPreCheckDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataCorrectPreCheckDBResponse
		var err error
		defer close(result)
		response, err = client.ListDataCorrectPreCheckDB(request)
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

// ListDataCorrectPreCheckDBRequest is the request struct for api ListDataCorrectPreCheckDB
type ListDataCorrectPreCheckDBRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	OrderId    requests.Integer `position:"Query" name:"OrderId"`
}

// ListDataCorrectPreCheckDBResponse is the response struct for api ListDataCorrectPreCheckDB
type ListDataCorrectPreCheckDBResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	ErrorMessage   string       `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string       `json:"ErrorCode" xml:"ErrorCode"`
	PreCheckDBList []PreCheckDB `json:"PreCheckDBList" xml:"PreCheckDBList"`
}

// CreateListDataCorrectPreCheckDBRequest creates a request to invoke ListDataCorrectPreCheckDB API
func CreateListDataCorrectPreCheckDBRequest() (request *ListDataCorrectPreCheckDBRequest) {
	request = &ListDataCorrectPreCheckDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListDataCorrectPreCheckDB", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDataCorrectPreCheckDBResponse creates a response to parse from ListDataCorrectPreCheckDB response
func CreateListDataCorrectPreCheckDBResponse() (response *ListDataCorrectPreCheckDBResponse) {
	response = &ListDataCorrectPreCheckDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
