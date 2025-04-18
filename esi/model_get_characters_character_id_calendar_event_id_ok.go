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

/* A list of GetCharactersCharacterIdCalendarEventIdOk. */
//easyjson:json
type GetCharactersCharacterIdCalendarEventIdOkList []GetCharactersCharacterIdCalendarEventIdOk

/* Full details of a specific event */
//easyjson:json
type GetCharactersCharacterIdCalendarEventIdOk struct {
	Date       time.Time `json:"date,omitempty"`       /* date string */
	Duration   int32     `json:"duration,omitempty"`   /* Length in minutes */
	EventId    int32     `json:"event_id,omitempty"`   /* event_id integer */
	Importance int32     `json:"importance,omitempty"` /* importance integer */
	OwnerId    int32     `json:"owner_id,omitempty"`   /* owner_id integer */
	OwnerName  string    `json:"owner_name,omitempty"` /* owner_name string */
	OwnerType  string    `json:"owner_type,omitempty"` /* owner_type string */
	Response   string    `json:"response,omitempty"`   /* response string */
	Text       string    `json:"text,omitempty"`       /* text string */
	Title      string    `json:"title,omitempty"`      /* title string */
}
