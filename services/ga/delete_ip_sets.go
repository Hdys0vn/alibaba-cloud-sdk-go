package ga

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

// DeleteIpSets invokes the ga.DeleteIpSets API synchronously
func (client *Client) DeleteIpSets(request *DeleteIpSetsRequest) (response *DeleteIpSetsResponse, err error) {
	response = CreateDeleteIpSetsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteIpSetsWithChan invokes the ga.DeleteIpSets API asynchronously
func (client *Client) DeleteIpSetsWithChan(request *DeleteIpSetsRequest) (<-chan *DeleteIpSetsResponse, <-chan error) {
	responseChan := make(chan *DeleteIpSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteIpSets(request)
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

// DeleteIpSetsWithCallback invokes the ga.DeleteIpSets API asynchronously
func (client *Client) DeleteIpSetsWithCallback(request *DeleteIpSetsRequest, callback func(response *DeleteIpSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteIpSetsResponse
		var err error
		defer close(result)
		response, err = client.DeleteIpSets(request)
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

// DeleteIpSetsRequest is the request struct for api DeleteIpSets
type DeleteIpSetsRequest struct {
	*requests.RpcRequest
	IpSetIds *[]string `position:"Query" name:"IpSetIds"  type:"Repeated"`
}

// DeleteIpSetsResponse is the response struct for api DeleteIpSets
type DeleteIpSetsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteIpSetsRequest creates a request to invoke DeleteIpSets API
func CreateDeleteIpSetsRequest() (request *DeleteIpSetsRequest) {
	request = &DeleteIpSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "DeleteIpSets", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteIpSetsResponse creates a response to parse from DeleteIpSets response
func CreateDeleteIpSetsResponse() (response *DeleteIpSetsResponse) {
	response = &DeleteIpSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}