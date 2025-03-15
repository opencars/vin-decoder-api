package service

import (
	"errors"
	"strings"
	"time"

	"github.com/opencars/seedwork/logger"
	"github.com/opencars/vin-decoder-api/pkg/domain"
	"github.com/opencars/vin-decoder-api/pkg/domain/model"
)

const (
	chars   = "ABCDEFGHIJKLMNOPRSTUVWXYZ1234567890"
	yearSym = "ABCDEFGHJKLMNPRSTVWXY123456789"
)

var weights = []int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}

func Parse(value string) *VIN {
	return &VIN{
		wmi: value[0:3],
		vds: value[3:9],
		vis: value[9:17],
	}
}

type VIN struct {
	wmi string // The World Manufacturer Identifier (WMI) code.
	vds string // The Vehicle Descriptor Section (VDS) code.
	vis string // The Vehicle Identifier Section (VIS) code.
}

func IndexOf(lexeme string) int {
	return strings.IndexByte(chars, lexeme[0])*len(chars) + strings.IndexByte(chars, lexeme[1])
}

func (vin *VIN) WMI() string {
	return vin.wmi
}

func (vin *VIN) VDS() string {
	return vin.vds
}

func (vin *VIN) VIS() string {
	return vin.vis
}

func (vin *VIN) String() string {
	return vin.wmi + vin.vds + vin.vis
}

// Obtain the 2-character region code for the manufacturing region.
func (vin VIN) Region() model.Region {
	region := vin.wmi[0]

	switch {
	case region >= 'A' && region <= 'H':
		return model.Africa
	case region >= 'J' && region <= 'R':
		return model.Asia
	case region >= 'S' && region <= 'Z':
		return model.Europe
	case region >= '1' && region <= '5':
		return model.NorthAmerica
	case region >= '6' && region <= '7':
		return model.Oceania
	case region >= '8' && region <= '9':
		return model.SouthAmerica
	}

	return "Unknown"
}

// Extract the single-character model year from the [number].
func (vin VIN) ModelYear() rune {
	return []rune(vin.String())[9]
}

// Extract the single-character assembly plant designator from the [number].
func (vin VIN) AssemblyPlant() rune {
	return []rune(vin.vis)[0]
}

// Extract the serial number from the [number].
func (vin VIN) SerialNumber() string {
	return vin.vis[2:]
}

func value(b byte) int {
	if b >= '0' && b <= '9' {
		return int(b) - '0'
	}

	switch b {
	case 'A', 'J':
		return 1
	case 'B', 'K', 'S':
		return 2
	case 'C', 'L', 'T':
		return 3
	case 'D', 'M', 'U':
		return 4
	case 'E', 'N', 'V':
		return 5
	case 'F', 'W':
		return 6
	case 'G', 'P', 'X':
		return 7
	case 'H', 'Y':
		return 8
	case 'R', 'Z':
		return 9
	}

	return -1
}

func (vin VIN) Check() bool {
	if vin.Region() != model.NorthAmerica {
		return true
	}

	sum := 0
	str := vin.String()
	for i := range str {
		prod := value(str[i]) * weights[i]
		sum += prod
	}

	res := sum % 11
	if res == 10 {
		return str[8] == 'X'
	} else {
		return res+'0' == int(str[8])
	}
}

func (vin VIN) Year() *uint {
	year := time.Now().Year() + 1
	yearIndex := (year - 2010) % len(yearSym)

	i := strings.IndexRune(yearSym, vin.ModelYear())
	if i == -1 {
		return nil
	}

	if i <= yearIndex {
		res := uint(year - (yearIndex - i))
		return &res
	}

	res := uint(2010 - len(yearSym) + i)
	return &res
}

func (vin VIN) Manufacturer(repo domain.ManufacturerRepository) string {
	manufacturer, err := repo.FindByWMI(vin.wmi)
	if errors.Is(err, model.ErrManufacturerNotFound) {
		return "Unknown"
	}

	if err != nil {
		logger.Errorf("failed to retrive manufacturer name: %s", err)
		return "Unknown"
	}

	return manufacturer.Name
}

func (vin VIN) Country() *Country {
	qi := IndexOf(vin.wmi[:2])
	for _, country := range countries {
		i := IndexOf(country.From)
		j := IndexOf(country.To)
		if qi >= i && qi <= j {
			return &country
		}
	}

	return nil
}
