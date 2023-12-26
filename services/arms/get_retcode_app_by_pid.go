package arms

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

// GetRetcodeAppByPid invokes the arms.GetRetcodeAppByPid API synchronously
func (client *Client) GetRetcodeAppByPid(request *GetRetcodeAppByPidRequest) (response *GetRetcodeAppByPidResponse, err error) {
	response = CreateGetRetcodeAppByPidResponse()
	err = client.DoAction(request, response)
	return
}

// GetRetcodeAppByPidWithChan invokes the arms.GetRetcodeAppByPid API asynchronously
func (client *Client) GetRetcodeAppByPidWithChan(request *GetRetcodeAppByPidRequest) (<-chan *GetRetcodeAppByPidResponse, <-chan error) {
	responseChan := make(chan *GetRetcodeAppByPidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRetcodeAppByPid(request)
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

// GetRetcodeAppByPidWithCallback invokes the arms.GetRetcodeAppByPid API asynchronously
func (client *Client) GetRetcodeAppByPidWithCallback(request *GetRetcodeAppByPidRequest, callback func(response *GetRetcodeAppByPidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRetcodeAppByPidResponse
		var err error
		defer close(result)
		response, err = client.GetRetcodeAppByPid(request)
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

// GetRetcodeAppByPidRequest is the request struct for api GetRetcodeAppByPid
type GetRetcodeAppByPidRequest struct {
	*requests.RpcRequest
	Pid string `position:"Query" name:"Pid"`
}

// GetRetcodeAppByPidResponse is the response struct for api GetRetcodeAppByPid
type GetRetcodeAppByPidResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	RetcodeApp RetcodeApp `json:"RetcodeApp" xml:"RetcodeApp"`
}

// CreateGetRetcodeAppByPidRequest creates a request to invoke GetRetcodeAppByPid API
func CreateGetRetcodeAppByPidRequest() (request *GetRetcodeAppByPidRequest) {
	request = &GetRetcodeAppByPidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetRetcodeAppByPid", "arms", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetRetcodeAppByPidResponse creates a response to parse from GetRetcodeAppByPid response
func CreateGetRetcodeAppByPidResponse() (response *GetRetcodeAppByPidResponse) {
	response = &GetRetcodeAppByPidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}