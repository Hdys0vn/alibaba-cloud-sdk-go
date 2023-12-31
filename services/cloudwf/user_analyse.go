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

// UserAnalyse invokes the cloudwf.UserAnalyse API synchronously
// api document: https://help.aliyun.com/api/cloudwf/useranalyse.html
func (client *Client) UserAnalyse(request *UserAnalyseRequest) (response *UserAnalyseResponse, err error) {
	response = CreateUserAnalyseResponse()
	err = client.DoAction(request, response)
	return
}

// UserAnalyseWithChan invokes the cloudwf.UserAnalyse API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/useranalyse.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UserAnalyseWithChan(request *UserAnalyseRequest) (<-chan *UserAnalyseResponse, <-chan error) {
	responseChan := make(chan *UserAnalyseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UserAnalyse(request)
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

// UserAnalyseWithCallback invokes the cloudwf.UserAnalyse API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/useranalyse.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UserAnalyseWithCallback(request *UserAnalyseRequest, callback func(response *UserAnalyseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UserAnalyseResponse
		var err error
		defer close(result)
		response, err = client.UserAnalyse(request)
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

// UserAnalyseRequest is the request struct for api UserAnalyse
type UserAnalyseRequest struct {
	*requests.RpcRequest
	Gsid requests.Integer `position:"Query" name:"Gsid"`
}

// UserAnalyseResponse is the response struct for api UserAnalyse
type UserAnalyseResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateUserAnalyseRequest creates a request to invoke UserAnalyse API
func CreateUserAnalyseRequest() (request *UserAnalyseRequest) {
	request = &UserAnalyseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "UserAnalyse", "cloudwf", "openAPI")
	return
}

// CreateUserAnalyseResponse creates a response to parse from UserAnalyse response
func CreateUserAnalyseResponse() (response *UserAnalyseResponse) {
	response = &UserAnalyseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
