package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Promotion struct {
	XMLName xml.Name `xml:"Root"`
	Text    string   `xml:",chardata"`
	Entry   []struct {
		Text            string `xml:",chardata"`
		ZType           string `xml:"zType"`
		Name            string `xml:"Name"`
		PromotionPrereq string `xml:"PromotionPrereq"`
		NationPrereq    string `xml:"NationPrereq"`
		EffectUnit      string `xml:"EffectUnit"`
		BPriority       string `xml:"bPriority"`
		BUpgrade        string `xml:"bUpgrade"`
	} `xml:"Entry"`
}

func readPromotions() Promotion {

	xmlFile, err := os.Open("promotion.xml")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var promotions Promotion
	xml.Unmarshal(byteValue, &promotions)

	defer xmlFile.Close()

	return promotions
}

type EffectUnit struct {
	XMLName xml.Name `xml:"Root"`
	Text    string   `xml:",chardata"`
	Entry   []struct {
		Text                         string `xml:",chardata"`
		ZType                        string `xml:"zType"`
		Name                         string `xml:"Name"`
		ZIconName                    string `xml:"zIconName"`
		Class                        string `xml:"Class"`
		SourceUnitTrait              string `xml:"SourceUnitTrait"`
		IClassNum                    string `xml:"iClassNum"`
		IStrengthModifier            string `xml:"iStrengthModifier"`
		IAttackModifier              string `xml:"iAttackModifier"`
		IDefenseModifier             string `xml:"iDefenseModifier"`
		ICriticalChance              string `xml:"iCriticalChance"`
		IEnlistOnKillChance          string `xml:"iEnlistOnKillChance"`
		IVisionExtra                 string `xml:"iVisionExtra"`
		IMovementExtra               string `xml:"iMovementExtra"`
		IFatigueExtra                string `xml:"iFatigueExtra"`
		IRangeExtra                  string `xml:"iRangeExtra"`
		IWaterControlExtra           string `xml:"iWaterControlExtra"`
		IPillageYieldModifier        string `xml:"iPillageYieldModifier"`
		IRiverAttackModifier         string `xml:"iRiverAttackModifier"`
		IWaterLandAttackModifier     string `xml:"iWaterLandAttackModifier"`
		IHomeModifier                string `xml:"iHomeModifier"`
		ISettlementAttackModifier    string `xml:"iSettlementAttackModifier"`
		IUrbanAttackModifier         string `xml:"iUrbanAttackModifier"`
		IUrbanDefenseModifier        string `xml:"iUrbanDefenseModifier"`
		IDamagedUsModifier           string `xml:"iDamagedUsModifier"`
		IDamagedThemModifier         string `xml:"iDamagedThemModifier"`
		IGeneralModifier             string `xml:"iGeneralModifier"`
		IFlankingAttackModifier      string `xml:"iFlankingAttackModifier"`
		IAdjacentSameModifier        string `xml:"iAdjacentSameModifier"`
		IAdjacentSameDefenseModifier string `xml:"iAdjacentSameDefenseModifier"`
		IHealExtra                   string `xml:"iHealExtra"`
		IHealAlways                  string `xml:"iHealAlways"`
		BSkipIcon                    string `xml:"bSkipIcon"`
		BMultiTeams                  string `xml:"bMultiTeams"`
		BMeleeCounter                string `xml:"bMeleeCounter"`
		BRout                        string `xml:"bRout"`
		BPush                        string `xml:"bPush"`
		BPushWater                   string `xml:"bPushWater"`
		BStun                        string `xml:"bStun"`
		BStunWater                   string `xml:"bStunWater"`
		BLastStand                   string `xml:"bLastStand"`
		BIgnoresDistance             string `xml:"bIgnoresDistance"`
		BCriticalImmune              string `xml:"bCriticalImmune"`
		BHealNeutral                 string `xml:"bHealNeutral"`
		BLaunchOffensive             string `xml:"bLaunchOffensive"`
		BNoRoadCooldown              string `xml:"bNoRoadCooldown"`
		AiAttackValue                struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiAttackValue"`
		AiAttackPercent struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiAttackPercent"`
		AiTerrainFromModifier struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiTerrainFromModifier"`
		AiHeightFromModifier struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiHeightFromModifier"`
		AiVegetationFromModifier struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiVegetationFromModifier"`
		AiUnitTraitModifier struct {
			Text string `xml:",chardata"`
			Pair []struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiUnitTraitModifier"`
		AiUnitTraitModifierAttack struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiUnitTraitModifierAttack"`
		AiUnitTraitModifierDefense struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiUnitTraitModifierDefense"`
		AiUnitTraitModifierMelee struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiUnitTraitModifierMelee"`
		AiMilitaryKillYield struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				IValue string `xml:"iValue"`
			} `xml:"Pair"`
		} `xml:"aiMilitaryKillYield"`
		AbHideVegetation struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				BValue string `xml:"bValue"`
			} `xml:"Pair"`
		} `xml:"abHideVegetation"`
		AbUnitTraitValid struct {
			Text string `xml:",chardata"`
			Pair []struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				BValue string `xml:"bValue"`
			} `xml:"Pair"`
		} `xml:"abUnitTraitValid"`
		AbUnitTraitInvalid struct {
			Text string `xml:",chardata"`
			Pair struct {
				Text   string `xml:",chardata"`
				ZIndex string `xml:"zIndex"`
				BValue string `xml:"bValue"`
			} `xml:"Pair"`
		} `xml:"abUnitTraitInvalid"`
		AeUnitTraitZOC struct {
			Text   string `xml:",chardata"`
			ZValue string `xml:"zValue"`
		} `xml:"aeUnitTraitZOC"`
	} `xml:"Entry"`
}

