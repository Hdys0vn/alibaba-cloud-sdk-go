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

// StopDiskReplicaPair invokes the ebs.StopDiskReplicaPair API synchronously
func (client *Client) StopDiskReplicaPair(request *StopDiskReplicaPairRequest) (response *StopDiskReplicaPairResponse, err error) {
	response = CreateStopDiskReplicaPairResponse()
	err = client.DoAction(request, response)
	return
}

// StopDiskReplicaPairWithChan invokes the ebs.StopDiskReplicaPair API asynchronously
func (client *Client) StopDiskReplicaPairWithChan(request *StopDiskReplicaPairRequest) (<-chan *StopDiskReplicaPairResponse, <-chan error) {
	responseChan := make(chan *StopDiskReplicaPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDiskReplicaPair(request)
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

// StopDiskReplicaPairWithCallback invokes the ebs.StopDiskReplicaPair API asynchronously
func (client *Client) StopDiskReplicaPairWithCallback(request *StopDiskReplicaPairRequest, callback func(response *StopDiskReplicaPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDiskReplicaPairResponse
		var err error
		defer close(result)
		response, err = client.StopDiskReplicaPair(request)
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

// StopDiskReplicaPairRequest is the request struct for api StopDiskReplicaPair
type StopDiskReplicaPairRequest struct {
	*requests.RpcRequest
	ClientToken   string `position:"Query" name:"ClientToken"`
	ReplicaPairId string `position:"Query" name:"ReplicaPairId"`
}

// StopDiskReplicaPairResponse is the response struct for api StopDiskReplicaPair
type StopDiskReplicaPairResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopDiskReplicaPairRequest creates a request to invoke StopDiskReplicaPair API
func CreateStopDiskReplicaPairRequest() (request *StopDiskReplicaPairRequest) {
	request = &StopDiskReplicaPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ebs", "2021-07-30", "StopDiskReplicaPair", "ebs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopDiskReplicaPairResponse creates a response to parse from StopDiskReplicaPair response
func CreateStopDiskReplicaPairResponse() (response *StopDiskReplicaPairResponse) {
	response = &StopDiskReplicaPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
