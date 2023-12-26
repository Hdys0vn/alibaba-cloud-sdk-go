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

// ValidateShrinkNodes invokes the elasticsearch.ValidateShrinkNodes API synchronously
func (client *Client) ValidateShrinkNodes(request *ValidateShrinkNodesRequest) (response *ValidateShrinkNodesResponse, err error) {
	response = CreateValidateShrinkNodesResponse()
	err = client.DoAction(request, response)
	return
}

// ValidateShrinkNodesWithChan invokes the elasticsearch.ValidateShrinkNodes API asynchronously
func (client *Client) ValidateShrinkNodesWithChan(request *ValidateShrinkNodesRequest) (<-chan *ValidateShrinkNodesResponse, <-chan error) {
	responseChan := make(chan *ValidateShrinkNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ValidateShrinkNodes(request)
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

// ValidateShrinkNodesWithCallback invokes the elasticsearch.ValidateShrinkNodes API asynchronously
func (client *Client) ValidateShrinkNodesWithCallback(request *ValidateShrinkNodesRequest, callback func(response *ValidateShrinkNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ValidateShrinkNodesResponse
		var err error
		defer close(result)
		response, err = client.ValidateShrinkNodes(request)
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

// ValidateShrinkNodesRequest is the request struct for api ValidateShrinkNodes
type ValidateShrinkNodesRequest struct {
	*requests.RoaRequest
	IgnoreStatus requests.Boolean `position:"Query" name:"ignoreStatus"`
	InstanceId   string           `position:"Path" name:"InstanceId"`
	NodeType     string           `position:"Query" name:"nodeType"`
}

// ValidateShrinkNodesResponse is the response struct for api ValidateShrinkNodes
type ValidateShrinkNodesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateValidateShrinkNodesRequest creates a request to invoke ValidateShrinkNodes API
func CreateValidateShrinkNodesRequest() (request *ValidateShrinkNodesRequest) {
	request = &ValidateShrinkNodesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ValidateShrinkNodes", "/openapi/instances/[InstanceId]/validate-shrink-nodes", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateValidateShrinkNodesResponse creates a response to parse from ValidateShrinkNodes response
func CreateValidateShrinkNodesResponse() (response *ValidateShrinkNodesResponse) {
	response = &ValidateShrinkNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
