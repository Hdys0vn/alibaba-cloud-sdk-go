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

// ModifyADInfo invokes the rds.ModifyADInfo API synchronously
func (client *Client) ModifyADInfo(request *ModifyADInfoRequest) (response *ModifyADInfoResponse, err error) {
	response = CreateModifyADInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyADInfoWithChan invokes the rds.ModifyADInfo API asynchronously
func (client *Client) ModifyADInfoWithChan(request *ModifyADInfoRequest) (<-chan *ModifyADInfoResponse, <-chan error) {
	responseChan := make(chan *ModifyADInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyADInfo(request)
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

// ModifyADInfoWithCallback invokes the rds.ModifyADInfo API asynchronously
func (client *Client) ModifyADInfoWithCallback(request *ModifyADInfoRequest, callback func(response *ModifyADInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyADInfoResponse
		var err error
		defer close(result)
		response, err = client.ModifyADInfo(request)
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

// ModifyADInfoRequest is the request struct for api ModifyADInfo
type ModifyADInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ADAccountName        string           `position:"Query" name:"ADAccountName"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ADDNS                string           `position:"Query" name:"ADDNS"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ADPassword           string           `position:"Query" name:"ADPassword"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ADServerIpAddress    string           `position:"Query" name:"ADServerIpAddress"`
}

// ModifyADInfoResponse is the response struct for api ModifyADInfo
type ModifyADInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyADInfoRequest creates a request to invoke ModifyADInfo API
func CreateModifyADInfoRequest() (request *ModifyADInfoRequest) {
	request = &ModifyADInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyADInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyADInfoResponse creates a response to parse from ModifyADInfo response
func CreateModifyADInfoResponse() (response *ModifyADInfoResponse) {
	response = &ModifyADInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}