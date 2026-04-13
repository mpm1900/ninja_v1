package game

type Nature string

const (
	NatureFire      Nature = "fire"
	NatureWind      Nature = "wind"
	NatureLightning Nature = "lightning"
	NatureEarth     Nature = "earth"
	NatureWater     Nature = "water"
	NatureYin       Nature = "yin"
	NatureYang      Nature = "yang"
)

type NatureSet string

const (
	NsTai       NatureSet = "tai"
	NsFire      NatureSet = NatureSet(NatureFire)
	NsWind      NatureSet = NatureSet(NatureWind)
	NsLightning NatureSet = NatureSet(NatureLightning)
	NsEarth     NatureSet = NatureSet(NatureEarth)
	NsWater     NatureSet = NatureSet(NatureWater)
	NsYin       NatureSet = NatureSet(NatureYin)
	NsYang      NatureSet = NatureSet(NatureYang)

	NsScorch    NatureSet = "scorch"
	NsLava      NatureSet = "lava"
	NsBoil      NatureSet = "boil"
	NsGale      NatureSet = "gale"
	NsMagnet    NatureSet = "magnet"
	NsIce       NatureSet = "ice"
	NsExplosion NatureSet = "explosion"
	NsStorm     NatureSet = "storm"
	NsWood      NatureSet = "wood"
	NsYinYang   NatureSet = "yinyang"
	NsParticle  NatureSet = "particle"
	NsPure      NatureSet = "pure"
	NsJashin    NatureSet = "jashin"
)

var NATURES = map[NatureSet][]Nature{
	NsTai:       {},
	NsPure:      {},
	NsFire:      {NatureFire},
	NsWind:      {NatureWind},
	NsLightning: {NatureLightning},
	NsEarth:     {NatureEarth},
	NsWater:     {NatureWater},
	NsYin:       {NatureYin},
	NsYang:      {NatureYang},

	NsScorch: {NatureFire, NatureWind},
	// ??? {NatureFire, NatureLightning}
	NsLava:      {NatureFire, NatureEarth},
	NsBoil:      {NatureFire, NatureWater},
	NsGale:      {NatureWind, NatureLightning},
	NsMagnet:    {NatureWind, NatureEarth},
	NsIce:       {NatureWind, NatureWater},
	NsExplosion: {NatureLightning, NatureEarth},
	NsStorm:     {NatureLightning, NatureWater},
	NsWood:      {NatureEarth, NatureWater},
	NsYinYang:   {NatureYin, NatureYang},
	NsParticle:  {NatureFire, NatureEarth, NatureWind},
	NsJashin:    {},
}

var ElementalCycle = map[Nature]Nature{
	NatureFire:      NatureWind,
	NatureWind:      NatureLightning,
	NatureLightning: NatureEarth,
	NatureEarth:     NatureWater,
	NatureWater:     NatureFire,
}

func NewNatureSetValues() map[Nature]float64 {
	return map[Nature]float64{
		NatureFire:      1.00,
		NatureWind:      1.00,
		NatureLightning: 1.00,
		NatureEarth:     1.00,
		NatureWater:     1.00,
		NatureYin:       1.00,
		NatureYang:      1.00,
	}
}

func GetEffectiveness(moveNature Nature, targetNature Nature) float64 {
	if ElementalCycle[moveNature] == targetNature {
		return 2.0
	}
	if ElementalCycle[targetNature] == moveNature {
		return 0.5
	}
	return 1.0
}

func MapNatures(keys []NatureSet) map[NatureSet][]Nature {
	natures := make(map[NatureSet][]Nature)
	for _, key := range keys {
		natures[key] = NATURES[key]
	}
	return natures
}

func ResolveNatures(
	input []Nature,
	damages map[Nature]float64,
	resistances map[Nature]float64,
	natures map[NatureSet][]Nature,
) float64 {
	targetNatures := make(map[Nature]struct{})
	for _, group := range natures {
		for _, nature := range group {
			targetNatures[nature] = struct{}{}
		}
	}

	total_effectiveness := 0.0
	for _, moveNature := range input {
		effectiveness := 1.0
		for targetNature := range targetNatures {
			effectiveness *= GetEffectiveness(moveNature, targetNature)
		}

		total_effectiveness += effectiveness
	}

	avg_effectiveness := total_effectiveness / float64(len(input))
	if len(input) == 0 {
		avg_effectiveness = 1
	}

	dr_ratio := 1.0
	for _, nature := range input {
		res := resistances[nature]

		if res == 0 {
			dr_ratio = 0
			break
		}
		dr_ratio = dr_ratio * damages[nature] / res
	}

	return dr_ratio * avg_effectiveness
}
