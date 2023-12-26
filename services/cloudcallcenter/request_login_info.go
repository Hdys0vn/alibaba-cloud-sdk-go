package cloudcallcenter

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

// RequestLoginInfo invokes the cloudcallcenter.RequestLoginInfo API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/requestlogininfo.html
func (client *Client) RequestLoginInfo(request *RequestLoginInfoRequest) (response *RequestLoginInfoResponse, err error) {
	response = CreateRequestLoginInfoResponse()
	err = client.DoAction(request, response)
	return
}

// RequestLoginInfoWithChan invokes the cloudcallcenter.RequestLoginInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/requestlogininfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RequestLoginInfoWithChan(request *RequestLoginInfoRequest) (<-chan *RequestLoginInfoResponse, <-chan error) {
	responseChan := make(chan *RequestLoginInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RequestLoginInfo(request)
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

// RequestLoginInfoWithCallback invokes the cloudcallcenter.RequestLoginInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/requestlogininfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RequestLoginInfoWithCallback(request *RequestLoginInfoRequest, callback func(response *RequestLoginInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RequestLoginInfoResponse
		var err error
		defer close(result)
		response, err = client.RequestLoginInfo(request)
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

// RequestLoginInfoRequest is the request struct for api RequestLoginInfo
type RequestLoginInfoRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	UserId     string `position:"Query" name:"UserId"`
}

// RequestLoginInfoResponse is the response struct for api RequestLoginInfo
type RequestLoginInfoResponse struct {
	*responses.BaseResponse
	RequestId      string    `json:"RequestId" xml:"RequestId"`
	Success        bool      `json:"Success" xml:"Success"`
	Code           string    `json:"Code" xml:"Code"`
	Message        string    `json:"Message" xml:"Message"`
	HttpStatusCode int       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	LoginInfo      LoginInfo `json:"LoginInfo" xml:"LoginInfo"`
}

// CreateRequestLoginInfoRequest creates a request to invoke RequestLoginInfo API
func CreateRequestLoginInfoRequest() (request *RequestLoginInfoRequest) {
	request = &RequestLoginInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "RequestLoginInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateRequestLoginInfoResponse creates a response to parse from RequestLoginInfo response
func CreateRequestLoginInfoResponse() (response *RequestLoginInfoResponse) {
	response = &RequestLoginInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
