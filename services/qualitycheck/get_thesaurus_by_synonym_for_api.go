package qualitycheck

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

// GetThesaurusBySynonymForApi invokes the qualitycheck.GetThesaurusBySynonymForApi API synchronously
func (client *Client) GetThesaurusBySynonymForApi(request *GetThesaurusBySynonymForApiRequest) (response *GetThesaurusBySynonymForApiResponse, err error) {
	response = CreateGetThesaurusBySynonymForApiResponse()
	err = client.DoAction(request, response)
	return
}

// GetThesaurusBySynonymForApiWithChan invokes the qualitycheck.GetThesaurusBySynonymForApi API asynchronously
func (client *Client) GetThesaurusBySynonymForApiWithChan(request *GetThesaurusBySynonymForApiRequest) (<-chan *GetThesaurusBySynonymForApiResponse, <-chan error) {
	responseChan := make(chan *GetThesaurusBySynonymForApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetThesaurusBySynonymForApi(request)
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

// GetThesaurusBySynonymForApiWithCallback invokes the qualitycheck.GetThesaurusBySynonymForApi API asynchronously
func (client *Client) GetThesaurusBySynonymForApiWithCallback(request *GetThesaurusBySynonymForApiRequest, callback func(response *GetThesaurusBySynonymForApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetThesaurusBySynonymForApiResponse
		var err error
		defer close(result)
		response, err = client.GetThesaurusBySynonymForApi(request)
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

// GetThesaurusBySynonymForApiRequest is the request struct for api GetThesaurusBySynonymForApi
type GetThesaurusBySynonymForApiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetThesaurusBySynonymForApiResponse is the response struct for api GetThesaurusBySynonymForApi
type GetThesaurusBySynonymForApiResponse struct {
	*responses.BaseResponse
	Code      string                            `json:"Code" xml:"Code"`
	Message   string                            `json:"Message" xml:"Message"`
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Success   bool                              `json:"Success" xml:"Success"`
	Data      DataInGetThesaurusBySynonymForApi `json:"Data" xml:"Data"`
}

// CreateGetThesaurusBySynonymForApiRequest creates a request to invoke GetThesaurusBySynonymForApi API
func CreateGetThesaurusBySynonymForApiRequest() (request *GetThesaurusBySynonymForApiRequest) {
	request = &GetThesaurusBySynonymForApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetThesaurusBySynonymForApi", "", "")
	request.Method = requests.POST
	return
}

// CreateGetThesaurusBySynonymForApiResponse creates a response to parse from GetThesaurusBySynonymForApi response
func CreateGetThesaurusBySynonymForApiResponse() (response *GetThesaurusBySynonymForApiResponse) {
	response = &GetThesaurusBySynonymForApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
