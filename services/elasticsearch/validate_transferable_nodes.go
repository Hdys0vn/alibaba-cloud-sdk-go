package elasticsearch

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

// ValidateTransferableNodes invokes the elasticsearch.ValidateTransferableNodes API synchronously
func (client *Client) ValidateTransferableNodes(request *ValidateTransferableNodesRequest) (response *ValidateTransferableNodesResponse, err error) {
	response = CreateValidateTransferableNodesResponse()
	err = client.DoAction(request, response)
	return
}

// ValidateTransferableNodesWithChan invokes the elasticsearch.ValidateTransferableNodes API asynchronously
func (client *Client) ValidateTransferableNodesWithChan(request *ValidateTransferableNodesRequest) (<-chan *ValidateTransferableNodesResponse, <-chan error) {
	responseChan := make(chan *ValidateTransferableNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ValidateTransferableNodes(request)
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

// ValidateTransferableNodesWithCallback invokes the elasticsearch.ValidateTransferableNodes API asynchronously
func (client *Client) ValidateTransferableNodesWithCallback(request *ValidateTransferableNodesRequest, callback func(response *ValidateTransferableNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ValidateTransferableNodesResponse
		var err error
		defer close(result)
		response, err = client.ValidateTransferableNodes(request)
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

// ValidateTransferableNodesRequest is the request struct for api ValidateTransferableNodes
type ValidateTransferableNodesRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	NodeType   string `position:"Query" name:"nodeType"`
}

// ValidateTransferableNodesResponse is the response struct for api ValidateTransferableNodes
type ValidateTransferableNodesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateValidateTransferableNodesRequest creates a request to invoke ValidateTransferableNodes API
func CreateValidateTransferableNodesRequest() (request *ValidateTransferableNodesRequest) {
	request = &ValidateTransferableNodesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ValidateTransferableNodes", "/openapi/instances/[InstanceId]/validate-transfer-nodes", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateValidateTransferableNodesResponse creates a response to parse from ValidateTransferableNodes response
func CreateValidateTransferableNodesResponse() (response *ValidateTransferableNodesResponse) {
	response = &ValidateTransferableNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
