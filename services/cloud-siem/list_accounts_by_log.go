package cloud_siem

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

// ListAccountsByLog invokes the cloud_siem.ListAccountsByLog API synchronously
func (client *Client) ListAccountsByLog(request *ListAccountsByLogRequest) (response *ListAccountsByLogResponse, err error) {
	response = CreateListAccountsByLogResponse()
	err = client.DoAction(request, response)
	return
}

// ListAccountsByLogWithChan invokes the cloud_siem.ListAccountsByLog API asynchronously
func (client *Client) ListAccountsByLogWithChan(request *ListAccountsByLogRequest) (<-chan *ListAccountsByLogResponse, <-chan error) {
	responseChan := make(chan *ListAccountsByLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccountsByLog(request)
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

// ListAccountsByLogWithCallback invokes the cloud_siem.ListAccountsByLog API asynchronously
func (client *Client) ListAccountsByLogWithCallback(request *ListAccountsByLogRequest, callback func(response *ListAccountsByLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAccountsByLogResponse
		var err error
		defer close(result)
		response, err = client.ListAccountsByLog(request)
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

// ListAccountsByLogRequest is the request struct for api ListAccountsByLog
type ListAccountsByLogRequest struct {
	*requests.RpcRequest
	CloudCode string    `position:"Body" name:"CloudCode"`
	LogCodes  *[]string `position:"Body" name:"LogCodes"  type:"Repeated"`
	ProdCode  string    `position:"Body" name:"ProdCode"`
}

// ListAccountsByLogResponse is the response struct for api ListAccountsByLog
type ListAccountsByLogResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListAccountsByLogRequest creates a request to invoke ListAccountsByLog API
func CreateListAccountsByLogRequest() (request *ListAccountsByLogRequest) {
	request = &ListAccountsByLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "ListAccountsByLog", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAccountsByLogResponse creates a response to parse from ListAccountsByLog response
func CreateListAccountsByLogResponse() (response *ListAccountsByLogResponse) {
	response = &ListAccountsByLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
