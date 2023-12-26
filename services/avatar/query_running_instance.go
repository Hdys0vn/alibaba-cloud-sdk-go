package avatar

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

// QueryRunningInstance invokes the avatar.QueryRunningInstance API synchronously
func (client *Client) QueryRunningInstance(request *QueryRunningInstanceRequest) (response *QueryRunningInstanceResponse, err error) {
	response = CreateQueryRunningInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRunningInstanceWithChan invokes the avatar.QueryRunningInstance API asynchronously
func (client *Client) QueryRunningInstanceWithChan(request *QueryRunningInstanceRequest) (<-chan *QueryRunningInstanceResponse, <-chan error) {
	responseChan := make(chan *QueryRunningInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRunningInstance(request)
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

// QueryRunningInstanceWithCallback invokes the avatar.QueryRunningInstance API asynchronously
func (client *Client) QueryRunningInstanceWithCallback(request *QueryRunningInstanceRequest, callback func(response *QueryRunningInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRunningInstanceResponse
		var err error
		defer close(result)
		response, err = client.QueryRunningInstance(request)
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

// QueryRunningInstanceRequest is the request struct for api QueryRunningInstance
type QueryRunningInstanceRequest struct {
	*requests.RpcRequest
	App       QueryRunningInstanceApp `position:"Query" name:"App"  type:"Struct"`
	TenantId  requests.Integer        `position:"Query" name:"TenantId"`
	SessionId string                  `position:"Query" name:"SessionId"`
}

// QueryRunningInstanceApp is a repeated param struct in QueryRunningInstanceRequest
type QueryRunningInstanceApp struct {
	AppId string `name:"AppId"`
}

// QueryRunningInstanceResponse is the response struct for api QueryRunningInstance
type QueryRunningInstanceResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	Success   bool       `json:"Success" xml:"Success"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateQueryRunningInstanceRequest creates a request to invoke QueryRunningInstance API
func CreateQueryRunningInstanceRequest() (request *QueryRunningInstanceRequest) {
	request = &QueryRunningInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "QueryRunningInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryRunningInstanceResponse creates a response to parse from QueryRunningInstance response
func CreateQueryRunningInstanceResponse() (response *QueryRunningInstanceResponse) {
	response = &QueryRunningInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}