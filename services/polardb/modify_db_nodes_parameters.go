package polardb

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

// ModifyDBNodesParameters invokes the polardb.ModifyDBNodesParameters API synchronously
func (client *Client) ModifyDBNodesParameters(request *ModifyDBNodesParametersRequest) (response *ModifyDBNodesParametersResponse, err error) {
	response = CreateModifyDBNodesParametersResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBNodesParametersWithChan invokes the polardb.ModifyDBNodesParameters API asynchronously
func (client *Client) ModifyDBNodesParametersWithChan(request *ModifyDBNodesParametersRequest) (<-chan *ModifyDBNodesParametersResponse, <-chan error) {
	responseChan := make(chan *ModifyDBNodesParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBNodesParameters(request)
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

// ModifyDBNodesParametersWithCallback invokes the polardb.ModifyDBNodesParameters API asynchronously
func (client *Client) ModifyDBNodesParametersWithCallback(request *ModifyDBNodesParametersRequest, callback func(response *ModifyDBNodesParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBNodesParametersResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBNodesParameters(request)
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

// ModifyDBNodesParametersRequest is the request struct for api ModifyDBNodesParameters
type ModifyDBNodesParametersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PlannedEndTime       string           `position:"Query" name:"PlannedEndTime"`
	DBNodeIds            string           `position:"Query" name:"DBNodeIds"`
	ParameterGroupId     string           `position:"Query" name:"ParameterGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PlannedStartTime     string           `position:"Query" name:"PlannedStartTime"`
	Parameters           string           `position:"Query" name:"Parameters"`
	FromTimeService      requests.Boolean `position:"Query" name:"FromTimeService"`
}

// ModifyDBNodesParametersResponse is the response struct for api ModifyDBNodesParameters
type ModifyDBNodesParametersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBNodesParametersRequest creates a request to invoke ModifyDBNodesParameters API
func CreateModifyDBNodesParametersRequest() (request *ModifyDBNodesParametersRequest) {
	request = &ModifyDBNodesParametersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBNodesParameters", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBNodesParametersResponse creates a response to parse from ModifyDBNodesParameters response
func CreateModifyDBNodesParametersResponse() (response *ModifyDBNodesParametersResponse) {
	response = &ModifyDBNodesParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
