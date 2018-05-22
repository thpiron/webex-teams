package webexteams

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
	"gopkg.in/resty.v1"
)

// LicensesService is the service to communicate with the Licenses API endpoint
type LicensesService service

// License is the License definition
type License struct {
	ID            string `json:"id,omitempty"`            // License ID.
	Name          string `json:"name,omitempty"`          // License Display Name.
	TotalUnits    int    `json:"totalUnits,omitempty"`    // License quantity total.
	ConsumedUnits int    `json:"consumedUnits,omitempty"` // License quantity consumed.
}

// Licenses is the List of Licenses
type Licenses struct {
	Items []License `json:"items,omitempty"`
}

// GetLicense Shows details for a license, by ID.
/* Shows details for a license, by ID.
Specify the license ID in the licenseID parameter in the URI.

 @param licenseID License ID.
 @return License
*/
func (s *LicensesService) GetLicense(licenseID string) (*License, *resty.Response, error) {

	path := "/licenses/{licenseId}"
	path = strings.Replace(path, "{"+"licenseId"+"}", fmt.Sprintf("%v", licenseID), -1)

	response, err := RestyClient.R().
		SetResult(&License{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*License)
	return result, response, err

}

// ListLicensesQueryParams are the query params for the ListLicenses API Call
type ListLicensesQueryParams struct {
	Max int `url:"max,omitempty"` // Limit the maximum number of items in the response.
}

// ListLicenses List all licenses for a given organization.
/* List all licenses for a given organization.
If no orgID is specified, the default is the organization of the authenticated user.

 @param "max" (int) Limit the maximum number of items in the response.
 @return Licenses
*/
func (s *LicensesService) ListLicenses(queryParams *ListLicensesQueryParams) (*Licenses, *resty.Response, error) {

	path := "/licenses/"

	queryParamsString, _ := query.Values(queryParams)

	response, err := RestyClient.R().
		SetQueryString(queryParamsString.Encode()).
		SetResult(&Licenses{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Licenses)
	return result, response, err

}
