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

// GetCabJobByPhoneNumber invokes the cloudcallcenter.GetCabJobByPhoneNumber API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabjobbyphonenumber.html
func (client *Client) GetCabJobByPhoneNumber(request *GetCabJobByPhoneNumberRequest) (response *GetCabJobByPhoneNumberResponse, err error) {
	response = CreateGetCabJobByPhoneNumberResponse()
	err = client.DoAction(request, response)
	return
}

// GetCabJobByPhoneNumberWithChan invokes the cloudcallcenter.GetCabJobByPhoneNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabjobbyphonenumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCabJobByPhoneNumberWithChan(request *GetCabJobByPhoneNumberRequest) (<-chan *GetCabJobByPhoneNumberResponse, <-chan error) {
	responseChan := make(chan *GetCabJobByPhoneNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCabJobByPhoneNumber(request)
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

// GetCabJobByPhoneNumberWithCallback invokes the cloudcallcenter.GetCabJobByPhoneNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabjobbyphonenumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCabJobByPhoneNumberWithCallback(request *GetCabJobByPhoneNumberRequest, callback func(response *GetCabJobByPhoneNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCabJobByPhoneNumberResponse
		var err error
		defer close(result)
		response, err = client.GetCabJobByPhoneNumber(request)
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

// GetCabJobByPhoneNumberRequest is the request struct for api GetCabJobByPhoneNumber
type GetCabJobByPhoneNumberRequest struct {
	*requests.RpcRequest
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	JobGroupId  string `position:"Query" name:"JobGroupId"`
}

// GetCabJobByPhoneNumberResponse is the response struct for api GetCabJobByPhoneNumber
type GetCabJobByPhoneNumberResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Jobs           []Job  `json:"Jobs" xml:"Jobs"`
}

// CreateGetCabJobByPhoneNumberRequest creates a request to invoke GetCabJobByPhoneNumber API
func CreateGetCabJobByPhoneNumberRequest() (request *GetCabJobByPhoneNumberRequest) {
	request = &GetCabJobByPhoneNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetCabJobByPhoneNumber", "", "")
	request.Method = requests.POST
	return
}

// CreateGetCabJobByPhoneNumberResponse creates a response to parse from GetCabJobByPhoneNumber response
func CreateGetCabJobByPhoneNumberResponse() (response *GetCabJobByPhoneNumberResponse) {
	response = &GetCabJobByPhoneNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
