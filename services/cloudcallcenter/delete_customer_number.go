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

// DeleteCustomerNumber invokes the cloudcallcenter.DeleteCustomerNumber API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletecustomernumber.html
func (client *Client) DeleteCustomerNumber(request *DeleteCustomerNumberRequest) (response *DeleteCustomerNumberResponse, err error) {
	response = CreateDeleteCustomerNumberResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCustomerNumberWithChan invokes the cloudcallcenter.DeleteCustomerNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletecustomernumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCustomerNumberWithChan(request *DeleteCustomerNumberRequest) (<-chan *DeleteCustomerNumberResponse, <-chan error) {
	responseChan := make(chan *DeleteCustomerNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCustomerNumber(request)
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

// DeleteCustomerNumberWithCallback invokes the cloudcallcenter.DeleteCustomerNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletecustomernumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCustomerNumberWithCallback(request *DeleteCustomerNumberRequest, callback func(response *DeleteCustomerNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCustomerNumberResponse
		var err error
		defer close(result)
		response, err = client.DeleteCustomerNumber(request)
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

// DeleteCustomerNumberRequest is the request struct for api DeleteCustomerNumber
type DeleteCustomerNumberRequest struct {
	*requests.RpcRequest
	Provider    string `position:"Query" name:"Provider"`
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
}

// DeleteCustomerNumberResponse is the response struct for api DeleteCustomerNumber
type DeleteCustomerNumberResponse struct {
	*responses.BaseResponse
	RequestId      string                             `json:"RequestId" xml:"RequestId"`
	Success        bool                               `json:"Success" xml:"Success"`
	Code           string                             `json:"Code" xml:"Code"`
	Message        string                             `json:"Message" xml:"Message"`
	HttpStatusCode int                                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumbers   PhoneNumbersInDeleteCustomerNumber `json:"PhoneNumbers" xml:"PhoneNumbers"`
}

// CreateDeleteCustomerNumberRequest creates a request to invoke DeleteCustomerNumber API
func CreateDeleteCustomerNumberRequest() (request *DeleteCustomerNumberRequest) {
	request = &DeleteCustomerNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeleteCustomerNumber", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCustomerNumberResponse creates a response to parse from DeleteCustomerNumber response
func CreateDeleteCustomerNumberResponse() (response *DeleteCustomerNumberResponse) {
	response = &DeleteCustomerNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
