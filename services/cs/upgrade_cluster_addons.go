package cs

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

// UpgradeClusterAddons invokes the cs.UpgradeClusterAddons API synchronously
func (client *Client) UpgradeClusterAddons(request *UpgradeClusterAddonsRequest) (response *UpgradeClusterAddonsResponse, err error) {
	response = CreateUpgradeClusterAddonsResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeClusterAddonsWithChan invokes the cs.UpgradeClusterAddons API asynchronously
func (client *Client) UpgradeClusterAddonsWithChan(request *UpgradeClusterAddonsRequest) (<-chan *UpgradeClusterAddonsResponse, <-chan error) {
	responseChan := make(chan *UpgradeClusterAddonsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeClusterAddons(request)
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

// UpgradeClusterAddonsWithCallback invokes the cs.UpgradeClusterAddons API asynchronously
func (client *Client) UpgradeClusterAddonsWithCallback(request *UpgradeClusterAddonsRequest, callback func(response *UpgradeClusterAddonsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeClusterAddonsResponse
		var err error
		defer close(result)
		response, err = client.UpgradeClusterAddons(request)
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

// UpgradeClusterAddonsRequest is the request struct for api UpgradeClusterAddons
type UpgradeClusterAddonsRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// UpgradeClusterAddonsResponse is the response struct for api UpgradeClusterAddons
type UpgradeClusterAddonsResponse struct {
	*responses.BaseResponse
}

// CreateUpgradeClusterAddonsRequest creates a request to invoke UpgradeClusterAddons API
func CreateUpgradeClusterAddonsRequest() (request *UpgradeClusterAddonsRequest) {
	request = &UpgradeClusterAddonsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "UpgradeClusterAddons", "/clusters/[ClusterId]/components/upgrade", "", "")
	request.Method = requests.POST
	return
}

// CreateUpgradeClusterAddonsResponse creates a response to parse from UpgradeClusterAddons response
func CreateUpgradeClusterAddonsResponse() (response *UpgradeClusterAddonsResponse) {
	response = &UpgradeClusterAddonsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
