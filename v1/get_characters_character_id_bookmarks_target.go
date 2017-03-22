/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev11
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

/* target object */
type GetCharactersCharacterIdBookmarksTarget struct {

	Coordinates GetCharactersCharacterIdBookmarksCoordinates `json:"coordinates,omitempty"`

	Item GetCharactersCharacterIdBookmarksItem `json:"item,omitempty"`

	/* location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
}
