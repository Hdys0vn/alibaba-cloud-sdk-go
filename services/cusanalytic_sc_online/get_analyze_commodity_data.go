package cusanalytic_sc_online

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

// GetAnalyzeCommodityData invokes the cusanalytic_sc_online.GetAnalyzeCommodityData API synchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/getanalyzecommoditydata.html
func (client *Client) GetAnalyzeCommodityData(request *GetAnalyzeCommodityDataRequest) (response *GetAnalyzeCommodityDataResponse, err error) {
	response = CreateGetAnalyzeCommodityDataResponse()
	err = client.DoAction(request, response)
	return
}

// GetAnalyzeCommodityDataWithChan invokes the cusanalytic_sc_online.GetAnalyzeCommodityData API asynchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/getanalyzecommoditydata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAnalyzeCommodityDataWithChan(request *GetAnalyzeCommodityDataRequest) (<-chan *GetAnalyzeCommodityDataResponse, <-chan error) {
	responseChan := make(chan *GetAnalyzeCommodityDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAnalyzeCommodityData(request)
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

// GetAnalyzeCommodityDataWithCallback invokes the cusanalytic_sc_online.GetAnalyzeCommodityData API asynchronously
// api document: https://help.aliyun.com/api/cusanalytic_sc_online/getanalyzecommoditydata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAnalyzeCommodityDataWithCallback(request *GetAnalyzeCommodityDataRequest, callback func(response *GetAnalyzeCommodityDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAnalyzeCommodityDataResponse
		var err error
		defer close(result)
		response, err = client.GetAnalyzeCommodityData(request)
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

// GetAnalyzeCommodityDataRequest is the request struct for api GetAnalyzeCommodityData
type GetAnalyzeCommodityDataRequest struct {
	*requests.RpcRequest
	StoreId         requests.Integer `position:"Body" name:"StoreId"`
	StartDate       string           `position:"Body" name:"StartDate"`
	EndUserCount    requests.Integer `position:"Body" name:"EndUserCount"`
	PageSize        requests.Integer `position:"Body" name:"PageSize"`
	PageIndex       requests.Integer `position:"Body" name:"PageIndex"`
	StayPeriod      requests.Integer `position:"Body" name:"StayPeriod"`
	StartUserCount  requests.Integer `position:"Body" name:"StartUserCount"`
	MinSupportCount requests.Integer `position:"Body" name:"MinSupportCount"`
	EndDate         string           `position:"Body" name:"EndDate"`
}

// GetAnalyzeCommodityDataResponse is the response struct for api GetAnalyzeCommodityData
type GetAnalyzeCommodityDataResponse struct {
	*responses.BaseResponse
	PageIndex             int                   `json:"PageIndex" xml:"PageIndex"`
	Total                 int                   `json:"Total" xml:"Total"`
	PageSize              int                   `json:"PageSize" xml:"PageSize"`
	AnalyzeCommodityItems AnalyzeCommodityItems `json:"AnalyzeCommodityItems" xml:"AnalyzeCommodityItems"`
}

// CreateGetAnalyzeCommodityDataRequest creates a request to invoke GetAnalyzeCommodityData API
func CreateGetAnalyzeCommodityDataRequest() (request *GetAnalyzeCommodityDataRequest) {
	request = &GetAnalyzeCommodityDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cusanalytic_sc_online", "2019-05-24", "GetAnalyzeCommodityData", "", "")
	return
}

// CreateGetAnalyzeCommodityDataResponse creates a response to parse from GetAnalyzeCommodityData response
func CreateGetAnalyzeCommodityDataResponse() (response *GetAnalyzeCommodityDataResponse) {
	response = &GetAnalyzeCommodityDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
