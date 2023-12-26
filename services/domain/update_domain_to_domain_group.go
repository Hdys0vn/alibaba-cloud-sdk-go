package domain

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

// UpdateDomainToDomainGroup invokes the domain.UpdateDomainToDomainGroup API synchronously
func (client *Client) UpdateDomainToDomainGroup(request *UpdateDomainToDomainGroupRequest) (response *UpdateDomainToDomainGroupResponse, err error) {
	response = CreateUpdateDomainToDomainGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDomainToDomainGroupWithChan invokes the domain.UpdateDomainToDomainGroup API asynchronously
func (client *Client) UpdateDomainToDomainGroupWithChan(request *UpdateDomainToDomainGroupRequest) (<-chan *UpdateDomainToDomainGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateDomainToDomainGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDomainToDomainGroup(request)
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

// UpdateDomainToDomainGroupWithCallback invokes the domain.UpdateDomainToDomainGroup API asynchronously
func (client *Client) UpdateDomainToDomainGroupWithCallback(request *UpdateDomainToDomainGroupRequest, callback func(response *UpdateDomainToDomainGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDomainToDomainGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateDomainToDomainGroup(request)
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

// UpdateDomainToDomainGroupRequest is the request struct for api UpdateDomainToDomainGroup
type UpdateDomainToDomainGroupRequest struct {
	*requests.RpcRequest
	FileToUpload  string           `position:"Body" name:"FileToUpload"`
	Replace       requests.Boolean `position:"Query" name:"Replace"`
	DomainName    *[]string        `position:"Query" name:"DomainName"  type:"Repeated"`
	DomainGroupId requests.Integer `position:"Query" name:"DomainGroupId"`
	DataSource    requests.Integer `position:"Query" name:"DataSource"`
	UserClientIp  string           `position:"Query" name:"UserClientIp"`
	Lang          string           `position:"Query" name:"Lang"`
}

// UpdateDomainToDomainGroupResponse is the response struct for api UpdateDomainToDomainGroup
type UpdateDomainToDomainGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDomainToDomainGroupRequest creates a request to invoke UpdateDomainToDomainGroup API
func CreateUpdateDomainToDomainGroupRequest() (request *UpdateDomainToDomainGroupRequest) {
	request = &UpdateDomainToDomainGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "UpdateDomainToDomainGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDomainToDomainGroupResponse creates a response to parse from UpdateDomainToDomainGroup response
func CreateUpdateDomainToDomainGroupResponse() (response *UpdateDomainToDomainGroupResponse) {
	response = &UpdateDomainToDomainGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}