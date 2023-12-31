package tesladam

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

// ActionDiskCheck invokes the tesladam.ActionDiskCheck API synchronously
// api document: https://help.aliyun.com/api/tesladam/actiondiskcheck.html
func (client *Client) ActionDiskCheck(request *ActionDiskCheckRequest) (response *ActionDiskCheckResponse, err error) {
	response = CreateActionDiskCheckResponse()
	err = client.DoAction(request, response)
	return
}

// ActionDiskCheckWithChan invokes the tesladam.ActionDiskCheck API asynchronously
// api document: https://help.aliyun.com/api/tesladam/actiondiskcheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ActionDiskCheckWithChan(request *ActionDiskCheckRequest) (<-chan *ActionDiskCheckResponse, <-chan error) {
	responseChan := make(chan *ActionDiskCheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ActionDiskCheck(request)
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

// ActionDiskCheckWithCallback invokes the tesladam.ActionDiskCheck API asynchronously
// api document: https://help.aliyun.com/api/tesladam/actiondiskcheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ActionDiskCheckWithCallback(request *ActionDiskCheckRequest, callback func(response *ActionDiskCheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ActionDiskCheckResponse
		var err error
		defer close(result)
		response, err = client.ActionDiskCheck(request)
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

// ActionDiskCheckRequest is the request struct for api ActionDiskCheck
type ActionDiskCheckRequest struct {
	*requests.RpcRequest
	DiskMount string `position:"Query" name:"DiskMount"`
	Ip        string `position:"Query" name:"Ip"`
}

// ActionDiskCheckResponse is the response struct for api ActionDiskCheck
type ActionDiskCheckResponse struct {
	*responses.BaseResponse
	Status  bool   `json:"Status" xml:"Status"`
	Message string `json:"Message" xml:"Message"`
	Result  string `json:"Result" xml:"Result"`
}

// CreateActionDiskCheckRequest creates a request to invoke ActionDiskCheck API
func CreateActionDiskCheckRequest() (request *ActionDiskCheckRequest) {
	request = &ActionDiskCheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaDam", "2018-01-18", "ActionDiskCheck", "tesladam", "openAPI")
	return
}

// CreateActionDiskCheckResponse creates a response to parse from ActionDiskCheck response
func CreateActionDiskCheckResponse() (response *ActionDiskCheckResponse) {
	response = &ActionDiskCheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
