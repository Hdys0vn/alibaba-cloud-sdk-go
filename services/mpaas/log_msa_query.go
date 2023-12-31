package mpaas

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

// LogMsaQuery invokes the mpaas.LogMsaQuery API synchronously
func (client *Client) LogMsaQuery(request *LogMsaQueryRequest) (response *LogMsaQueryResponse, err error) {
	response = CreateLogMsaQueryResponse()
	err = client.DoAction(request, response)
	return
}

// LogMsaQueryWithChan invokes the mpaas.LogMsaQuery API asynchronously
func (client *Client) LogMsaQueryWithChan(request *LogMsaQueryRequest) (<-chan *LogMsaQueryResponse, <-chan error) {
	responseChan := make(chan *LogMsaQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LogMsaQuery(request)
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

// LogMsaQueryWithCallback invokes the mpaas.LogMsaQuery API asynchronously
func (client *Client) LogMsaQueryWithCallback(request *LogMsaQueryRequest, callback func(response *LogMsaQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LogMsaQueryResponse
		var err error
		defer close(result)
		response, err = client.LogMsaQuery(request)
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

// LogMsaQueryRequest is the request struct for api LogMsaQuery
type LogMsaQueryRequest struct {
	*requests.RpcRequest
	OnexFlag    string `position:"Body" name:"OnexFlag"`
	TenantId    string `position:"Body" name:"TenantId"`
	Id          string `position:"Body" name:"Id"`
	AppId       string `position:"Body" name:"AppId"`
	WorkspaceId string `position:"Body" name:"WorkspaceId"`
}

// LogMsaQueryResponse is the response struct for api LogMsaQuery
type LogMsaQueryResponse struct {
	*responses.BaseResponse
	ResultMessage string                     `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string                     `json:"ResultCode" xml:"ResultCode"`
	RequestId     string                     `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContentInLogMsaQuery `json:"ResultContent" xml:"ResultContent"`
}

// CreateLogMsaQueryRequest creates a request to invoke LogMsaQuery API
func CreateLogMsaQueryRequest() (request *LogMsaQueryRequest) {
	request = &LogMsaQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "LogMsaQuery", "", "")
	request.Method = requests.POST
	return
}

// CreateLogMsaQueryResponse creates a response to parse from LogMsaQuery response
func CreateLogMsaQueryResponse() (response *LogMsaQueryResponse) {
	response = &LogMsaQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
