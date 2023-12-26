package ebs

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

// ModifyDiskReplicaPair invokes the ebs.ModifyDiskReplicaPair API synchronously
func (client *Client) ModifyDiskReplicaPair(request *ModifyDiskReplicaPairRequest) (response *ModifyDiskReplicaPairResponse, err error) {
	response = CreateModifyDiskReplicaPairResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDiskReplicaPairWithChan invokes the ebs.ModifyDiskReplicaPair API asynchronously
func (client *Client) ModifyDiskReplicaPairWithChan(request *ModifyDiskReplicaPairRequest) (<-chan *ModifyDiskReplicaPairResponse, <-chan error) {
	responseChan := make(chan *ModifyDiskReplicaPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDiskReplicaPair(request)
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

// ModifyDiskReplicaPairWithCallback invokes the ebs.ModifyDiskReplicaPair API asynchronously
func (client *Client) ModifyDiskReplicaPairWithCallback(request *ModifyDiskReplicaPairRequest, callback func(response *ModifyDiskReplicaPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDiskReplicaPairResponse
		var err error
		defer close(result)
		response, err = client.ModifyDiskReplicaPair(request)
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

// ModifyDiskReplicaPairRequest is the request struct for api ModifyDiskReplicaPair
type ModifyDiskReplicaPairRequest struct {
	*requests.RpcRequest
	PairName      string           `position:"Query" name:"PairName"`
	ClientToken   string           `position:"Query" name:"ClientToken"`
	Description   string           `position:"Query" name:"Description"`
	Bandwidth     requests.Integer `position:"Query" name:"Bandwidth"`
	ReplicaPairId string           `position:"Query" name:"ReplicaPairId"`
	RPO           requests.Integer `position:"Query" name:"RPO"`
}

// ModifyDiskReplicaPairResponse is the response struct for api ModifyDiskReplicaPair
type ModifyDiskReplicaPairResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDiskReplicaPairRequest creates a request to invoke ModifyDiskReplicaPair API
func CreateModifyDiskReplicaPairRequest() (request *ModifyDiskReplicaPairRequest) {
	request = &ModifyDiskReplicaPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ebs", "2021-07-30", "ModifyDiskReplicaPair", "ebs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDiskReplicaPairResponse creates a response to parse from ModifyDiskReplicaPair response
func CreateModifyDiskReplicaPairResponse() (response *ModifyDiskReplicaPairResponse) {
	response = &ModifyDiskReplicaPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}