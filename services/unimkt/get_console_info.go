package unimkt

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

// GetConsoleInfo invokes the unimkt.GetConsoleInfo API synchronously
func (client *Client) GetConsoleInfo(request *GetConsoleInfoRequest) (response *GetConsoleInfoResponse, err error) {
	response = CreateGetConsoleInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetConsoleInfoWithChan invokes the unimkt.GetConsoleInfo API asynchronously
func (client *Client) GetConsoleInfoWithChan(request *GetConsoleInfoRequest) (<-chan *GetConsoleInfoResponse, <-chan error) {
	responseChan := make(chan *GetConsoleInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetConsoleInfo(request)
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

// GetConsoleInfoWithCallback invokes the unimkt.GetConsoleInfo API asynchronously
func (client *Client) GetConsoleInfoWithCallback(request *GetConsoleInfoRequest, callback func(response *GetConsoleInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetConsoleInfoResponse
		var err error
		defer close(result)
		response, err = client.GetConsoleInfo(request)
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

// GetConsoleInfoRequest is the request struct for api GetConsoleInfo
type GetConsoleInfoRequest struct {
	*requests.RpcRequest
	Message string `position:"Body" name:"Message"`
}

// GetConsoleInfoResponse is the response struct for api GetConsoleInfo
type GetConsoleInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    int    `json:"Status" xml:"Status"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetConsoleInfoRequest creates a request to invoke GetConsoleInfo API
func CreateGetConsoleInfoRequest() (request *GetConsoleInfoRequest) {
	request = &GetConsoleInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-07", "GetConsoleInfo", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetConsoleInfoResponse creates a response to parse from GetConsoleInfo response
func CreateGetConsoleInfoResponse() (response *GetConsoleInfoResponse) {
	response = &GetConsoleInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