func readEffectUnits() EffectUnit {

	xmlFile, err := os.Open("effectUnit.xml")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var effectUnits EffectUnit
	xml.Unmarshal(byteValue, &effectUnits)

	defer xmlFile.Close()

	return effectUnits
}

type TextEffectUnit struct {
	XMLName xml.Name `xml:"Root"`
	Text    string   `xml:",chardata"`
	Entry   []struct {
		Text    string `xml:",chardata"`
		ZType   string `xml:"zType"`
		English string `xml:"English"`
	} `xml:"Entry"`
}

func readTextEffectUnits() TextEffectUnit {

	xmlFile, err := os.Open("text-effectUnit.xml")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var textEffectUnits TextEffectUnit
	xml.Unmarshal(byteValue, &textEffectUnits)

	defer xmlFile.Close()

	return textEffectUnits
}

type TextPromotion struct {
	XMLName xml.Name `xml:"Root"`
	Text    string   `xml:",chardata"`
	Entry   []struct {
		Text    string `xml:",chardata"`
		ZType   string `xml:"zType"`
		English string `xml:"English"`
	} `xml:"Entry"`
}

func readTextPromotions() TextPromotion {

	xmlFile, err := os.Open("text-promotion.xml")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var textPromotions TextPromotion
	xml.Unmarshal(byteValue, &textPromotions)

	defer xmlFile.Close()

	return textPromotions
}

func lookupTableTextPromotion(tp TextPromotion) map[string]string {
	m := make(map[string]string)

	for _, v := range tp.Entry {
		m[v.ZType] = v.English
	}

	return m
}

func lookupTableTextEffectUnit(te TextEffectUnit) map[string]string {
	m := make(map[string]string)

	for _, v := range te.Entry {
		m[v.ZType] = v.English
	}

	return m
}

type UnitEffects struct {
	Infantry string
	Melee    string
	Ranged   string
	Siege    string
	Ship     string
	Effects  string
}

func lookupTableEffectUnit(eu EffectUnit) map[string]UnitEffects {
	m := make(map[string]UnitEffects)

	for _, v := range eu.Entry {
		var u UnitEffects
		for _, w := range v.AbUnitTraitValid.Pair {
			switch w.ZIndex {
			case "UNITTRAIT_INFANTRY":
				u.Infantry = "Infantry"
			case "UNITTRAIT_MELEE":
				u.Melee = "Melee"
			case "UNITTRAIT_RANGED":
				u.Ranged = "Ranged"
			case "UNITTRAIT_SIEGE":
				u.Siege = "Siege"
			case "UNITTRAIT_SHIP":
				u.Ship = "Ship"
			default:
			}
		}

		if v.IStrengthModifier != "" {
			d, err := strconv.Atoi(v.IStrengthModifier)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d%% Strength", d)
			}
		}

		if v.IAttackModifier != "" {
			d, err := strconv.Atoi(v.IAttackModifier)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d%% Attack", d)
			}
		}

		if v.IDefenseModifier != "" {
			d, err := strconv.Atoi(v.IDefenseModifier)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d%% Defense", d)
			}
		}

		if v.ICriticalChance != "" {
			d, err := strconv.Atoi(v.ICriticalChance)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d%% Crit", d)
			}
		}

		if v.IVisionExtra != "" {
			d, err := strconv.Atoi(v.IVisionExtra)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d Vision", d)
			}
		}

		if v.IMovementExtra != "" {
			d, err := strconv.Atoi(v.IMovementExtra)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d Movement", d)
			}
		}

		if v.BIgnoresDistance != "" {
			d, err := strconv.Atoi(v.BIgnoresDistance)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("Ignores distance penalty")
			}
		}

		if v.IRangeExtra != "" {
			d, err := strconv.Atoi(v.IRangeExtra)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d Range", d)
			}
		}

		if v.IRiverAttackModifier != "" {
			d, err := strconv.Atoi(v.IRiverAttackModifier)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d%% Cross-River Attack", d)
			}
		}

		if v.IWaterLandAttackModifier != "" {
			d, err := strconv.Atoi(v.IWaterLandAttackModifier)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d%% Water-to-Land Attack", d)
			}
		}

		if v.IUrbanAttackModifier != "" {
			d, err := strconv.Atoi(v.IUrbanAttackModifier)
			if err != nil {
				panic(err)
			}
			if d > 0 {
				u.Effects += fmt.Sprintf("+%d%% Into Urban Attack", d)
			}
		}

		m[v.ZType] = u
	}

	return m

}

func main() {
	promos := readPromotions()
	tp := lookupTableTextPromotion(readTextPromotions())

	eu := lookupTableEffectUnit(readEffectUnits())
	//	te := lookupTableTextEffectUnit(readTextEffectUnits())

	// Promotion
	// Prereq (if any)
	// Infantry
	// Melee
	// Ranged
	// Siege
	// Ship
	// Effects

	fmt.Printf("Promotion,Prereq,Infantry,Melee,Ranged,Siege,Ship,Effects\n")

	for _, v := range promos.Entry {
		fmt.Printf("[[%s]],", tp[v.Name])
		fmt.Printf("%s,", tp["TEXT_"+v.PromotionPrereq])
		fmt.Printf("%s,", eu[v.EffectUnit].Infantry)
		fmt.Printf("%s,", eu[v.EffectUnit].Melee)
		fmt.Printf("%s,", eu[v.EffectUnit].Ranged)
		fmt.Printf("%s,", eu[v.EffectUnit].Siege)
		fmt.Printf("%s,", eu[v.EffectUnit].Ship)
		fmt.Printf("%s", eu[v.EffectUnit].Effects)
		fmt.Printf("\n")
	}

}
