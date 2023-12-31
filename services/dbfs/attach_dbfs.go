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

// AttachDbfs invokes the dbfs.AttachDbfs API synchronously
func (client *Client) AttachDbfs(request *AttachDbfsRequest) (response *AttachDbfsResponse, err error) {
	response = CreateAttachDbfsResponse()
	err = client.DoAction(request, response)
	return
}

// AttachDbfsWithChan invokes the dbfs.AttachDbfs API asynchronously
func (client *Client) AttachDbfsWithChan(request *AttachDbfsRequest) (<-chan *AttachDbfsResponse, <-chan error) {
	responseChan := make(chan *AttachDbfsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachDbfs(request)
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

// AttachDbfsWithCallback invokes the dbfs.AttachDbfs API asynchronously
func (client *Client) AttachDbfsWithCallback(request *AttachDbfsRequest, callback func(response *AttachDbfsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachDbfsResponse
		var err error
		defer close(result)
		response, err = client.AttachDbfs(request)
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

// AttachDbfsRequest is the request struct for api AttachDbfs
type AttachDbfsRequest struct {
	*requests.RpcRequest
	ECSInstanceId string `position:"Query" name:"ECSInstanceId"`
	AttachPoint   string `position:"Query" name:"AttachPoint"`
	ServerUrl     string `position:"Query" name:"ServerUrl"`
	FsId          string `position:"Query" name:"FsId"`
	AttachMode    string `position:"Query" name:"AttachMode"`
}

// AttachDbfsResponse is the response struct for api AttachDbfs
type AttachDbfsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachDbfsRequest creates a request to invoke AttachDbfs API
func CreateAttachDbfsRequest() (request *AttachDbfsRequest) {
	request = &AttachDbfsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "AttachDbfs", "dbfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachDbfsResponse creates a response to parse from AttachDbfs response
func CreateAttachDbfsResponse() (response *AttachDbfsResponse) {
	response = &AttachDbfsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
