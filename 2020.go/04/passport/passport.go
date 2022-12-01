package passport

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	id             string
	birthYear      int
	issueYear      int
	expirationYear int
	height         string
	hairColor      string
	eyeColor       string
	countryID      int
}

func (p *Passport) SetID(id string) error {
	match := regexp.MustCompile(`^[0-9]{9}$`).MatchString(id)

	if !match {
		return errors.New("invalid passport id")
	}

	p.id = id
	return nil
}

func (p *Passport) SetBirthYear(year int) error {
	if year > 2002 || year < 1920 {
		return errors.New("invalid birth year")
	}

	p.birthYear = year
	return nil
}

func (p *Passport) SetExpirationYear(year int) error {
	if year > 2030 || year < 2020 {
		return errors.New("invalid expiration year")
	}

	p.expirationYear = year
	return nil
}

func (p *Passport) SetIssueYear(year int) error {
	if year > 2020 || year < 2010 {
		return errors.New("invalid issue year")
	}

	p.issueYear = year
	return nil
}

func (p *Passport) SetHeight(height string) error {
	invalidHeightError := errors.New("invalid height")

	parts := regexp.MustCompile(`^(\d{2,3})(in|cm)$`).FindStringSubmatch(height)

	if len(parts) != 3 {
		return invalidHeightError
	}

	unitAmount := parts[1]
	unit := parts[2]

	amount, err := strconv.Atoi(unitAmount)
	if err != nil {
		return err
	}

	switch unit {
	case "in":
		if amount < 59 || amount > 76 {
			return invalidHeightError
		}
		break

	case "cm":
		if amount < 150 || amount > 193 {
			return invalidHeightError
		}
		break
	default:
		return invalidHeightError
	}

	p.height = height
	return nil
}

func (p *Passport) SetHairColor(color string) error {
	match := regexp.MustCompile(`^#[a-f0-9]{6}$`).MatchString(color)

	if !match {
		return errors.New("invalid hair color")
	}

	p.hairColor = color
	return nil

}

func (p *Passport) SetEyeColor(color string) error {
	validColors := []string{
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
	}

	for _, validColor := range validColors {
		if color == validColor {
			p.eyeColor = color
			return nil
		}
	}

	return errors.New("invalid eye color")
}

func (p *Passport) IsValid() bool {
	requiredNumbers := []int{
		p.birthYear, p.issueYear, p.expirationYear,
	}

	requiredStrings := []string{
		p.id, p.height, p.hairColor, p.eyeColor,
	}

	for _, n := range requiredNumbers {
		if n == 0 {
			return false
		}
	}

	for _, s := range requiredStrings {
		if s == "" {
			return false
		}
	}

	return true
}

var BatchEntryParseError = errors.New("failed to parse batch entry")
var InvalidPassport = errors.New("invalid passport")

// Create password from batch entry. This should be the only path to construction
func FromBatchEntry(entry string) (*Passport, error) {
	result := &Passport{}
	for _, kv := range strings.Split(entry, " ") {
		parts := strings.Split(kv, ":")

		if len(parts) != 2 {
			panic(BatchEntryParseError)
		}

		key := parts[0]
		value := parts[1]

		switch key {
		case "pid":
			err := result.SetID(value)
			if err != nil {
				return nil, err
			}
			break

		case "byr":
			birthYear, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			err = result.SetBirthYear(birthYear)
			if err != nil {
				return nil, err
			}
			break

		case "eyr":
			expYear, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			err = result.SetExpirationYear(expYear)
			if err != nil {
				return nil, err
			}
			break
		case "iyr":
			issueYear, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			err = result.SetIssueYear(issueYear)

			if err != nil {
				return nil, err
			}

			break
		case "hgt":
			err := result.SetHeight(value)
			if err != nil {
				return nil, err
			}
			break
		case "hcl":
			err := result.SetHairColor(value)
			if err != nil {
				return nil, err
			}
			break
		case "ecl":
			err := result.SetEyeColor(value)
			if err != nil {
				return nil, err
			}
			break
		}
	}
	if !result.IsValid() {
		return nil, InvalidPassport
	}
	return result, nil
}
