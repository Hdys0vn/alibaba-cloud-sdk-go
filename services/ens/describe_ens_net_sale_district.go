package ens

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

// DescribeEnsNetSaleDistrict invokes the ens.DescribeEnsNetSaleDistrict API synchronously
func (client *Client) DescribeEnsNetSaleDistrict(request *DescribeEnsNetSaleDistrictRequest) (response *DescribeEnsNetSaleDistrictResponse, err error) {
	response = CreateDescribeEnsNetSaleDistrictResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEnsNetSaleDistrictWithChan invokes the ens.DescribeEnsNetSaleDistrict API asynchronously
func (client *Client) DescribeEnsNetSaleDistrictWithChan(request *DescribeEnsNetSaleDistrictRequest) (<-chan *DescribeEnsNetSaleDistrictResponse, <-chan error) {
	responseChan := make(chan *DescribeEnsNetSaleDistrictResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEnsNetSaleDistrict(request)
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

// DescribeEnsNetSaleDistrictWithCallback invokes the ens.DescribeEnsNetSaleDistrict API asynchronously
func (client *Client) DescribeEnsNetSaleDistrictWithCallback(request *DescribeEnsNetSaleDistrictRequest, callback func(response *DescribeEnsNetSaleDistrictResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEnsNetSaleDistrictResponse
		var err error
		defer close(result)
		response, err = client.DescribeEnsNetSaleDistrict(request)
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

// DescribeEnsNetSaleDistrictRequest is the request struct for api DescribeEnsNetSaleDistrict
type DescribeEnsNetSaleDistrictRequest struct {
	*requests.RpcRequest
	NetLevelCode    string `position:"Query" name:"NetLevelCode"`
	NetDistrictCode string `position:"Query" name:"NetDistrictCode"`
}

// DescribeEnsNetSaleDistrictResponse is the response struct for api DescribeEnsNetSaleDistrict
type DescribeEnsNetSaleDistrictResponse struct {
	*responses.BaseResponse
	Code            int                                         `json:"Code" xml:"Code"`
	RequestId       string                                      `json:"RequestId" xml:"RequestId"`
	EnsNetDistricts EnsNetDistrictsInDescribeEnsNetSaleDistrict `json:"EnsNetDistricts" xml:"EnsNetDistricts"`
}

// CreateDescribeEnsNetSaleDistrictRequest creates a request to invoke DescribeEnsNetSaleDistrict API
func CreateDescribeEnsNetSaleDistrictRequest() (request *DescribeEnsNetSaleDistrictRequest) {
	request = &DescribeEnsNetSaleDistrictRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeEnsNetSaleDistrict", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEnsNetSaleDistrictResponse creates a response to parse from DescribeEnsNetSaleDistrict response
func CreateDescribeEnsNetSaleDistrictResponse() (response *DescribeEnsNetSaleDistrictResponse) {
	response = &DescribeEnsNetSaleDistrictResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
