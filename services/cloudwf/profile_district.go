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

// ProfileDistrict invokes the cloudwf.ProfileDistrict API synchronously
// api document: https://help.aliyun.com/api/cloudwf/profiledistrict.html
func (client *Client) ProfileDistrict(request *ProfileDistrictRequest) (response *ProfileDistrictResponse, err error) {
	response = CreateProfileDistrictResponse()
	err = client.DoAction(request, response)
	return
}

// ProfileDistrictWithChan invokes the cloudwf.ProfileDistrict API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/profiledistrict.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProfileDistrictWithChan(request *ProfileDistrictRequest) (<-chan *ProfileDistrictResponse, <-chan error) {
	responseChan := make(chan *ProfileDistrictResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ProfileDistrict(request)
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

// ProfileDistrictWithCallback invokes the cloudwf.ProfileDistrict API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/profiledistrict.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ProfileDistrictWithCallback(request *ProfileDistrictRequest, callback func(response *ProfileDistrictResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ProfileDistrictResponse
		var err error
		defer close(result)
		response, err = client.ProfileDistrict(request)
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

// ProfileDistrictRequest is the request struct for api ProfileDistrict
type ProfileDistrictRequest struct {
	*requests.RpcRequest
	BeginDate string           `position:"Query" name:"BeginDate"`
	EndDate   string           `position:"Query" name:"EndDate"`
	DataType  requests.Integer `position:"Query" name:"DataType"`
	Gsid      requests.Integer `position:"Query" name:"Gsid"`
}

// ProfileDistrictResponse is the response struct for api ProfileDistrict
type ProfileDistrictResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateProfileDistrictRequest creates a request to invoke ProfileDistrict API
func CreateProfileDistrictRequest() (request *ProfileDistrictRequest) {
	request = &ProfileDistrictRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileDistrict", "cloudwf", "openAPI")
	return
}

// CreateProfileDistrictResponse creates a response to parse from ProfileDistrict response
func CreateProfileDistrictResponse() (response *ProfileDistrictResponse) {
	response = &ProfileDistrictResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
