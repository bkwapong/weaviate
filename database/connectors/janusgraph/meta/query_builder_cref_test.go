/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */
package meta

import (
	"testing"

	gm "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
)

func Test_QueryBuilder_RefProps(t *testing.T) {
	tests := testCases{
		testCase{
			name: "with ref prop and only count",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name:                "InCountry",
					StatisticalAnalyses: []gm.StatisticalAnalysis{gm.Count},
				},
			},
			expectedQuery: `
				.union(
					union(
						outE("inCountry").count().as("refcount").project("count").by(select("refcount"))
					)
					.as("InCountry").project("InCountry").by(select("InCountry"))
				)
			`,
		},

		testCase{
			name: "with ref prop and: pointingTo, count",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name:                "InCountry",
					StatisticalAnalyses: []gm.StatisticalAnalysis{gm.PointingTo, gm.Count},
				},
			},
			expectedQuery: `
				.union(
					union(
						outE("inCountry").count().as("refcount").project("count").by(select("refcount"))
					)
					.as("InCountry").project("InCountry").by(select("InCountry"))
				)
			`,
		},

		testCase{
			name: "with only pointingTo",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name:                "InCountry",
					StatisticalAnalyses: []gm.StatisticalAnalysis{gm.PointingTo},
				},
			},
			expectedQuery: `
				.union()
			`,
		},
	}

	tests.AssertQuery(t, nil)
}
