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

// DeleteDiskReplicaPair invokes the ebs.DeleteDiskReplicaPair API synchronously
func (client *Client) DeleteDiskReplicaPair(request *DeleteDiskReplicaPairRequest) (response *DeleteDiskReplicaPairResponse, err error) {
	response = CreateDeleteDiskReplicaPairResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDiskReplicaPairWithChan invokes the ebs.DeleteDiskReplicaPair API asynchronously
func (client *Client) DeleteDiskReplicaPairWithChan(request *DeleteDiskReplicaPairRequest) (<-chan *DeleteDiskReplicaPairResponse, <-chan error) {
	responseChan := make(chan *DeleteDiskReplicaPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDiskReplicaPair(request)
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

// DeleteDiskReplicaPairWithCallback invokes the ebs.DeleteDiskReplicaPair API asynchronously
func (client *Client) DeleteDiskReplicaPairWithCallback(request *DeleteDiskReplicaPairRequest, callback func(response *DeleteDiskReplicaPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDiskReplicaPairResponse
		var err error
		defer close(result)
		response, err = client.DeleteDiskReplicaPair(request)
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

// DeleteDiskReplicaPairRequest is the request struct for api DeleteDiskReplicaPair
type DeleteDiskReplicaPairRequest struct {
	*requests.RpcRequest
	ClientToken   string `position:"Query" name:"ClientToken"`
	ReplicaPairId string `position:"Query" name:"ReplicaPairId"`
}

// DeleteDiskReplicaPairResponse is the response struct for api DeleteDiskReplicaPair
type DeleteDiskReplicaPairResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDiskReplicaPairRequest creates a request to invoke DeleteDiskReplicaPair API
func CreateDeleteDiskReplicaPairRequest() (request *DeleteDiskReplicaPairRequest) {
	request = &DeleteDiskReplicaPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ebs", "2021-07-30", "DeleteDiskReplicaPair", "ebs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDiskReplicaPairResponse creates a response to parse from DeleteDiskReplicaPair response
func CreateDeleteDiskReplicaPairResponse() (response *DeleteDiskReplicaPairResponse) {
	response = &DeleteDiskReplicaPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
