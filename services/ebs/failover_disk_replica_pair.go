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

// FailoverDiskReplicaPair invokes the ebs.FailoverDiskReplicaPair API synchronously
func (client *Client) FailoverDiskReplicaPair(request *FailoverDiskReplicaPairRequest) (response *FailoverDiskReplicaPairResponse, err error) {
	response = CreateFailoverDiskReplicaPairResponse()
	err = client.DoAction(request, response)
	return
}

// FailoverDiskReplicaPairWithChan invokes the ebs.FailoverDiskReplicaPair API asynchronously
func (client *Client) FailoverDiskReplicaPairWithChan(request *FailoverDiskReplicaPairRequest) (<-chan *FailoverDiskReplicaPairResponse, <-chan error) {
	responseChan := make(chan *FailoverDiskReplicaPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FailoverDiskReplicaPair(request)
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

// FailoverDiskReplicaPairWithCallback invokes the ebs.FailoverDiskReplicaPair API asynchronously
func (client *Client) FailoverDiskReplicaPairWithCallback(request *FailoverDiskReplicaPairRequest, callback func(response *FailoverDiskReplicaPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FailoverDiskReplicaPairResponse
		var err error
		defer close(result)
		response, err = client.FailoverDiskReplicaPair(request)
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

// FailoverDiskReplicaPairRequest is the request struct for api FailoverDiskReplicaPair
type FailoverDiskReplicaPairRequest struct {
	*requests.RpcRequest
	ClientToken   string `position:"Query" name:"ClientToken"`
	ReplicaPairId string `position:"Query" name:"ReplicaPairId"`
}

// FailoverDiskReplicaPairResponse is the response struct for api FailoverDiskReplicaPair
type FailoverDiskReplicaPairResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateFailoverDiskReplicaPairRequest creates a request to invoke FailoverDiskReplicaPair API
func CreateFailoverDiskReplicaPairRequest() (request *FailoverDiskReplicaPairRequest) {
	request = &FailoverDiskReplicaPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ebs", "2021-07-30", "FailoverDiskReplicaPair", "ebs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFailoverDiskReplicaPairResponse creates a response to parse from FailoverDiskReplicaPair response
func CreateFailoverDiskReplicaPairResponse() (response *FailoverDiskReplicaPairResponse) {
	response = &FailoverDiskReplicaPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
