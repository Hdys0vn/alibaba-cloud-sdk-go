package ens

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

// UpgradeAICInstanceImage invokes the ens.UpgradeAICInstanceImage API synchronously
func (client *Client) UpgradeAICInstanceImage(request *UpgradeAICInstanceImageRequest) (response *UpgradeAICInstanceImageResponse, err error) {
	response = CreateUpgradeAICInstanceImageResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeAICInstanceImageWithChan invokes the ens.UpgradeAICInstanceImage API asynchronously
func (client *Client) UpgradeAICInstanceImageWithChan(request *UpgradeAICInstanceImageRequest) (<-chan *UpgradeAICInstanceImageResponse, <-chan error) {
	responseChan := make(chan *UpgradeAICInstanceImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeAICInstanceImage(request)
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

// UpgradeAICInstanceImageWithCallback invokes the ens.UpgradeAICInstanceImage API asynchronously
func (client *Client) UpgradeAICInstanceImageWithCallback(request *UpgradeAICInstanceImageRequest, callback func(response *UpgradeAICInstanceImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeAICInstanceImageResponse
		var err error
		defer close(result)
		response, err = client.UpgradeAICInstanceImage(request)
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

// UpgradeAICInstanceImageRequest is the request struct for api UpgradeAICInstanceImage
type UpgradeAICInstanceImageRequest struct {
	*requests.RpcRequest
	ImageId   string           `position:"Query" name:"ImageId"`
	Timeout   requests.Integer `position:"Query" name:"Timeout"`
	ServerIds *[]string        `position:"Query" name:"ServerIds"  type:"Json"`
}

// UpgradeAICInstanceImageResponse is the response struct for api UpgradeAICInstanceImage
type UpgradeAICInstanceImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpgradeAICInstanceImageRequest creates a request to invoke UpgradeAICInstanceImage API
func CreateUpgradeAICInstanceImageRequest() (request *UpgradeAICInstanceImageRequest) {
	request = &UpgradeAICInstanceImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "UpgradeAICInstanceImage", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateUpgradeAICInstanceImageResponse creates a response to parse from UpgradeAICInstanceImage response
func CreateUpgradeAICInstanceImageResponse() (response *UpgradeAICInstanceImageResponse) {
	response = &UpgradeAICInstanceImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
