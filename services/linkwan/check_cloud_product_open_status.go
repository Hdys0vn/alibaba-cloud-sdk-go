package linkwan

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

// CheckCloudProductOpenStatus invokes the linkwan.CheckCloudProductOpenStatus API synchronously
func (client *Client) CheckCloudProductOpenStatus(request *CheckCloudProductOpenStatusRequest) (response *CheckCloudProductOpenStatusResponse, err error) {
	response = CreateCheckCloudProductOpenStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckCloudProductOpenStatusWithChan invokes the linkwan.CheckCloudProductOpenStatus API asynchronously
func (client *Client) CheckCloudProductOpenStatusWithChan(request *CheckCloudProductOpenStatusRequest) (<-chan *CheckCloudProductOpenStatusResponse, <-chan error) {
	responseChan := make(chan *CheckCloudProductOpenStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckCloudProductOpenStatus(request)
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

// CheckCloudProductOpenStatusWithCallback invokes the linkwan.CheckCloudProductOpenStatus API asynchronously
func (client *Client) CheckCloudProductOpenStatusWithCallback(request *CheckCloudProductOpenStatusRequest, callback func(response *CheckCloudProductOpenStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckCloudProductOpenStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckCloudProductOpenStatus(request)
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

// CheckCloudProductOpenStatusRequest is the request struct for api CheckCloudProductOpenStatus
type CheckCloudProductOpenStatusRequest struct {
	*requests.RpcRequest
	ServiceCode string `position:"Query" name:"ServiceCode"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// CheckCloudProductOpenStatusResponse is the response struct for api CheckCloudProductOpenStatus
type CheckCloudProductOpenStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateCheckCloudProductOpenStatusRequest creates a request to invoke CheckCloudProductOpenStatus API
func CreateCheckCloudProductOpenStatusRequest() (request *CheckCloudProductOpenStatusRequest) {
	request = &CheckCloudProductOpenStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "CheckCloudProductOpenStatus", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckCloudProductOpenStatusResponse creates a response to parse from CheckCloudProductOpenStatus response
func CreateCheckCloudProductOpenStatusResponse() (response *CheckCloudProductOpenStatusResponse) {
	response = &CheckCloudProductOpenStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
