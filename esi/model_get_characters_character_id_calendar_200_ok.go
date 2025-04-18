/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.30
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

package esi

import (
	"time"
)

/* A list of GetCharactersCharacterIdCalendar200Ok. */
//easyjson:json
type GetCharactersCharacterIdCalendar200OkList []GetCharactersCharacterIdCalendar200Ok

/* event */
//easyjson:json
type GetCharactersCharacterIdCalendar200Ok struct {
	EventDate     time.Time `json:"event_date,omitempty"`     /* event_date string */
	EventId       int32     `json:"event_id,omitempty"`       /* event_id integer */
	EventResponse string    `json:"event_response,omitempty"` /* event_response string */
	Importance    int32     `json:"importance,omitempty"`     /* importance integer */
	Title         string    `json:"title,omitempty"`          /* title string */
}
