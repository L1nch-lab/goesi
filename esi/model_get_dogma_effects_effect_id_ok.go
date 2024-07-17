/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.25
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

/* A list of GetDogmaEffectsEffectIdOk. */
//easyjson:json
type GetDogmaEffectsEffectIdOkList []GetDogmaEffectsEffectIdOk

/* 200 ok object */
//easyjson:json
type GetDogmaEffectsEffectIdOk struct {
	Description              string                            `json:"description,omitempty"`                 /* description string */
	DisallowAutoRepeat       bool                              `json:"disallow_auto_repeat,omitempty"`        /* disallow_auto_repeat boolean */
	DischargeAttributeId     int32                             `json:"discharge_attribute_id,omitempty"`      /* discharge_attribute_id integer */
	DisplayName              string                            `json:"display_name,omitempty"`                /* display_name string */
	DurationAttributeId      int32                             `json:"duration_attribute_id,omitempty"`       /* duration_attribute_id integer */
	EffectCategory           int32                             `json:"effect_category,omitempty"`             /* effect_category integer */
	EffectId                 int32                             `json:"effect_id,omitempty"`                   /* effect_id integer */
	ElectronicChance         bool                              `json:"electronic_chance,omitempty"`           /* electronic_chance boolean */
	FalloffAttributeId       int32                             `json:"falloff_attribute_id,omitempty"`        /* falloff_attribute_id integer */
	IconId                   int32                             `json:"icon_id,omitempty"`                     /* icon_id integer */
	IsAssistance             bool                              `json:"is_assistance,omitempty"`               /* is_assistance boolean */
	IsOffensive              bool                              `json:"is_offensive,omitempty"`                /* is_offensive boolean */
	IsWarpSafe               bool                              `json:"is_warp_safe,omitempty"`                /* is_warp_safe boolean */
	Modifiers                []GetDogmaEffectsEffectIdModifier `json:"modifiers,omitempty"`                   /* modifiers array */
	Name                     string                            `json:"name,omitempty"`                        /* name string */
	PostExpression           int32                             `json:"post_expression,omitempty"`             /* post_expression integer */
	PreExpression            int32                             `json:"pre_expression,omitempty"`              /* pre_expression integer */
	Published                bool                              `json:"published,omitempty"`                   /* published boolean */
	RangeAttributeId         int32                             `json:"range_attribute_id,omitempty"`          /* range_attribute_id integer */
	RangeChance              bool                              `json:"range_chance,omitempty"`                /* range_chance boolean */
	TrackingSpeedAttributeId int32                             `json:"tracking_speed_attribute_id,omitempty"` /* tracking_speed_attribute_id integer */
}
