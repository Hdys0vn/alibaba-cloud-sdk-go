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

// GetTransferableNodes invokes the elasticsearch.GetTransferableNodes API synchronously
func (client *Client) GetTransferableNodes(request *GetTransferableNodesRequest) (response *GetTransferableNodesResponse, err error) {
	response = CreateGetTransferableNodesResponse()
	err = client.DoAction(request, response)
	return
}

// GetTransferableNodesWithChan invokes the elasticsearch.GetTransferableNodes API asynchronously
func (client *Client) GetTransferableNodesWithChan(request *GetTransferableNodesRequest) (<-chan *GetTransferableNodesResponse, <-chan error) {
	responseChan := make(chan *GetTransferableNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTransferableNodes(request)
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

// GetTransferableNodesWithCallback invokes the elasticsearch.GetTransferableNodes API asynchronously
func (client *Client) GetTransferableNodesWithCallback(request *GetTransferableNodesRequest, callback func(response *GetTransferableNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTransferableNodesResponse
		var err error
		defer close(result)
		response, err = client.GetTransferableNodes(request)
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

// GetTransferableNodesRequest is the request struct for api GetTransferableNodes
type GetTransferableNodesRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	NodeType   string           `position:"Query" name:"nodeType"`
	Count      requests.Integer `position:"Query" name:"count"`
}

// GetTransferableNodesResponse is the response struct for api GetTransferableNodes
type GetTransferableNodesResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateGetTransferableNodesRequest creates a request to invoke GetTransferableNodes API
func CreateGetTransferableNodesRequest() (request *GetTransferableNodesRequest) {
	request = &GetTransferableNodesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "GetTransferableNodes", "/openapi/instances/[InstanceId]/transferable-nodes", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetTransferableNodesResponse creates a response to parse from GetTransferableNodes response
func CreateGetTransferableNodesResponse() (response *GetTransferableNodesResponse) {
	response = &GetTransferableNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
