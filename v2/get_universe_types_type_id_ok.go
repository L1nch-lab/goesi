/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.2.dev16
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
type GetUniverseTypesTypeIdOk struct {

	/* capacity number */
	Capacity float32 `json:"capacity,omitempty"`

	/* description string */
	Description string `json:"description,omitempty"`

	/* dogma_attributes array */
	DogmaAttributes []GetUniverseTypesTypeIdDogmaAttribute `json:"dogma_attributes,omitempty"`

	/* dogma_effects array */
	DogmaEffects []GetUniverseTypesTypeIdDogmaEffect `json:"dogma_effects,omitempty"`

	/* graphic_id integer */
	GraphicId int32 `json:"graphic_id,omitempty"`

	/* group_id integer */
	GroupId int32 `json:"group_id,omitempty"`

	/* icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`

	/* mass number */
	Mass float32 `json:"mass,omitempty"`

	/* name string */
	Name string `json:"name,omitempty"`

	/* portion_size integer */
	PortionSize int32 `json:"portion_size,omitempty"`

	/* published boolean */
	Published bool `json:"published,omitempty"`

	/* radius number */
	Radius float32 `json:"radius,omitempty"`

	/* type_id integer */
	TypeId int32 `json:"type_id,omitempty"`

	/* volume number */
	Volume float32 `json:"volume,omitempty"`
}
