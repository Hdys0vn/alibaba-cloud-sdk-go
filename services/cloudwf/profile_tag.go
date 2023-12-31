package cloudwf

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

// ProfileTag invokes the cloudwf.ProfileTag API synchronously
// api document: https://help.aliyun.com/api/cloudwf/profiletag.html
func (client *Client) ProfileTag(request *ProfileTagRequest) (response *ProfileTagResponse, err error) {
	response = CreateProfileTagResponse()
	err = client.DoAction(request, response)
	return
}

// ProfileTagWithChan invokes the cloudwf.ProfileTag API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/profiletag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProfileTagWithChan(request *ProfileTagRequest) (<-chan *ProfileTagResponse, <-chan error) {
	responseChan := make(chan *ProfileTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProfileTag(request)
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

// ProfileTagWithCallback invokes the cloudwf.ProfileTag API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/profiletag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProfileTagWithCallback(request *ProfileTagRequest, callback func(response *ProfileTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProfileTagResponse
		var err error
		defer close(result)
		response, err = client.ProfileTag(request)
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

// ProfileTagRequest is the request struct for api ProfileTag
type ProfileTagRequest struct {
	*requests.RpcRequest
	Idtype     requests.Integer `position:"Query" name:"Idtype"`
	BeginDate  string           `position:"Query" name:"BeginDate"`
	EndDate    string           `position:"Query" name:"EndDate"`
	AppType    requests.Integer `position:"Query" name:"AppType"`
	Tag        string           `position:"Query" name:"Tag"`
	Agsid      requests.Integer `position:"Query" name:"Agsid"`
	AreaNumber requests.Integer `position:"Query" name:"AreaNumber"`
}

// ProfileTagResponse is the response struct for api ProfileTag
type ProfileTagResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateProfileTagRequest creates a request to invoke ProfileTag API
func CreateProfileTagRequest() (request *ProfileTagRequest) {
	request = &ProfileTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileTag", "cloudwf", "openAPI")
	return
}

// CreateProfileTagResponse creates a response to parse from ProfileTag response
func CreateProfileTagResponse() (response *ProfileTagResponse) {
	response = &ProfileTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
