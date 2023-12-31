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

// CreatePrivacyNumberAccount invokes the cloudcallcenter.CreatePrivacyNumberAccount API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createprivacynumberaccount.html
func (client *Client) CreatePrivacyNumberAccount(request *CreatePrivacyNumberAccountRequest) (response *CreatePrivacyNumberAccountResponse, err error) {
	response = CreateCreatePrivacyNumberAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePrivacyNumberAccountWithChan invokes the cloudcallcenter.CreatePrivacyNumberAccount API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createprivacynumberaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePrivacyNumberAccountWithChan(request *CreatePrivacyNumberAccountRequest) (<-chan *CreatePrivacyNumberAccountResponse, <-chan error) {
	responseChan := make(chan *CreatePrivacyNumberAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePrivacyNumberAccount(request)
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

// CreatePrivacyNumberAccountWithCallback invokes the cloudcallcenter.CreatePrivacyNumberAccount API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createprivacynumberaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePrivacyNumberAccountWithCallback(request *CreatePrivacyNumberAccountRequest, callback func(response *CreatePrivacyNumberAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePrivacyNumberAccountResponse
		var err error
		defer close(result)
		response, err = client.CreatePrivacyNumberAccount(request)
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

// CreatePrivacyNumberAccountRequest is the request struct for api CreatePrivacyNumberAccount
type CreatePrivacyNumberAccountRequest struct {
	*requests.RpcRequest
	AccountId  string `position:"Query" name:"AccountId"`
	ProviderId string `position:"Query" name:"ProviderId"`
	AuthToken  string `position:"Query" name:"AuthToken"`
}

// CreatePrivacyNumberAccountResponse is the response struct for api CreatePrivacyNumberAccount
type CreatePrivacyNumberAccountResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreatePrivacyNumberAccountRequest creates a request to invoke CreatePrivacyNumberAccount API
func CreateCreatePrivacyNumberAccountRequest() (request *CreatePrivacyNumberAccountRequest) {
	request = &CreatePrivacyNumberAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreatePrivacyNumberAccount", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePrivacyNumberAccountResponse creates a response to parse from CreatePrivacyNumberAccount response
func CreateCreatePrivacyNumberAccountResponse() (response *CreatePrivacyNumberAccountResponse) {
	response = &CreatePrivacyNumberAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
