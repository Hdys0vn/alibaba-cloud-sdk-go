package ddoscoo

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

// DescribePortViewSourceCountries invokes the ddoscoo.DescribePortViewSourceCountries API synchronously
func (client *Client) DescribePortViewSourceCountries(request *DescribePortViewSourceCountriesRequest) (response *DescribePortViewSourceCountriesResponse, err error) {
	response = CreateDescribePortViewSourceCountriesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePortViewSourceCountriesWithChan invokes the ddoscoo.DescribePortViewSourceCountries API asynchronously
func (client *Client) DescribePortViewSourceCountriesWithChan(request *DescribePortViewSourceCountriesRequest) (<-chan *DescribePortViewSourceCountriesResponse, <-chan error) {
	responseChan := make(chan *DescribePortViewSourceCountriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePortViewSourceCountries(request)
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

// DescribePortViewSourceCountriesWithCallback invokes the ddoscoo.DescribePortViewSourceCountries API asynchronously
func (client *Client) DescribePortViewSourceCountriesWithCallback(request *DescribePortViewSourceCountriesRequest, callback func(response *DescribePortViewSourceCountriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePortViewSourceCountriesResponse
		var err error
		defer close(result)
		response, err = client.DescribePortViewSourceCountries(request)
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

// DescribePortViewSourceCountriesRequest is the request struct for api DescribePortViewSourceCountries
type DescribePortViewSourceCountriesRequest struct {
	*requests.RpcRequest
	EndTime         requests.Integer `position:"Query" name:"EndTime"`
	StartTime       requests.Integer `position:"Query" name:"StartTime"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	InstanceIds     *[]string        `position:"Query" name:"InstanceIds"  type:"Repeated"`
}

// DescribePortViewSourceCountriesResponse is the response struct for api DescribePortViewSourceCountries
type DescribePortViewSourceCountriesResponse struct {
	*responses.BaseResponse
	RequestId      string    `json:"RequestId" xml:"RequestId"`
	SourceCountrys []Country `json:"SourceCountrys" xml:"SourceCountrys"`
}

// CreateDescribePortViewSourceCountriesRequest creates a request to invoke DescribePortViewSourceCountries API
func CreateDescribePortViewSourceCountriesRequest() (request *DescribePortViewSourceCountriesRequest) {
	request = &DescribePortViewSourceCountriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribePortViewSourceCountries", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePortViewSourceCountriesResponse creates a response to parse from DescribePortViewSourceCountries response
func CreateDescribePortViewSourceCountriesResponse() (response *DescribePortViewSourceCountriesResponse) {
	response = &DescribePortViewSourceCountriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
