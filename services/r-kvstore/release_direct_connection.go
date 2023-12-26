package r_kvstore

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

// ReleaseDirectConnection invokes the r_kvstore.ReleaseDirectConnection API synchronously
func (client *Client) ReleaseDirectConnection(request *ReleaseDirectConnectionRequest) (response *ReleaseDirectConnectionResponse, err error) {
	response = CreateReleaseDirectConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseDirectConnectionWithChan invokes the r_kvstore.ReleaseDirectConnection API asynchronously
func (client *Client) ReleaseDirectConnectionWithChan(request *ReleaseDirectConnectionRequest) (<-chan *ReleaseDirectConnectionResponse, <-chan error) {
	responseChan := make(chan *ReleaseDirectConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseDirectConnection(request)
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

// ReleaseDirectConnectionWithCallback invokes the r_kvstore.ReleaseDirectConnection API asynchronously
func (client *Client) ReleaseDirectConnectionWithCallback(request *ReleaseDirectConnectionRequest, callback func(response *ReleaseDirectConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseDirectConnectionResponse
		var err error
		defer close(result)
		response, err = client.ReleaseDirectConnection(request)
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

// ReleaseDirectConnectionRequest is the request struct for api ReleaseDirectConnection
type ReleaseDirectConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ReleaseDirectConnectionResponse is the response struct for api ReleaseDirectConnection
type ReleaseDirectConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseDirectConnectionRequest creates a request to invoke ReleaseDirectConnection API
func CreateReleaseDirectConnectionRequest() (request *ReleaseDirectConnectionRequest) {
	request = &ReleaseDirectConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ReleaseDirectConnection", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleaseDirectConnectionResponse creates a response to parse from ReleaseDirectConnection response
func CreateReleaseDirectConnectionResponse() (response *ReleaseDirectConnectionResponse) {
	response = &ReleaseDirectConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
