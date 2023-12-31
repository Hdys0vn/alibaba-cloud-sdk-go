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

// GetCabTask invokes the cloudcallcenter.GetCabTask API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabtask.html
func (client *Client) GetCabTask(request *GetCabTaskRequest) (response *GetCabTaskResponse, err error) {
	response = CreateGetCabTaskResponse()
	err = client.DoAction(request, response)
	return
}

// GetCabTaskWithChan invokes the cloudcallcenter.GetCabTask API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCabTaskWithChan(request *GetCabTaskRequest) (<-chan *GetCabTaskResponse, <-chan error) {
	responseChan := make(chan *GetCabTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCabTask(request)
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

// GetCabTaskWithCallback invokes the cloudcallcenter.GetCabTask API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCabTaskWithCallback(request *GetCabTaskRequest, callback func(response *GetCabTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCabTaskResponse
		var err error
		defer close(result)
		response, err = client.GetCabTask(request)
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

// GetCabTaskRequest is the request struct for api GetCabTask
type GetCabTaskRequest struct {
	*requests.RpcRequest
	AccTaskId string `position:"Query" name:"AccTaskId"`
}

// GetCabTaskResponse is the response struct for api GetCabTask
type GetCabTaskResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	TaskId         string `json:"TaskId" xml:"TaskId"`
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
	Owner          int64  `json:"Owner" xml:"Owner"`
	Callee         string `json:"Callee" xml:"Callee"`
	Caller         string `json:"Caller" xml:"Caller"`
	ScenarioId     string `json:"ScenarioId" xml:"ScenarioId"`
}

// CreateGetCabTaskRequest creates a request to invoke GetCabTask API
func CreateGetCabTaskRequest() (request *GetCabTaskRequest) {
	request = &GetCabTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetCabTask", "", "")
	request.Method = requests.POST
	return
}

// CreateGetCabTaskResponse creates a response to parse from GetCabTask response
func CreateGetCabTaskResponse() (response *GetCabTaskResponse) {
	response = &GetCabTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
