package domain

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

// QueryTaskList invokes the domain.QueryTaskList API synchronously
func (client *Client) QueryTaskList(request *QueryTaskListRequest) (response *QueryTaskListResponse, err error) {
	response = CreateQueryTaskListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryTaskListWithChan invokes the domain.QueryTaskList API asynchronously
func (client *Client) QueryTaskListWithChan(request *QueryTaskListRequest) (<-chan *QueryTaskListResponse, <-chan error) {
	responseChan := make(chan *QueryTaskListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTaskList(request)
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

// QueryTaskListWithCallback invokes the domain.QueryTaskList API asynchronously
func (client *Client) QueryTaskListWithCallback(request *QueryTaskListRequest, callback func(response *QueryTaskListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskListResponse
		var err error
		defer close(result)
		response, err = client.QueryTaskList(request)
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

// QueryTaskListRequest is the request struct for api QueryTaskList
type QueryTaskListRequest struct {
	*requests.RpcRequest
	EndCreateTime   requests.Integer `position:"Query" name:"EndCreateTime"`
	PageNum         requests.Integer `position:"Query" name:"PageNum"`
	BeginCreateTime requests.Integer `position:"Query" name:"BeginCreateTime"`
	UserClientIp    string           `position:"Query" name:"UserClientIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
}

// QueryTaskListResponse is the response struct for api QueryTaskList
type QueryTaskListResponse struct {
	*responses.BaseResponse
	PrePage        bool                `json:"PrePage" xml:"PrePage"`
	CurrentPageNum int                 `json:"CurrentPageNum" xml:"CurrentPageNum"`
	RequestId      string              `json:"RequestId" xml:"RequestId"`
	PageSize       int                 `json:"PageSize" xml:"PageSize"`
	TotalPageNum   int                 `json:"TotalPageNum" xml:"TotalPageNum"`
	TotalItemNum   int                 `json:"TotalItemNum" xml:"TotalItemNum"`
	NextPage       bool                `json:"NextPage" xml:"NextPage"`
	Data           DataInQueryTaskList `json:"Data" xml:"Data"`
}

// CreateQueryTaskListRequest creates a request to invoke QueryTaskList API
func CreateQueryTaskListRequest() (request *QueryTaskListRequest) {
	request = &QueryTaskListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryTaskList", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryTaskListResponse creates a response to parse from QueryTaskList response
func CreateQueryTaskListResponse() (response *QueryTaskListResponse) {
	response = &QueryTaskListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
