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

// ListPhoneNumbersByRamId invokes the cloudcallcenter.ListPhoneNumbersByRamId API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listphonenumbersbyramid.html
func (client *Client) ListPhoneNumbersByRamId(request *ListPhoneNumbersByRamIdRequest) (response *ListPhoneNumbersByRamIdResponse, err error) {
	response = CreateListPhoneNumbersByRamIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListPhoneNumbersByRamIdWithChan invokes the cloudcallcenter.ListPhoneNumbersByRamId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listphonenumbersbyramid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPhoneNumbersByRamIdWithChan(request *ListPhoneNumbersByRamIdRequest) (<-chan *ListPhoneNumbersByRamIdResponse, <-chan error) {
	responseChan := make(chan *ListPhoneNumbersByRamIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPhoneNumbersByRamId(request)
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

// ListPhoneNumbersByRamIdWithCallback invokes the cloudcallcenter.ListPhoneNumbersByRamId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listphonenumbersbyramid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPhoneNumbersByRamIdWithCallback(request *ListPhoneNumbersByRamIdRequest, callback func(response *ListPhoneNumbersByRamIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPhoneNumbersByRamIdResponse
		var err error
		defer close(result)
		response, err = client.ListPhoneNumbersByRamId(request)
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

// ListPhoneNumbersByRamIdRequest is the request struct for api ListPhoneNumbersByRamId
type ListPhoneNumbersByRamIdRequest struct {
	*requests.RpcRequest
	RamId requests.Integer `position:"Query" name:"RamId"`
}

// ListPhoneNumbersByRamIdResponse is the response struct for api ListPhoneNumbersByRamId
type ListPhoneNumbersByRamIdResponse struct {
	*responses.BaseResponse
	RequestId      string                                `json:"RequestId" xml:"RequestId"`
	Success        bool                                  `json:"Success" xml:"Success"`
	Code           string                                `json:"Code" xml:"Code"`
	Message        string                                `json:"Message" xml:"Message"`
	HttpStatusCode int                                   `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumbers   PhoneNumbersInListPhoneNumbersByRamId `json:"PhoneNumbers" xml:"PhoneNumbers"`
}

// CreateListPhoneNumbersByRamIdRequest creates a request to invoke ListPhoneNumbersByRamId API
func CreateListPhoneNumbersByRamIdRequest() (request *ListPhoneNumbersByRamIdRequest) {
	request = &ListPhoneNumbersByRamIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListPhoneNumbersByRamId", "", "")
	request.Method = requests.POST
	return
}

// CreateListPhoneNumbersByRamIdResponse creates a response to parse from ListPhoneNumbersByRamId response
func CreateListPhoneNumbersByRamIdResponse() (response *ListPhoneNumbersByRamIdResponse) {
	response = &ListPhoneNumbersByRamIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
