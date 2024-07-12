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

/* A list of GetFwStats200Ok. */
//easyjson:json
type GetFwStats200OkList []GetFwStats200Ok

/* 200 ok object */
//easyjson:json
type GetFwStats200Ok struct {
	FactionId         int32                   `json:"faction_id,omitempty"` /* faction_id integer */
	Kills             GetFwStatsKills         `json:"kills,omitempty"`
	Pilots            int32                   `json:"pilots,omitempty"`             /* How many pilots fight for the given faction */
	SystemsControlled int32                   `json:"systems_controlled,omitempty"` /* The number of solar systems controlled by the given faction */
	VictoryPoints     GetFwStatsVictoryPoints `json:"victory_points,omitempty"`
}
