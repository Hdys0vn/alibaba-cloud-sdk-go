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

// CreateDnatEntry invokes the uis.CreateDnatEntry API synchronously
// api document: https://help.aliyun.com/api/uis/creatednatentry.html
func (client *Client) CreateDnatEntry(request *CreateDnatEntryRequest) (response *CreateDnatEntryResponse, err error) {
	response = CreateCreateDnatEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDnatEntryWithChan invokes the uis.CreateDnatEntry API asynchronously
// api document: https://help.aliyun.com/api/uis/creatednatentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDnatEntryWithChan(request *CreateDnatEntryRequest) (<-chan *CreateDnatEntryResponse, <-chan error) {
	responseChan := make(chan *CreateDnatEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDnatEntry(request)
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

// CreateDnatEntryWithCallback invokes the uis.CreateDnatEntry API asynchronously
// api document: https://help.aliyun.com/api/uis/creatednatentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDnatEntryWithCallback(request *CreateDnatEntryRequest, callback func(response *CreateDnatEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDnatEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateDnatEntry(request)
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

// CreateDnatEntryRequest is the request struct for api CreateDnatEntry
type CreateDnatEntryRequest struct {
	*requests.RpcRequest
	DestinationIp        string           `position:"Query" name:"DestinationIp"`
	DestinationPort      requests.Integer `position:"Query" name:"DestinationPort"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UisNodeId            string           `position:"Query" name:"UisNodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string           `position:"Query" name:"IpProtocol"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	OriginalPort         requests.Integer `position:"Query" name:"OriginalPort"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OriginalIp           string           `position:"Query" name:"OriginalIp"`
}

// CreateDnatEntryResponse is the response struct for api CreateDnatEntry
type CreateDnatEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	UisDnatId string `json:"UisDnatId" xml:"UisDnatId"`
}

// CreateCreateDnatEntryRequest creates a request to invoke CreateDnatEntry API
func CreateCreateDnatEntryRequest() (request *CreateDnatEntryRequest) {
	request = &CreateDnatEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "CreateDnatEntry", "uis", "openAPI")
	return
}

// CreateCreateDnatEntryResponse creates a response to parse from CreateDnatEntry response
func CreateCreateDnatEntryResponse() (response *CreateDnatEntryResponse) {
	response = &CreateDnatEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
