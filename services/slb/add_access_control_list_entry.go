package slb

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

// AddAccessControlListEntry invokes the slb.AddAccessControlListEntry API synchronously
func (client *Client) AddAccessControlListEntry(request *AddAccessControlListEntryRequest) (response *AddAccessControlListEntryResponse, err error) {
	response = CreateAddAccessControlListEntryResponse()
	err = client.DoAction(request, response)
	return
}

// AddAccessControlListEntryWithChan invokes the slb.AddAccessControlListEntry API asynchronously
func (client *Client) AddAccessControlListEntryWithChan(request *AddAccessControlListEntryRequest) (<-chan *AddAccessControlListEntryResponse, <-chan error) {
	responseChan := make(chan *AddAccessControlListEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddAccessControlListEntry(request)
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

// AddAccessControlListEntryWithCallback invokes the slb.AddAccessControlListEntry API asynchronously
func (client *Client) AddAccessControlListEntryWithCallback(request *AddAccessControlListEntryRequest, callback func(response *AddAccessControlListEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddAccessControlListEntryResponse
		var err error
		defer close(result)
		response, err = client.AddAccessControlListEntry(request)
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

// AddAccessControlListEntryRequest is the request struct for api AddAccessControlListEntry
type AddAccessControlListEntryRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AclEntrys            string           `position:"Query" name:"AclEntrys"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// AddAccessControlListEntryResponse is the response struct for api AddAccessControlListEntry
type AddAccessControlListEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddAccessControlListEntryRequest creates a request to invoke AddAccessControlListEntry API
func CreateAddAccessControlListEntryRequest() (request *AddAccessControlListEntryRequest) {
	request = &AddAccessControlListEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "AddAccessControlListEntry", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddAccessControlListEntryResponse creates a response to parse from AddAccessControlListEntry response
func CreateAddAccessControlListEntryResponse() (response *AddAccessControlListEntryResponse) {
	response = &AddAccessControlListEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
