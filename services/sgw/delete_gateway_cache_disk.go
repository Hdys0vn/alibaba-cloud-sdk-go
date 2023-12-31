package sgw

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

// DeleteGatewayCacheDisk invokes the sgw.DeleteGatewayCacheDisk API synchronously
func (client *Client) DeleteGatewayCacheDisk(request *DeleteGatewayCacheDiskRequest) (response *DeleteGatewayCacheDiskResponse, err error) {
	response = CreateDeleteGatewayCacheDiskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGatewayCacheDiskWithChan invokes the sgw.DeleteGatewayCacheDisk API asynchronously
func (client *Client) DeleteGatewayCacheDiskWithChan(request *DeleteGatewayCacheDiskRequest) (<-chan *DeleteGatewayCacheDiskResponse, <-chan error) {
	responseChan := make(chan *DeleteGatewayCacheDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGatewayCacheDisk(request)
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

// DeleteGatewayCacheDiskWithCallback invokes the sgw.DeleteGatewayCacheDisk API asynchronously
func (client *Client) DeleteGatewayCacheDiskWithCallback(request *DeleteGatewayCacheDiskRequest, callback func(response *DeleteGatewayCacheDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGatewayCacheDiskResponse
		var err error
		defer close(result)
		response, err = client.DeleteGatewayCacheDisk(request)
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

// DeleteGatewayCacheDiskRequest is the request struct for api DeleteGatewayCacheDisk
type DeleteGatewayCacheDiskRequest struct {
	*requests.RpcRequest
	LocalFilePath string `position:"Query" name:"LocalFilePath"`
	CacheId       string `position:"Query" name:"CacheId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// DeleteGatewayCacheDiskResponse is the response struct for api DeleteGatewayCacheDisk
type DeleteGatewayCacheDiskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateDeleteGatewayCacheDiskRequest creates a request to invoke DeleteGatewayCacheDisk API
func CreateDeleteGatewayCacheDiskRequest() (request *DeleteGatewayCacheDiskRequest) {
	request = &DeleteGatewayCacheDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DeleteGatewayCacheDisk", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGatewayCacheDiskResponse creates a response to parse from DeleteGatewayCacheDisk response
func CreateDeleteGatewayCacheDiskResponse() (response *DeleteGatewayCacheDiskResponse) {
	response = &DeleteGatewayCacheDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
