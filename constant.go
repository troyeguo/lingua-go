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

import "regexp"

const maxNgramLength = 5

var japaneseCharacterSet = regexp.MustCompile(`^[\p{Hiragana}\p{Katakana}\p{Han}]+$`)
var multipleWhitespace = regexp.MustCompile(`\s+`)
var numbers = regexp.MustCompile(`\p{N}`)
var punctuation = regexp.MustCompile(`\p{P}`)
var letters = regexp.MustCompile(`\p{Han}|\p{Hangul}|\p{Hiragana}|\p{Katakana}|\p{L}+`)
var tokensWithOptionalWhitespace = regexp.MustCompile(`\s*(?:\p{Han}|\p{Hangul}|\p{Hiragana}|\p{Katakana}|[\p{L}'-]+)[\p{N}\p{P}]*\s*`)
var tokensWithoutWhitespace = regexp.MustCompile(`\p{Han}|\p{Hangul}|\p{Hiragana}|\p{Katakana}|\p{L}+`)

var charsToLanguagesMapping = map[string][]Language{
	"Îî":     {French},
	"Ññ":     {Spanish},
	"Ûû":     {French},
	"Ìì":     {Italian},
	"Ëë":     {French},
	"ÈèÙù":   {French, Italian},
	"Êê":     {French, Portuguese},
	"Õõ":     {Portuguese},
	"Ôô":     {French, Portuguese},
	"ЁёЫыЭэ": {Russian},
	"ЩщЪъ":   {Russian},
	"Òò":     {Italian},
	"Ää":     {German},
	"Àà":     {French, Italian, Portuguese},
	"Ââ":     {French, Portuguese},
	"Üü":     {German, Spanish},
	"Çç":     {French, Portuguese},
	"Öö":     {German},
	"Óó":     {Portuguese, Spanish},
	"ÁáÍíÚú": {Portuguese, Spanish},
	"Éé":     {French, Italian, Portuguese, Spanish},
}
