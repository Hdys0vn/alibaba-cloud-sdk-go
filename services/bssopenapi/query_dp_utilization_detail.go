package bssopenapi

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

// QueryDPUtilizationDetail invokes the bssopenapi.QueryDPUtilizationDetail API synchronously
func (client *Client) QueryDPUtilizationDetail(request *QueryDPUtilizationDetailRequest) (response *QueryDPUtilizationDetailResponse, err error) {
	response = CreateQueryDPUtilizationDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDPUtilizationDetailWithChan invokes the bssopenapi.QueryDPUtilizationDetail API asynchronously
func (client *Client) QueryDPUtilizationDetailWithChan(request *QueryDPUtilizationDetailRequest) (<-chan *QueryDPUtilizationDetailResponse, <-chan error) {
	responseChan := make(chan *QueryDPUtilizationDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDPUtilizationDetail(request)
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

// QueryDPUtilizationDetailWithCallback invokes the bssopenapi.QueryDPUtilizationDetail API asynchronously
func (client *Client) QueryDPUtilizationDetailWithCallback(request *QueryDPUtilizationDetailRequest, callback func(response *QueryDPUtilizationDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDPUtilizationDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryDPUtilizationDetail(request)
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

// QueryDPUtilizationDetailRequest is the request struct for api QueryDPUtilizationDetail
type QueryDPUtilizationDetailRequest struct {
	*requests.RpcRequest
	DeductedInstanceId string           `position:"Query" name:"DeductedInstanceId"`
	LastToken          string           `position:"Query" name:"LastToken"`
	InstanceSpec       string           `position:"Query" name:"InstanceSpec"`
	ProdCode           string           `position:"Query" name:"ProdCode"`
	EndTime            string           `position:"Query" name:"EndTime"`
	IncludeShare       requests.Boolean `position:"Query" name:"IncludeShare"`
	CommodityCode      string           `position:"Query" name:"CommodityCode"`
	StartTime          string           `position:"Query" name:"StartTime"`
	InstanceId         string           `position:"Query" name:"InstanceId"`
	Limit              requests.Integer `position:"Query" name:"Limit"`
}

// QueryDPUtilizationDetailResponse is the response struct for api QueryDPUtilizationDetail
type QueryDPUtilizationDetailResponse struct {
	*responses.BaseResponse
	Code      string                         `json:"Code" xml:"Code"`
	Message   string                         `json:"Message" xml:"Message"`
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Success   bool                           `json:"Success" xml:"Success"`
	Data      DataInQueryDPUtilizationDetail `json:"Data" xml:"Data"`
}

// CreateQueryDPUtilizationDetailRequest creates a request to invoke QueryDPUtilizationDetail API
func CreateQueryDPUtilizationDetailRequest() (request *QueryDPUtilizationDetailRequest) {
	request = &QueryDPUtilizationDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryDPUtilizationDetail", "bssopenapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDPUtilizationDetailResponse creates a response to parse from QueryDPUtilizationDetail response
func CreateQueryDPUtilizationDetailResponse() (response *QueryDPUtilizationDetailResponse) {
	response = &QueryDPUtilizationDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
