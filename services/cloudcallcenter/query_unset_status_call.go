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

// QueryUnsetStatusCall invokes the cloudcallcenter.QueryUnsetStatusCall API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryunsetstatuscall.html
func (client *Client) QueryUnsetStatusCall(request *QueryUnsetStatusCallRequest) (response *QueryUnsetStatusCallResponse, err error) {
	response = CreateQueryUnsetStatusCallResponse()
	err = client.DoAction(request, response)
	return
}

// QueryUnsetStatusCallWithChan invokes the cloudcallcenter.QueryUnsetStatusCall API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryunsetstatuscall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUnsetStatusCallWithChan(request *QueryUnsetStatusCallRequest) (<-chan *QueryUnsetStatusCallResponse, <-chan error) {
	responseChan := make(chan *QueryUnsetStatusCallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryUnsetStatusCall(request)
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

// QueryUnsetStatusCallWithCallback invokes the cloudcallcenter.QueryUnsetStatusCall API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryunsetstatuscall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryUnsetStatusCallWithCallback(request *QueryUnsetStatusCallRequest, callback func(response *QueryUnsetStatusCallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryUnsetStatusCallResponse
		var err error
		defer close(result)
		response, err = client.QueryUnsetStatusCall(request)
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

// QueryUnsetStatusCallRequest is the request struct for api QueryUnsetStatusCall
type QueryUnsetStatusCallRequest struct {
	*requests.RpcRequest
	QueryUpstream requests.Boolean `position:"Query" name:"QueryUpstream"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
}

// QueryUnsetStatusCallResponse is the response struct for api QueryUnsetStatusCall
type QueryUnsetStatusCallResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	CallStatusList CallStatusList `json:"CallStatusList" xml:"CallStatusList"`
}

// CreateQueryUnsetStatusCallRequest creates a request to invoke QueryUnsetStatusCall API
func CreateQueryUnsetStatusCallRequest() (request *QueryUnsetStatusCallRequest) {
	request = &QueryUnsetStatusCallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "QueryUnsetStatusCall", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryUnsetStatusCallResponse creates a response to parse from QueryUnsetStatusCall response
func CreateQueryUnsetStatusCallResponse() (response *QueryUnsetStatusCallResponse) {
	response = &QueryUnsetStatusCallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
