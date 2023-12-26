package dysmsapi

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

// QuerySmsSignList invokes the dysmsapi.QuerySmsSignList API synchronously
func (client *Client) QuerySmsSignList(request *QuerySmsSignListRequest) (response *QuerySmsSignListResponse, err error) {
	response = CreateQuerySmsSignListResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySmsSignListWithChan invokes the dysmsapi.QuerySmsSignList API asynchronously
func (client *Client) QuerySmsSignListWithChan(request *QuerySmsSignListRequest) (<-chan *QuerySmsSignListResponse, <-chan error) {
	responseChan := make(chan *QuerySmsSignListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySmsSignList(request)
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

// QuerySmsSignListWithCallback invokes the dysmsapi.QuerySmsSignList API asynchronously
func (client *Client) QuerySmsSignListWithCallback(request *QuerySmsSignListRequest, callback func(response *QuerySmsSignListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySmsSignListResponse
		var err error
		defer close(result)
		response, err = client.QuerySmsSignList(request)
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

// QuerySmsSignListRequest is the request struct for api QuerySmsSignList
type QuerySmsSignListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageIndex            requests.Integer `position:"Query" name:"PageIndex"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// QuerySmsSignListResponse is the response struct for api QuerySmsSignList
type QuerySmsSignListResponse struct {
	*responses.BaseResponse
	RequestId   string            `json:"RequestId" xml:"RequestId"`
	Code        string            `json:"Code" xml:"Code"`
	Message     string            `json:"Message" xml:"Message"`
	SmsSignList []QuerySmsSignDTO `json:"SmsSignList" xml:"SmsSignList"`
}

// CreateQuerySmsSignListRequest creates a request to invoke QuerySmsSignList API
func CreateQuerySmsSignListRequest() (request *QuerySmsSignListRequest) {
	request = &QuerySmsSignListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "QuerySmsSignList", "", "")
	request.Method = requests.POST
	return
}

// CreateQuerySmsSignListResponse creates a response to parse from QuerySmsSignList response
func CreateQuerySmsSignListResponse() (response *QuerySmsSignListResponse) {
	response = &QuerySmsSignListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
