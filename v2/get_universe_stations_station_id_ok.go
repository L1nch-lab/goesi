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

package goesiv2

/* 200 ok object */
type GetUniverseStationsStationIdOk struct {

	/* max_dockable_ship_volume number */
	MaxDockableShipVolume float32 `json:"max_dockable_ship_volume,omitempty"`

	/* name string */
	Name string `json:"name,omitempty"`

	/* office_rental_cost number */
	OfficeRentalCost float32 `json:"office_rental_cost,omitempty"`

	/* ID of the corporation that controls this station */
	Owner int32 `json:"owner,omitempty"`

	Position GetUniverseStationsStationIdPosition `json:"position,omitempty"`

	/* race_id integer */
	RaceId int32 `json:"race_id,omitempty"`

	/* reprocessing_efficiency number */
	ReprocessingEfficiency float32 `json:"reprocessing_efficiency,omitempty"`

	/* reprocessing_stations_take number */
	ReprocessingStationsTake float32 `json:"reprocessing_stations_take,omitempty"`

	/* services array */
	Services []string `json:"services,omitempty"`

	/* station_id integer */
	StationId int32 `json:"station_id,omitempty"`

	/* The solar system this station is in */
	SystemId int32 `json:"system_id,omitempty"`

	/* type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
