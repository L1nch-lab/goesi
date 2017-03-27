/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev16
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goesiv3

import (
	"time"
)

/* 200 ok object */
type GetCorporationsCorporationIdOk struct {

	/* id of alliance that corporation is a member of, if any */
	AllianceId int32 `json:"alliance_id,omitempty"`

	/* ceo_id integer */
	CeoId int32 `json:"ceo_id,omitempty"`

	/* corporation_description string */
	CorporationDescription string `json:"corporation_description,omitempty"`

	/* the full name of the corporation */
	CorporationName string `json:"corporation_name,omitempty"`

	/* creation_date string */
	CreationDate time.Time `json:"creation_date,omitempty"`

	/* creator_id integer */
	CreatorId int32 `json:"creator_id,omitempty"`

	/* faction string */
	Faction string `json:"faction,omitempty"`

	/* member_count integer */
	MemberCount int32 `json:"member_count,omitempty"`

	/* tax_rate number */
	TaxRate float32 `json:"tax_rate,omitempty"`

	/* the short name of the corporation */
	Ticker string `json:"ticker,omitempty"`

	/* url string */
	Url string `json:"url,omitempty"`
}
