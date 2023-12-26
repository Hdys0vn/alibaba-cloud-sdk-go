package companyreg

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

// QueryInstance invokes the companyreg.QueryInstance API synchronously
func (client *Client) QueryInstance(request *QueryInstanceRequest) (response *QueryInstanceResponse, err error) {
	response = CreateQueryInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInstanceWithChan invokes the companyreg.QueryInstance API asynchronously
func (client *Client) QueryInstanceWithChan(request *QueryInstanceRequest) (<-chan *QueryInstanceResponse, <-chan error) {
	responseChan := make(chan *QueryInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInstance(request)
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

// QueryInstanceWithCallback invokes the companyreg.QueryInstance API asynchronously
func (client *Client) QueryInstanceWithCallback(request *QueryInstanceRequest, callback func(response *QueryInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInstanceResponse
		var err error
		defer close(result)
		response, err = client.QueryInstance(request)
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

// QueryInstanceRequest is the request struct for api QueryInstance
type QueryInstanceRequest struct {
	*requests.RpcRequest
	BizType    string `position:"Query" name:"BizType"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// QueryInstanceResponse is the response struct for api QueryInstance
type QueryInstanceResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateQueryInstanceRequest creates a request to invoke QueryInstance API
func CreateQueryInstanceRequest() (request *QueryInstanceRequest) {
	request = &QueryInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "QueryInstance", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryInstanceResponse creates a response to parse from QueryInstance response
func CreateQueryInstanceResponse() (response *QueryInstanceResponse) {
	response = &QueryInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
