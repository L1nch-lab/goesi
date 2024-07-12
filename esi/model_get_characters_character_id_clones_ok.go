/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.24
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

/* A list of GetCharactersCharacterIdClonesOk. */
//easyjson:json
type GetCharactersCharacterIdClonesOkList []GetCharactersCharacterIdClonesOk

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdClonesOk struct {
	HomeLocation          GetCharactersCharacterIdClonesHomeLocation `json:"home_location,omitempty"`
	JumpClones            []GetCharactersCharacterIdClonesJumpClone  `json:"jump_clones,omitempty"`              /* jump_clones array */
	LastCloneJumpDate     time.Time                                  `json:"last_clone_jump_date,omitempty"`     /* last_clone_jump_date string */
	LastStationChangeDate time.Time                                  `json:"last_station_change_date,omitempty"` /* last_station_change_date string */
}
