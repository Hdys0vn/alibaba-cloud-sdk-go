package uis

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

// DeleteUis invokes the uis.DeleteUis API synchronously
// api document: https://help.aliyun.com/api/uis/deleteuis.html
func (client *Client) DeleteUis(request *DeleteUisRequest) (response *DeleteUisResponse, err error) {
	response = CreateDeleteUisResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUisWithChan invokes the uis.DeleteUis API asynchronously
// api document: https://help.aliyun.com/api/uis/deleteuis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUisWithChan(request *DeleteUisRequest) (<-chan *DeleteUisResponse, <-chan error) {
	responseChan := make(chan *DeleteUisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUis(request)
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

// DeleteUisWithCallback invokes the uis.DeleteUis API asynchronously
// api document: https://help.aliyun.com/api/uis/deleteuis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUisWithCallback(request *DeleteUisRequest, callback func(response *DeleteUisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUisResponse
		var err error
		defer close(result)
		response, err = client.DeleteUis(request)
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

// DeleteUisRequest is the request struct for api DeleteUis
type DeleteUisRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UisId                string           `position:"Query" name:"UisId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteUisResponse is the response struct for api DeleteUis
type DeleteUisResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUisRequest creates a request to invoke DeleteUis API
func CreateDeleteUisRequest() (request *DeleteUisRequest) {
	request = &DeleteUisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "DeleteUis", "uis", "openAPI")
	return
}

// CreateDeleteUisResponse creates a response to parse from DeleteUis response
func CreateDeleteUisResponse() (response *DeleteUisResponse) {
	response = &DeleteUisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
