package live

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

// DeleteLiveSnapshotNotifyConfig invokes the live.DeleteLiveSnapshotNotifyConfig API synchronously
func (client *Client) DeleteLiveSnapshotNotifyConfig(request *DeleteLiveSnapshotNotifyConfigRequest) (response *DeleteLiveSnapshotNotifyConfigResponse, err error) {
	response = CreateDeleteLiveSnapshotNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveSnapshotNotifyConfigWithChan invokes the live.DeleteLiveSnapshotNotifyConfig API asynchronously
func (client *Client) DeleteLiveSnapshotNotifyConfigWithChan(request *DeleteLiveSnapshotNotifyConfigRequest) (<-chan *DeleteLiveSnapshotNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveSnapshotNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveSnapshotNotifyConfig(request)
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

// DeleteLiveSnapshotNotifyConfigWithCallback invokes the live.DeleteLiveSnapshotNotifyConfig API asynchronously
func (client *Client) DeleteLiveSnapshotNotifyConfigWithCallback(request *DeleteLiveSnapshotNotifyConfigRequest, callback func(response *DeleteLiveSnapshotNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveSnapshotNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveSnapshotNotifyConfig(request)
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

// DeleteLiveSnapshotNotifyConfigRequest is the request struct for api DeleteLiveSnapshotNotifyConfig
type DeleteLiveSnapshotNotifyConfigRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLiveSnapshotNotifyConfigResponse is the response struct for api DeleteLiveSnapshotNotifyConfig
type DeleteLiveSnapshotNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveSnapshotNotifyConfigRequest creates a request to invoke DeleteLiveSnapshotNotifyConfig API
func CreateDeleteLiveSnapshotNotifyConfigRequest() (request *DeleteLiveSnapshotNotifyConfigRequest) {
	request = &DeleteLiveSnapshotNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveSnapshotNotifyConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveSnapshotNotifyConfigResponse creates a response to parse from DeleteLiveSnapshotNotifyConfig response
func CreateDeleteLiveSnapshotNotifyConfigResponse() (response *DeleteLiveSnapshotNotifyConfigResponse) {
	response = &DeleteLiveSnapshotNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
