package cloudcallcenter

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

// UpdateTenant invokes the cloudcallcenter.UpdateTenant API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/updatetenant.html
func (client *Client) UpdateTenant(request *UpdateTenantRequest) (response *UpdateTenantResponse, err error) {
	response = CreateUpdateTenantResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTenantWithChan invokes the cloudcallcenter.UpdateTenant API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/updatetenant.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateTenantWithChan(request *UpdateTenantRequest) (<-chan *UpdateTenantResponse, <-chan error) {
	responseChan := make(chan *UpdateTenantResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTenant(request)
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

// UpdateTenantWithCallback invokes the cloudcallcenter.UpdateTenant API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/updatetenant.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateTenantWithCallback(request *UpdateTenantRequest, callback func(response *UpdateTenantResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTenantResponse
		var err error
		defer close(result)
		response, err = client.UpdateTenant(request)
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

// UpdateTenantRequest is the request struct for api UpdateTenant
type UpdateTenantRequest struct {
	*requests.RpcRequest
	UpdateTenantData *[]UpdateTenantUpdateTenantData `position:"Query" name:"UpdateTenantData"  type:"Repeated"`
}

// UpdateTenantUpdateTenantData is a repeated param struct in UpdateTenantRequest
type UpdateTenantUpdateTenantData struct {
	TenantId string `name:"tenantId"`
	RamId    string `name:"ramId"`
	Status   string `name:"status"`
}

// UpdateTenantResponse is the response struct for api UpdateTenant
type UpdateTenantResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	Success                bool                   `json:"Success" xml:"Success"`
	Code                   string                 `json:"Code" xml:"Code"`
	Message                string                 `json:"Message" xml:"Message"`
	HttpStatusCode         int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	UpdateTenantResultList UpdateTenantResultList `json:"updateTenantResultList" xml:"updateTenantResultList"`
}

// CreateUpdateTenantRequest creates a request to invoke UpdateTenant API
func CreateUpdateTenantRequest() (request *UpdateTenantRequest) {
	request = &UpdateTenantRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "UpdateTenant", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateTenantResponse creates a response to parse from UpdateTenant response
func CreateUpdateTenantResponse() (response *UpdateTenantResponse) {
	response = &UpdateTenantResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
