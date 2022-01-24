package helpers

import (
	"encoding/json"
	"io/ioutil"
)

// LoadData laadt het JSON bestand op bestandslocatie `f` en converteert de structuur naar variabele `vv` welke van een bijpassend type is.
// De functie levert een error wanneer er iets misgaat met het uitlezen van het bestand of het converteren van de JSON structuur naar de gewenste type.
func LoadData(f string, vv interface{}) error {

	// Alle data uit het gegeven bestandspath (variabele `f`) wordt opgehaald (`ioutil.ReadFile`)
	raw, err := ioutil.ReadFile(f)

	// Geeft error terug indien bestand niet is uitgelezen.
	if err != nil {
		return err
	}

	// JSON data wordt omgezet met `json.Unmarshal` naar een gewenste data structuur.
	// Het type interface{} van variabele `vv` (zoals `[]StructNaam`) wordt omgezet in het gespecificeerde type.
	err = json.Unmarshal(raw, &vv)


	// Geeft error terug indien JSON structuur niet kon worden omgezet.
	if err != nil {
		return err
	}

	// Geen errors.
	return nil
}