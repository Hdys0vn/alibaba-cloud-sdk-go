package ledgerdb

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

// ModifyMemberKey invokes the ledgerdb.ModifyMemberKey API synchronously
// api document: https://help.aliyun.com/api/ledgerdb/modifymemberkey.html
func (client *Client) ModifyMemberKey(request *ModifyMemberKeyRequest) (response *ModifyMemberKeyResponse, err error) {
	response = CreateModifyMemberKeyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMemberKeyWithChan invokes the ledgerdb.ModifyMemberKey API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/modifymemberkey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMemberKeyWithChan(request *ModifyMemberKeyRequest) (<-chan *ModifyMemberKeyResponse, <-chan error) {
	responseChan := make(chan *ModifyMemberKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMemberKey(request)
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

// ModifyMemberKeyWithCallback invokes the ledgerdb.ModifyMemberKey API asynchronously
// api document: https://help.aliyun.com/api/ledgerdb/modifymemberkey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMemberKeyWithCallback(request *ModifyMemberKeyRequest, callback func(response *ModifyMemberKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMemberKeyResponse
		var err error
		defer close(result)
		response, err = client.ModifyMemberKey(request)
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

// ModifyMemberKeyRequest is the request struct for api ModifyMemberKey
type ModifyMemberKeyRequest struct {
	*requests.RpcRequest
	PublicKey string `position:"Body" name:"PublicKey"`
	KeyType   string `position:"Body" name:"KeyType"`
	LedgerId  string `position:"Body" name:"LedgerId"`
	MemberId  string `position:"Body" name:"MemberId"`
}

// ModifyMemberKeyResponse is the response struct for api ModifyMemberKey
type ModifyMemberKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyMemberKeyRequest creates a request to invoke ModifyMemberKey API
func CreateModifyMemberKeyRequest() (request *ModifyMemberKeyRequest) {
	request = &ModifyMemberKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ledgerdb", "2019-11-22", "ModifyMemberKey", "ledgerdb", "openAPI")
	return
}

// CreateModifyMemberKeyResponse creates a response to parse from ModifyMemberKey response
func CreateModifyMemberKeyResponse() (response *ModifyMemberKeyResponse) {
	response = &ModifyMemberKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
