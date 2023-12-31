package dbfs

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

// DetachDbfs invokes the dbfs.DetachDbfs API synchronously
func (client *Client) DetachDbfs(request *DetachDbfsRequest) (response *DetachDbfsResponse, err error) {
	response = CreateDetachDbfsResponse()
	err = client.DoAction(request, response)
	return
}

// DetachDbfsWithChan invokes the dbfs.DetachDbfs API asynchronously
func (client *Client) DetachDbfsWithChan(request *DetachDbfsRequest) (<-chan *DetachDbfsResponse, <-chan error) {
	responseChan := make(chan *DetachDbfsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachDbfs(request)
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

// DetachDbfsWithCallback invokes the dbfs.DetachDbfs API asynchronously
func (client *Client) DetachDbfsWithCallback(request *DetachDbfsRequest, callback func(response *DetachDbfsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachDbfsResponse
		var err error
		defer close(result)
		response, err = client.DetachDbfs(request)
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

// DetachDbfsRequest is the request struct for api DetachDbfs
type DetachDbfsRequest struct {
	*requests.RpcRequest
	ECSInstanceId string `position:"Query" name:"ECSInstanceId"`
	FsId          string `position:"Query" name:"FsId"`
}

// DetachDbfsResponse is the response struct for api DetachDbfs
type DetachDbfsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachDbfsRequest creates a request to invoke DetachDbfs API
func CreateDetachDbfsRequest() (request *DetachDbfsRequest) {
	request = &DetachDbfsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "DetachDbfs", "dbfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachDbfsResponse creates a response to parse from DetachDbfs response
func CreateDetachDbfsResponse() (response *DetachDbfsResponse) {
	response = &DetachDbfsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
