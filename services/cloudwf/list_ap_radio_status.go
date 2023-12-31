package cloudwf

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

// ListApRadioStatus invokes the cloudwf.ListApRadioStatus API synchronously
// api document: https://help.aliyun.com/api/cloudwf/listapradiostatus.html
func (client *Client) ListApRadioStatus(request *ListApRadioStatusRequest) (response *ListApRadioStatusResponse, err error) {
	response = CreateListApRadioStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ListApRadioStatusWithChan invokes the cloudwf.ListApRadioStatus API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listapradiostatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListApRadioStatusWithChan(request *ListApRadioStatusRequest) (<-chan *ListApRadioStatusResponse, <-chan error) {
	responseChan := make(chan *ListApRadioStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApRadioStatus(request)
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

// ListApRadioStatusWithCallback invokes the cloudwf.ListApRadioStatus API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listapradiostatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListApRadioStatusWithCallback(request *ListApRadioStatusRequest, callback func(response *ListApRadioStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApRadioStatusResponse
		var err error
		defer close(result)
		response, err = client.ListApRadioStatus(request)
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

// ListApRadioStatusRequest is the request struct for api ListApRadioStatus
type ListApRadioStatusRequest struct {
	*requests.RpcRequest
	SearchDisabled      requests.Integer `position:"Query" name:"SearchDisabled"`
	OrderCol            string           `position:"Query" name:"OrderCol"`
	SearchName          string           `position:"Query" name:"SearchName"`
	SearchChannelEquals requests.Integer `position:"Query" name:"SearchChannelEquals"`
	Length              requests.Integer `position:"Query" name:"Length"`
	SearchMac           string           `position:"Query" name:"SearchMac"`
	SearchApgroupName   string           `position:"Query" name:"SearchApgroupName"`
	PageIndex           requests.Integer `position:"Query" name:"PageIndex"`
	OrderDir            string           `position:"Query" name:"OrderDir"`
	SearchApStatus      requests.Integer `position:"Query" name:"SearchApStatus"`
}

// ListApRadioStatusResponse is the response struct for api ListApRadioStatus
type ListApRadioStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateListApRadioStatusRequest creates a request to invoke ListApRadioStatus API
func CreateListApRadioStatusRequest() (request *ListApRadioStatusRequest) {
	request = &ListApRadioStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ListApRadioStatus", "cloudwf", "openAPI")
	return
}

// CreateListApRadioStatusResponse creates a response to parse from ListApRadioStatus response
func CreateListApRadioStatusResponse() (response *ListApRadioStatusResponse) {
	response = &ListApRadioStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
