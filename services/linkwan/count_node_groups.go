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

// CountNodeGroups invokes the linkwan.CountNodeGroups API synchronously
func (client *Client) CountNodeGroups(request *CountNodeGroupsRequest) (response *CountNodeGroupsResponse, err error) {
	response = CreateCountNodeGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// CountNodeGroupsWithChan invokes the linkwan.CountNodeGroups API asynchronously
func (client *Client) CountNodeGroupsWithChan(request *CountNodeGroupsRequest) (<-chan *CountNodeGroupsResponse, <-chan error) {
	responseChan := make(chan *CountNodeGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CountNodeGroups(request)
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

// CountNodeGroupsWithCallback invokes the linkwan.CountNodeGroups API asynchronously
func (client *Client) CountNodeGroupsWithCallback(request *CountNodeGroupsRequest, callback func(response *CountNodeGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CountNodeGroupsResponse
		var err error
		defer close(result)
		response, err = client.CountNodeGroups(request)
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

// CountNodeGroupsRequest is the request struct for api CountNodeGroups
type CountNodeGroupsRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	FuzzyJoinEui  string `position:"Query" name:"FuzzyJoinEui"`
	FuzzyDevEui   string `position:"Query" name:"FuzzyDevEui"`
	FuzzyName     string `position:"Query" name:"FuzzyName"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// CountNodeGroupsResponse is the response struct for api CountNodeGroups
type CountNodeGroupsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateCountNodeGroupsRequest creates a request to invoke CountNodeGroups API
func CreateCountNodeGroupsRequest() (request *CountNodeGroupsRequest) {
	request = &CountNodeGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "CountNodeGroups", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCountNodeGroupsResponse creates a response to parse from CountNodeGroups response
func CreateCountNodeGroupsResponse() (response *CountNodeGroupsResponse) {
	response = &CountNodeGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
