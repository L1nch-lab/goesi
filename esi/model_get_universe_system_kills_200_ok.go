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

/* A list of GetUniverseSystemKills200Ok. */
//easyjson:json
type GetUniverseSystemKills200OkList []GetUniverseSystemKills200Ok

/* 200 ok object */
//easyjson:json
type GetUniverseSystemKills200Ok struct {
	NpcKills  int32 `json:"npc_kills,omitempty"`  /* Number of NPC ships killed in this system */
	PodKills  int32 `json:"pod_kills,omitempty"`  /* Number of pods killed in this system */
	ShipKills int32 `json:"ship_kills,omitempty"` /* Number of player ships killed in this system */
	SystemId  int32 `json:"system_id,omitempty"`  /* system_id integer */
}
