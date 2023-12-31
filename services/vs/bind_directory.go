package vs

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

// BindDirectory invokes the vs.BindDirectory API synchronously
func (client *Client) BindDirectory(request *BindDirectoryRequest) (response *BindDirectoryResponse, err error) {
	response = CreateBindDirectoryResponse()
	err = client.DoAction(request, response)
	return
}

// BindDirectoryWithChan invokes the vs.BindDirectory API asynchronously
func (client *Client) BindDirectoryWithChan(request *BindDirectoryRequest) (<-chan *BindDirectoryResponse, <-chan error) {
	responseChan := make(chan *BindDirectoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindDirectory(request)
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

// BindDirectoryWithCallback invokes the vs.BindDirectory API asynchronously
func (client *Client) BindDirectoryWithCallback(request *BindDirectoryRequest, callback func(response *BindDirectoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindDirectoryResponse
		var err error
		defer close(result)
		response, err = client.BindDirectory(request)
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

// BindDirectoryRequest is the request struct for api BindDirectory
type BindDirectoryRequest struct {
	*requests.RpcRequest
	DirectoryId string           `position:"Query" name:"DirectoryId"`
	ShowLog     string           `position:"Query" name:"ShowLog"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	DeviceId    string           `position:"Query" name:"DeviceId"`
}

// BindDirectoryResponse is the response struct for api BindDirectory
type BindDirectoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBindDirectoryRequest creates a request to invoke BindDirectory API
func CreateBindDirectoryRequest() (request *BindDirectoryRequest) {
	request = &BindDirectoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "BindDirectory", "", "")
	request.Method = requests.POST
	return
}

// CreateBindDirectoryResponse creates a response to parse from BindDirectory response
func CreateBindDirectoryResponse() (response *BindDirectoryResponse) {
	response = &BindDirectoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
