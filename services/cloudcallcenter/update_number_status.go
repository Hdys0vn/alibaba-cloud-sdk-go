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

// UpdateNumberStatus invokes the cloudcallcenter.UpdateNumberStatus API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/updatenumberstatus.html
func (client *Client) UpdateNumberStatus(request *UpdateNumberStatusRequest) (response *UpdateNumberStatusResponse, err error) {
	response = CreateUpdateNumberStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateNumberStatusWithChan invokes the cloudcallcenter.UpdateNumberStatus API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/updatenumberstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateNumberStatusWithChan(request *UpdateNumberStatusRequest) (<-chan *UpdateNumberStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateNumberStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateNumberStatus(request)
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

// UpdateNumberStatusWithCallback invokes the cloudcallcenter.UpdateNumberStatus API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/updatenumberstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateNumberStatusWithCallback(request *UpdateNumberStatusRequest, callback func(response *UpdateNumberStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateNumberStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateNumberStatus(request)
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

// UpdateNumberStatusRequest is the request struct for api UpdateNumberStatus
type UpdateNumberStatusRequest struct {
	*requests.RpcRequest
	NumberStatus *[]UpdateNumberStatusNumberStatus `position:"Query" name:"numberStatus"  type:"Repeated"`
}

// UpdateNumberStatusNumberStatus is a repeated param struct in UpdateNumberStatusRequest
type UpdateNumberStatusNumberStatus struct {
	Number  string `name:"Number"`
	Inbound string `name:"Inbound"`
	Enable  string `name:"Enable"`
	RamId   string `name:"RamId"`
	Tenant  string `name:"Tenant"`
}

// UpdateNumberStatusResponse is the response struct for api UpdateNumberStatus
type UpdateNumberStatusResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ResultList     ResultList `json:"resultList" xml:"resultList"`
}

// CreateUpdateNumberStatusRequest creates a request to invoke UpdateNumberStatus API
func CreateUpdateNumberStatusRequest() (request *UpdateNumberStatusRequest) {
	request = &UpdateNumberStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "UpdateNumberStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateNumberStatusResponse creates a response to parse from UpdateNumberStatus response
func CreateUpdateNumberStatusResponse() (response *UpdateNumberStatusResponse) {
	response = &UpdateNumberStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
