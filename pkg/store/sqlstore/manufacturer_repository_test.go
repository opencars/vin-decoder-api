package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/vin-decoder-api/pkg/domain/model"
	"github.com/opencars/vin-decoder-api/pkg/store/sqlstore"
)

func TestManufacturerRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("manufacturers")

	manufacturer := model.TestManufacturer(t)
	assert.NoError(t, s.Manufacturer().Create(manufacturer))
}

func TestManufacturerRepository_FindByWMI(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("manufacturers")

	manufacturer := model.TestManufacturer(t)
	assert.NoError(t, s.Manufacturer().Create(manufacturer))

	actual, err := s.Manufacturer().FindByWMI(manufacturer.WMI)
	assert.NoError(t, err)
	assert.Equal(t, manufacturer, actual)
}
