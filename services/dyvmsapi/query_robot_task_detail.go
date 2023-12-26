package dyvmsapi

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

// QueryRobotTaskDetail invokes the dyvmsapi.QueryRobotTaskDetail API synchronously
func (client *Client) QueryRobotTaskDetail(request *QueryRobotTaskDetailRequest) (response *QueryRobotTaskDetailResponse, err error) {
	response = CreateQueryRobotTaskDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRobotTaskDetailWithChan invokes the dyvmsapi.QueryRobotTaskDetail API asynchronously
func (client *Client) QueryRobotTaskDetailWithChan(request *QueryRobotTaskDetailRequest) (<-chan *QueryRobotTaskDetailResponse, <-chan error) {
	responseChan := make(chan *QueryRobotTaskDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRobotTaskDetail(request)
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

// QueryRobotTaskDetailWithCallback invokes the dyvmsapi.QueryRobotTaskDetail API asynchronously
func (client *Client) QueryRobotTaskDetailWithCallback(request *QueryRobotTaskDetailRequest, callback func(response *QueryRobotTaskDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRobotTaskDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryRobotTaskDetail(request)
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

// QueryRobotTaskDetailRequest is the request struct for api QueryRobotTaskDetail
type QueryRobotTaskDetailRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Id                   requests.Integer `position:"Query" name:"Id"`
}

// QueryRobotTaskDetailResponse is the response struct for api QueryRobotTaskDetail
type QueryRobotTaskDetailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateQueryRobotTaskDetailRequest creates a request to invoke QueryRobotTaskDetail API
func CreateQueryRobotTaskDetailRequest() (request *QueryRobotTaskDetailRequest) {
	request = &QueryRobotTaskDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "QueryRobotTaskDetail", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryRobotTaskDetailResponse creates a response to parse from QueryRobotTaskDetail response
func CreateQueryRobotTaskDetailResponse() (response *QueryRobotTaskDetailResponse) {
	response = &QueryRobotTaskDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
