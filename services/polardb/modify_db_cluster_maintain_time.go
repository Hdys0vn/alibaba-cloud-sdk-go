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

// ModifyDBClusterMaintainTime invokes the polardb.ModifyDBClusterMaintainTime API synchronously
func (client *Client) ModifyDBClusterMaintainTime(request *ModifyDBClusterMaintainTimeRequest) (response *ModifyDBClusterMaintainTimeResponse, err error) {
	response = CreateModifyDBClusterMaintainTimeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterMaintainTimeWithChan invokes the polardb.ModifyDBClusterMaintainTime API asynchronously
func (client *Client) ModifyDBClusterMaintainTimeWithChan(request *ModifyDBClusterMaintainTimeRequest) (<-chan *ModifyDBClusterMaintainTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterMaintainTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterMaintainTime(request)
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

// ModifyDBClusterMaintainTimeWithCallback invokes the polardb.ModifyDBClusterMaintainTime API asynchronously
func (client *Client) ModifyDBClusterMaintainTimeWithCallback(request *ModifyDBClusterMaintainTimeRequest, callback func(response *ModifyDBClusterMaintainTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterMaintainTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterMaintainTime(request)
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

// ModifyDBClusterMaintainTimeRequest is the request struct for api ModifyDBClusterMaintainTime
type ModifyDBClusterMaintainTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MaintainTime         string           `position:"Query" name:"MaintainTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDBClusterMaintainTimeResponse is the response struct for api ModifyDBClusterMaintainTime
type ModifyDBClusterMaintainTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterMaintainTimeRequest creates a request to invoke ModifyDBClusterMaintainTime API
func CreateModifyDBClusterMaintainTimeRequest() (request *ModifyDBClusterMaintainTimeRequest) {
	request = &ModifyDBClusterMaintainTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterMaintainTime", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterMaintainTimeResponse creates a response to parse from ModifyDBClusterMaintainTime response
func CreateModifyDBClusterMaintainTimeResponse() (response *ModifyDBClusterMaintainTimeResponse) {
	response = &ModifyDBClusterMaintainTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
