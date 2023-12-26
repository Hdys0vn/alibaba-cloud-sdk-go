package ddoscoo

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

// DeletePort invokes the ddoscoo.DeletePort API synchronously
func (client *Client) DeletePort(request *DeletePortRequest) (response *DeletePortResponse, err error) {
	response = CreateDeletePortResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePortWithChan invokes the ddoscoo.DeletePort API asynchronously
func (client *Client) DeletePortWithChan(request *DeletePortRequest) (<-chan *DeletePortResponse, <-chan error) {
	responseChan := make(chan *DeletePortResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePort(request)
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

// DeletePortWithCallback invokes the ddoscoo.DeletePort API asynchronously
func (client *Client) DeletePortWithCallback(request *DeletePortRequest, callback func(response *DeletePortResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePortResponse
		var err error
		defer close(result)
		response, err = client.DeletePort(request)
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

// DeletePortRequest is the request struct for api DeletePort
type DeletePortRequest struct {
	*requests.RpcRequest
	SourceIp         string    `position:"Query" name:"SourceIp"`
	BackendPort      string    `position:"Query" name:"BackendPort"`
	FrontendProtocol string    `position:"Query" name:"FrontendProtocol"`
	InstanceId       string    `position:"Query" name:"InstanceId"`
	RealServers      *[]string `position:"Query" name:"RealServers"  type:"Repeated"`
	FrontendPort     string    `position:"Query" name:"FrontendPort"`
}

// DeletePortResponse is the response struct for api DeletePort
type DeletePortResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePortRequest creates a request to invoke DeletePort API
func CreateDeletePortRequest() (request *DeletePortRequest) {
	request = &DeletePortRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DeletePort", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeletePortResponse creates a response to parse from DeletePort response
func CreateDeletePortResponse() (response *DeletePortResponse) {
	response = &DeletePortResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}