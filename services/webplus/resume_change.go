package webplus

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

// ResumeChange invokes the webplus.ResumeChange API synchronously
// api document: https://help.aliyun.com/api/webplus/resumechange.html
func (client *Client) ResumeChange(request *ResumeChangeRequest) (response *ResumeChangeResponse, err error) {
	response = CreateResumeChangeResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeChangeWithChan invokes the webplus.ResumeChange API asynchronously
// api document: https://help.aliyun.com/api/webplus/resumechange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeChangeWithChan(request *ResumeChangeRequest) (<-chan *ResumeChangeResponse, <-chan error) {
	responseChan := make(chan *ResumeChangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeChange(request)
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

// ResumeChangeWithCallback invokes the webplus.ResumeChange API asynchronously
// api document: https://help.aliyun.com/api/webplus/resumechange.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeChangeWithCallback(request *ResumeChangeRequest, callback func(response *ResumeChangeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeChangeResponse
		var err error
		defer close(result)
		response, err = client.ResumeChange(request)
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

// ResumeChangeRequest is the request struct for api ResumeChange
type ResumeChangeRequest struct {
	*requests.RoaRequest
	ChangeId string `position:"Body" name:"ChangeId"`
}

// ResumeChangeResponse is the response struct for api ResumeChange
type ResumeChangeResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Code      string    `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	EnvChange EnvChange `json:"EnvChange" xml:"EnvChange"`
}

// CreateResumeChangeRequest creates a request to invoke ResumeChange API
func CreateResumeChangeRequest() (request *ResumeChangeRequest) {
	request = &ResumeChangeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "ResumeChange", "/pop/v1/wam/change/resume", "", "")
	request.Method = requests.POST
	return
}

// CreateResumeChangeResponse creates a response to parse from ResumeChange response
func CreateResumeChangeResponse() (response *ResumeChangeResponse) {
	response = &ResumeChangeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
