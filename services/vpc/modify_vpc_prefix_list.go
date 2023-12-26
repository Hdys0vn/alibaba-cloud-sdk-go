package vpc

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

// ModifyVpcPrefixList invokes the vpc.ModifyVpcPrefixList API synchronously
func (client *Client) ModifyVpcPrefixList(request *ModifyVpcPrefixListRequest) (response *ModifyVpcPrefixListResponse, err error) {
	response = CreateModifyVpcPrefixListResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpcPrefixListWithChan invokes the vpc.ModifyVpcPrefixList API asynchronously
func (client *Client) ModifyVpcPrefixListWithChan(request *ModifyVpcPrefixListRequest) (<-chan *ModifyVpcPrefixListResponse, <-chan error) {
	responseChan := make(chan *ModifyVpcPrefixListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpcPrefixList(request)
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

// ModifyVpcPrefixListWithCallback invokes the vpc.ModifyVpcPrefixList API asynchronously
func (client *Client) ModifyVpcPrefixListWithCallback(request *ModifyVpcPrefixListRequest, callback func(response *ModifyVpcPrefixListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpcPrefixListResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpcPrefixList(request)
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

// ModifyVpcPrefixListRequest is the request struct for api ModifyVpcPrefixList
type ModifyVpcPrefixListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer                            `position:"Query" name:"ResourceOwnerId"`
	ClientToken           string                                      `position:"Query" name:"ClientToken"`
	MaxEntries            requests.Integer                            `position:"Query" name:"MaxEntries"`
	RemovePrefixListEntry *[]ModifyVpcPrefixListRemovePrefixListEntry `position:"Query" name:"RemovePrefixListEntry"  type:"Repeated"`
	PrefixListId          string                                      `position:"Query" name:"PrefixListId"`
	DryRun                requests.Boolean                            `position:"Query" name:"DryRun"`
	ResourceOwnerAccount  string                                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string                                      `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer                            `position:"Query" name:"OwnerId"`
	AddPrefixListEntry    *[]ModifyVpcPrefixListAddPrefixListEntry    `position:"Query" name:"AddPrefixListEntry"  type:"Repeated"`
	PrefixListName        string                                      `position:"Query" name:"PrefixListName"`
	PrefixListDescription string                                      `position:"Query" name:"PrefixListDescription"`
}

// ModifyVpcPrefixListRemovePrefixListEntry is a repeated param struct in ModifyVpcPrefixListRequest
type ModifyVpcPrefixListRemovePrefixListEntry struct {
	Cidr        string `name:"Cidr"`
	Description string `name:"Description"`
}

// ModifyVpcPrefixListAddPrefixListEntry is a repeated param struct in ModifyVpcPrefixListRequest
type ModifyVpcPrefixListAddPrefixListEntry struct {
	Cidr        string `name:"Cidr"`
	Description string `name:"Description"`
}

// ModifyVpcPrefixListResponse is the response struct for api ModifyVpcPrefixList
type ModifyVpcPrefixListResponse struct {
	*responses.BaseResponse
	PrefixListId string `json:"PrefixListId" xml:"PrefixListId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVpcPrefixListRequest creates a request to invoke ModifyVpcPrefixList API
func CreateModifyVpcPrefixListRequest() (request *ModifyVpcPrefixListRequest) {
	request = &ModifyVpcPrefixListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVpcPrefixList", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyVpcPrefixListResponse creates a response to parse from ModifyVpcPrefixList response
func CreateModifyVpcPrefixListResponse() (response *ModifyVpcPrefixListResponse) {
	response = &ModifyVpcPrefixListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
