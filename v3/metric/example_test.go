package metric_test

import (
	"fmt"

	"github.com/goark/go-cvss/v3/metric"
)

func ExampleBase_Decode() {
	m, err := metric.NewBase().Decode("CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N") //CVE-2015-8252
	if err != nil {
		return
	}
	fmt.Println("Score =", m.Score())
	fmt.Println("Severity =", m.Severity())
	//Output
	//Score = 7.5
	//Severity = High
}

/* Contributed by Florent Viel, 2020 */
/* Copyright 2018-2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
