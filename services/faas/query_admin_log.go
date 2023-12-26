package faas

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

// QueryAdminLog invokes the faas.QueryAdminLog API synchronously
// api document: https://help.aliyun.com/api/faas/queryadminlog.html
func (client *Client) QueryAdminLog(request *QueryAdminLogRequest) (response *QueryAdminLogResponse, err error) {
	response = CreateQueryAdminLogResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAdminLogWithChan invokes the faas.QueryAdminLog API asynchronously
// api document: https://help.aliyun.com/api/faas/queryadminlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAdminLogWithChan(request *QueryAdminLogRequest) (<-chan *QueryAdminLogResponse, <-chan error) {
	responseChan := make(chan *QueryAdminLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAdminLog(request)
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

// QueryAdminLogWithCallback invokes the faas.QueryAdminLog API asynchronously
// api document: https://help.aliyun.com/api/faas/queryadminlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAdminLogWithCallback(request *QueryAdminLogRequest, callback func(response *QueryAdminLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAdminLogResponse
		var err error
		defer close(result)
		response, err = client.QueryAdminLog(request)
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

// QueryAdminLogRequest is the request struct for api QueryAdminLog
type QueryAdminLogRequest struct {
	*requests.RpcRequest
	FpgaImageUniqueId string `position:"Query" name:"FpgaImageUniqueId"`
}

// QueryAdminLogResponse is the response struct for api QueryAdminLog
type QueryAdminLogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	AdminLog  string `json:"AdminLog" xml:"AdminLog"`
}

// CreateQueryAdminLogRequest creates a request to invoke QueryAdminLog API
func CreateQueryAdminLogRequest() (request *QueryAdminLogRequest) {
	request = &QueryAdminLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "QueryAdminLog", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAdminLogResponse creates a response to parse from QueryAdminLog response
func CreateQueryAdminLogResponse() (response *QueryAdminLogResponse) {
	response = &QueryAdminLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}