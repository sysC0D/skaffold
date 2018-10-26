/*
Copyright 2018 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package schema

import (
	"fmt"
	"log"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/latest"
)

// Select cluster
func SelectCluster(c *latest.SkaffoldPipeline, clusters []string) error {
	byName := clustersByName(c.MultiClusters)
	for _, name := range clusters {
		cluster, present := byName[name]
		if !present {
			return fmt.Errorf("couldn't find cluster %s", name)
		}

		connexionToCluster(c, cluster)
	}

	return nil
}

func clustersByName(clusters []latest.ClusterConfig) map[string]latest.ClusterConfig {
	byName := make(map[string]latest.ClusterConfig)
	for _, cluster := range clusters {
		byName[cluster.Name] = cluster
	}
	return byName
}

func connexionToCluster(config *latest.SkaffoldPipeline, cluster latest.ClusterConfig) {
	//logrus.isOneOf("using cluster: %s", cluster.Name)
	log.Printf("Name %s", cluster.Name)
	log.Printf("Name %s", cluster.Project)
	log.Printf("Name %s", cluster.Zone)
}
