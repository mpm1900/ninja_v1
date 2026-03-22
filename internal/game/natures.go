package game

type Nature string

const (
	NaturePure		Nature = "pure"
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
	NsPure 		NatureSet = NatureSet(NaturePure)
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
	NsMagnet    NatureSet = "magnet"
	NsIce       NatureSet = "ice"
	NsExplosion NatureSet = "explosion"
	NsStorm     NatureSet = "storm"
	NsWood      NatureSet = "wood"
	NsYinYang   NatureSet = "yinyang"
	NsDust      NatureSet = "dust"
)

var NATURES = map[NatureSet][]Nature{
	NsFire:      {NatureFire},
	NsWind:      {NatureWind},
	NsLightning: {NatureLightning},
	NsEarth:     {NatureEarth},
	NsWater:     {NatureWater},
	NsYin:       {NatureYin},
	NsYang:      {NatureYang},

	NsScorch: {NatureFire, NatureWind},
	// ??? {NatureFire, NatureLightning}
	NsLava: {NatureFire, NatureEarth},
	NsBoil: {NatureFire, NatureWater},
	// ??? {NatureWind, NatureLightning}
	NsMagnet:    {NatureWind, NatureEarth},
	NsIce:       {NatureWind, NatureWater},
	NsExplosion: {NatureLightning, NatureEarth},
	NsStorm:     {NatureLightning, NatureWater},
	NsWood:      {NatureEarth, NatureWater},
	NsYinYang:   {NatureYin, NatureYang},
	NsDust:      {NatureFire, NatureEarth, NatureWind},
}

func MapNatures(keys []NatureSet) map[NatureSet][]Nature {
	natures := make(map[NatureSet][]Nature)
	for _, key := range keys {
		natures[key] = NATURES[key]
	}
	return natures
}
