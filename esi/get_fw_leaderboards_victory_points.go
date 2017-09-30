/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.6.0
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

/* A list of GetFwLeaderboardsVictoryPoints. */
//easyjson:json
type GetFwLeaderboardsVictoryPointsList []GetFwLeaderboardsVictoryPoints

/* Top 4 rankings of factions by victory points from yesterday, last week and in total */
//easyjson:json
type GetFwLeaderboardsVictoryPoints struct {
	ActiveTotal []GetFwLeaderboardsActiveTotal1 `json:"active_total,omitempty"` /* Top 4 ranking of factions active in faction warfare by total victory points. A faction is considered \"active\" if they have participated in faction warfare in the past 14 days. */
	LastWeek    []GetFwLeaderboardsLastWeek1    `json:"last_week,omitempty"`    /* Top 4 ranking of factions by victory points in the past week */
	Yesterday   []GetFwLeaderboardsYesterday1   `json:"yesterday,omitempty"`    /* Top 4 ranking of factions by victory points in the past day */
}
