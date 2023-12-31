package ddospro

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

// DeleteBlackHole invokes the ddospro.DeleteBlackHole API synchronously
// api document: https://help.aliyun.com/api/ddospro/deleteblackhole.html
func (client *Client) DeleteBlackHole(request *DeleteBlackHoleRequest) (response *DeleteBlackHoleResponse, err error) {
	response = CreateDeleteBlackHoleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBlackHoleWithChan invokes the ddospro.DeleteBlackHole API asynchronously
// api document: https://help.aliyun.com/api/ddospro/deleteblackhole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBlackHoleWithChan(request *DeleteBlackHoleRequest) (<-chan *DeleteBlackHoleResponse, <-chan error) {
	responseChan := make(chan *DeleteBlackHoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBlackHole(request)
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

// DeleteBlackHoleWithCallback invokes the ddospro.DeleteBlackHole API asynchronously
// api document: https://help.aliyun.com/api/ddospro/deleteblackhole.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBlackHoleWithCallback(request *DeleteBlackHoleRequest, callback func(response *DeleteBlackHoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBlackHoleResponse
		var err error
		defer close(result)
		response, err = client.DeleteBlackHole(request)
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

// DeleteBlackHoleRequest is the request struct for api DeleteBlackHole
type DeleteBlackHoleRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Vip             string           `position:"Query" name:"Vip"`
}

// DeleteBlackHoleResponse is the response struct for api DeleteBlackHole
type DeleteBlackHoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"data" xml:"data"`
}

// CreateDeleteBlackHoleRequest creates a request to invoke DeleteBlackHole API
func CreateDeleteBlackHoleRequest() (request *DeleteBlackHoleRequest) {
	request = &DeleteBlackHoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DeleteBlackHole", "", "")
	return
}

// CreateDeleteBlackHoleResponse creates a response to parse from DeleteBlackHole response
func CreateDeleteBlackHoleResponse() (response *DeleteBlackHoleResponse) {
	response = &DeleteBlackHoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
