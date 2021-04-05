package kubernetes

import "testing"

func TestSplit(t *testing.T) {
	key := "@foo=bar%boo=baa%test=value@/registry/namespaces/mynamespace"

	labels, realKey, withLabels := ExtractLabels(key)

	if !withLabels {
		t.Error("Expected labels, none found")
	}

	if realKey != "/registry/namespaces/mynamespace" {
		t.Errorf("key is wrong. Received %s", realKey)
	}

	if len(labels) != 3 {
		t.Errorf("Wrong number of labels. Have %d values", len(labels))
	}

	if labels[0] != "foo=bar" || labels[1] != "boo=baa" || labels[2] != "test=value" {
		t.Errorf("Bad labels returned")
	}
}
