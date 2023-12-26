package ubsms

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

// SetUserBusinessStatus invokes the ubsms.SetUserBusinessStatus API synchronously
// api document: https://help.aliyun.com/api/ubsms/setuserbusinessstatus.html
func (client *Client) SetUserBusinessStatus(request *SetUserBusinessStatusRequest) (response *SetUserBusinessStatusResponse, err error) {
	response = CreateSetUserBusinessStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SetUserBusinessStatusWithChan invokes the ubsms.SetUserBusinessStatus API asynchronously
// api document: https://help.aliyun.com/api/ubsms/setuserbusinessstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUserBusinessStatusWithChan(request *SetUserBusinessStatusRequest) (<-chan *SetUserBusinessStatusResponse, <-chan error) {
	responseChan := make(chan *SetUserBusinessStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetUserBusinessStatus(request)
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

// SetUserBusinessStatusWithCallback invokes the ubsms.SetUserBusinessStatus API asynchronously
// api document: https://help.aliyun.com/api/ubsms/setuserbusinessstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetUserBusinessStatusWithCallback(request *SetUserBusinessStatusRequest, callback func(response *SetUserBusinessStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetUserBusinessStatusResponse
		var err error
		defer close(result)
		response, err = client.SetUserBusinessStatus(request)
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

// SetUserBusinessStatusRequest is the request struct for api SetUserBusinessStatus
type SetUserBusinessStatusRequest struct {
	*requests.RpcRequest
	Uid         string `position:"Query" name:"Uid"`
	StatusValue string `position:"Query" name:"StatusValue"`
	Service     string `position:"Query" name:"Service"`
	StatusKey   string `position:"Query" name:"StatusKey"`
}

// SetUserBusinessStatusResponse is the response struct for api SetUserBusinessStatus
type SetUserBusinessStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSetUserBusinessStatusRequest creates a request to invoke SetUserBusinessStatus API
func CreateSetUserBusinessStatusRequest() (request *SetUserBusinessStatusRequest) {
	request = &SetUserBusinessStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ubsms", "2015-06-23", "SetUserBusinessStatus", "ubsms", "openAPI")
	return
}

// CreateSetUserBusinessStatusResponse creates a response to parse from SetUserBusinessStatus response
func CreateSetUserBusinessStatusResponse() (response *SetUserBusinessStatusResponse) {
	response = &SetUserBusinessStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}