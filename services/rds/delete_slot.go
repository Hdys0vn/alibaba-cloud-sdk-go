package rds

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

// DeleteSlot invokes the rds.DeleteSlot API synchronously
func (client *Client) DeleteSlot(request *DeleteSlotRequest) (response *DeleteSlotResponse, err error) {
	response = CreateDeleteSlotResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSlotWithChan invokes the rds.DeleteSlot API asynchronously
func (client *Client) DeleteSlotWithChan(request *DeleteSlotRequest) (<-chan *DeleteSlotResponse, <-chan error) {
	responseChan := make(chan *DeleteSlotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSlot(request)
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

// DeleteSlotWithCallback invokes the rds.DeleteSlot API asynchronously
func (client *Client) DeleteSlotWithCallback(request *DeleteSlotRequest, callback func(response *DeleteSlotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSlotResponse
		var err error
		defer close(result)
		response, err = client.DeleteSlot(request)
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

// DeleteSlotRequest is the request struct for api DeleteSlot
type DeleteSlotRequest struct {
	*requests.RpcRequest
	SlotName             string           `position:"Query" name:"SlotName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SlotStatus           string           `position:"Query" name:"SlotStatus"`
}

// DeleteSlotResponse is the response struct for api DeleteSlot
type DeleteSlotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SlotName  string `json:"SlotName" xml:"SlotName"`
}

// CreateDeleteSlotRequest creates a request to invoke DeleteSlot API
func CreateDeleteSlotRequest() (request *DeleteSlotRequest) {
	request = &DeleteSlotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteSlot", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSlotResponse creates a response to parse from DeleteSlot response
func CreateDeleteSlotResponse() (response *DeleteSlotResponse) {
	response = &DeleteSlotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}