package ehpc

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

// DescribeContainerApp invokes the ehpc.DescribeContainerApp API synchronously
func (client *Client) DescribeContainerApp(request *DescribeContainerAppRequest) (response *DescribeContainerAppResponse, err error) {
	response = CreateDescribeContainerAppResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeContainerAppWithChan invokes the ehpc.DescribeContainerApp API asynchronously
func (client *Client) DescribeContainerAppWithChan(request *DescribeContainerAppRequest) (<-chan *DescribeContainerAppResponse, <-chan error) {
	responseChan := make(chan *DescribeContainerAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeContainerApp(request)
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

// DescribeContainerAppWithCallback invokes the ehpc.DescribeContainerApp API asynchronously
func (client *Client) DescribeContainerAppWithCallback(request *DescribeContainerAppRequest, callback func(response *DescribeContainerAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeContainerAppResponse
		var err error
		defer close(result)
		response, err = client.DescribeContainerApp(request)
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

// DescribeContainerAppRequest is the request struct for api DescribeContainerApp
type DescribeContainerAppRequest struct {
	*requests.RpcRequest
	ContainerId string `position:"Query" name:"ContainerId"`
}

// DescribeContainerAppResponse is the response struct for api DescribeContainerApp
type DescribeContainerAppResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	ContainerAppInfo ContainerAppInfo `json:"ContainerAppInfo" xml:"ContainerAppInfo"`
}

// CreateDescribeContainerAppRequest creates a request to invoke DescribeContainerApp API
func CreateDescribeContainerAppRequest() (request *DescribeContainerAppRequest) {
	request = &DescribeContainerAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "DescribeContainerApp", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeContainerAppResponse creates a response to parse from DescribeContainerApp response
func CreateDescribeContainerAppResponse() (response *DescribeContainerAppResponse) {
	response = &DescribeContainerAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
