package csas

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

// DeletePrivateAccessTag invokes the csas.DeletePrivateAccessTag API synchronously
func (client *Client) DeletePrivateAccessTag(request *DeletePrivateAccessTagRequest) (response *DeletePrivateAccessTagResponse, err error) {
	response = CreateDeletePrivateAccessTagResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePrivateAccessTagWithChan invokes the csas.DeletePrivateAccessTag API asynchronously
func (client *Client) DeletePrivateAccessTagWithChan(request *DeletePrivateAccessTagRequest) (<-chan *DeletePrivateAccessTagResponse, <-chan error) {
	responseChan := make(chan *DeletePrivateAccessTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePrivateAccessTag(request)
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

// DeletePrivateAccessTagWithCallback invokes the csas.DeletePrivateAccessTag API asynchronously
func (client *Client) DeletePrivateAccessTagWithCallback(request *DeletePrivateAccessTagRequest, callback func(response *DeletePrivateAccessTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePrivateAccessTagResponse
		var err error
		defer close(result)
		response, err = client.DeletePrivateAccessTag(request)
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

// DeletePrivateAccessTagRequest is the request struct for api DeletePrivateAccessTag
type DeletePrivateAccessTagRequest struct {
	*requests.RpcRequest
	TagId    string `position:"Body" name:"TagId"`
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DeletePrivateAccessTagResponse is the response struct for api DeletePrivateAccessTag
type DeletePrivateAccessTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeletePrivateAccessTagRequest creates a request to invoke DeletePrivateAccessTag API
func CreateDeletePrivateAccessTagRequest() (request *DeletePrivateAccessTagRequest) {
	request = &DeletePrivateAccessTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "DeletePrivateAccessTag", "", "")
	request.Method = requests.POST
	return
}

// CreateDeletePrivateAccessTagResponse creates a response to parse from DeletePrivateAccessTag response
func CreateDeletePrivateAccessTagResponse() (response *DeletePrivateAccessTagResponse) {
	response = &DeletePrivateAccessTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}