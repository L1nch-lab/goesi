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

/* A list of GetFwLeaderboardsCorporationsKills. */
//easyjson:json
type GetFwLeaderboardsCorporationsKillsList []GetFwLeaderboardsCorporationsKills

/* Top 10 rankings of corporations by number of kills from yesterday, last week and in total */
//easyjson:json
type GetFwLeaderboardsCorporationsKills struct {
	ActiveTotal []GetFwLeaderboardsCorporationsActiveTotalActiveTotal `json:"active_total,omitempty"` /* Top 10 ranking of corporations active in faction warfare by total kills. A corporation is considered \"active\" if they have participated in faction warfare in the past 14 days */
	LastWeek    []GetFwLeaderboardsCorporationsLastWeekLastWeek       `json:"last_week,omitempty"`    /* Top 10 ranking of corporations by kills in the past week */
	Yesterday   []GetFwLeaderboardsCorporationsYesterdayYesterday     `json:"yesterday,omitempty"`    /* Top 10 ranking of corporations by kills in the past day */
}
