/*
 * Copyright © 2021-present Peter M. Stahl pemistahl@gmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either expressed or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lingua

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllLanguages(t *testing.T) {
	assert.Equal(
		t,
		[]Language{
			Chinese,
			English,
			French,
			German,
			Italian,
			Japanese,
			Korean,
			Portuguese,
			Russian,
			Spanish,
		},
		AllLanguages())
}

func TestAllSpokenLanguages(t *testing.T) {
	assert.Equal(
		t,
		[]Language{
			Chinese,
			English,
			French,
			German,
			Italian,
			Japanese,
			Korean,
			Portuguese,
			Russian,
			Spanish,
		},
		AllSpokenLanguages())
}

func TestAllLanguagesWithArabicScript(t *testing.T) {
	assert.Equal(t, []Language{}, AllLanguagesWithArabicScript())
}

func TestAllLanguagesWithCyrillicScript(t *testing.T) {
	assert.Equal(
		t,
		[]Language{Russian},
		AllLanguagesWithCyrillicScript())
}

func TestAllLanguagesWithDevanagariScript(t *testing.T) {
	assert.Equal(t, []Language{}, AllLanguagesWithDevanagariScript())
}

func TestAllLanguagesWithLatinScript(t *testing.T) {
	assert.Equal(
		t,
		[]Language{
			English,
			French,
			German,
			Italian,
			Portuguese,
			Spanish,
		},
		AllLanguagesWithLatinScript())
}

func TestGetLanguageFromIsoCode639_1(t *testing.T) {
	assert.Equal(t, Chinese, GetLanguageFromIsoCode639_1(ZH))
	assert.Equal(t, English, GetLanguageFromIsoCode639_1(EN))
	assert.Equal(t, Unknown, GetLanguageFromIsoCode639_1(UnknownIsoCode639_1))
}

func TestGetLanguageFromIsoCode639_3(t *testing.T) {
	assert.Equal(t, Chinese, GetLanguageFromIsoCode639_3(ZHO))
	assert.Equal(t, English, GetLanguageFromIsoCode639_3(ENG))
	assert.Equal(t, Unknown, GetLanguageFromIsoCode639_3(UnknownIsoCode639_3))
}
