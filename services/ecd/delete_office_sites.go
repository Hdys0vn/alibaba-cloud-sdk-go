package ecd

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

// DeleteOfficeSites invokes the ecd.DeleteOfficeSites API synchronously
func (client *Client) DeleteOfficeSites(request *DeleteOfficeSitesRequest) (response *DeleteOfficeSitesResponse, err error) {
	response = CreateDeleteOfficeSitesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteOfficeSitesWithChan invokes the ecd.DeleteOfficeSites API asynchronously
func (client *Client) DeleteOfficeSitesWithChan(request *DeleteOfficeSitesRequest) (<-chan *DeleteOfficeSitesResponse, <-chan error) {
	responseChan := make(chan *DeleteOfficeSitesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteOfficeSites(request)
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

// DeleteOfficeSitesWithCallback invokes the ecd.DeleteOfficeSites API asynchronously
func (client *Client) DeleteOfficeSitesWithCallback(request *DeleteOfficeSitesRequest, callback func(response *DeleteOfficeSitesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteOfficeSitesResponse
		var err error
		defer close(result)
		response, err = client.DeleteOfficeSites(request)
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

// DeleteOfficeSitesRequest is the request struct for api DeleteOfficeSites
type DeleteOfficeSitesRequest struct {
	*requests.RpcRequest
	OfficeSiteId *[]string `position:"Query" name:"OfficeSiteId"  type:"Repeated"`
}

// DeleteOfficeSitesResponse is the response struct for api DeleteOfficeSites
type DeleteOfficeSitesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteOfficeSitesRequest creates a request to invoke DeleteOfficeSites API
func CreateDeleteOfficeSitesRequest() (request *DeleteOfficeSitesRequest) {
	request = &DeleteOfficeSitesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DeleteOfficeSites", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteOfficeSitesResponse creates a response to parse from DeleteOfficeSites response
func CreateDeleteOfficeSitesResponse() (response *DeleteOfficeSitesResponse) {
	response = &DeleteOfficeSitesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
