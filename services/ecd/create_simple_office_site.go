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

// CreateSimpleOfficeSite invokes the ecd.CreateSimpleOfficeSite API synchronously
func (client *Client) CreateSimpleOfficeSite(request *CreateSimpleOfficeSiteRequest) (response *CreateSimpleOfficeSiteResponse, err error) {
	response = CreateCreateSimpleOfficeSiteResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSimpleOfficeSiteWithChan invokes the ecd.CreateSimpleOfficeSite API asynchronously
func (client *Client) CreateSimpleOfficeSiteWithChan(request *CreateSimpleOfficeSiteRequest) (<-chan *CreateSimpleOfficeSiteResponse, <-chan error) {
	responseChan := make(chan *CreateSimpleOfficeSiteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSimpleOfficeSite(request)
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

// CreateSimpleOfficeSiteWithCallback invokes the ecd.CreateSimpleOfficeSite API asynchronously
func (client *Client) CreateSimpleOfficeSiteWithCallback(request *CreateSimpleOfficeSiteRequest, callback func(response *CreateSimpleOfficeSiteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSimpleOfficeSiteResponse
		var err error
		defer close(result)
		response, err = client.CreateSimpleOfficeSite(request)
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

// CreateSimpleOfficeSiteRequest is the request struct for api CreateSimpleOfficeSite
type CreateSimpleOfficeSiteRequest struct {
	*requests.RpcRequest
	CenId                string           `position:"Query" name:"CenId"`
	CenOwnerId           requests.Integer `position:"Query" name:"CenOwnerId"`
	EnableInternetAccess requests.Boolean `position:"Query" name:"EnableInternetAccess"`
	VerifyCode           string           `position:"Query" name:"VerifyCode"`
	NeedVerifyZeroDevice requests.Boolean `position:"Query" name:"NeedVerifyZeroDevice"`
	EnableAdminAccess    requests.Boolean `position:"Query" name:"EnableAdminAccess"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	DesktopAccessType    string           `position:"Query" name:"DesktopAccessType"`
	OfficeSiteName       string           `position:"Query" name:"OfficeSiteName"`
	CloudBoxOfficeSite   requests.Boolean `position:"Query" name:"CloudBoxOfficeSite"`
	VSwitchId            *[]string        `position:"Query" name:"VSwitchId"  type:"Repeated"`
	CidrBlock            string           `position:"Query" name:"CidrBlock"`
}

// CreateSimpleOfficeSiteResponse is the response struct for api CreateSimpleOfficeSite
type CreateSimpleOfficeSiteResponse struct {
	*responses.BaseResponse
	OfficeSiteId string `json:"OfficeSiteId" xml:"OfficeSiteId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateSimpleOfficeSiteRequest creates a request to invoke CreateSimpleOfficeSite API
func CreateCreateSimpleOfficeSiteRequest() (request *CreateSimpleOfficeSiteRequest) {
	request = &CreateSimpleOfficeSiteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateSimpleOfficeSite", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSimpleOfficeSiteResponse creates a response to parse from CreateSimpleOfficeSite response
func CreateCreateSimpleOfficeSiteResponse() (response *CreateSimpleOfficeSiteResponse) {
	response = &CreateSimpleOfficeSiteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
