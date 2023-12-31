package dm

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

// CreateMailAddress invokes the dm.CreateMailAddress API synchronously
// api document: https://help.aliyun.com/api/dm/createmailaddress.html
func (client *Client) CreateMailAddress(request *CreateMailAddressRequest) (response *CreateMailAddressResponse, err error) {
	response = CreateCreateMailAddressResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMailAddressWithChan invokes the dm.CreateMailAddress API asynchronously
// api document: https://help.aliyun.com/api/dm/createmailaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMailAddressWithChan(request *CreateMailAddressRequest) (<-chan *CreateMailAddressResponse, <-chan error) {
	responseChan := make(chan *CreateMailAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMailAddress(request)
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

// CreateMailAddressWithCallback invokes the dm.CreateMailAddress API asynchronously
// api document: https://help.aliyun.com/api/dm/createmailaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMailAddressWithCallback(request *CreateMailAddressRequest, callback func(response *CreateMailAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMailAddressResponse
		var err error
		defer close(result)
		response, err = client.CreateMailAddress(request)
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

// CreateMailAddressRequest is the request struct for api CreateMailAddress
type CreateMailAddressRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	ReplyAddress         string           `position:"Query" name:"ReplyAddress"`
	Sendtype             string           `position:"Query" name:"Sendtype"`
}

// CreateMailAddressResponse is the response struct for api CreateMailAddress
type CreateMailAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateMailAddressRequest creates a request to invoke CreateMailAddress API
func CreateCreateMailAddressRequest() (request *CreateMailAddressRequest) {
	request = &CreateMailAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "CreateMailAddress", "", "")
	return
}

// CreateCreateMailAddressResponse creates a response to parse from CreateMailAddress response
func CreateCreateMailAddressResponse() (response *CreateMailAddressResponse) {
	response = &CreateMailAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
