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

// GetQccSearchCertification invokes the dt_oc_info.GetQccSearchCertification API synchronously
func (client *Client) GetQccSearchCertification(request *GetQccSearchCertificationRequest) (response *GetQccSearchCertificationResponse, err error) {
	response = CreateGetQccSearchCertificationResponse()
	err = client.DoAction(request, response)
	return
}

// GetQccSearchCertificationWithChan invokes the dt_oc_info.GetQccSearchCertification API asynchronously
func (client *Client) GetQccSearchCertificationWithChan(request *GetQccSearchCertificationRequest) (<-chan *GetQccSearchCertificationResponse, <-chan error) {
	responseChan := make(chan *GetQccSearchCertificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQccSearchCertification(request)
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

// GetQccSearchCertificationWithCallback invokes the dt_oc_info.GetQccSearchCertification API asynchronously
func (client *Client) GetQccSearchCertificationWithCallback(request *GetQccSearchCertificationRequest, callback func(response *GetQccSearchCertificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQccSearchCertificationResponse
		var err error
		defer close(result)
		response, err = client.GetQccSearchCertification(request)
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

// GetQccSearchCertificationRequest is the request struct for api GetQccSearchCertification
type GetQccSearchCertificationRequest struct {
	*requests.RpcRequest
	CertCategory string           `position:"Body" name:"CertCategory"`
	EntName      string           `position:"Body" name:"EntName"`
	PageNo       requests.Integer `position:"Body" name:"PageNo"`
	PageSize     requests.Integer `position:"Body" name:"PageSize"`
}

// GetQccSearchCertificationResponse is the response struct for api GetQccSearchCertification
type GetQccSearchCertificationResponse struct {
	*responses.BaseResponse
	Code      string                   `json:"Code" xml:"Code"`
	Success   bool                     `json:"Success" xml:"Success"`
	Message   string                   `json:"Message" xml:"Message"`
	TotalNum  int                      `json:"TotalNum" xml:"TotalNum"`
	PageIndex int                      `json:"PageIndex" xml:"PageIndex"`
	PageNum   int                      `json:"PageNum" xml:"PageNum"`
	RequestId string                   `json:"RequestId" xml:"RequestId"`
	Data      []map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateGetQccSearchCertificationRequest creates a request to invoke GetQccSearchCertification API
func CreateGetQccSearchCertificationRequest() (request *GetQccSearchCertificationRequest) {
	request = &GetQccSearchCertificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dt-oc-info", "2022-08-29", "GetQccSearchCertification", "", "")
	request.Method = requests.POST
	return
}

// CreateGetQccSearchCertificationResponse creates a response to parse from GetQccSearchCertification response
func CreateGetQccSearchCertificationResponse() (response *GetQccSearchCertificationResponse) {
	response = &GetQccSearchCertificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}