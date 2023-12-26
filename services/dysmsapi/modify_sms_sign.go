package dysmsapi

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

// ModifySmsSign invokes the dysmsapi.ModifySmsSign API synchronously
func (client *Client) ModifySmsSign(request *ModifySmsSignRequest) (response *ModifySmsSignResponse, err error) {
	response = CreateModifySmsSignResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySmsSignWithChan invokes the dysmsapi.ModifySmsSign API asynchronously
func (client *Client) ModifySmsSignWithChan(request *ModifySmsSignRequest) (<-chan *ModifySmsSignResponse, <-chan error) {
	responseChan := make(chan *ModifySmsSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySmsSign(request)
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

// ModifySmsSignWithCallback invokes the dysmsapi.ModifySmsSign API asynchronously
func (client *Client) ModifySmsSignWithCallback(request *ModifySmsSignRequest, callback func(response *ModifySmsSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySmsSignResponse
		var err error
		defer close(result)
		response, err = client.ModifySmsSign(request)
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

// ModifySmsSignRequest is the request struct for api ModifySmsSign
type ModifySmsSignRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer             `position:"Query" name:"ResourceOwnerId"`
	Remark               string                       `position:"Query" name:"Remark"`
	SignName             string                       `position:"Query" name:"SignName"`
	SignFileList         *[]ModifySmsSignSignFileList `position:"Body" name:"SignFileList"  type:"Repeated"`
	ResourceOwnerAccount string                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer             `position:"Query" name:"OwnerId"`
	SignSource           requests.Integer             `position:"Query" name:"SignSource"`
}

// ModifySmsSignSignFileList is a repeated param struct in ModifySmsSignRequest
type ModifySmsSignSignFileList struct {
	FileContents string `name:"FileContents"`
	FileSuffix   string `name:"FileSuffix"`
}

// ModifySmsSignResponse is the response struct for api ModifySmsSign
type ModifySmsSignResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SignName  string `json:"SignName" xml:"SignName"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateModifySmsSignRequest creates a request to invoke ModifySmsSign API
func CreateModifySmsSignRequest() (request *ModifySmsSignRequest) {
	request = &ModifySmsSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "ModifySmsSign", "", "")
	request.Method = requests.POST
	return
}

// CreateModifySmsSignResponse creates a response to parse from ModifySmsSign response
func CreateModifySmsSignResponse() (response *ModifySmsSignResponse) {
	response = &ModifySmsSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
