package nrf

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/free5gc/openapi/nrf/NFManagement"
)

func TestSetHBTimer(t *testing.T) {
	configuration := NFManagement.NewConfiguration()
	configuration.SetBasePath("nrfUri1")
	client := NFManagement.NewAPIClient(configuration)
	log := logrus.New().WithField("test", "HB")

	startHBTimer(client, "nrfUri1", "instId1", "NF", 3, log)
	err := resetHBTimer("nrfUri1", 1, log)
	require.NoError(t, err)
	err = resetHBTimer("nrfUri1", 0, log)
	require.NoError(t, err)
	err = resetHBTimer("nrfUri1", -1, log)
	require.Error(t, err)
	err = resetHBTimer("nrfUri2", 1, log)
	require.Error(t, err)
}
