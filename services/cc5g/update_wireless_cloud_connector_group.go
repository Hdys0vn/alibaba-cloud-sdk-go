package cc5g

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

// UpdateWirelessCloudConnectorGroup invokes the cc5g.UpdateWirelessCloudConnectorGroup API synchronously
func (client *Client) UpdateWirelessCloudConnectorGroup(request *UpdateWirelessCloudConnectorGroupRequest) (response *UpdateWirelessCloudConnectorGroupResponse, err error) {
	response = CreateUpdateWirelessCloudConnectorGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWirelessCloudConnectorGroupWithChan invokes the cc5g.UpdateWirelessCloudConnectorGroup API asynchronously
func (client *Client) UpdateWirelessCloudConnectorGroupWithChan(request *UpdateWirelessCloudConnectorGroupRequest) (<-chan *UpdateWirelessCloudConnectorGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateWirelessCloudConnectorGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWirelessCloudConnectorGroup(request)
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

// UpdateWirelessCloudConnectorGroupWithCallback invokes the cc5g.UpdateWirelessCloudConnectorGroup API asynchronously
func (client *Client) UpdateWirelessCloudConnectorGroupWithCallback(request *UpdateWirelessCloudConnectorGroupRequest, callback func(response *UpdateWirelessCloudConnectorGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWirelessCloudConnectorGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateWirelessCloudConnectorGroup(request)
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

// UpdateWirelessCloudConnectorGroupRequest is the request struct for api UpdateWirelessCloudConnectorGroup
type UpdateWirelessCloudConnectorGroupRequest struct {
	*requests.RpcRequest
	WirelessCloudConnectorGroupId string           `position:"Query" name:"WirelessCloudConnectorGroupId"`
	DryRun                        requests.Boolean `position:"Query" name:"DryRun"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	Description                   string           `position:"Query" name:"Description"`
	Name                          string           `position:"Query" name:"Name"`
}

// UpdateWirelessCloudConnectorGroupResponse is the response struct for api UpdateWirelessCloudConnectorGroup
type UpdateWirelessCloudConnectorGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateWirelessCloudConnectorGroupRequest creates a request to invoke UpdateWirelessCloudConnectorGroup API
func CreateUpdateWirelessCloudConnectorGroupRequest() (request *UpdateWirelessCloudConnectorGroupRequest) {
	request = &UpdateWirelessCloudConnectorGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "UpdateWirelessCloudConnectorGroup", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateWirelessCloudConnectorGroupResponse creates a response to parse from UpdateWirelessCloudConnectorGroup response
func CreateUpdateWirelessCloudConnectorGroupResponse() (response *UpdateWirelessCloudConnectorGroupResponse) {
	response = &UpdateWirelessCloudConnectorGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
