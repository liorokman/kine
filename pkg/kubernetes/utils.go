package kubernetes

import (
	"strings"

	"github.com/sirupsen/logrus"
)

/*
  ExtractLabels receives a key from the doctored version of the APIServer that appears like this:

  @k=v%k=v%k=v...@returnedkey

  The first section is an associated set of labels to go with the key.

  For insert/update operations, it's the values to be inserted for the record.
  For list operations, it's the label selector to be used.
  For delete operations, it's ignored.

  As select operators, currently only simple selectors are supported. Only '=', no '||'.
  TODO: support all relevant operations.

*/
func ExtractLabels(key string) (labels []string, returnedKey string, keyModified bool) {
	if !strings.Contains(key, "@") {
		return nil, key, false
	}
	parts := strings.Split(key, "@")
	if len(parts) == 1 {
		return nil, key, false
	}
	logrus.Infof("Extracted labels: %s from key %s", parts[1], parts[2])
	return strings.Split(parts[1], "%"), parts[2], true
}
