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

// ListPhoneNumbersEx invokes the cloudcallcenter.ListPhoneNumbersEx API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listphonenumbersex.html
func (client *Client) ListPhoneNumbersEx(request *ListPhoneNumbersExRequest) (response *ListPhoneNumbersExResponse, err error) {
	response = CreateListPhoneNumbersExResponse()
	err = client.DoAction(request, response)
	return
}

// ListPhoneNumbersExWithChan invokes the cloudcallcenter.ListPhoneNumbersEx API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listphonenumbersex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPhoneNumbersExWithChan(request *ListPhoneNumbersExRequest) (<-chan *ListPhoneNumbersExResponse, <-chan error) {
	responseChan := make(chan *ListPhoneNumbersExResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPhoneNumbersEx(request)
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

// ListPhoneNumbersExWithCallback invokes the cloudcallcenter.ListPhoneNumbersEx API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listphonenumbersex.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPhoneNumbersExWithCallback(request *ListPhoneNumbersExRequest, callback func(response *ListPhoneNumbersExResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPhoneNumbersExResponse
		var err error
		defer close(result)
		response, err = client.ListPhoneNumbersEx(request)
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

// ListPhoneNumbersExRequest is the request struct for api ListPhoneNumbersEx
type ListPhoneNumbersExRequest struct {
	*requests.RpcRequest
	OutboundOnly   requests.Boolean `position:"Query" name:"OutboundOnly"`
	Number         string           `position:"Query" name:"Number"`
	NumberGroupIds *[]string        `position:"Query" name:"NumberGroupIds"  type:"Repeated"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
}

// ListPhoneNumbersExResponse is the response struct for api ListPhoneNumbersEx
type ListPhoneNumbersExResponse struct {
	*responses.BaseResponse
	RequestId      string       `json:"RequestId" xml:"RequestId"`
	Success        bool         `json:"Success" xml:"Success"`
	Code           string       `json:"Code" xml:"Code"`
	Message        string       `json:"Message" xml:"Message"`
	HttpStatusCode int          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumbers   PhoneNumbers `json:"PhoneNumbers" xml:"PhoneNumbers"`
}

// CreateListPhoneNumbersExRequest creates a request to invoke ListPhoneNumbersEx API
func CreateListPhoneNumbersExRequest() (request *ListPhoneNumbersExRequest) {
	request = &ListPhoneNumbersExRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListPhoneNumbersEx", "", "")
	request.Method = requests.POST
	return
}

// CreateListPhoneNumbersExResponse creates a response to parse from ListPhoneNumbersEx response
func CreateListPhoneNumbersExResponse() (response *ListPhoneNumbersExResponse) {
	response = &ListPhoneNumbersExResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
