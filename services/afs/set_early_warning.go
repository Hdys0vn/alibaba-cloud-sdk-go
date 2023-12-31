package afs

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

// SetEarlyWarning invokes the afs.SetEarlyWarning API synchronously
// api document: https://help.aliyun.com/api/afs/setearlywarning.html
func (client *Client) SetEarlyWarning(request *SetEarlyWarningRequest) (response *SetEarlyWarningResponse, err error) {
	response = CreateSetEarlyWarningResponse()
	err = client.DoAction(request, response)
	return
}

// SetEarlyWarningWithChan invokes the afs.SetEarlyWarning API asynchronously
// api document: https://help.aliyun.com/api/afs/setearlywarning.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetEarlyWarningWithChan(request *SetEarlyWarningRequest) (<-chan *SetEarlyWarningResponse, <-chan error) {
	responseChan := make(chan *SetEarlyWarningResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetEarlyWarning(request)
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

// SetEarlyWarningWithCallback invokes the afs.SetEarlyWarning API asynchronously
// api document: https://help.aliyun.com/api/afs/setearlywarning.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetEarlyWarningWithCallback(request *SetEarlyWarningRequest, callback func(response *SetEarlyWarningResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetEarlyWarningResponse
		var err error
		defer close(result)
		response, err = client.SetEarlyWarning(request)
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

// SetEarlyWarningRequest is the request struct for api SetEarlyWarning
type SetEarlyWarningRequest struct {
	*requests.RpcRequest
	TimeEnd   string           `position:"Query" name:"TimeEnd"`
	WarnOpen  requests.Boolean `position:"Query" name:"WarnOpen"`
	SourceIp  string           `position:"Query" name:"SourceIp"`
	Channel   string           `position:"Query" name:"Channel"`
	Title     string           `position:"Query" name:"Title"`
	TimeOpen  requests.Boolean `position:"Query" name:"TimeOpen"`
	TimeBegin string           `position:"Query" name:"TimeBegin"`
	Frequency string           `position:"Query" name:"Frequency"`
}

// SetEarlyWarningResponse is the response struct for api SetEarlyWarning
type SetEarlyWarningResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	BizCode   string `json:"BizCode" xml:"BizCode"`
}

// CreateSetEarlyWarningRequest creates a request to invoke SetEarlyWarning API
func CreateSetEarlyWarningRequest() (request *SetEarlyWarningRequest) {
	request = &SetEarlyWarningRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("afs", "2018-01-12", "SetEarlyWarning", "afs", "openAPI")
	return
}

// CreateSetEarlyWarningResponse creates a response to parse from SetEarlyWarning response
func CreateSetEarlyWarningResponse() (response *SetEarlyWarningResponse) {
	response = &SetEarlyWarningResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
