package dt_oc_info

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

// GetOcJusticeCourtNotice invokes the dt_oc_info.GetOcJusticeCourtNotice API synchronously
func (client *Client) GetOcJusticeCourtNotice(request *GetOcJusticeCourtNoticeRequest) (response *GetOcJusticeCourtNoticeResponse, err error) {
	response = CreateGetOcJusticeCourtNoticeResponse()
	err = client.DoAction(request, response)
	return
}

// GetOcJusticeCourtNoticeWithChan invokes the dt_oc_info.GetOcJusticeCourtNotice API asynchronously
func (client *Client) GetOcJusticeCourtNoticeWithChan(request *GetOcJusticeCourtNoticeRequest) (<-chan *GetOcJusticeCourtNoticeResponse, <-chan error) {
	responseChan := make(chan *GetOcJusticeCourtNoticeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOcJusticeCourtNotice(request)
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

// GetOcJusticeCourtNoticeWithCallback invokes the dt_oc_info.GetOcJusticeCourtNotice API asynchronously
func (client *Client) GetOcJusticeCourtNoticeWithCallback(request *GetOcJusticeCourtNoticeRequest, callback func(response *GetOcJusticeCourtNoticeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOcJusticeCourtNoticeResponse
		var err error
		defer close(result)
		response, err = client.GetOcJusticeCourtNotice(request)
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

// GetOcJusticeCourtNoticeRequest is the request struct for api GetOcJusticeCourtNotice
type GetOcJusticeCourtNoticeRequest struct {
	*requests.RpcRequest
	RequestId string           `position:"Body" name:"RequestId"`
	PageNo    requests.Integer `position:"Body" name:"PageNo"`
	PageSize  requests.Integer `position:"Body" name:"PageSize"`
	SearchKey string           `position:"Body" name:"SearchKey"`
}

// GetOcJusticeCourtNoticeResponse is the response struct for api GetOcJusticeCourtNotice
type GetOcJusticeCourtNoticeResponse struct {
	*responses.BaseResponse
	Code      string     `json:"Code" xml:"Code"`
	Success   bool       `json:"Success" xml:"Success"`
	Message   string     `json:"Message" xml:"Message"`
	TotalNum  int        `json:"TotalNum" xml:"TotalNum"`
	PageIndex int        `json:"PageIndex" xml:"PageIndex"`
	PageNum   int        `json:"PageNum" xml:"PageNum"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetOcJusticeCourtNoticeRequest creates a request to invoke GetOcJusticeCourtNotice API
func CreateGetOcJusticeCourtNoticeRequest() (request *GetOcJusticeCourtNoticeRequest) {
	request = &GetOcJusticeCourtNoticeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dt-oc-info", "2022-08-29", "GetOcJusticeCourtNotice", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOcJusticeCourtNoticeResponse creates a response to parse from GetOcJusticeCourtNotice response
func CreateGetOcJusticeCourtNoticeResponse() (response *GetOcJusticeCourtNoticeResponse) {
	response = &GetOcJusticeCourtNoticeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
