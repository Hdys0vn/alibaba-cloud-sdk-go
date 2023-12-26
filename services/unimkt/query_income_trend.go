package unimkt

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

// QueryIncomeTrend invokes the unimkt.QueryIncomeTrend API synchronously
func (client *Client) QueryIncomeTrend(request *QueryIncomeTrendRequest) (response *QueryIncomeTrendResponse, err error) {
	response = CreateQueryIncomeTrendResponse()
	err = client.DoAction(request, response)
	return
}

// QueryIncomeTrendWithChan invokes the unimkt.QueryIncomeTrend API asynchronously
func (client *Client) QueryIncomeTrendWithChan(request *QueryIncomeTrendRequest) (<-chan *QueryIncomeTrendResponse, <-chan error) {
	responseChan := make(chan *QueryIncomeTrendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryIncomeTrend(request)
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

// QueryIncomeTrendWithCallback invokes the unimkt.QueryIncomeTrend API asynchronously
func (client *Client) QueryIncomeTrendWithCallback(request *QueryIncomeTrendRequest, callback func(response *QueryIncomeTrendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryIncomeTrendResponse
		var err error
		defer close(result)
		response, err = client.QueryIncomeTrend(request)
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

// QueryIncomeTrendRequest is the request struct for api QueryIncomeTrend
type QueryIncomeTrendRequest struct {
	*requests.RpcRequest
	AdSlotType       string           `position:"Query" name:"AdSlotType"`
	StartTime        requests.Integer `position:"Query" name:"StartTime"`
	Slot             requests.Integer `position:"Query" name:"Slot"`
	UserId           string           `position:"Query" name:"UserId"`
	OriginSiteUserId string           `position:"Query" name:"OriginSiteUserId"`
	PageNumber       requests.Integer `position:"Query" name:"PageNumber"`
	MediaName        string           `position:"Query" name:"MediaName"`
	SlotDimension    string           `position:"Query" name:"SlotDimension"`
	AppName          string           `position:"Query" name:"AppName"`
	TenantId         string           `position:"Query" name:"TenantId"`
	AdSlotId         string           `position:"Query" name:"AdSlotId"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	Dimension        string           `position:"Query" name:"Dimension"`
	QueryType        string           `position:"Query" name:"QueryType"`
	Business         string           `position:"Query" name:"Business"`
	EndTime          requests.Integer `position:"Query" name:"EndTime"`
	MediaId          string           `position:"Query" name:"MediaId"`
	Environment      string           `position:"Query" name:"Environment"`
	UserSite         string           `position:"Query" name:"UserSite"`
	AdSlotName       string           `position:"Query" name:"AdSlotName"`
}

// QueryIncomeTrendResponse is the response struct for api QueryIncomeTrend
type QueryIncomeTrendResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Success    bool   `json:"Success" xml:"Success"`
	Message    string `json:"Message" xml:"Message"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Total      int64  `json:"Total" xml:"Total"`
	Model      []Data `json:"Model" xml:"Model"`
}

// CreateQueryIncomeTrendRequest creates a request to invoke QueryIncomeTrend API
func CreateQueryIncomeTrendRequest() (request *QueryIncomeTrendRequest) {
	request = &QueryIncomeTrendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "QueryIncomeTrend", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryIncomeTrendResponse creates a response to parse from QueryIncomeTrend response
func CreateQueryIncomeTrendResponse() (response *QueryIncomeTrendResponse) {
	response = &QueryIncomeTrendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
