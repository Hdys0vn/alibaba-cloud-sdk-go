package vpc

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

// DeleteGlobalAccelerationInstance invokes the vpc.DeleteGlobalAccelerationInstance API synchronously
func (client *Client) DeleteGlobalAccelerationInstance(request *DeleteGlobalAccelerationInstanceRequest) (response *DeleteGlobalAccelerationInstanceResponse, err error) {
	response = CreateDeleteGlobalAccelerationInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGlobalAccelerationInstanceWithChan invokes the vpc.DeleteGlobalAccelerationInstance API asynchronously
func (client *Client) DeleteGlobalAccelerationInstanceWithChan(request *DeleteGlobalAccelerationInstanceRequest) (<-chan *DeleteGlobalAccelerationInstanceResponse, <-chan error) {
	responseChan := make(chan *DeleteGlobalAccelerationInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGlobalAccelerationInstance(request)
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

// DeleteGlobalAccelerationInstanceWithCallback invokes the vpc.DeleteGlobalAccelerationInstance API asynchronously
func (client *Client) DeleteGlobalAccelerationInstanceWithCallback(request *DeleteGlobalAccelerationInstanceRequest, callback func(response *DeleteGlobalAccelerationInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGlobalAccelerationInstanceResponse
		var err error
		defer close(result)
		response, err = client.DeleteGlobalAccelerationInstance(request)
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

// DeleteGlobalAccelerationInstanceRequest is the request struct for api DeleteGlobalAccelerationInstance
type DeleteGlobalAccelerationInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	GlobalAccelerationInstanceId string           `position:"Query" name:"GlobalAccelerationInstanceId"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteGlobalAccelerationInstanceResponse is the response struct for api DeleteGlobalAccelerationInstance
type DeleteGlobalAccelerationInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteGlobalAccelerationInstanceRequest creates a request to invoke DeleteGlobalAccelerationInstance API
func CreateDeleteGlobalAccelerationInstanceRequest() (request *DeleteGlobalAccelerationInstanceRequest) {
	request = &DeleteGlobalAccelerationInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteGlobalAccelerationInstance", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGlobalAccelerationInstanceResponse creates a response to parse from DeleteGlobalAccelerationInstance response
func CreateDeleteGlobalAccelerationInstanceResponse() (response *DeleteGlobalAccelerationInstanceResponse) {
	response = &DeleteGlobalAccelerationInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
