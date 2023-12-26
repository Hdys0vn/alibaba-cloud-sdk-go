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

// StartDiskReplicaPair invokes the ebs.StartDiskReplicaPair API synchronously
func (client *Client) StartDiskReplicaPair(request *StartDiskReplicaPairRequest) (response *StartDiskReplicaPairResponse, err error) {
	response = CreateStartDiskReplicaPairResponse()
	err = client.DoAction(request, response)
	return
}

// StartDiskReplicaPairWithChan invokes the ebs.StartDiskReplicaPair API asynchronously
func (client *Client) StartDiskReplicaPairWithChan(request *StartDiskReplicaPairRequest) (<-chan *StartDiskReplicaPairResponse, <-chan error) {
	responseChan := make(chan *StartDiskReplicaPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartDiskReplicaPair(request)
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

// StartDiskReplicaPairWithCallback invokes the ebs.StartDiskReplicaPair API asynchronously
func (client *Client) StartDiskReplicaPairWithCallback(request *StartDiskReplicaPairRequest, callback func(response *StartDiskReplicaPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartDiskReplicaPairResponse
		var err error
		defer close(result)
		response, err = client.StartDiskReplicaPair(request)
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

// StartDiskReplicaPairRequest is the request struct for api StartDiskReplicaPair
type StartDiskReplicaPairRequest struct {
	*requests.RpcRequest
	ClientToken   string           `position:"Query" name:"ClientToken"`
	OneShot       requests.Boolean `position:"Query" name:"OneShot"`
	ReplicaPairId string           `position:"Query" name:"ReplicaPairId"`
}

// StartDiskReplicaPairResponse is the response struct for api StartDiskReplicaPair
type StartDiskReplicaPairResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartDiskReplicaPairRequest creates a request to invoke StartDiskReplicaPair API
func CreateStartDiskReplicaPairRequest() (request *StartDiskReplicaPairRequest) {
	request = &StartDiskReplicaPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ebs", "2021-07-30", "StartDiskReplicaPair", "ebs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartDiskReplicaPairResponse creates a response to parse from StartDiskReplicaPair response
func CreateStartDiskReplicaPairResponse() (response *StartDiskReplicaPairResponse) {
	response = &StartDiskReplicaPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
