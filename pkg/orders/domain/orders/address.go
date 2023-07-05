package orders

import "errors"

type Address struct{
	name string
	street	string
	city	string
	postCode	string
	country	string
}

func NewAddress(name string, street string, city string, postCode string , country string) (*Address,error) {
	if len(name) == 0{
		return nil,errors.New("Empty name")
	}

	if len(street) == 0{
		return nil,errors.New("Empty street")
	}

	if len(city) == 0{
		return nil,errors.New("Empty city")
	}

	if len(postCode) == 0{
		return nil,errors.New("Empty postCode")
	}

	if len(country) == 0{
		return nil,errors.New("Empty country")
	}
	return &Address{
		name: name,
		street: street,
		city: city,
		postCode: postCode,
		country: country,
	},nil
}

func (a *Address) Name() string{
	return a.name
}

func (a *Address) City() string{
	return a.city
}

func (a *Address) Street() string{
	return a.street
}

func (a *Address) PostCode() string{
	return a.postCode
}

func (a *Address) Country() string{
	return a.country
}