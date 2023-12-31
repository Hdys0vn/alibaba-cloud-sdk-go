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

// GetInstanceState invokes the cloudcallcenter.GetInstanceState API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancestate.html
func (client *Client) GetInstanceState(request *GetInstanceStateRequest) (response *GetInstanceStateResponse, err error) {
	response = CreateGetInstanceStateResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceStateWithChan invokes the cloudcallcenter.GetInstanceState API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancestate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceStateWithChan(request *GetInstanceStateRequest) (<-chan *GetInstanceStateResponse, <-chan error) {
	responseChan := make(chan *GetInstanceStateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceState(request)
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

// GetInstanceStateWithCallback invokes the cloudcallcenter.GetInstanceState API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getinstancestate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceStateWithCallback(request *GetInstanceStateRequest, callback func(response *GetInstanceStateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceStateResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceState(request)
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

// GetInstanceStateRequest is the request struct for api GetInstanceState
type GetInstanceStateRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetInstanceStateResponse is the response struct for api GetInstanceState
type GetInstanceStateResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	Success               bool                  `json:"Success" xml:"Success"`
	Code                  string                `json:"Code" xml:"Code"`
	Message               string                `json:"Message" xml:"Message"`
	HttpStatusCode        int                   `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RealTimeInstanceState RealTimeInstanceState `json:"RealTimeInstanceState" xml:"RealTimeInstanceState"`
}

// CreateGetInstanceStateRequest creates a request to invoke GetInstanceState API
func CreateGetInstanceStateRequest() (request *GetInstanceStateRequest) {
	request = &GetInstanceStateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetInstanceState", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInstanceStateResponse creates a response to parse from GetInstanceState response
func CreateGetInstanceStateResponse() (response *GetInstanceStateResponse) {
	response = &GetInstanceStateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
