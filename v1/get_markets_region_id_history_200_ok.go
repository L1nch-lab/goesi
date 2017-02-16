/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.1.dev1
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

package goesiv1

/* 200 ok object */
type GetMarketsRegionIdHistory200Ok struct {

	/* average number */
	Average float32 `json:"average,omitempty"`

	/* The date of this historical statistic entry */
	Date SwaggerDateType `json:"date,omitempty"`

	/* highest number */
	Highest float32 `json:"highest,omitempty"`

	/* lowest number */
	Lowest float32 `json:"lowest,omitempty"`

	/* Total number of orders happened that day */
	OrderCount int64 `json:"order_count,omitempty"`

	/* Total */
	Volume int64 `json:"volume,omitempty"`
}
