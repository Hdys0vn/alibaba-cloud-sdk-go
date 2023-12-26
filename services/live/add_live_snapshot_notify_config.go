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

// AddLiveSnapshotNotifyConfig invokes the live.AddLiveSnapshotNotifyConfig API synchronously
func (client *Client) AddLiveSnapshotNotifyConfig(request *AddLiveSnapshotNotifyConfigRequest) (response *AddLiveSnapshotNotifyConfigResponse, err error) {
	response = CreateAddLiveSnapshotNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AddLiveSnapshotNotifyConfigWithChan invokes the live.AddLiveSnapshotNotifyConfig API asynchronously
func (client *Client) AddLiveSnapshotNotifyConfigWithChan(request *AddLiveSnapshotNotifyConfigRequest) (<-chan *AddLiveSnapshotNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *AddLiveSnapshotNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveSnapshotNotifyConfig(request)
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

// AddLiveSnapshotNotifyConfigWithCallback invokes the live.AddLiveSnapshotNotifyConfig API asynchronously
func (client *Client) AddLiveSnapshotNotifyConfigWithCallback(request *AddLiveSnapshotNotifyConfigRequest, callback func(response *AddLiveSnapshotNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveSnapshotNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.AddLiveSnapshotNotifyConfig(request)
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

// AddLiveSnapshotNotifyConfigRequest is the request struct for api AddLiveSnapshotNotifyConfig
type AddLiveSnapshotNotifyConfigRequest struct {
	*requests.RpcRequest
	NotifyReqAuth string           `position:"Query" name:"NotifyReqAuth"`
	NotifyUrl     string           `position:"Query" name:"NotifyUrl"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	NotifyAuthKey string           `position:"Query" name:"NotifyAuthKey"`
}

// AddLiveSnapshotNotifyConfigResponse is the response struct for api AddLiveSnapshotNotifyConfig
type AddLiveSnapshotNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLiveSnapshotNotifyConfigRequest creates a request to invoke AddLiveSnapshotNotifyConfig API
func CreateAddLiveSnapshotNotifyConfigRequest() (request *AddLiveSnapshotNotifyConfigRequest) {
	request = &AddLiveSnapshotNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddLiveSnapshotNotifyConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddLiveSnapshotNotifyConfigResponse creates a response to parse from AddLiveSnapshotNotifyConfig response
func CreateAddLiveSnapshotNotifyConfigResponse() (response *AddLiveSnapshotNotifyConfigResponse) {
	response = &AddLiveSnapshotNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}